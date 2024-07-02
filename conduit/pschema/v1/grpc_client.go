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

	schemav1 "github.com/conduitio/conduit-commons/proto/schema/v1"
	cschema "github.com/conduitio/conduit-commons/schema"
	"github.com/conduitio/conduit-connector-protocol/conduit/pschema"
	"github.com/conduitio/conduit-connector-protocol/internal"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ pschema.Service = (*Client)(nil)

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

func (c *Client) Create(ctx context.Context, request pschema.CreateRequest) (pschema.CreateResponse, error) {
	// request is a pschema.CreateRequest and I need to change to proto so I can create a schema request with that proto
	resp, err := c.grpcClient.Create(ctx, &conduitv1.CreateSchemaRequest{
		Subject: request.Subject,
		Type:    schemav1.Schema_Type(request.Type),
		Bytes:   request.Bytes,
	})
	if err != nil {
		return pschema.CreateResponse{}, internal.UnwrapGRPCError(err)
	}

	// var schema cschema.Schema
	// schema.FromProto(resp.Schema) // resp.Schema is a CreateSchemaResponse, not a proto Schema

	return pschema.CreateResponse{
		Schema: cschema.Schema{
			Subject: resp.Schema.Subject,
			Version: int(resp.Schema.Version),
			Type:    cschema.Type(resp.Schema.Type),
			Bytes:   resp.Schema.Bytes,
		},
	}, nil
}

func (c *Client) Get(ctx context.Context, request pschema.GetRequest) (pschema.GetResponse, error) {
	resp, err := c.grpcClient.Get(ctx, &conduitv1.GetSchemaRequest{
		Subject: request.Subject,
		Version: int32(request.Version),
	})

	if err != nil {
		return pschema.GetResponse{}, fmt.Errorf("failed creating schema: %w", err)
	}

	return pschema.GetResponse{
		Schema: cschema.Schema{
			Subject: resp.Schema.Subject,
			Version: int(resp.Schema.Version),
			Type:    cschema.Type(resp.Schema.Type),
			Bytes:   resp.Schema.Bytes,
		},
	}, nil
}
