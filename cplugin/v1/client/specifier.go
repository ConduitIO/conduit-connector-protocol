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

	"github.com/conduitio/conduit-connector-protocol/cplugin"
	"github.com/conduitio/conduit-connector-protocol/cplugin/v1/fromproto"
	"github.com/conduitio/conduit-connector-protocol/cplugin/v1/toproto"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
	"google.golang.org/grpc"
)

type SpecifierPluginClient struct {
	grpcClient connectorv1.SpecifierPluginClient
}

var _ cplugin.SpecifierPlugin = (*SpecifierPluginClient)(nil)

func NewSpecifierPluginClient(cc *grpc.ClientConn) *SpecifierPluginClient {
	return &SpecifierPluginClient{grpcClient: connectorv1.NewSpecifierPluginClient(cc)}
}

func (s *SpecifierPluginClient) Specify(ctx context.Context, goReq cplugin.SpecifierSpecifyRequest) (cplugin.SpecifierSpecifyResponse, error) {
	protoReq := toproto.SpecifierSpecifyRequest(goReq)
	protoResp, err := s.grpcClient.Specify(ctx, protoReq)
	if err != nil {
		return cplugin.SpecifierSpecifyResponse{}, unwrapGRPCError(err)
	}
	goResp, err := fromproto.SpecifierSpecifyResponse(protoResp)
	if err != nil {
		return cplugin.SpecifierSpecifyResponse{}, err
	}
	return goResp, nil
}
