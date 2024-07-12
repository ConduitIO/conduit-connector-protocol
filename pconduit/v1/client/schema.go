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

	"github.com/conduitio/conduit-connector-protocol/pconduit"
	"github.com/conduitio/conduit-connector-protocol/pconduit/internal"
	"github.com/conduitio/conduit-connector-protocol/pconduit/v1/fromproto"
	"github.com/conduitio/conduit-connector-protocol/pconduit/v1/toproto"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
	"google.golang.org/grpc"
)

type SchemaServiceClient struct {
	grpcClient conduitv1.SchemaServiceClient
}

var _ pconduit.SchemaService = (*SchemaServiceClient)(nil)

func NewSchemaServiceClient(cc *grpc.ClientConn) *SchemaServiceClient {
	return &SchemaServiceClient{
		grpcClient: conduitv1.NewSchemaServiceClient(cc),
	}
}

func (c *SchemaServiceClient) CreateSchema(ctx context.Context, request pconduit.CreateSchemaRequest) (pconduit.CreateSchemaResponse, error) {
	createCtx := internal.RepackConnectorTokenOutgoingContext(ctx)
	protoReq := toproto.CreateSchemaRequest(request)
	protoResp, err := c.grpcClient.CreateSchema(createCtx, protoReq)
	if err != nil {
		return pconduit.CreateSchemaResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.CreateSchemaResponse(protoResp)
}

func (c *SchemaServiceClient) GetSchema(ctx context.Context, request pconduit.GetSchemaRequest) (pconduit.GetSchemaResponse, error) {
	getCtx := internal.RepackConnectorTokenOutgoingContext(ctx)
	protoReq := toproto.GetSchemaRequest(request)
	protoResp, err := c.grpcClient.GetSchema(getCtx, protoReq)
	if err != nil {
		return pconduit.GetSchemaResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.GetSchemaResponse(protoResp)
}
