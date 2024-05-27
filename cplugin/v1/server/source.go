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

func NewSourcePluginServer(impl cplugin.SourcePlugin) connectorv1.SourcePluginServer {
	return &sourcePluginServer{impl: impl}
}

type sourcePluginServer struct {
	connectorv1.UnimplementedSourcePluginServer
	impl cplugin.SourcePlugin
}

func (s *sourcePluginServer) Configure(ctx context.Context, protoReq *connectorv1.Source_Configure_Request) (*connectorv1.Source_Configure_Response, error) {
	goReq := fromproto.SourceConfigureRequest(protoReq)
	goResp, err := s.impl.Configure(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceConfigureResponse(goResp), nil
}
func (s *sourcePluginServer) Start(ctx context.Context, protoReq *connectorv1.Source_Start_Request) (*connectorv1.Source_Start_Response, error) {
	goReq := fromproto.SourceStartRequest(protoReq)
	goResp, err := s.impl.Open(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceStartResponse(goResp), nil
}
func (s *sourcePluginServer) Run(stream connectorv1.SourcePlugin_RunServer) error {
	err := s.impl.Run(stream.Context(), &SourceRunStream{stream: stream})
	if err != nil {
		return err
	}
	return nil
}
func (s *sourcePluginServer) Stop(ctx context.Context, protoReq *connectorv1.Source_Stop_Request) (*connectorv1.Source_Stop_Response, error) {
	goReq := fromproto.SourceStopRequest(protoReq)
	goResp, err := s.impl.Stop(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceStopResponse(goResp), nil
}
func (s *sourcePluginServer) Teardown(ctx context.Context, protoReq *connectorv1.Source_Teardown_Request) (*connectorv1.Source_Teardown_Response, error) {
	goReq := fromproto.SourceTeardownRequest(protoReq)
	goResp, err := s.impl.Teardown(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceTeardownResponse(goResp), nil
}
func (s *sourcePluginServer) LifecycleOnCreated(ctx context.Context, protoReq *connectorv1.Source_Lifecycle_OnCreated_Request) (*connectorv1.Source_Lifecycle_OnCreated_Response, error) {
	goReq := fromproto.SourceLifecycleOnCreatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnCreated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceLifecycleOnCreatedResponse(goResp), nil
}
func (s *sourcePluginServer) LifecycleOnUpdated(ctx context.Context, protoReq *connectorv1.Source_Lifecycle_OnUpdated_Request) (*connectorv1.Source_Lifecycle_OnUpdated_Response, error) {
	goReq := fromproto.SourceLifecycleOnUpdatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnUpdated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceLifecycleOnUpdatedResponse(goResp), nil
}
func (s *sourcePluginServer) LifecycleOnDeleted(ctx context.Context, protoReq *connectorv1.Source_Lifecycle_OnDeleted_Request) (*connectorv1.Source_Lifecycle_OnDeleted_Response, error) {
	goReq := fromproto.SourceLifecycleOnDeletedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnDeleted(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SourceLifecycleOnDeletedResponse(goResp), nil
}

// SourceRunStream is the server-side implementation of the
// cplugin.SourceRunStream interface.
type SourceRunStream struct {
	stream connectorv1.SourcePlugin_RunServer
}

func (s *SourceRunStream) Client() cplugin.SourceRunStreamClient {
	panic("invalid use of server.SourceRunStream - it is a server-side type only")
}
func (s *SourceRunStream) Server() cplugin.SourceRunStreamServer {
	if s.stream == nil {
		panic("invalid use of server.SourceRunStream - stream has not been initialized using SourcePluginServer.Run")
	}
	return s
}

func (s *SourceRunStream) Send(in cplugin.SourceRunResponse) error {
	out, err := toproto.SourceRunResponse(in)
	if err != nil {
		return err
	}
	for _, out := range out {
		err := s.stream.Send(out)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SourceRunStream) Recv() (cplugin.SourceRunRequest, error) {
	in, err := s.stream.Recv()
	if err != nil {
		return cplugin.SourceRunRequest{}, err
	}
	out := fromproto.SourceRunRequest(in)
	return out, nil
}
