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

package pconduit

import (
	"context"
	"testing"

	"github.com/conduitio/conduit-commons/schema"
	"github.com/google/go-cmp/cmp"
	"github.com/matryer/is"
)

func TestInMemoryService(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	underTest := NewInMemoryService()

	want1 := schema.Schema{
		Subject: "test-subject",
		Version: 1,
		Type:    schema.TypeAvro,
		Bytes:   []byte("irrelevant 1"),
	}

	// CreateSchema first version
	got1, err := underTest.CreateSchema(ctx, CreateSchemaRequest{
		Subject: want1.Subject,
		Type:    want1.Type,
		Bytes:   want1.Bytes,
	})
	is.NoErr(err)
	is.Equal("", cmp.Diff(want1, got1.Schema))

	// CreateSchema second version
	want2 := schema.Schema{
		Subject: want1.Subject,
		Version: 2,
		Type:    want1.Type,
		Bytes:   []byte("irrelevant 2"),
	}
	got2, err := underTest.CreateSchema(ctx, CreateSchemaRequest{
		Subject: want2.Subject,
		Type:    want2.Type,
		Bytes:   want2.Bytes,
	})
	is.NoErr(err)
	is.Equal("", cmp.Diff(want2, got2.Schema))

	// GetSchema first version
	getResp1, err := underTest.GetSchema(ctx, GetSchemaRequest{Subject: want1.Subject, Version: 1})
	is.NoErr(err)
	is.Equal("", cmp.Diff(want1, getResp1.Schema))

	// GetSchema second version
	getResp2, err := underTest.GetSchema(ctx, GetSchemaRequest{Subject: want2.Subject, Version: 2})
	is.NoErr(err)
	is.Equal("", cmp.Diff(want2, getResp2.Schema))
}
