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

	"github.com/conduitio/conduit-connector-protocol/conduit/schema"
	"github.com/conduitio/conduit-connector-protocol/conduit/schema/v1/fromproto"
	"github.com/conduitio/conduit-connector-protocol/conduit/schema/v1/toproto"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
	"google.golang.org/grpc"
)

var _ schema.Service = (*Client)(nil)

type Client struct {
	grpcClient conduitv1.SchemaServiceClient
}

func NewClient(conn *grpc.ClientConn) (*Client, error) {
	return &Client{
		grpcClient: conduitv1.NewSchemaServiceClient(conn),
	}, nil
}

func (c *Client) Create(ctx context.Context, request schema.CreateRequest) (schema.CreateResponse, error) {
	resp, err := c.grpcClient.Create(ctx, toproto.CreateSchemaRequest(request))
	if err != nil {
		return schema.CreateResponse{}, fmt.Errorf("failed creating schema: %w", err)
	}

	return fromproto.CreateResponse(resp), nil
}

func (c *Client) Get(ctx context.Context, request schema.GetRequest) (schema.GetResponse, error) {
	resp, err := c.grpcClient.Get(ctx, toproto.GetSchemaRequest(request))
	if err != nil {
		return schema.GetResponse{}, fmt.Errorf("failed creating schema: %w", err)
	}

	return fromproto.GetResponse(resp), nil
}
