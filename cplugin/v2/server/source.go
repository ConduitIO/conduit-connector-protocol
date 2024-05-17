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

	"github.com/conduitio/conduit-connector-protocol/cplugin"
	"github.com/conduitio/conduit-connector-protocol/cplugin/v2/fromproto"
	"github.com/conduitio/conduit-connector-protocol/cplugin/v2/toproto"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
)

func NewSourcePluginServer(impl cplugin.SourcePlugin) connectorv2.SourcePluginServer {
	return &sourcePluginServer{impl: impl}
}

type sourcePluginServer struct {
	connectorv2.UnimplementedSourcePluginServer
	impl cplugin.SourcePlugin
}

func (s *sourcePluginServer) Configure(ctx context.Context, protoReq *connectorv2.Source_Configure_Request) (*connectorv2.Source_Configure_Response, error) {
	goReq := fromproto.SourceConfigureRequest(protoReq)
	goResp, err := s.impl.Configure(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.SourceConfigureResponse(goResp)
	return protoResp, nil
}
func (s *sourcePluginServer) Start(ctx context.Context, protoReq *connectorv2.Source_Start_Request) (*connectorv2.Source_Start_Response, error) {
	goReq := fromproto.SourceStartRequest(protoReq)
	goResp, err := s.impl.Start(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.SourceStartResponse(goResp)
	return protoResp, nil
}
func (s *sourcePluginServer) Run(stream connectorv2.SourcePlugin_RunServer) error {
	err := s.impl.Run(stream.Context(), &sourceRunStream{impl: stream})
	if err != nil {
		return err
	}
	return nil
}
func (s *sourcePluginServer) Stop(ctx context.Context, protoReq *connectorv2.Source_Stop_Request) (*connectorv2.Source_Stop_Response, error) {
	goReq := fromproto.SourceStopRequest(protoReq)
	goResp, err := s.impl.Stop(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.SourceStopResponse(goResp)
	return protoResp, nil
}
func (s *sourcePluginServer) Teardown(ctx context.Context, protoReq *connectorv2.Source_Teardown_Request) (*connectorv2.Source_Teardown_Response, error) {
	goReq := fromproto.SourceTeardownRequest(protoReq)
	goResp, err := s.impl.Teardown(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.SourceTeardownResponse(goResp)
	return protoResp, nil
}
func (s *sourcePluginServer) LifecycleOnCreated(ctx context.Context, protoReq *connectorv2.Source_Lifecycle_OnCreated_Request) (*connectorv2.Source_Lifecycle_OnCreated_Response, error) {
	goReq := fromproto.SourceLifecycleOnCreatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnCreated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.SourceLifecycleOnCreatedResponse(goResp)
	return protoResp, nil
}
func (s *sourcePluginServer) LifecycleOnUpdated(ctx context.Context, protoReq *connectorv2.Source_Lifecycle_OnUpdated_Request) (*connectorv2.Source_Lifecycle_OnUpdated_Response, error) {
	goReq := fromproto.SourceLifecycleOnUpdatedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnUpdated(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.SourceLifecycleOnUpdatedResponse(goResp)
	return protoResp, nil
}
func (s *sourcePluginServer) LifecycleOnDeleted(ctx context.Context, protoReq *connectorv2.Source_Lifecycle_OnDeleted_Request) (*connectorv2.Source_Lifecycle_OnDeleted_Response, error) {
	goReq := fromproto.SourceLifecycleOnDeletedRequest(protoReq)
	goResp, err := s.impl.LifecycleOnDeleted(ctx, goReq)
	if err != nil {
		return nil, err
	}
	protoResp := toproto.SourceLifecycleOnDeletedResponse(goResp)
	return protoResp, nil
}

type sourceRunStream struct {
	impl connectorv2.SourcePlugin_RunServer
}

func (s *sourceRunStream) Send(in cplugin.SourceRunResponse) error {
	out, err := toproto.SourceRunResponse(in)
	if err != nil {
		return err
	}
	return s.impl.Send(out)
}

func (s *sourceRunStream) Recv() (cplugin.SourceRunRequest, error) {
	in, err := s.impl.Recv()
	if err != nil {
		return cplugin.SourceRunRequest{}, err
	}
	out := fromproto.SourceRunRequest(in)
	if err != nil {
		return cplugin.SourceRunRequest{}, err
	}
	return out, nil
}
