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

	"github.com/conduitio/conduit-commons/opencdc"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/matryer/is"
)

func TestSourceConfigureRequest(t *testing.T) {
	have := &connectorv1.Source_Configure_Request{
		Config: map[string]string{
			"foo": "test_config",
		},
	}
	want := pconnector.SourceConfigureRequest{
		Config: map[string]string{
			"foo": "test_config",
		},
	}

	is := is.New(t)
	got := SourceConfigureRequest(have)
	is.Equal("", cmp.Diff(want, got))
}

func TestSourceStartRequest(t *testing.T) {
	have := &connectorv1.Source_Start_Request{
		Position: []byte("test_position"),
	}
	want := pconnector.SourceOpenRequest{
		Position: []byte("test_position"),
	}

	is := is.New(t)
	got := SourceStartRequest(have)
	is.Equal("", cmp.Diff(want, got))
}

func TestSourceRunRequest(t *testing.T) {
	have := &connectorv1.Source_Run_Request{
		AckPosition: []byte("1"),
	}
	want := pconnector.SourceRunRequest{
		AckPositions: []opencdc.Position{opencdc.Position("1")},
	}

	is := is.New(t)
	got := SourceRunRequest(have)
	is.Equal("", cmp.Diff(want, got))
}
