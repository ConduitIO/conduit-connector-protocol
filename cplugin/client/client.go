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
	"fmt"
	"net"
	"os/exec"

	"github.com/conduitio/conduit-connector-protocol/cplugin"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewClient creates a new plugin client. The provided context is used to kill
// the process (by calling os.Process.Kill) if the context becomes done before
// the plugin completes on its own. Path should point to the plugin executable.
func NewClient(
	logger hclog.Logger,
	path string,
	opts ...ClientOption,
) (*plugin.Client, error) {
	cmd := exec.Command(path)
	// NB: we give cmd a clean env here by setting Env to an empty slice
	cmd.Env = make([]string, 0)

	clientConfig := &plugin.ClientConfig{
		HandshakeConfig: cplugin.HandshakeConfig,
		VersionedPlugins: map[int]plugin.PluginSet{
			1: {
				"specifier":   &GRPCSpecifierPlugin{},
				"source":      &GRPCSourcePlugin{},
				"destination": &GRPCDestinationPlugin{},
			},
			2: {
				"specifier":   &GRPCSpecifierPlugin{},
				"source":      &GRPCSourcePlugin{},
				"destination": &GRPCDestinationPlugin{},
			},
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
		err := opt.ApplyClientOption(clientConfig)
		if err != nil {
			return nil, err
		}
	}

	return plugin.NewClient(clientConfig), nil
}

// ClientOption is an interface for defining options that can be passed to the
// NewClient function. Each implementation modifies the ClientConfig being
// generated. A slice of ClientOptions then, cumulatively applied, render a full
// ClientConfig.
type ClientOption interface {
	ApplyClientOption(*plugin.ClientConfig) error
}

type serveConfigFunc func(*plugin.ClientConfig) error

func (s serveConfigFunc) ApplyClientOption(in *plugin.ClientConfig) error {
	return s(in)
}

// WithReattachConfig returns a ClientOption that will set the client into debug
// mode, using the passed options to populate the go-plugin ReattachConfig.
func WithReattachConfig(config *plugin.ReattachConfig) ClientOption {
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
func WithDelve(port int) ClientOption {
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

	return l.Addr().(*net.TCPAddr).Port
}

// knownErrors contains known error messages that are mapped to internal error
// types. gRPC does not retain error types, so we have to resort to relying on
// the error message itself.
var knownErrors = map[string]error{
	"context canceled":          context.Canceled,
	"context deadline exceeded": context.DeadlineExceeded,
}

// unwrapGRPCError removes the gRPC wrapper from the error and returns a known
// error if possible, otherwise creates an internal error.
func unwrapGRPCError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	if st.Code() == codes.Unimplemented {
		return fmt.Errorf("%s: %w", st.Message(), cplugin.ErrUnimplemented)
	}
	if knownErr, ok := knownErrors[st.Message()]; ok {
		return knownErr
	}
	return errors.New(st.Message())
}
