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

	"github.com/conduitio/conduit-connector-protocol/pconnector"
	v1 "github.com/conduitio/conduit-connector-protocol/pconnector/v1"              //nolint:staticcheck // v1 is used for backwards compatibility
	serverv1 "github.com/conduitio/conduit-connector-protocol/pconnector/v1/server" //nolint:staticcheck // v1 is used for backwards compatibility
	v2 "github.com/conduitio/conduit-connector-protocol/pconnector/v2"
	serverv2 "github.com/conduitio/conduit-connector-protocol/pconnector/v2/server"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// Serve starts a go-plugin server with the given factories for creating
// Specifier, Source, and Destination plugins. The server will support both v1
// and v2 of the connector protocol.
func Serve(
	specifierFactory func() pconnector.SpecifierPlugin,
	sourceFactory func() pconnector.SourcePlugin,
	destinationFactory func() pconnector.DestinationPlugin,
	opts ...Option,
) error {
	serveConfig := &plugin.ServeConfig{
		HandshakeConfig: pconnector.HandshakeConfig,
		VersionedPlugins: map[int]plugin.PluginSet{
			v1.Version: v1PluginSetFactory(specifierFactory, sourceFactory, destinationFactory),
			v2.Version: v2PluginSetFactory(specifierFactory, sourceFactory, destinationFactory),
		},
	}
	for _, opt := range opts {
		err := opt.ApplyOption(serveConfig)
		if err != nil {
			return err
		}
	}

	plugin.Serve(serveConfig)
	return nil
}

var (
	// v1PluginSetFactory is a factory for creating a PluginSet for v1 of the
	// connector protocol.
	v1PluginSetFactory = newPluginSetFactory(
		connectorv1.RegisterSpecifierPluginServer,
		serverv1.NewSpecifierPluginServer,
		connectorv1.RegisterSourcePluginServer,
		serverv1.NewSourcePluginServer,
		connectorv1.RegisterDestinationPluginServer,
		serverv1.NewDestinationPluginServer,
	)
	// v2PluginSetFactory is a factory for creating a PluginSet for v2 of the
	// connector protocol.
	v2PluginSetFactory = newPluginSetFactory(
		connectorv2.RegisterSpecifierPluginServer,
		serverv2.NewSpecifierPluginServer,
		connectorv2.RegisterSourcePluginServer,
		serverv2.NewSourcePluginServer,
		connectorv2.RegisterDestinationPluginServer,
		serverv2.NewDestinationPluginServer,
	)
)

// pluginSetFactory is a function that takes pconnector factories and creates a PluginSet.
type pluginSetFactory func(
	specifierFactory func() pconnector.SpecifierPlugin,
	sourceFactory func() pconnector.SourcePlugin,
	destinationFactory func() pconnector.DestinationPlugin,
) plugin.PluginSet

func newPluginSetFactory[SPEC, SRC, DST any](
	registerSpecifierPluginServer func(grpc.ServiceRegistrar, SPEC),
	newSpecifierPluginServer func(pconnector.SpecifierPlugin) SPEC,
	registerSourcePluginServer func(grpc.ServiceRegistrar, SRC),
	newSourcePluginServer func(pconnector.SourcePlugin) SRC,
	registerDestinationPluginServer func(grpc.ServiceRegistrar, DST),
	newDestinationPluginServer func(pconnector.DestinationPlugin) DST,
) pluginSetFactory {
	return func(
		specifierFactory func() pconnector.SpecifierPlugin,
		sourceFactory func() pconnector.SourcePlugin,
		destinationFactory func() pconnector.DestinationPlugin,
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

var (
	_ plugin.Plugin     = (*grpcPlugin[any])(nil)
	_ plugin.GRPCPlugin = (*grpcPlugin[any])(nil)
)

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
