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

	"github.com/conduitio/conduit-connector-protocol/pconduit"
	"github.com/conduitio/conduit-connector-protocol/pconduit/internal"
	"github.com/conduitio/conduit-connector-protocol/pconduit/v1/fromproto"
	"github.com/conduitio/conduit-connector-protocol/pconduit/v1/toproto"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
)

func NewSchemaServiceServer(impl pconduit.SchemaService) conduitv1.SchemaServiceServer {
	return &schemaServiceServer{impl: impl}
}

type schemaServiceServer struct {
	conduitv1.UnimplementedSchemaServiceServer
	impl pconduit.SchemaService
}

func (s *schemaServiceServer) CreateSchema(ctx context.Context, protoReq *conduitv1.CreateSchemaRequest) (*conduitv1.CreateSchemaResponse, error) {
	ctx = internal.RepackConnectorTokenIncomingContext(ctx)
	goReq := fromproto.CreateSchemaRequest(protoReq)
	goResp, err := s.impl.CreateSchema(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.CreateSchemaResponse(goResp)
}

func (s *schemaServiceServer) GetSchema(ctx context.Context, protoReq *conduitv1.GetSchemaRequest) (*conduitv1.GetSchemaResponse, error) {
	ctx = internal.RepackConnectorTokenIncomingContext(ctx)
	goReq := fromproto.GetSchemaRequest(protoReq)
	goResp, err := s.impl.GetSchema(ctx, goReq)
	if err != nil {
		return nil, err
	}
	return toproto.GetSchemaResponse(goResp)
}
