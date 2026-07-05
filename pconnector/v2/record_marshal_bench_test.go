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

// Package v2_test benchmarks the record conversion + (un)marshal hot path
// that runs once per record on every Conduit <-> connector plugin gRPC call.
//
// The v1 protocol (pconnector/v1) is deprecated in favor of v2, so these
// benchmarks target the v2 wire types, which is where future performance
// work on this seam should focus.
package v2_test

import (
	"testing"

	"github.com/conduitio/conduit-commons/opencdc"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v2/fromproto"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v2/toproto"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
	"google.golang.org/protobuf/proto"
)

// Sinks that benchmark results are written to, so the compiler can't optimize
// away the work being measured.
var (
	benchBytesSink   []byte
	benchRecordsSink []opencdc.Record
)

// benchJSONPayload is a realistic raw-data payload (a serialized order row),
// large enough that its marshaling cost isn't dwarfed by fixed overhead.
const benchJSONPayload = `{"id":48219,"customer":"Ada Lovelace","email":"ada@example.com","total_cents":542399,"currency":"USD","status":"paid","items":[{"sku":"BABBAGE-1","qty":2,"price_cents":129999},{"sku":"LOVELACE-NOTES","qty":1,"price_cents":282401}],"shipping_address":{"street":"12 Analytical Engine Way","city":"London","postal_code":"EC1A 1BB","country":"UK"},"created_at":"2026-07-04T12:00:00Z"}`

// benchRecord returns a fixed, realistic opencdc.Record as it would arrive
// from a CDC source: a position, standard + custom metadata, a key, and an
// After payload. Structured data is far more expensive to convert/marshal
// than raw bytes (it goes through structpb reflection), so callers pass the
// payload shape they want to measure.
func benchRecord(after opencdc.Data) opencdc.Record {
	return opencdc.Record{
		Position:  opencdc.Position("postgres/wal/000000010000000000000042"),
		Operation: opencdc.OperationCreate,
		Metadata: opencdc.Metadata{
			opencdc.MetadataCreatedAt:               "1735689600000000000",
			opencdc.MetadataCollection:              "public.orders",
			opencdc.MetadataConduitSourcePluginName: "builtin:postgres",
			"postgres.table":                        "orders",
			"postgres.wal.lsn":                      "0/1708A50",
		},
		Key: opencdc.RawData(`{"id":48219}`),
		Payload: opencdc.Change{
			Before: nil,
			After:  after,
		},
	}
}

func benchStructuredPayload() opencdc.StructuredData {
	return opencdc.StructuredData{
		"id":          48219,
		"customer":    "Ada Lovelace",
		"email":       "ada@example.com",
		"total_cents": 542399,
		"currency":    "USD",
		"status":      "paid",
		"items": []any{
			map[string]any{"sku": "BABBAGE-1", "qty": 2, "price_cents": 129999},
			map[string]any{"sku": "LOVELACE-NOTES", "qty": 1, "price_cents": 282401},
		},
		"shipping_address": map[string]any{
			"street":      "12 Analytical Engine Way",
			"city":        "London",
			"postal_code": "EC1A 1BB",
			"country":     "UK",
		},
		"created_at": "2026-07-04T12:00:00Z",
	}
}

// benchCase is one payload shape to run every benchmark below against.
type benchCase struct {
	name string
	rec  opencdc.Record
}

func benchCases() []benchCase {
	return []benchCase{
		{name: "raw", rec: benchRecord(opencdc.RawData(benchJSONPayload))},
		{name: "structured", rec: benchRecord(benchStructuredPayload())},
	}
}

// BenchmarkSourceRunResponseMarshal benchmarks the outbound hot path a source
// plugin executes once per record: converting an opencdc.Record to its wire
// proto representation (toproto.SourceRunResponse) and marshaling the
// resulting message to bytes before it crosses the gRPC seam to the engine.
func BenchmarkSourceRunResponseMarshal(b *testing.B) {
	for _, bc := range benchCases() {
		resp := pconnector.SourceRunResponse{Records: []opencdc.Record{bc.rec}}

		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				wire, err := toproto.SourceRunResponse(resp)
				if err != nil {
					b.Fatal(err)
				}
				data, err := proto.Marshal(wire)
				if err != nil {
					b.Fatal(err)
				}
				benchBytesSink = data
			}
		})
	}
}

// BenchmarkSourceRunResponseUnmarshal benchmarks the inbound hot path the
// engine executes once per record: unmarshaling wire bytes received from a
// source plugin and converting the proto message back into an opencdc.Record
// (fromproto.SourceRunResponse).
func BenchmarkSourceRunResponseUnmarshal(b *testing.B) {
	for _, bc := range benchCases() {
		resp := pconnector.SourceRunResponse{Records: []opencdc.Record{bc.rec}}
		wire, err := toproto.SourceRunResponse(resp)
		if err != nil {
			b.Fatal(err)
		}
		data, err := proto.Marshal(wire)
		if err != nil {
			b.Fatal(err)
		}

		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var msg connectorv2.Source_Run_Response
				if err := proto.Unmarshal(data, &msg); err != nil {
					b.Fatal(err)
				}
				out, err := fromproto.SourceRunResponse(&msg)
				if err != nil {
					b.Fatal(err)
				}
				benchRecordsSink = out.Records
			}
		})
	}
}

// BenchmarkDestinationRunRequestMarshal benchmarks the outbound hot path the
// engine executes once per record: converting an opencdc.Record to its wire
// proto representation (toproto.DestinationRunRequest) and marshaling it
// before sending it to a destination plugin.
func BenchmarkDestinationRunRequestMarshal(b *testing.B) {
	for _, bc := range benchCases() {
		req := pconnector.DestinationRunRequest{Records: []opencdc.Record{bc.rec}}

		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				wire, err := toproto.DestinationRunRequest(req)
				if err != nil {
					b.Fatal(err)
				}
				data, err := proto.Marshal(wire)
				if err != nil {
					b.Fatal(err)
				}
				benchBytesSink = data
			}
		})
	}
}

// BenchmarkDestinationRunRequestUnmarshal benchmarks the inbound hot path a
// destination plugin executes once per record: unmarshaling wire bytes
// received from the engine and converting the proto message back into an
// opencdc.Record (fromproto.DestinationRunRequest).
func BenchmarkDestinationRunRequestUnmarshal(b *testing.B) {
	for _, bc := range benchCases() {
		req := pconnector.DestinationRunRequest{Records: []opencdc.Record{bc.rec}}
		wire, err := toproto.DestinationRunRequest(req)
		if err != nil {
			b.Fatal(err)
		}
		data, err := proto.Marshal(wire)
		if err != nil {
			b.Fatal(err)
		}

		b.Run(bc.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var msg connectorv2.Destination_Run_Request
				if err := proto.Unmarshal(data, &msg); err != nil {
					b.Fatal(err)
				}
				out, err := fromproto.DestinationRunRequest(&msg)
				if err != nil {
					b.Fatal(err)
				}
				benchRecordsSink = out.Records
			}
		})
	}
}
