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

func NewDestinationPluginServer(impl pconnector.DestinationPlugin) connectorv2.DestinationPluginServer {
	return &DestinationPluginServer{impl: impl}
}

type DestinationPluginServer struct {
	connectorv2.UnimplementedDestinationPluginServer
	impl pconnector.DestinationPlugin
}

func (s *DestinationPluginServer) Configure(ctx context.Context, protoReq *connectorv2.Destination_Configure_Request) (*connectorv2.Destination_Configure_Response, error) {
	goReq := fromproto.DestinationConfigureRequest(protoReq)
	goResp, err := s.impl.Configure(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationConfigureResponse(goResp), nil
}

func (s *DestinationPluginServer) Open(ctx context.Context, protoReq *connectorv2.Destination_Open_Request) (*connectorv2.Destination_Open_Response, error) {
	goReq := fromproto.DestinationOpenRequest(protoReq)
	goResp, err := s.impl.Open(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationOpenResponse(goResp), nil
}

func (s *DestinationPluginServer) Run(stream connectorv2.DestinationPlugin_RunServer) error {
	return s.impl.Run(stream.Context(), &DestinationRunStream{stream: stream})
}

func (s *DestinationPluginServer) Stop(ctx context.Context, protoReq *connectorv2.Destination_Stop_Request) (*connectorv2.Destination_Stop_Response, error) {
	goReq := fromproto.DestinationStopRequest(protoReq)
	goResp, err := s.impl.Stop(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationStopResponse(goResp), nil
}

func (s *DestinationPluginServer) Teardown(ctx context.Context, protoReq *connectorv2.Destination_Teardown_Request) (*connectorv2.Destination_Teardown_Response, error) {
	goReq := fromproto.DestinationTeardownRequest(protoReq)
	goResp, err := s.impl.Teardown(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationTeardownResponse(goResp), nil
}

func (s *DestinationPluginServer) LifecycleOnCreated(ctx context.Context, protoReq *connectorv2.Destination_Lifecycle_OnCreated_Request) (*connectorv2.Destination_Lifecycle_OnCreated_Response, error) {
	goReq := fromproto.DestinationLifecycleOnCreatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnCreated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationLifecycleOnCreatedResponse(goResp), nil
}

func (s *DestinationPluginServer) LifecycleOnUpdated(ctx context.Context, protoReq *connectorv2.Destination_Lifecycle_OnUpdated_Request) (*connectorv2.Destination_Lifecycle_OnUpdated_Response, error) {
	goReq := fromproto.DestinationLifecycleOnUpdatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnUpdated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationLifecycleOnUpdatedResponse(goResp), nil
}

func (s *DestinationPluginServer) LifecycleOnDeleted(ctx context.Context, protoReq *connectorv2.Destination_Lifecycle_OnDeleted_Request) (*connectorv2.Destination_Lifecycle_OnDeleted_Response, error) {
	goReq := fromproto.DestinationLifecycleOnDeletedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnDeleted(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.DestinationLifecycleOnDeletedResponse(goResp), nil
}

// DestinationRunStream is the server-side implementation of the
// pconnector.DestinationRunStream interface.
type DestinationRunStream struct {
	stream connectorv2.DestinationPlugin_RunServer
}

func (s *DestinationRunStream) Client() pconnector.DestinationRunStreamClient {
	panic("invalid use of server.DestinationRunStream - it is a server-side type only")
}

func (s *DestinationRunStream) Server() pconnector.DestinationRunStreamServer {
	if s.stream == nil {
		panic("invalid use of server.DestinationRunStream - stream has not been initialized using DestinationPluginServer.Run")
	}
	return s
}

func (s *DestinationRunStream) Send(in pconnector.DestinationRunResponse) error {
	out := toproto.DestinationRunResponse(in)
	return s.stream.Send(out)
}

func (s *DestinationRunStream) Recv() (pconnector.DestinationRunRequest, error) {
	in, err := s.stream.Recv()
	if err != nil {
		return pconnector.DestinationRunRequest{}, err
	}
	out, err := fromproto.DestinationRunRequest(in)
	if err != nil {
		return pconnector.DestinationRunRequest{}, err
	}
	return out, nil
}
