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
	connectorv1 "github.com/conduitio/conduit-plugin/internal/connector/v1"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type grpcSpecifierPlugin struct {
	plugin.NetRPCUnsupportedPlugin
}

var _ plugin.Plugin = (*grpcSpecifierPlugin)(nil)

func (p *grpcSpecifierPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
	return &specifierPlugin{client: connectorv1.NewSpecifierPluginClient(cc)}, nil
}

// GRPCServer always returns an error; we're only implementing the client half
// of the interface.
func (p *grpcSpecifierPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	return errors.New("this package only implements gRPC clients")
}

func NewSpecifierPlugin(client connectorv1.SpecifierPluginClient) cpluginv1.SpecifierPlugin {
	return &specifierPlugin{client: client}
}

type specifierPlugin struct {
	client connectorv1.SpecifierPluginClient
}

func (s *specifierPlugin) Specify(ctx context.Context, goReq cpluginv1.SpecifierSpecifyRequest) (cpluginv1.SpecifierSpecifyResponse, error) {
	protoReq, err := toproto.SpecifierSpecifyRequest(goReq)
	if err != nil {
		return cpluginv1.SpecifierSpecifyResponse{}, err
	}
	protoResp, err := s.client.Specify(ctx, protoReq)
	if err != nil {
		return cpluginv1.SpecifierSpecifyResponse{}, err
	}
	goResp, err := fromproto.SpecifierSpecifyResponse(protoResp)
	if err != nil {
		return cpluginv1.SpecifierSpecifyResponse{}, err
	}
	return goResp, nil
}
