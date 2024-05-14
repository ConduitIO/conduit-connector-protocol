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

	"github.com/conduitio/conduit-connector-protocol/cpluginv2"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/matryer/is"
)

func TestSourceConfigureRequest(t *testing.T) {
	have := &connectorv2.Source_Configure_Request{
		Config: map[string]string{
			"foo": "test_config",
		},
	}
	want := cpluginv2.SourceConfigureRequest{
		Config: map[string]string{
			"foo": "test_config",
		},
	}

	is := is.New(t)
	got, err := SourceConfigureRequest(have)
	is.NoErr(err)
	is.Equal(
		"",
		cmp.Diff(want, got,
			cmpopts.IgnoreUnexported(cpluginv2.SourceConfigureRequest{}),
		),
	)
}

func TestSourceStartRequest(t *testing.T) {
	have := &connectorv2.Source_Start_Request{
		Position: []byte("test_position"),
	}
	want := cpluginv2.SourceStartRequest{
		Position: []byte("test_position"),
	}

	is := is.New(t)
	got, err := SourceStartRequest(have)
	is.NoErr(err)
	is.Equal(
		"",
		cmp.Diff(want, got,
			cmpopts.IgnoreUnexported(cpluginv2.SourceStartRequest{}),
		),
	)
}

func TestSourceRunRequest(t *testing.T) {
	have := &connectorv2.Source_Run_Request{
		AckPosition: []byte("test_ack_position"),
	}
	want := cpluginv2.SourceRunRequest{
		AckPosition: []byte("test_ack_position"),
	}

	is := is.New(t)
	got, err := SourceRunRequest(have)
	is.NoErr(err)
	is.Equal(
		"",
		cmp.Diff(want, got,
			cmpopts.IgnoreUnexported(cpluginv2.SourceRunRequest{}),
		),
	)
}
