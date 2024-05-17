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

	"github.com/conduitio/conduit-connector-protocol/cplugin"
	"github.com/conduitio/conduit-connector-protocol/cplugin/v1/fromproto"
	"github.com/conduitio/conduit-connector-protocol/cplugin/v1/toproto"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
)

func NewDestinationPluginServer(impl cplugin.DestinationPlugin) connectorv1.DestinationPluginServer {
	return &destinationPluginServer{impl: impl}
}

type destinationPluginServer struct {
	connectorv1.UnimplementedDestinationPluginServer
	impl cplugin.DestinationPlugin
}

func (s *destinationPluginServer) Configure(ctx context.Context, protoReq *connectorv1.Destination_Configure_Request) (*connectorv1.Destination_Configure_Response, error) {
	goReq := fromproto.DestinationConfigureRequest(protoReq)
	goResp, err := s.impl.Configure(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.DestinationConfigureResponse(goResp)
	return protoResp, nil
}
func (s *destinationPluginServer) Start(ctx context.Context, protoReq *connectorv1.Destination_Start_Request) (*connectorv1.Destination_Start_Response, error) {
	goReq := fromproto.DestinationStartRequest(protoReq)
	goResp, err := s.impl.Start(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.DestinationStartResponse(goResp)
	return protoResp, nil
}
func (s *destinationPluginServer) Run(stream connectorv1.DestinationPlugin_RunServer) error {
	err := s.impl.Run(stream.Context(), &destinationRunStream{impl: stream})
	if err != nil {
		return err
	}
	return nil
}
func (s *destinationPluginServer) Stop(ctx context.Context, protoReq *connectorv1.Destination_Stop_Request) (*connectorv1.Destination_Stop_Response, error) {
	goReq := fromproto.DestinationStopRequest(protoReq)
	goResp, err := s.impl.Stop(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.DestinationStopResponse(goResp)
	return protoResp, nil
}
func (s *destinationPluginServer) Teardown(ctx context.Context, protoReq *connectorv1.Destination_Teardown_Request) (*connectorv1.Destination_Teardown_Response, error) {
	goReq := fromproto.DestinationTeardownRequest(protoReq)
	goResp, err := s.impl.Teardown(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.DestinationTeardownResponse(goResp)
	return protoResp, nil
}
func (s *destinationPluginServer) LifecycleOnCreated(ctx context.Context, protoReq *connectorv1.Destination_Lifecycle_OnCreated_Request) (*connectorv1.Destination_Lifecycle_OnCreated_Response, error) {
	goReq := fromproto.DestinationLifecycleOnCreatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnCreated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.DestinationLifecycleOnCreatedResponse(goResp)
	return protoResp, nil
}
func (s *destinationPluginServer) LifecycleOnUpdated(ctx context.Context, protoReq *connectorv1.Destination_Lifecycle_OnUpdated_Request) (*connectorv1.Destination_Lifecycle_OnUpdated_Response, error) {
	goReq := fromproto.DestinationLifecycleOnUpdatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnUpdated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.DestinationLifecycleOnUpdatedResponse(goResp)
	return protoResp, nil
}
func (s *destinationPluginServer) LifecycleOnDeleted(ctx context.Context, protoReq *connectorv1.Destination_Lifecycle_OnDeleted_Request) (*connectorv1.Destination_Lifecycle_OnDeleted_Response, error) {
	goReq := fromproto.DestinationLifecycleOnDeletedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnDeleted(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.DestinationLifecycleOnDeletedResponse(goResp)
	return protoResp, nil
}

type destinationRunStream struct {
	impl connectorv1.DestinationPlugin_RunServer
}

func (s *destinationRunStream) Send(in cplugin.DestinationRunResponse) error {
	out := toproto.DestinationRunResponse(in)
	return s.impl.Send(out)
}

func (s *destinationRunStream) Recv() (cplugin.DestinationRunRequest, error) {
	in, err := s.impl.Recv()
	if err != nil {
		return cplugin.DestinationRunRequest{}, err
	}
	out, err := fromproto.DestinationRunRequest(in)
	if err != nil {
		return cplugin.DestinationRunRequest{}, err
	}
	return out, nil
}
