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
	"github.com/conduitio/conduit-plugin/cpluginv1/internal/fromproto"
	"github.com/conduitio/conduit-plugin/cpluginv1/internal/toproto"
	cproto "github.com/conduitio/conduit-plugin/proto/gen/go/conduitio/cplugin/v1"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// grpcSpecifierPlugin is an implementation of the
// github.com/hashicorp/go-plugin#Plugin and
// github.com/hashicorp/go-plugin#GRPCPlugin interfaces, it's using
// cpluginv1.SpecifierPluginServer.
type grpcSpecifierPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	SpecifierPluginServer func() cpluginv1.SpecifierPluginServer
}

var _ plugin.Plugin = (*grpcSpecifierPlugin)(nil)

// GRPCClient always returns an error; we're only implementing the server half
// of the interface.
func (p *grpcSpecifierPlugin) GRPCClient(context.Context, *plugin.GRPCBroker, *grpc.ClientConn) (interface{}, error) {
	return nil, errors.New("conduit-plugin only implements gRPC servers")
}

// GRPCServer registers the gRPC specifier plugin server with the gRPC server that
// go-plugin is standing up.
func (p *grpcSpecifierPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	cproto.RegisterSpecifierPluginServer(s, NewSpecifierPluginServer(p.SpecifierPluginServer()))
	return nil
}

func NewSpecifierPluginServer(impl cpluginv1.SpecifierPluginServer) cproto.SpecifierPluginServer {
	return &specifierPluginServer{impl: impl}
}

type specifierPluginServer struct {
	cproto.UnimplementedSpecifierPluginServer
	impl cpluginv1.SpecifierPluginServer
}

func (s specifierPluginServer) Specify(ctx context.Context, req *cproto.Specifier_Specify_Request) (*cproto.Specifier_Specify_Response, error) {
	r, err := fromproto.SpecifierSpecifyRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.impl.Specify(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.SpecifierSpecifyResponse(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
