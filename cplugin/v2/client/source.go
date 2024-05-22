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

type SourcePluginClient struct {
	grpcClient connectorv2.SourcePluginClient
}

var _ cplugin.SourcePlugin = (*SourcePluginClient)(nil)

func NewSourcePluginClient(cc *grpc.ClientConn) cplugin.SourcePlugin {
	return &SourcePluginClient{grpcClient: connectorv2.NewSourcePluginClient(cc)}
}

func (s *SourcePluginClient) Configure(ctx context.Context, goReq cplugin.SourceConfigureRequest) (cplugin.SourceConfigureResponse, error) {
	protoReq := toproto.SourceConfigureRequest(goReq)
	protoResp, err := s.grpcClient.Configure(ctx, protoReq)
	if err != nil {
		return cplugin.SourceConfigureResponse{}, unwrapGRPCError(err)
	}
	return fromproto.SourceConfigureResponse(protoResp), nil
}

func (s *SourcePluginClient) Start(ctx context.Context, goReq cplugin.SourceStartRequest) (cplugin.SourceStartResponse, error) {
	protoReq := toproto.SourceStartRequest(goReq)
	protoResp, err := s.grpcClient.Start(ctx, protoReq)
	if err != nil {
		return cplugin.SourceStartResponse{}, unwrapGRPCError(err)
	}
	return fromproto.SourceStartResponse(protoResp), nil
}

// Run initializes a stream for the source plugin to send and receive messages.
// It only accepts a SourceRunStream as the stream type. If the stream has already
// been initialized, it will return an error.
// If the function returns no error, the stream is ready to be used.
func (s *SourcePluginClient) Run(ctx context.Context, stream cplugin.SourceRunStream) error {
	clientStream, ok := stream.(*SourceRunStream)
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

func (s *SourcePluginClient) Stop(ctx context.Context, goReq cplugin.SourceStopRequest) (cplugin.SourceStopResponse, error) {
	protoReq := toproto.SourceStopRequest(goReq)
	protoResp, err := s.grpcClient.Stop(ctx, protoReq)
	if err != nil {
		return cplugin.SourceStopResponse{}, unwrapGRPCError(err)
	}
	return fromproto.SourceStopResponse(protoResp), nil
}

func (s *SourcePluginClient) Teardown(ctx context.Context, goReq cplugin.SourceTeardownRequest) (cplugin.SourceTeardownResponse, error) {
	protoReq := toproto.SourceTeardownRequest(goReq)
	protoResp, err := s.grpcClient.Teardown(ctx, protoReq)
	if err != nil {
		return cplugin.SourceTeardownResponse{}, unwrapGRPCError(err)
	}
	return fromproto.SourceTeardownResponse(protoResp), nil
}

func (s *SourcePluginClient) LifecycleOnCreated(ctx context.Context, goReq cplugin.SourceLifecycleOnCreatedRequest) (cplugin.SourceLifecycleOnCreatedResponse, error) {
	protoReq := toproto.SourceLifecycleOnCreatedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnCreated(ctx, protoReq)
	if err != nil {
		return cplugin.SourceLifecycleOnCreatedResponse{}, unwrapGRPCError(err)
	}
	return fromproto.SourceLifecycleOnCreatedResponse(protoResp), nil
}

func (s *SourcePluginClient) LifecycleOnUpdated(ctx context.Context, goReq cplugin.SourceLifecycleOnUpdatedRequest) (cplugin.SourceLifecycleOnUpdatedResponse, error) {
	protoReq := toproto.SourceLifecycleOnUpdatedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnUpdated(ctx, protoReq)
	if err != nil {
		return cplugin.SourceLifecycleOnUpdatedResponse{}, unwrapGRPCError(err)
	}
	return fromproto.SourceLifecycleOnUpdatedResponse(protoResp), nil
}

func (s *SourcePluginClient) LifecycleOnDeleted(ctx context.Context, goReq cplugin.SourceLifecycleOnDeletedRequest) (cplugin.SourceLifecycleOnDeletedResponse, error) {
	protoReq := toproto.SourceLifecycleOnDeletedRequest(goReq)
	protoResp, err := s.grpcClient.LifecycleOnDeleted(ctx, protoReq)
	if err != nil {
		return cplugin.SourceLifecycleOnDeletedResponse{}, unwrapGRPCError(err)
	}
	return fromproto.SourceLifecycleOnDeletedResponse(protoResp), nil
}

func (s *SourcePluginClient) NewStream() cplugin.SourceRunStream {
	return &SourceRunStream{}
}

// SourceRunStream is the client-side implementation of the
// cplugin.SourceRunStream interface.
type SourceRunStream struct {
	client connectorv2.SourcePlugin_RunClient
}

func (s *SourceRunStream) Client() cplugin.SourceRunStreamClient {
	if s.client == nil {
		panic("invalid use of client.SourceRunStream - stream has not been initialized using SourcePluginClient.Run")
	}
	return s
}

func (s *SourceRunStream) Server() cplugin.SourceRunStreamServer {
	panic("invalid use of client.SourceRunStream - it is a client-side type only")
}

func (s *SourceRunStream) Send(goReq cplugin.SourceRunRequest) error {
	protoReq := toproto.SourceRunRequest(goReq)
	err := s.client.Send(protoReq)
	if err != nil {
		return unwrapGRPCError(err)
	}
	return nil
}

func (s *SourceRunStream) Recv() (cplugin.SourceRunResponse, error) {
	protoResp, err := s.client.Recv()
	if err != nil {
		return cplugin.SourceRunResponse{}, unwrapGRPCError(err)
	}
	goResp, err := fromproto.SourceRunResponse(protoResp)
	if err != nil {
		return cplugin.SourceRunResponse{}, err
	}
	return goResp, nil
}
