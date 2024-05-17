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
	"errors"

	"github.com/conduitio/conduit-connector-protocol/cplugin"
	v1 "github.com/conduitio/conduit-connector-protocol/cplugin/v1"
	serverv1 "github.com/conduitio/conduit-connector-protocol/cplugin/v1/server"
	v2 "github.com/conduitio/conduit-connector-protocol/cplugin/v2"
	serverv2 "github.com/conduitio/conduit-connector-protocol/cplugin/v2/server"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var HandshakeConfig = plugin.HandshakeConfig{
	MagicCookieKey:   "CONDUIT_PLUGIN_MAGIC_COOKIE",
	MagicCookieValue: "204e8e812c3a1bb73b838928c575b42a105dd2e9aa449be481bc4590486df53f",
}

var (
	v1PluginSetFactory = newPluginSetFactory(
		connectorv1.RegisterSpecifierPluginServer,
		serverv1.NewSpecifierPluginServer,
		connectorv1.RegisterSourcePluginServer,
		serverv1.NewSourcePluginServer,
		connectorv1.RegisterDestinationPluginServer,
		serverv1.NewDestinationPluginServer,
	)
	v2PluginSetFactory = newPluginSetFactory(
		connectorv2.RegisterSpecifierPluginServer,
		serverv2.NewSpecifierPluginServer,
		connectorv2.RegisterSourcePluginServer,
		serverv2.NewSourcePluginServer,
		connectorv2.RegisterDestinationPluginServer,
		serverv2.NewDestinationPluginServer,
	)
)

func Serve(
	specifierFactory func() cplugin.SpecifierPlugin,
	sourceFactory func() cplugin.SourcePlugin,
	destinationFactory func() cplugin.DestinationPlugin,
	opts ...ServeOption,
) error {
	serveConfig := &plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		VersionedPlugins: map[int]plugin.PluginSet{
			v1.Version: v1PluginSetFactory(specifierFactory, sourceFactory, destinationFactory),
			v2.Version: v2PluginSetFactory(specifierFactory, sourceFactory, destinationFactory),
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
		in.Test = &plugin.ServeTestConfig{
			Context:          ctx,
			ReattachConfigCh: config,
			CloseCh:          closeCh,
		}
		return nil
	})
}

func WithGRPCServerOptions(opt ...grpc.ServerOption) ServeOption {
	return serveConfigFunc(func(in *plugin.ServeConfig) error {
		in.GRPCServer = func(opts []grpc.ServerOption) *grpc.Server {
			return plugin.DefaultGRPCServer(append(opts, opt...))
		}
		return nil
	})
}

// pluginSetFactory is a function that takes cplugin factories and creates a PluginSet.
type pluginSetFactory func(
	specifierFactory func() cplugin.SpecifierPlugin,
	sourceFactory func() cplugin.SourcePlugin,
	destinationFactory func() cplugin.DestinationPlugin,
) plugin.PluginSet

func newPluginSetFactory[SPEC, SRC, DST any](
	registerSpecifierPluginServer func(grpc.ServiceRegistrar, SPEC),
	newSpecifierPluginServer func(cplugin.SpecifierPlugin) SPEC,
	registerSourcePluginServer func(grpc.ServiceRegistrar, SRC),
	newSourcePluginServer func(cplugin.SourcePlugin) SRC,
	registerDestinationPluginServer func(grpc.ServiceRegistrar, DST),
	newDestinationPluginServer func(cplugin.DestinationPlugin) DST,
) pluginSetFactory {
	return func(
		specifierFactory func() cplugin.SpecifierPlugin,
		sourceFactory func() cplugin.SourcePlugin,
		destinationFactory func() cplugin.DestinationPlugin,
	) plugin.PluginSet {
		return plugin.PluginSet{
			"specifier": newGRPCPlugin(
				registerSpecifierPluginServer,
				newSpecifierPluginServer,
				specifierFactory,
			),
			"source": newGRPCPlugin(
				registerSourcePluginServer,
				newSourcePluginServer,
				sourceFactory,
			),
			"destination": newGRPCPlugin(
				registerDestinationPluginServer,
				newDestinationPluginServer,
				destinationFactory,
			),
		}
	}
}

// grpcPlugin is a generic helper for creating gRPC plugins.
type grpcPlugin[T any] struct {
	plugin.NetRPCUnsupportedPlugin
	registerPluginServer func(grpc.ServiceRegistrar, T)
	newPluginServer      func() T
}

func newGRPCPlugin[GRPC, GO any](
	registerPluginServer func(grpc.ServiceRegistrar, GRPC),
	newPluginServer func(GO) GRPC,
	pluginFactory func() GO,
) *grpcPlugin[GRPC] {
	return &grpcPlugin[GRPC]{
		registerPluginServer: registerPluginServer,
		newPluginServer: func() GRPC {
			return newPluginServer(pluginFactory())
		},
	}
}

var _ plugin.Plugin = (*grpcPlugin[any])(nil)
var _ plugin.GRPCPlugin = (*grpcPlugin[any])(nil)

// GRPCClient always returns an error; we're only implementing the server half
// of the interface.
func (p *grpcPlugin[T]) GRPCClient(context.Context, *plugin.GRPCBroker, *grpc.ClientConn) (any, error) {
	return nil, errors.New("this package only implements gRPC servers")
}

// GRPCServer registers the gRPC specifier plugin server with the gRPC server that
// go-plugin is standing up.
func (p *grpcPlugin[T]) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	p.registerPluginServer(s, p.newPluginServer())
	return nil
}
