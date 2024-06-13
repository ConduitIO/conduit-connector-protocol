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
	v1 "github.com/conduitio/conduit-connector-protocol/conduit/schema/v1"
	"testing"

	"github.com/conduitio/conduit-connector-protocol/conduit/schema/mock"
	"github.com/matryer/is"
	"go.uber.org/mock/gomock"
)

func TestClient_NewFromContext(t *testing.T) {
	is := is.New(t)
	service := mock.NewService(gomock.NewController(t))
	ctx := WithSchemaService(context.Background(), service)

	s, err := New(ctx)
	is.NoErr(err)
	is.Equal(s, service)
}

func TestClient_NewGRPC(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	s, err := New(ctx)
	is.NoErr(err)
	_, ok := s.(*v1.Client)
	is.True(ok)
}
