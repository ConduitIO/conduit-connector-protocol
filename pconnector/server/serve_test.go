// Copyright © 2026 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/conduitio/conduit-connector-protocol/pconnector"
	v1 "github.com/conduitio/conduit-connector-protocol/pconnector/v1" //nolint:staticcheck // v1 is used for backwards compatibility
	v2 "github.com/conduitio/conduit-connector-protocol/pconnector/v2"
	v3 "github.com/conduitio/conduit-connector-protocol/pconnector/v3"
	"github.com/hashicorp/go-plugin"
	"github.com/matryer/is"
)

func nilSpecifierFactory() pconnector.SpecifierPlugin     { return nil }
func nilSourceFactory() pconnector.SourcePlugin           { return nil }
func nilDestinationFactory() pconnector.DestinationPlugin { return nil }

// TestVersionedPlugins asserts the server offers exactly {1, 2, 3}, each with
// the full specifier/source/destination plugin set.
func TestVersionedPlugins(t *testing.T) {
	is := is.New(t)

	vp := versionedPlugins(nilSpecifierFactory, nilSourceFactory, nilDestinationFactory)

	is.Equal(3, len(vp))
	for _, version := range []int{v1.Version, v2.Version, v3.Version} {
		ps, ok := vp[version]
		is.True(ok) // version %d must be registered

		is.Equal(3, len(ps))
		for _, name := range []string{"specifier", "source", "destination"} {
			is.True(ps[name] != nil) // plugin set for version %d is missing %s
		}
	}
}

// TestNegotiation drives go-plugin's server-side version negotiation (the same
// code path Serve uses) and asserts the highest version both peers share is
// selected. The client's offer is passed exactly as go-plugin's client does:
// the PLUGIN_PROTOCOL_VERSIONS env var. v3 must be selected only when both
// sides offer it; an old peer (no v3) falls back to v2, never v3.
func TestNegotiation(t *testing.T) {
	newServer := versionedPlugins(nilSpecifierFactory, nilSourceFactory, nilDestinationFactory) // {1, 2, 3}
	oldServer := map[int]plugin.PluginSet{1: newServer[1], 2: newServer[2]}                     // {1, 2}

	tests := []struct {
		name           string
		clientVersions []int
		serverVersions map[int]plugin.PluginSet
		want           int
	}{
		{
			name:           "old client, new server selects v2",
			clientVersions: []int{1, 2},
			serverVersions: newServer,
			want:           2,
		},
		{
			name:           "new client, old server selects v2",
			clientVersions: []int{1, 2, 3},
			serverVersions: oldServer,
			want:           2,
		},
		{
			name:           "new client, new server selects v3",
			clientVersions: []int{1, 2, 3},
			serverVersions: newServer,
			want:           3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := negotiateVersion(t, tt.clientVersions, tt.serverVersions)
			is.Equal(tt.want, got)
		})
	}
}

// negotiateVersion starts the plugin server in-process and returns the protocol
// version it negotiated for the given client offer. It mirrors what Serve does
// (VersionedPlugins + DefaultGRPCServer) but runs in go-plugin's test mode so it
// does not take over the process.
func negotiateVersion(t *testing.T, clientVersions []int, serverVersions map[int]plugin.PluginSet) int {
	t.Helper()

	// go-plugin's client sends its supported versions through this env var
	// (see go-plugin client.go). Order does not matter; the server sorts it.
	t.Setenv("PLUGIN_PROTOCOL_VERSIONS", intsToString(clientVersions))

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	reattachCh := make(chan *plugin.ReattachConfig, 1)
	serveConfig := &plugin.ServeConfig{
		HandshakeConfig:  pconnector.HandshakeConfig,
		VersionedPlugins: serverVersions,
		GRPCServer:       plugin.DefaultGRPCServer,
		Test: &plugin.ServeTestConfig{
			Context:          ctx,
			ReattachConfigCh: reattachCh,
		},
	}

	go plugin.Serve(serveConfig)

	select {
	case rc := <-reattachCh:
		return rc.ProtocolVersion
	case <-time.After(10 * time.Second):
		t.Fatal("timed out waiting for the plugin server to negotiate a version")
		return 0
	}
}

func intsToString(vs []int) string {
	s := make([]string, len(vs))
	for i, v := range vs {
		s[i] = strconv.Itoa(v)
	}
	return strings.Join(s, ",")
}
