// Copyright © 2022 Meroxa, Inc.
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

	"buf.build/gen/go/conduitio/conduit-connector-protocol/grpc/go/connector/v1/connectorv1grpc"
	connectorv1 "buf.build/gen/go/conduitio/conduit-connector-protocol/protocolbuffers/go/connector/v1"
	"github.com/conduitio/conduit-connector-protocol/cpluginv1"
	"github.com/conduitio/conduit-connector-protocol/cpluginv1/internal/fromproto"
	"github.com/conduitio/conduit-connector-protocol/cpluginv1/internal/toproto"
)

func NewSourcePluginServer(impl cpluginv1.SourcePlugin) connectorv1grpc.SourcePluginServer {
	return &sourcePluginServer{impl: impl}
}

type sourcePluginServer struct {
	connectorv1grpc.UnimplementedSourcePluginServer
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
func (s *sourcePluginServer) Run(stream connectorv1grpc.SourcePlugin_RunServer) error {
	err := s.impl.Run(stream.Context(), &sourceRunStream{impl: stream})
	if err != nil {
		return err
	}
	return nil
}

type sourceRunStream struct {
	impl connectorv1grpc.SourcePlugin_RunServer
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
