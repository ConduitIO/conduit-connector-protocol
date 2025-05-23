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

package server

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// DefaultMaxReceiveRecordSize is the default maximum size of a record that can
// be received by the connector. This is set to 4MB by default, which
// corresponds to defaultServerMaxReceiveMessageSize = 1024 * 1024 * 4 defined in
// https://github.com/grpc/grpc-go/blob/master/server.go#L59
const DefaultMaxReceiveRecordSize = 4 * 1024 * 1024 // 4MB

// Option is an interface for defining options that can be passed to the
// Serve function. Each implementation modifies the ServeConfig being
// generated. A slice of Options then, cumulatively applied, render a full
// ServeConfig.
type Option interface {
	ApplyOption(*plugin.ServeConfig) error
}

type serveConfigFunc func(*plugin.ServeConfig) error

func (s serveConfigFunc) ApplyOption(in *plugin.ServeConfig) error {
	return s(in)
}

// WithDebug returns an Option that will set the server into debug mode, using
// the passed options to populate the go-plugin ServeTestConfig.
func WithDebug(ctx context.Context, config chan *plugin.ReattachConfig, closeCh chan struct{}) Option {
	return serveConfigFunc(func(in *plugin.ServeConfig) error {
		in.Test = &plugin.ServeTestConfig{
			Context:          ctx,
			ReattachConfigCh: config,
			CloseCh:          closeCh,
		}
		return nil
	})
}

// WithGRPCServerOptions returns an Option that will set the server's gRPC
// server options (appended to the default options).
// Can be used to configure grpc.MaxRecvMsgSize for example.
func WithGRPCServerOptions(opt ...grpc.ServerOption) Option {
	return serveConfigFunc(func(in *plugin.ServeConfig) error {
		in.GRPCServer = func(opts []grpc.ServerOption) *grpc.Server {
			return plugin.DefaultGRPCServer(append(opts, opt...))
		}
		return nil
	})
}
