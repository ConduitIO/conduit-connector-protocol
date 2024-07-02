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

	"github.com/conduitio/conduit-connector-protocol/internal"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v2/fromproto"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v2/toproto"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
	"google.golang.org/grpc"
)

type DestinationPluginClient struct {
	grpcClient connectorv2.DestinationPluginClient
}

var _ pconnector.DestinationPlugin = (*DestinationPluginClient)(nil)

func NewDestinationPluginClient(cc *grpc.ClientConn) *DestinationPluginClient {
	return &DestinationPluginClient{grpcClient: connectorv2.NewDestinationPluginClient(cc)}
}

func (s *DestinationPluginClient) Configure(ctx context.Context, goReq pconnector.DestinationConfigureRequest) (pconnector.DestinationConfigureResponse, error) {
	protoReq := toproto.DestinationConfigureRequest(goReq)
	protoResp, err := s.grpcClient.Configure(ctx, protoReq)
	if err != nil {
		return pconnector.DestinationConfigureResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.DestinationConfigureResponse(protoResp), nil
}

func (s *DestinationPluginClient) Open(ctx context.Context, goReq pconnector.DestinationOpenRequest) (pconnector.DestinationOpenResponse, error) {
	protoReq := toproto.DestinationOpenRequest(goReq)
	protoResp, err := s.grpcClient.Open(ctx, protoReq)
	if err != nil {
		return pconnector.DestinationOpenResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.DestinationOpenResponse(protoResp), nil
}

func (s *DestinationPluginClient) Run(ctx context.Context, stream pconnector.DestinationRunStream) error {
	clientStream, ok := stream.(*DestinationRunStream)
	if !ok {
		return fmt.Errorf("invalid stream type, expected %T, got %T", s.NewStream(), stream)
	}
	if clientStream.client != nil {
		return fmt.Errorf("stream has already been initialized")
	}

	grpcStream, err := s.grpcClient.Run(ctx)
	if err != nil {
		return internal.UnwrapGRPCError(err)
	}

	clientStream.client = grpcStream
	return nil
}

func (s *DestinationPluginClient) Stop(ctx context.Context, goReq pconnector.DestinationStopRequest) (pconnector.DestinationStopResponse, error) {
	protoReq := toproto.DestinationStopRequest(goReq)
	protoResp, err := s.grpcClient.Stop(ctx, protoReq)
	if err != nil {
		return pconnector.DestinationStopResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.DestinationStopResponse(protoResp), nil
}

func (s *DestinationPluginClient) Teardown(ctx context.Context, goReq pconnector.DestinationTeardownRequest) (pconnector.DestinationTeardownResponse, error) {
	protoReq := toproto.DestinationTeardownRequest(goReq)
	protoResp, err := s.grpcClient.Teardown(ctx, protoReq)
	if err != nil {
		return pconnector.DestinationTeardownResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.DestinationTeardownResponse(protoResp), nil
}

func (s *DestinationPluginClient) LifecycleOnCreated(ctx context.Context, goReq pconnector.DestinationLifecycleOnCreatedRequest) (pconnector.DestinationLifecycleOnCreatedResponse, error) {
	protoReq := toproto.DestinationLifecycleOnCreatedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnCreated(ctx, protoReq)
	if err != nil {
		return pconnector.DestinationLifecycleOnCreatedResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.DestinationLifecycleOnCreatedResponse(protoResp), nil
}

func (s *DestinationPluginClient) LifecycleOnUpdated(ctx context.Context, goReq pconnector.DestinationLifecycleOnUpdatedRequest) (pconnector.DestinationLifecycleOnUpdatedResponse, error) {
	protoReq := toproto.DestinationLifecycleOnUpdatedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnUpdated(ctx, protoReq)
	if err != nil {
		return pconnector.DestinationLifecycleOnUpdatedResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.DestinationLifecycleOnUpdatedResponse(protoResp), nil
}

func (s *DestinationPluginClient) LifecycleOnDeleted(ctx context.Context, goReq pconnector.DestinationLifecycleOnDeletedRequest) (pconnector.DestinationLifecycleOnDeletedResponse, error) {
	protoReq := toproto.DestinationLifecycleOnDeletedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnDeleted(ctx, protoReq)
	if err != nil {
		return pconnector.DestinationLifecycleOnDeletedResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.DestinationLifecycleOnDeletedResponse(protoResp), nil
}

func (s *DestinationPluginClient) NewStream() pconnector.DestinationRunStream {
	return &DestinationRunStream{}
}

// DestinationRunStream is the client-side implementation of the
// pconnector.DestinationRunStream interface.
type DestinationRunStream struct {
	client connectorv2.DestinationPlugin_RunClient
}

func (s *DestinationRunStream) Client() pconnector.DestinationRunStreamClient {
	if s.client == nil {
		panic("invalid use of client.DestinationRunStream - stream has not been initialized using DestinationPluginClient.Run")
	}
	return s
}

func (s *DestinationRunStream) Server() pconnector.DestinationRunStreamServer {
	panic("invalid use of client.DestinationRunStream - it is a client-side type only")
}

func (s *DestinationRunStream) Send(goReq pconnector.DestinationRunRequest) error {
	protoReq, err := toproto.DestinationRunRequest(goReq)
	if err != nil {
		return err
	}

	err = s.client.Send(protoReq)
	if err != nil {
		return internal.UnwrapGRPCError(err)
	}
	return nil
}

func (s *DestinationRunStream) Recv() (pconnector.DestinationRunResponse, error) {
	protoResp, err := s.client.Recv()
	if err != nil {
		return pconnector.DestinationRunResponse{}, internal.UnwrapGRPCError(err)
	}
	goResp := fromproto.DestinationRunResponse(protoResp)

	return goResp, nil
}
