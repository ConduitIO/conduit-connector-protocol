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

package schema

import (
	"context"

	"github.com/conduitio/conduit-commons/schema"
)

type Service interface {
	Create(context.Context, CreateRequest) (CreateResponse, error)
	Get(context.Context, GetRequest) (GetResponse, error)
}

type Type int32

const (
	TypeAvro Type = iota + 1
)

type CreateRequest struct {
	Name  string
	Type  Type
	Bytes []byte
}
type CreateResponse struct {
	schema.Instance
}

type GetRequest struct {
	ID string
}
type GetResponse struct {
	schema.Instance
}
