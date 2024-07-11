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
	"google.golang.org/grpc/metadata"
)

type SchemaServiceClient struct {
	grpcClient conduitv1.SchemaServiceClient
	token      string
}

var _ pconduit.SchemaService = (*SchemaServiceClient)(nil)

func NewSchemaServiceClient(cc *grpc.ClientConn, token string) *SchemaServiceClient {
	return &SchemaServiceClient{
		grpcClient: conduitv1.NewSchemaServiceClient(cc),
		token:      token,
	}
}

func (c *SchemaServiceClient) CreateSchema(ctx context.Context, request pconduit.CreateSchemaRequest) (pconduit.CreateSchemaResponse, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, internal.MetadataConnectorTokenKey, c.token)
	protoReq := toproto.CreateSchemaRequest(request)
	protoResp, err := c.grpcClient.CreateSchema(ctx, protoReq)
	if err != nil {
		return pconduit.CreateSchemaResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.CreateSchemaResponse(protoResp)
}

func (c *SchemaServiceClient) GetSchema(ctx context.Context, request pconduit.GetSchemaRequest) (pconduit.GetSchemaResponse, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, internal.MetadataConnectorTokenKey, c.token)
	protoReq := toproto.GetSchemaRequest(request)
	protoResp, err := c.grpcClient.GetSchema(ctx, protoReq)
	if err != nil {
		return pconduit.GetSchemaResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.GetSchemaResponse(protoResp)
}
