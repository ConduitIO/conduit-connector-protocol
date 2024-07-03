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

package fromproto

import (
	"testing"

	schemav1 "github.com/conduitio/conduit-commons/proto/schema/v1"
	"github.com/conduitio/conduit-commons/schema"
	"github.com/conduitio/conduit-connector-protocol/pconduit"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/matryer/is"
)

func TestCreateSchemaRequest(t *testing.T) {
	have := &conduitv1.CreateSchemaRequest{
		Subject: "foo",
		Type:    schemav1.Schema_TYPE_AVRO,
		Bytes:   []byte("bar"),
	}
	want := pconduit.CreateSchemaRequest{
		Subject: "foo",
		Type:    schema.TypeAvro,
		Bytes:   []byte("bar"),
	}

	is := is.New(t)
	got := CreateSchemaRequest(have)
	is.Equal("", cmp.Diff(want, got))
}

func TestGetSchemaRequest(t *testing.T) {
	have := &conduitv1.GetSchemaRequest{
		Subject: "foo",
		Version: 2,
	}
	want := pconduit.GetSchemaRequest{
		Subject: "foo",
		Version: 2,
	}

	is := is.New(t)
	got := GetSchemaRequest(have)
	is.Equal("", cmp.Diff(want, got))
}

func TestCreateSchemaResponse(t *testing.T) {
	have := &conduitv1.CreateSchemaResponse{
		Schema: &schemav1.Schema{
			Subject: "foo",
			Version: 2,
			Type:    schemav1.Schema_TYPE_AVRO,
			Bytes:   []byte("bar"),
		},
	}
	want := pconduit.CreateSchemaResponse{
		Schema: schema.Schema{
			Subject: "foo",
			Version: 2,
			Type:    schema.TypeAvro,
			Bytes:   []byte("bar"),
		},
	}

	is := is.New(t)
	got, err := CreateSchemaResponse(have)
	is.NoErr(err)
	is.Equal("", cmp.Diff(want, got))
}

func TestGetSchemaResponse(t *testing.T) {
	have := &conduitv1.GetSchemaResponse{
		Schema: &schemav1.Schema{
			Subject: "foo",
			Version: 2,
			Type:    schemav1.Schema_TYPE_AVRO,
			Bytes:   []byte("bar"),
		},
	}
	want := pconduit.GetSchemaResponse{
		Schema: schema.Schema{
			Subject: "foo",
			Version: 2,
			Type:    schema.TypeAvro,
			Bytes:   []byte("bar"),
		},
	}

	is := is.New(t)
	got, err := GetSchemaResponse(have)
	is.NoErr(err)
	is.Equal("", cmp.Diff(want, got))
}
