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

func NewSpecifierPluginServer(impl cplugin.SpecifierPlugin) connectorv1.SpecifierPluginServer {
	return &specifierPluginServer{impl: impl}
}

type specifierPluginServer struct {
	connectorv1.UnimplementedSpecifierPluginServer
	impl cplugin.SpecifierPlugin
}

func (s specifierPluginServer) Specify(ctx context.Context, protoReq *connectorv1.Specifier_Specify_Request) (*connectorv1.Specifier_Specify_Response, error) {
	goReq := fromproto.SpecifierSpecifyRequest(protoReq)
	goResp, err := s.impl.Specify(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.SpecifierSpecifyResponse(goResp), nil
}
