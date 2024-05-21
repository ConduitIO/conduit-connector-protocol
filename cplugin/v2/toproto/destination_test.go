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

package toproto

import (
	"testing"

	"github.com/conduitio/conduit-connector-protocol/cplugin"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/matryer/is"
)

func TestDestinationRunResponse(t *testing.T) {
	have := cplugin.DestinationRunResponse{
		Acks: []cplugin.DestinationRunResponseAck{{
			Position: []byte("test_position_1"),
			Error:    "test_error",
		}, {
			Position: []byte("test_position_2"),
		}},
	}
	want := &connectorv2.Destination_Run_Response{
		Acks: []*connectorv2.Destination_Run_Response_Ack{{
			Position: []byte("test_position_1"),
			Error:    "test_error",
		}, {
			Position: []byte("test_position_2"),
		}},
	}

	is := is.New(t)
	got := DestinationRunResponse(have)
	is.Equal(
		"",
		cmp.Diff(want, got,
			cmpopts.IgnoreUnexported(connectorv2.Destination_Run_Response{}),
			cmpopts.IgnoreUnexported(connectorv2.Destination_Run_Response_Ack{}),
		),
	)
}
