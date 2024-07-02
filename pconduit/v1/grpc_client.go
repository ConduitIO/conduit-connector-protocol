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

package v1

import (
	"context"
	"fmt"

	"github.com/conduitio/conduit-connector-protocol/internal"
	"github.com/conduitio/conduit-connector-protocol/pconduit"
	"github.com/conduitio/conduit-connector-protocol/pconduit/v1/fromproto"
	"github.com/conduitio/conduit-connector-protocol/pconduit/v1/toproto"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ pconduit.SchemaService = (*Client)(nil)

type Client struct {
	grpcClient conduitv1.SchemaServiceClient
}

func NewClient(target string) (*Client, error) {
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed creating gRPC client: %w", err)
	}

	return &Client{grpcClient: conduitv1.NewSchemaServiceClient(conn)}, nil
}

func (c *Client) CreateSchema(ctx context.Context, request pconduit.CreateSchemaRequest) (pconduit.CreateSchemaResponse, error) {
	protoReq := toproto.CreateSchemaRequest(request)
	protoResp, err := c.grpcClient.Create(ctx, protoReq)
	if err != nil {
		return pconduit.CreateSchemaResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.CreateSchemaResponse(protoResp)
}

func (c *Client) GetSchema(ctx context.Context, request pconduit.GetSchemaRequest) (pconduit.GetSchemaResponse, error) {
	protoReq := toproto.GetSchemaRequest(request)
	protoResp, err := c.grpcClient.Get(ctx, protoReq)
	if err != nil {
		return pconduit.GetSchemaResponse{}, internal.UnwrapGRPCError(err)
	}
	return fromproto.GetSchemaResponse(protoResp)
}
