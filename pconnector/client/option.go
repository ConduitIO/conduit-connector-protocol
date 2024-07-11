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
	"fmt"
	"github.com/conduitio/conduit-connector-protocol/pconduit"
	"net"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

// Option is an interface for defining options that can be passed to the
// NewClient function. Each implementation modifies the ClientConfig being
// generated. A slice of Options then, cumulatively applied, render a full
// ClientConfig.
type Option interface {
	ApplyOption(*plugin.ClientConfig) error
}

type serveConfigFunc func(*plugin.ClientConfig) error

func (s serveConfigFunc) ApplyOption(in *plugin.ClientConfig) error {
	return s(in)
}

// WithReattachConfig returns an Option that will set the client into debug
// mode, using the passed options to populate the go-plugin ReattachConfig.
func WithReattachConfig(config *plugin.ReattachConfig) Option {
	return serveConfigFunc(func(in *plugin.ClientConfig) error {
		in.Reattach = config
		in.Cmd = nil // only reattach or cmd can be set
		return nil
	})
}

// WithDelve runs the plugin with Delve listening on the supplied port. If the
// port is 0 it finds a random free port.
// For more information on how to use Delve refer to the official repository:
// https://github.com/go-delve/delve
func WithDelve(port int) Option {
	const delve = "dlv"
	return serveConfigFunc(func(in *plugin.ClientConfig) error {
		delvePath, err := exec.LookPath(delve)
		if err != nil {
			return err
		}

		if port == 0 {
			port = getFreePort()
		}

		pluginPath := in.Cmd.Path
		in.Cmd.Path = delvePath
		in.Cmd.Args = []string{
			delve,
			fmt.Sprintf("--listen=:%d", port),
			"--headless=true", "--api-version=2", "--accept-multiclient", "--log-dest=2",
			"exec", pluginPath,
		}

		in.Logger.Info("DELVE: starting plugin", "port", port)
		return nil
	})
}

func WithConnectorUtilsAddress(addr string) Option {
	return serveConfigFunc(func(in *plugin.ClientConfig) error {
		in.Cmd.Env = append(in.Cmd.Env, pconduit.EnvConduitConnectorUtilitiesGRPCTarget+"="+addr)
		return nil
	})
}

func WithConnectorUtilsToken(token string) Option {
	return serveConfigFunc(func(in *plugin.ClientConfig) error {
		in.Cmd.Env = append(in.Cmd.Env, pconduit.EnvConduitConnectorSchemaToken+"="+token)
		return nil
	})
}

func getFreePort() int {
	// Excerpt from net.Listen godoc:
	// If the port in the address parameter is empty or "0", as in
	// "127.0.0.1:" or "[::1]:0", a port number is automatically chosen.
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	err = l.Close()
	if err != nil {
		panic(err)
	}

	return l.Addr().(*net.TCPAddr).Port //nolint:forcetypeassert // we know it's a TCPAddr
}
