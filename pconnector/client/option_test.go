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
	"os/exec"
	"testing"

	"github.com/conduitio/conduit-connector-protocol/pconduit"
	"github.com/hashicorp/go-plugin"
	"github.com/matryer/is"
)

func TestOption_WithConnectorUtilsAddress(t *testing.T) {
	is := is.New(t)
	cmd := exec.Command("test-path")
	cmd.Env = []string{"FOO=BAR"}
	cc := &plugin.ClientConfig{Cmd: cmd}

	underTest := WithEnvVar(pconduit.EnvConduitConnectorUtilitiesGRPCTarget, "localhost:12345")

	err := underTest.ApplyOption(cc)
	is.NoErr(err)

	want := []string{
		"FOO=BAR",
		pconduit.EnvConduitConnectorUtilitiesGRPCTarget + "=localhost:12345",
	}
	is.Equal(want, cc.Cmd.Env)
}
