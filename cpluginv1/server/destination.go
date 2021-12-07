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

// grpcDestinationPlugin is an implementation of the
// github.com/hashicorp/go-plugin#Plugin and
// github.com/hashicorp/go-plugin#GRPCPlugin interfaces, it's using
// cpluginv1.DestinationPluginServer.
type grpcDestinationPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	DestinationPluginServer func() cpluginv1.DestinationPluginServer
}

var _ plugin.Plugin = (*grpcDestinationPlugin)(nil)

// GRPCClient always returns an error; we're only implementing the server half
// of the interface.
func (p *grpcDestinationPlugin) GRPCClient(context.Context, *plugin.GRPCBroker, *grpc.ClientConn) (interface{}, error) {
	return nil, errors.New("this package only implements gRPC servers")
}

// GRPCServer registers the gRPC destination plugin server with the gRPC server
// that go-plugin is standing up.
func (p *grpcDestinationPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	connectorv1.RegisterDestinationPluginServer(s, NewDestinationPluginServer(p.DestinationPluginServer()))
	return nil
}

func NewDestinationPluginServer(impl cpluginv1.DestinationPluginServer) connectorv1.DestinationPluginServer {
	return &destinationPluginServer{impl: impl}
}

type destinationPluginServer struct {
	connectorv1.UnimplementedDestinationPluginServer
	impl cpluginv1.DestinationPluginServer
}

func (s *destinationPluginServer) Configure(ctx context.Context, req *connectorv1.Destination_Configure_Request) (*connectorv1.Destination_Configure_Response, error) {
	r, err := fromproto.DestinationConfigureRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.impl.Configure(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.DestinationConfigureResponse(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
func (s *destinationPluginServer) Start(ctx context.Context, req *connectorv1.Destination_Start_Request) (*connectorv1.Destination_Start_Response, error) {
	r, err := fromproto.DestinationStartRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.impl.Start(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.DestinationStartResponse(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
func (s *destinationPluginServer) Run(stream connectorv1.DestinationPlugin_RunServer) error {
	err := s.impl.Run(stream.Context(), &destinationRunStream{impl: stream})
	if err != nil {
		return err
	}
	return nil
}
func (s *destinationPluginServer) Stop(ctx context.Context, req *connectorv1.Destination_Stop_Request) (*connectorv1.Destination_Stop_Response, error) {
	r, err := fromproto.DestinationStopRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.impl.Stop(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.DestinationStopResponse(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

type destinationRunStream struct {
	impl connectorv1.DestinationPlugin_RunServer
}

func (s *destinationRunStream) Send(in cpluginv1.DestinationRunResponse) error {
	out, err := toproto.DestinationRunResponse(in)
	if err != nil {
		return err
	}
	return s.impl.Send(out)
}

func (s *destinationRunStream) Recv() (cpluginv1.DestinationRunRequest, error) {
	in, err := s.impl.Recv()
	if err != nil {
		return cpluginv1.DestinationRunRequest{}, err
	}
	out, err := fromproto.DestinationRunRequest(in)
	if err != nil {
		return cpluginv1.DestinationRunRequest{}, err
	}
	return out, nil
}
