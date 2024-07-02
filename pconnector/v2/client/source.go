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

type SourcePluginClient struct {
	grpcClient connectorv2.SourcePluginClient
}

var _ pconnector.SourcePlugin = (*SourcePluginClient)(nil)

func NewSourcePluginClient(cc *grpc.ClientConn) *SourcePluginClient {
	return &SourcePluginClient{grpcClient: connectorv2.NewSourcePluginClient(cc)}
}

func (s *SourcePluginClient) Configure(ctx context.Context, goReq pconnector.SourceConfigureRequest) (pconnector.SourceConfigureResponse, error) {
	protoReq := toproto.SourceConfigureRequest(goReq)
	protoResp, err := s.grpcClient.Configure(ctx, protoReq)
	if err != nil {
		return pconnector.SourceConfigureResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.SourceConfigureResponse(protoResp), nil
}

func (s *SourcePluginClient) Open(ctx context.Context, goReq pconnector.SourceOpenRequest) (pconnector.SourceOpenResponse, error) {
	protoReq := toproto.SourceOpenRequest(goReq)
	protoResp, err := s.grpcClient.Open(ctx, protoReq)
	if err != nil {
		return pconnector.SourceOpenResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.SourceOpenResponse(protoResp), nil
}

// Run initializes a stream for the source plugin to send and receive messages.
// It only accepts a SourceRunStream as the stream type. If the stream has already
// been initialized, it will return an error.
// If the function returns no error, the stream is ready to be used.
func (s *SourcePluginClient) Run(ctx context.Context, stream pconnector.SourceRunStream) error {
	clientStream, ok := stream.(*SourceRunStream)
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

func (s *SourcePluginClient) Stop(ctx context.Context, goReq pconnector.SourceStopRequest) (pconnector.SourceStopResponse, error) {
	protoReq := toproto.SourceStopRequest(goReq)
	protoResp, err := s.grpcClient.Stop(ctx, protoReq)
	if err != nil {
		return pconnector.SourceStopResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.SourceStopResponse(protoResp), nil
}

func (s *SourcePluginClient) Teardown(ctx context.Context, goReq pconnector.SourceTeardownRequest) (pconnector.SourceTeardownResponse, error) {
	protoReq := toproto.SourceTeardownRequest(goReq)
	protoResp, err := s.grpcClient.Teardown(ctx, protoReq)
	if err != nil {
		return pconnector.SourceTeardownResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.SourceTeardownResponse(protoResp), nil
}

func (s *SourcePluginClient) LifecycleOnCreated(ctx context.Context, goReq pconnector.SourceLifecycleOnCreatedRequest) (pconnector.SourceLifecycleOnCreatedResponse, error) {
	protoReq := toproto.SourceLifecycleOnCreatedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnCreated(ctx, protoReq)
	if err != nil {
		return pconnector.SourceLifecycleOnCreatedResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.SourceLifecycleOnCreatedResponse(protoResp), nil
}

func (s *SourcePluginClient) LifecycleOnUpdated(ctx context.Context, goReq pconnector.SourceLifecycleOnUpdatedRequest) (pconnector.SourceLifecycleOnUpdatedResponse, error) {
	protoReq := toproto.SourceLifecycleOnUpdatedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnUpdated(ctx, protoReq)
	if err != nil {
		return pconnector.SourceLifecycleOnUpdatedResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.SourceLifecycleOnUpdatedResponse(protoResp), nil
}

func (s *SourcePluginClient) LifecycleOnDeleted(ctx context.Context, goReq pconnector.SourceLifecycleOnDeletedRequest) (pconnector.SourceLifecycleOnDeletedResponse, error) {
	protoReq := toproto.SourceLifecycleOnDeletedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnDeleted(ctx, protoReq)
	if err != nil {
		return pconnector.SourceLifecycleOnDeletedResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.SourceLifecycleOnDeletedResponse(protoResp), nil
}

func (s *SourcePluginClient) NewStream() pconnector.SourceRunStream {
	return &SourceRunStream{}
}

// SourceRunStream is the client-side implementation of the
// pconnector.SourceRunStream interface.
type SourceRunStream struct {
	client connectorv2.SourcePlugin_RunClient
}

func (s *SourceRunStream) Client() pconnector.SourceRunStreamClient {
	if s.client == nil {
		panic("invalid use of client.SourceRunStream - stream has not been initialized using SourcePluginClient.Run")
	}
	return s
}

func (s *SourceRunStream) Server() pconnector.SourceRunStreamServer {
	panic("invalid use of client.SourceRunStream - it is a client-side type only")
}

func (s *SourceRunStream) Send(goReq pconnector.SourceRunRequest) error {
	protoReq := toproto.SourceRunRequest(goReq)
	err := s.client.Send(protoReq)
	if err != nil {
		return internal.UnwrapGRPCError(err)
	}
	return nil
}

func (s *SourceRunStream) Recv() (pconnector.SourceRunResponse, error) {
	protoResp, err := s.client.Recv()
	if err != nil {
		return pconnector.SourceRunResponse{}, internal.UnwrapGRPCError(err)
	}
	goResp, err := fromproto.SourceRunResponse(protoResp)
	if err != nil {
		return pconnector.SourceRunResponse{}, err
	}
	return goResp, nil
}
