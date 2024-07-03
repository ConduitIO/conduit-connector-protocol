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

//go:generate mockgen -typed -destination=mock/pschema.go -package=mock -mock_names=SchemaService=SchemaService . SchemaService

package pconduit

import (
	"context"

	"github.com/conduitio/conduit-commons/schema"
)

type SchemaService interface {
	CreateSchema(context.Context, CreateSchemaRequest) (CreateSchemaResponse, error)
	GetSchema(context.Context, GetSchemaRequest) (GetSchemaResponse, error)
}

type CreateSchemaRequest struct {
	Subject string
	Type    schema.Type
	Bytes   []byte
}

type CreateSchemaResponse struct {
	Schema schema.Schema
}

type GetSchemaRequest struct {
	Subject string
	Version int
}
type GetSchemaResponse struct {
	Schema schema.Schema
}
