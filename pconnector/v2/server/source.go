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

	"github.com/conduitio/conduit-connector-protocol/pconnector"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v2/fromproto"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v2/toproto"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
)

func NewSourcePluginServer(impl pconnector.SourcePlugin) connectorv2.SourcePluginServer {
	return &SourcePluginServer{impl: impl}
}

type SourcePluginServer struct {
	connectorv2.UnimplementedSourcePluginServer
	impl pconnector.SourcePlugin
}

func (s *SourcePluginServer) Configure(ctx context.Context, protoReq *connectorv2.Source_Configure_Request) (*connectorv2.Source_Configure_Response, error) {
	goReq := fromproto.SourceConfigureRequest(protoReq)
	goResp, err := s.impl.Configure(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceConfigureResponse(goResp), nil
}

func (s *SourcePluginServer) Open(ctx context.Context, protoReq *connectorv2.Source_Open_Request) (*connectorv2.Source_Open_Response, error) {
	goReq := fromproto.SourceOpenRequest(protoReq)
	goResp, err := s.impl.Open(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceOpenResponse(goResp), nil
}

func (s *SourcePluginServer) Run(stream connectorv2.SourcePlugin_RunServer) error {
	return s.impl.Run(stream.Context(), &SourceRunStream{stream: stream})
}

func (s *SourcePluginServer) Stop(ctx context.Context, protoReq *connectorv2.Source_Stop_Request) (*connectorv2.Source_Stop_Response, error) {
	goReq := fromproto.SourceStopRequest(protoReq)
	goResp, err := s.impl.Stop(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceStopResponse(goResp), nil
}

func (s *SourcePluginServer) Teardown(ctx context.Context, protoReq *connectorv2.Source_Teardown_Request) (*connectorv2.Source_Teardown_Response, error) {
	goReq := fromproto.SourceTeardownRequest(protoReq)
	goResp, err := s.impl.Teardown(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceTeardownResponse(goResp), nil
}

func (s *SourcePluginServer) LifecycleOnCreated(ctx context.Context, protoReq *connectorv2.Source_Lifecycle_OnCreated_Request) (*connectorv2.Source_Lifecycle_OnCreated_Response, error) {
	goReq := fromproto.SourceLifecycleOnCreatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnCreated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceLifecycleOnCreatedResponse(goResp), nil
}

func (s *SourcePluginServer) LifecycleOnUpdated(ctx context.Context, protoReq *connectorv2.Source_Lifecycle_OnUpdated_Request) (*connectorv2.Source_Lifecycle_OnUpdated_Response, error) {
	goReq := fromproto.SourceLifecycleOnUpdatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnUpdated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceLifecycleOnUpdatedResponse(goResp), nil
}

func (s *SourcePluginServer) LifecycleOnDeleted(ctx context.Context, protoReq *connectorv2.Source_Lifecycle_OnDeleted_Request) (*connectorv2.Source_Lifecycle_OnDeleted_Response, error) {
	goReq := fromproto.SourceLifecycleOnDeletedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnDeleted(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceLifecycleOnDeletedResponse(goResp), nil
}

// SourceRunStream is the server-side implementation of the
// pconnector.SourceRunStream interface.
type SourceRunStream struct {
	stream connectorv2.SourcePlugin_RunServer
}

func (s *SourceRunStream) Client() pconnector.SourceRunStreamClient {
	panic("invalid use of server.SourceRunStream - it is a server-side type only")
}

func (s *SourceRunStream) Server() pconnector.SourceRunStreamServer {
	if s.stream == nil {
		panic("invalid use of server.SourceRunStream - stream has not been initialized using SourcePluginServer.Run")
	}
	return s
}

func (s *SourceRunStream) Send(in pconnector.SourceRunResponse) error {
	out, err := toproto.SourceRunResponse(in)
	if err != nil {
		return err
	}
	return s.stream.Send(out)
}

func (s *SourceRunStream) Recv() (pconnector.SourceRunRequest, error) {
	in, err := s.stream.Recv()
	if err != nil {
		return pconnector.SourceRunRequest{}, err
	}
	out := fromproto.SourceRunRequest(in)
	if err != nil {
		return pconnector.SourceRunRequest{}, err
	}
	return out, nil
}
