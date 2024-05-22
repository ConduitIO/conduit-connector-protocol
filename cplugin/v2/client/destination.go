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

package client

import (
	"context"
	"fmt"

	"github.com/conduitio/conduit-connector-protocol/cplugin"
	"github.com/conduitio/conduit-connector-protocol/cplugin/v2/fromproto"
	"github.com/conduitio/conduit-connector-protocol/cplugin/v2/toproto"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
	"google.golang.org/grpc"
)

type DestinationPluginClient struct {
	grpcClient connectorv2.DestinationPluginClient
}

var _ cplugin.DestinationPlugin = (*DestinationPluginClient)(nil)

func NewDestinationPluginClient(cc *grpc.ClientConn) cplugin.DestinationPlugin {
	return &DestinationPluginClient{grpcClient: connectorv2.NewDestinationPluginClient(cc)}
}

func (s *DestinationPluginClient) Configure(ctx context.Context, goReq cplugin.DestinationConfigureRequest) (cplugin.DestinationConfigureResponse, error) {
	protoReq := toproto.DestinationConfigureRequest(goReq)
	protoResp, err := s.grpcClient.Configure(ctx, protoReq)
	if err != nil {
		return cplugin.DestinationConfigureResponse{}, unwrapGRPCError(err)
	}
	return fromproto.DestinationConfigureResponse(protoResp), nil
}

func (s *DestinationPluginClient) Start(ctx context.Context, goReq cplugin.DestinationStartRequest) (cplugin.DestinationStartResponse, error) {
	protoReq := toproto.DestinationStartRequest(goReq)
	protoResp, err := s.grpcClient.Start(ctx, protoReq)
	if err != nil {
		return cplugin.DestinationStartResponse{}, unwrapGRPCError(err)
	}
	return fromproto.DestinationStartResponse(protoResp), nil
}

func (s *DestinationPluginClient) Run(ctx context.Context, stream cplugin.DestinationRunStream) error {
	clientStream, ok := stream.(*DestinationRunStream)
	if !ok {
		return fmt.Errorf("invalid stream type, expected %T, got %T", s.NewStream(), stream)
	}
	if clientStream.client != nil {
		return fmt.Errorf("stream has already been initialized")
	}

	grpcStream, err := s.grpcClient.Run(ctx)
	if err != nil {
		return unwrapGRPCError(err)
	}

	clientStream.client = grpcStream
	return nil
}

func (s *DestinationPluginClient) Stop(ctx context.Context, goReq cplugin.DestinationStopRequest) (cplugin.DestinationStopResponse, error) {
	protoReq := toproto.DestinationStopRequest(goReq)
	protoResp, err := s.grpcClient.Stop(ctx, protoReq)
	if err != nil {
		return cplugin.DestinationStopResponse{}, unwrapGRPCError(err)
	}
	return fromproto.DestinationStopResponse(protoResp), nil
}

func (s *DestinationPluginClient) Teardown(ctx context.Context, goReq cplugin.DestinationTeardownRequest) (cplugin.DestinationTeardownResponse, error) {
	protoReq := toproto.DestinationTeardownRequest(goReq)
	protoResp, err := s.grpcClient.Teardown(ctx, protoReq)
	if err != nil {
		return cplugin.DestinationTeardownResponse{}, unwrapGRPCError(err)
	}
	return fromproto.DestinationTeardownResponse(protoResp), nil
}

func (s *DestinationPluginClient) LifecycleOnCreated(ctx context.Context, goReq cplugin.DestinationLifecycleOnCreatedRequest) (cplugin.DestinationLifecycleOnCreatedResponse, error) {
	protoReq := toproto.DestinationLifecycleOnCreatedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnCreated(ctx, protoReq)
	if err != nil {
		return cplugin.DestinationLifecycleOnCreatedResponse{}, unwrapGRPCError(err)
	}
	return fromproto.DestinationLifecycleOnCreatedResponse(protoResp), nil
}
func (s *DestinationPluginClient) LifecycleOnUpdated(ctx context.Context, goReq cplugin.DestinationLifecycleOnUpdatedRequest) (cplugin.DestinationLifecycleOnUpdatedResponse, error) {
	protoReq := toproto.DestinationLifecycleOnUpdatedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnUpdated(ctx, protoReq)
	if err != nil {
		return cplugin.DestinationLifecycleOnUpdatedResponse{}, unwrapGRPCError(err)
	}
	return fromproto.DestinationLifecycleOnUpdatedResponse(protoResp), nil
}
func (s *DestinationPluginClient) LifecycleOnDeleted(ctx context.Context, goReq cplugin.DestinationLifecycleOnDeletedRequest) (cplugin.DestinationLifecycleOnDeletedResponse, error) {
	protoReq := toproto.DestinationLifecycleOnDeletedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnDeleted(ctx, protoReq)
	if err != nil {
		return cplugin.DestinationLifecycleOnDeletedResponse{}, unwrapGRPCError(err)
	}
	return fromproto.DestinationLifecycleOnDeletedResponse(protoResp), nil
}

func (s *DestinationPluginClient) NewStream() cplugin.DestinationRunStream {
	return &DestinationRunStream{}
}

// DestinationRunStream is the client-side implementation of the
// cplugin.DestinationRunStream interface.
type DestinationRunStream struct {
	client connectorv2.DestinationPlugin_RunClient
}

func (s *DestinationRunStream) Client() cplugin.DestinationRunStreamClient {
	if s.client == nil {
		panic("invalid use of client.DestinationRunStream - stream has not been initialized using DestinationPluginClient.Run")
	}
	return s
}

func (s *DestinationRunStream) Server() cplugin.DestinationRunStreamServer {
	panic("invalid use of client.DestinationRunStream - it is a client-side type only")
}

func (s *DestinationRunStream) Send(goReq cplugin.DestinationRunRequest) error {
	protoReq, err := toproto.DestinationRunRequest(goReq)
	if err != nil {
		return err
	}

	err = s.client.Send(protoReq)
	if err != nil {
		return unwrapGRPCError(err)
	}
	return nil
}

func (s *DestinationRunStream) Recv() (cplugin.DestinationRunResponse, error) {
	protoResp, err := s.client.Recv()
	if err != nil {
		return cplugin.DestinationRunResponse{}, unwrapGRPCError(err)
	}
	goResp := fromproto.DestinationRunResponse(protoResp)

	return goResp, nil
}
