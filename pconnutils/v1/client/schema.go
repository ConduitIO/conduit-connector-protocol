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

	"github.com/conduitio/conduit-connector-protocol/pconnutils"
	"github.com/conduitio/conduit-connector-protocol/pconnutils/internal"
	"github.com/conduitio/conduit-connector-protocol/pconnutils/v1/fromproto"
	"github.com/conduitio/conduit-connector-protocol/pconnutils/v1/toproto"
	connutilsv1 "github.com/conduitio/conduit-connector-protocol/proto/connutils/v1"
	"google.golang.org/grpc"
)

type SchemaServiceClient struct {
	grpcClient connutilsv1.SchemaServiceClient
}

var _ pconnutils.SchemaService = (*SchemaServiceClient)(nil)

func NewSchemaServiceClient(cc *grpc.ClientConn) *SchemaServiceClient {
	return &SchemaServiceClient{
		grpcClient: connutilsv1.NewSchemaServiceClient(cc),
	}
}

func (c *SchemaServiceClient) CreateSchema(ctx context.Context, request pconnutils.CreateSchemaRequest) (pconnutils.CreateSchemaResponse, error) {
	ctx = internal.RepackConnectorTokenOutgoingContext(ctx)
	protoReq := toproto.CreateSchemaRequest(request)
	protoResp, err := c.grpcClient.CreateSchema(ctx, protoReq)
	if err != nil {
		return pconnutils.CreateSchemaResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.CreateSchemaResponse(protoResp)
}

func (c *SchemaServiceClient) GetSchema(ctx context.Context, request pconnutils.GetSchemaRequest) (pconnutils.GetSchemaResponse, error) {
	ctx = internal.RepackConnectorTokenOutgoingContext(ctx)
	protoReq := toproto.GetSchemaRequest(request)
	protoResp, err := c.grpcClient.GetSchema(ctx, protoReq)
	if err != nil {
		return pconnutils.GetSchemaResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.GetSchemaResponse(protoResp)
}
