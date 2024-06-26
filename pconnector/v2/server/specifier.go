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

func NewSpecifierPluginServer(impl pconnector.SpecifierPlugin) connectorv2.SpecifierPluginServer {
	return &SpecifierPluginServer{impl: impl}
}

type SpecifierPluginServer struct {
	connectorv2.UnimplementedSpecifierPluginServer
	impl pconnector.SpecifierPlugin
}

func (s SpecifierPluginServer) Specify(ctx context.Context, protoReq *connectorv2.Specifier_Specify_Request) (*connectorv2.Specifier_Specify_Response, error) {
	goReq := fromproto.SpecifierSpecifyRequest(protoReq)
	goResp, err := s.impl.Specify(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SpecifierSpecifyResponse(goResp), nil
}
