// Copyright © 2021 Meroxa Inc
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
	"errors"

	"github.com/conduitio/conduit-plugin/cpluginv1"
	connectorv1 "github.com/conduitio/conduit-plugin/internal/connector/v1"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

func Serve(
	sourceFactory func() cpluginv1.SourcePlugin,
	// TODO specifier, destination
	opts ...ServeOption,
) error {
	serveConfig := &plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "CONDUIT_PLUGIN_MAGIC_COOKIE",
			MagicCookieValue: "204e8e812c3a1bb73b838928c575b42a105dd2e9aa449be481bc4590486df53f",
		},
		Plugins: plugin.PluginSet{
			"source": grpcSourcePlugin{factory: sourceFactory},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	}
	for _, opt := range opts {
		err := opt.ApplyServeOption(serveConfig)
		if err != nil {
			return err
		}
	}

	plugin.Serve(serveConfig)
	return nil
}

// ServeOption is an interface for defining options that can be passed to the
// Serve function. Each implementation modifies the ServeConfig being
// generated. A slice of ServeOptions then, cumulatively applied, render a full
// ServeConfig.
type ServeOption interface {
	ApplyServeOption(*plugin.ServeConfig) error
}

type serveConfigFunc func(*plugin.ServeConfig) error

func (s serveConfigFunc) ApplyServeOption(in *plugin.ServeConfig) error {
	return s(in)
}

// WithDebug returns a ServeOption that will set the server into debug mode, using
// the passed options to populate the go-plugin ServeTestConfig.
func WithDebug(ctx context.Context, config chan *plugin.ReattachConfig, closeCh chan struct{}) ServeOption {
	return serveConfigFunc(func(in *plugin.ServeConfig) error {
		in.Test.Context = ctx
		in.Test.ReattachConfigCh = config
		in.Test.CloseCh = closeCh
		return nil
	})
}

// grpcSourcePlugin is an implementation of the
// github.com/hashicorp/go-plugin#Plugin and
// github.com/hashicorp/go-plugin#GRPCPlugin interfaces, it's using
// SourcePlugin.
type grpcSourcePlugin struct {
	plugin.NetRPCUnsupportedPlugin
	factory func() cpluginv1.SourcePlugin
}

var _ plugin.Plugin = (*grpcSourcePlugin)(nil)

// GRPCClient always returns an error; we're only implementing the server half
// of the interface.
func (p *grpcSourcePlugin) GRPCClient(context.Context, *plugin.GRPCBroker, *grpc.ClientConn) (interface{}, error) {
	return nil, errors.New("this package only implements gRPC servers")
}

// GRPCServer registers the gRPC source plugin server with the gRPC server that
// go-plugin is standing up.
func (p *grpcSourcePlugin) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	connectorv1.RegisterSourcePluginServer(s, NewSourcePluginServer(p.factory()))
	return nil
}

// grpcDestinationPlugin is an implementation of the
// github.com/hashicorp/go-plugin#Plugin and
// github.com/hashicorp/go-plugin#GRPCPlugin interfaces, it's using
// cpluginv1.DestinationPlugin.
type grpcDestinationPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	factory func() cpluginv1.DestinationPlugin
}

var _ plugin.Plugin = (*grpcDestinationPlugin)(nil)

// GRPCClient always returns an error; we're only implementing the server half
// of the interface.
func (p *grpcDestinationPlugin) GRPCClient(context.Context, *plugin.GRPCBroker, *grpc.ClientConn) (interface{}, error) {
	return nil, errors.New("this package only implements gRPC servers")
}

// GRPCServer registers the gRPC destination plugin server with the gRPC server
// that go-plugin is standing up.
func (p *grpcDestinationPlugin) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	connectorv1.RegisterDestinationPluginServer(s, NewDestinationPluginServer(p.factory()))
	return nil
}

// grpcSpecifierPlugin is an implementation of the
// github.com/hashicorp/go-plugin#Plugin and
// github.com/hashicorp/go-plugin#GRPCPlugin interfaces, it's using
// cpluginv1.SpecifierPlugin.
type grpcSpecifierPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	factory func() cpluginv1.SpecifierPlugin
}

var _ plugin.Plugin = (*grpcSpecifierPlugin)(nil)

// GRPCClient always returns an error; we're only implementing the server half
// of the interface.
func (p *grpcSpecifierPlugin) GRPCClient(context.Context, *plugin.GRPCBroker, *grpc.ClientConn) (interface{}, error) {
	return nil, errors.New("this package only implements gRPC servers")
}

// GRPCServer registers the gRPC specifier plugin server with the gRPC server that
// go-plugin is standing up.
func (p *grpcSpecifierPlugin) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	connectorv1.RegisterSpecifierPluginServer(s, NewSpecifierPluginServer(p.factory()))
	return nil
}
