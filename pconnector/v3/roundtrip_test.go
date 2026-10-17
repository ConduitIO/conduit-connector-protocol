// Copyright © 2026 Meroxa, Inc.
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

// Package v3_test verifies that the v3 wire format carries structured values
// losslessly: int64 beyond 2^53, uint64, []byte, NaN/±Inf and decimals all
// survive an encode (toproto + marshal) and decode (unmarshal + fromproto)
// round trip with their Go type intact. The v1/v2 wire format (structpb) would
// round, refuse, or stringify each of these (see conduit#3100 and ADR
// 20261015-lossless-structured-values-protocol-v3.md).
package v3_test

import (
	"bytes"
	"encoding/json"
	"math"
	"testing"

	"github.com/conduitio/conduit-commons/opencdc"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v3/fromproto"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v3/toproto"
	connectorv3 "github.com/conduitio/conduit-connector-protocol/proto/connector/v3"
	"github.com/matryer/is"
	"google.golang.org/protobuf/proto"
)

const (
	// twoTo53PlusOne is 2^53 + 1, the first integer a float64 rounds.
	twoTo53PlusOne = int64(9007199254740993)
	// maxUint64 does not fit in a float64 at all.
	maxUint64 = uint64(18446744073709551615)
)

// lossyStructuredData is a payload exercising every value kind that the v1/v2
// structpb wire format cannot carry exactly.
func lossyStructuredData() opencdc.StructuredData {
	return opencdc.StructuredData{
		"int64":   twoTo53PlusOne,
		"uint64":  maxUint64,
		"bytes":   []byte{0x00, 0x01, 0x02, 0xfe, 0xff},
		"nan":     math.NaN(),
		"pos_inf": math.Inf(1),
		"neg_inf": math.Inf(-1),
		"decimal": json.Number("12345678901234567890.12345678901234567890"),
	}
}

func lossyRecord() opencdc.Record {
	return opencdc.Record{
		Position:  opencdc.Position("position"),
		Operation: opencdc.OperationUpdate,
		Metadata:  opencdc.Metadata{"foo": "bar"},
		Key:       opencdc.RawData("key"),
		Payload: opencdc.Change{
			Before: opencdc.StructuredData{"before": int64(1)},
			After:  lossyStructuredData(),
		},
	}
}

// TestSourceRunResponseRoundTrip proves a source record survives the v3
// encode/decode round trip with every lossless value kind intact.
func TestSourceRunResponseRoundTrip(t *testing.T) {
	is := is.New(t)

	wire, err := toproto.SourceRunResponse(pconnector.SourceRunResponse{
		Records: []opencdc.Record{lossyRecord()},
	})
	is.NoErr(err)

	data, err := proto.Marshal(wire)
	is.NoErr(err)

	var msg connectorv3.Source_Run_Response
	err = proto.Unmarshal(data, &msg)
	is.NoErr(err)

	got, err := fromproto.SourceRunResponse(&msg)
	is.NoErr(err)

	is.Equal(1, len(got.Records))
	assertLosslessAfter(t, got.Records[0].Payload.After)
}

// TestDestinationRunRequestRoundTrip proves a destination record survives the
// v3 encode/decode round trip with every lossless value kind intact.
func TestDestinationRunRequestRoundTrip(t *testing.T) {
	is := is.New(t)

	wire, err := toproto.DestinationRunRequest(pconnector.DestinationRunRequest{
		Records: []opencdc.Record{lossyRecord()},
	})
	is.NoErr(err)

	data, err := proto.Marshal(wire)
	is.NoErr(err)

	var msg connectorv3.Destination_Run_Request
	err = proto.Unmarshal(data, &msg)
	is.NoErr(err)

	got, err := fromproto.DestinationRunRequest(&msg)
	is.NoErr(err)

	is.Equal(1, len(got.Records))
	assertLosslessAfter(t, got.Records[0].Payload.After)
}

func assertLosslessAfter(t *testing.T, after opencdc.Data) {
	t.Helper()
	is := is.New(t)

	sd, ok := after.(opencdc.StructuredData)
	is.True(ok) // after must still be structured data

	// int64 > 2^53 must come back as the exact int64, not a rounded float64.
	is.Equal(twoTo53PlusOne, sd["int64"])
	// uint64 must come back as the exact uint64.
	is.Equal(maxUint64, sd["uint64"])
	// []byte must come back as []byte, not a base64 string.
	rawBytes, ok := sd["bytes"].([]byte)
	is.True(ok)
	is.True(bytes.Equal([]byte{0x00, 0x01, 0x02, 0xfe, 0xff}, rawBytes))

	// NaN and ±Inf must come back as float64 values, not the strings the
	// structpb wire format produces.
	nan, ok := sd["nan"].(float64)
	is.True(ok)
	is.True(math.IsNaN(nan))
	posInf, ok := sd["pos_inf"].(float64)
	is.True(ok)
	is.True(math.IsInf(posInf, 1))
	negInf, ok := sd["neg_inf"].(float64)
	is.True(ok)
	is.True(math.IsInf(negInf, -1))

	// The decimal must come back as a json.Number with its exact literal form.
	is.Equal(json.Number("12345678901234567890.12345678901234567890"), sd["decimal"])
}
