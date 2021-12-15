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

// grpcSourcePlugin is an implementation of the
// github.com/hashicorp/go-plugin#Plugin and
// github.com/hashicorp/go-plugin#GRPCPlugin interfaces, it's using
// cpluginv1.SourcePlugin.
type grpcSourcePlugin struct {
	plugin.NetRPCUnsupportedPlugin
	SourcePluginServer func() cpluginv1.SourcePlugin
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
	connectorv1.RegisterSourcePluginServer(s, NewSourcePluginServer(p.SourcePluginServer()))
	return nil
}

func NewSourcePluginServer(impl cpluginv1.SourcePlugin) connectorv1.SourcePluginServer {
	return &sourcePluginServer{impl: impl}
}

type sourcePluginServer struct {
	connectorv1.UnimplementedSourcePluginServer
	impl cpluginv1.SourcePlugin
}

func (s *sourcePluginServer) Configure(ctx context.Context, protoReq *connectorv1.Source_Configure_Request) (*connectorv1.Source_Configure_Response, error) {
	goReq, err := fromproto.SourceConfigureRequest(protoReq)
	if err != nil {
		return nil, err
	}
	goResp, err := s.impl.Configure(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp, err := toproto.SourceConfigureResponse(goResp)
	if err != nil {
		return nil, err
	}
	return protoResp, nil
}
func (s *sourcePluginServer) Start(ctx context.Context, protoReq *connectorv1.Source_Start_Request) (*connectorv1.Source_Start_Response, error) {
	goReq, err := fromproto.SourceStartRequest(protoReq)
	if err != nil {
		return nil, err
	}
	goResp, err := s.impl.Start(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp, err := toproto.SourceStartResponse(goResp)
	if err != nil {
		return nil, err
	}
	return protoResp, nil
}
func (s *sourcePluginServer) Stop(ctx context.Context, protoReq *connectorv1.Source_Stop_Request) (*connectorv1.Source_Stop_Response, error) {
	goReq, err := fromproto.SourceStopRequest(protoReq)
	if err != nil {
		return nil, err
	}
	goResp, err := s.impl.Stop(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp, err := toproto.SourceStopResponse(goResp)
	if err != nil {
		return nil, err
	}
	return protoResp, nil
}
func (s *sourcePluginServer) Teardown(ctx context.Context, protoReq *connectorv1.Source_Teardown_Request) (*connectorv1.Source_Teardown_Response, error) {
	goReq, err := fromproto.SourceTeardownRequest(protoReq)
	if err != nil {
		return nil, err
	}
	goResp, err := s.impl.Teardown(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp, err := toproto.SourceTeardownResponse(goResp)
	if err != nil {
		return nil, err
	}
	return protoResp, nil
}
func (s *sourcePluginServer) Run(stream connectorv1.SourcePlugin_RunServer) error {
	err := s.impl.Run(stream.Context(), &sourceRunStream{impl: stream})
	if err != nil {
		return err
	}
	return nil
}

type sourceRunStream struct {
	impl connectorv1.SourcePlugin_RunServer
}

func (s *sourceRunStream) Send(in cpluginv1.SourceRunResponse) error {
	out, err := toproto.SourceRunResponse(in)
	if err != nil {
		return err
	}
	return s.impl.Send(out)
}

func (s *sourceRunStream) Recv() (cpluginv1.SourceRunRequest, error) {
	in, err := s.impl.Recv()
	if err != nil {
		return cpluginv1.SourceRunRequest{}, err
	}
	out, err := fromproto.SourceRunRequest(in)
	if err != nil {
		return cpluginv1.SourceRunRequest{}, err
	}
	return out, nil
}
