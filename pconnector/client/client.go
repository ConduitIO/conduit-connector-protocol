// Copyright © 2024 Meroxa, Inc.
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

package client

import (
	"context"
	"errors"
	"os/exec"

	"github.com/conduitio/conduit-connector-protocol/pconduit"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	v1 "github.com/conduitio/conduit-connector-protocol/pconnector/v1"              //nolint:staticcheck // v1 is used for backwards compatibility
	clientv1 "github.com/conduitio/conduit-connector-protocol/pconnector/v1/client" //nolint:staticcheck // v1 is used for backwards compatibility
	v2 "github.com/conduitio/conduit-connector-protocol/pconnector/v2"
	clientv2 "github.com/conduitio/conduit-connector-protocol/pconnector/v2/client"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// New creates a new plugin client. Path should point to the plugin
// executable. The client will support both v1 and v2 of the connector protocol.
func New(
	logger hclog.Logger,
	path string,
	connUtilsAddr string,
	connUtilsToken string,
	opts ...Option,
) (*plugin.Client, error) {
	cmd := exec.Command(path)
	// NB: we give cmd a clean env here by setting Env to an empty slice
	cmd.Env = []string{
		pconduit.EnvConduitConnectorUtilitiesGRPCTarget + "=" + connUtilsAddr,
		pconduit.EnvConduitConnectorSchemaToken + "=" + connUtilsToken,
	}

	clientConfig := &plugin.ClientConfig{
		HandshakeConfig: pconnector.HandshakeConfig,
		VersionedPlugins: map[int]plugin.PluginSet{
			v1.Version: newPluginSet(
				clientv1.NewSpecifierPluginClient,
				clientv1.NewSourcePluginClient,
				clientv1.NewDestinationPluginClient,
			),
			v2.Version: newPluginSet(
				clientv2.NewSpecifierPluginClient,
				clientv2.NewSourcePluginClient,
				clientv2.NewDestinationPluginClient,
			),
		},
		Cmd: cmd,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
		SyncStderr: logger.StandardWriter(nil),
		SyncStdout: logger.StandardWriter(nil),
		Logger:     logger,
	}

	for _, opt := range opts {
		err := opt.ApplyOption(clientConfig)
		if err != nil {
			return nil, err
		}
	}

	return plugin.NewClient(clientConfig), nil
}

func newPluginSet[SPEC, SRC, DST any](
	newSpecifierPluginClient func(*grpc.ClientConn) SPEC,
	newSourcePluginClient func(*grpc.ClientConn) SRC,
	newDestinationPluginClient func(*grpc.ClientConn) DST,
) plugin.PluginSet {
	return plugin.PluginSet{
		"specifier":   newGRPCPlugin(newSpecifierPluginClient),
		"source":      newGRPCPlugin(newSourcePluginClient),
		"destination": newGRPCPlugin(newDestinationPluginClient),
	}
}

type grpcPlugin[T any] struct {
	plugin.NetRPCUnsupportedPlugin
	newPluginClient func(cc *grpc.ClientConn) T
}

func newGRPCPlugin[T any](
	newPluginClient func(cc *grpc.ClientConn) T,
) *grpcPlugin[T] {
	return &grpcPlugin[T]{
		newPluginClient: newPluginClient,
	}
}

var (
	_ plugin.Plugin     = (*grpcPlugin[any])(nil)
	_ plugin.GRPCPlugin = (*grpcPlugin[any])(nil)
)

func (p *grpcPlugin[T]) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, cc *grpc.ClientConn) (any, error) {
	return p.newPluginClient(cc), nil
}

// GRPCServer always returns an error; we're only implementing the client half
// of the interface.
func (p *grpcPlugin[T]) GRPCServer(*plugin.GRPCBroker, *grpc.Server) error {
	return errors.New("this package only implements gRPC clients")
}
