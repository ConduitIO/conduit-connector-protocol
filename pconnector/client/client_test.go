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

package client

import (
	"testing"

	v1 "github.com/conduitio/conduit-connector-protocol/pconnector/v1" //nolint:staticcheck // v1 is used for backwards compatibility
	v2 "github.com/conduitio/conduit-connector-protocol/pconnector/v2"
	v3 "github.com/conduitio/conduit-connector-protocol/pconnector/v3"
	"github.com/matryer/is"
)

// TestVersionedPlugins asserts the client offers exactly {1, 2, 3}, each with
// the full specifier/source/destination plugin set. go-plugin negotiates the
// highest version both peers share, so a server that only offers {1, 2} will
// never be driven with v3.
func TestVersionedPlugins(t *testing.T) {
	is := is.New(t)

	vp := versionedPlugins()

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
