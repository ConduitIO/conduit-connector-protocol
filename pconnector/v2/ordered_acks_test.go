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

package v2_test

import (
	"testing"

	"github.com/conduitio/conduit-connector-protocol/pconnector"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v2/fromproto"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v2/toproto"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
	"github.com/matryer/is"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// TestDestinationOpenResponse_OrderedAcks_RoundTrip checks that the ordered_acks
// declaration survives toproto, the wire and fromproto in both states.
func TestDestinationOpenResponse_OrderedAcks_RoundTrip(t *testing.T) {
	for _, want := range []bool{false, true} {
		is := is.New(t)
		wire, err := proto.Marshal(toproto.DestinationOpenResponse(pconnector.DestinationOpenResponse{OrderedAcks: want}))
		is.NoErr(err)

		var msg connectorv2.Destination_Open_Response
		is.NoErr(proto.Unmarshal(wire, &msg))
		is.Equal(fromproto.DestinationOpenResponse(&msg).OrderedAcks, want)
	}
}

// TestDestinationOpenResponse_OrderedAcks_WireCompat covers both directions of
// compatibility: an old plugin's empty response reads as false, and a new
// plugin's response is accepted by an old engine, which knows no fields
// (modeled by the empty message) and ignores the unknown one.
func TestDestinationOpenResponse_OrderedAcks_WireCompat(t *testing.T) {
	is := is.New(t)

	// Old plugin -> new engine.
	oldWire, err := proto.Marshal(&emptypb.Empty{})
	is.NoErr(err)
	var fromOld connectorv2.Destination_Open_Response
	is.NoErr(proto.Unmarshal(oldWire, &fromOld))
	is.Equal(fromproto.DestinationOpenResponse(&fromOld).OrderedAcks, false)

	// New plugin -> old engine.
	newWire, err := proto.Marshal(&connectorv2.Destination_Open_Response{OrderedAcks: true})
	is.NoErr(err)
	is.NoErr(proto.Unmarshal(newWire, &emptypb.Empty{}))

	// The field number is part of the contract: never renumber or reuse it.
	is.Equal(newWire, []byte{0x08, 0x01}) // field 1, varint, true
}
