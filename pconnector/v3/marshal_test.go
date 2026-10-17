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

package v3_test

import (
	"bytes"
	"testing"

	"github.com/conduitio/conduit-commons/opencdc"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v3/fromproto"
	"github.com/conduitio/conduit-connector-protocol/pconnector/v3/toproto"
	connectorv3 "github.com/conduitio/conduit-connector-protocol/proto/connector/v3"
	"github.com/matryer/is"
	"google.golang.org/protobuf/proto"
)

// representativeRecords returns a set of records that exercises every field and
// every data shape the wire encoder must handle: the zero record, raw data,
// structured data (including the lossless value kinds), tombstones, and nested
// containers.
func representativeRecords() []opencdc.Record {
	return []opencdc.Record{
		{}, // zero record
		{
			Position:  opencdc.Position("postgres/wal/000000010000000000000042"),
			Operation: opencdc.OperationCreate,
			Metadata: opencdc.Metadata{
				opencdc.MetadataCollection: "public.orders",
				"custom":                   "value",
			},
			Key:     opencdc.RawData(`{"id":48219}`),
			Payload: opencdc.Change{After: opencdc.RawData(`{"id":48219,"total":542399}`)},
		},
		lossyRecord(), // structured with int64/uint64/[]byte/NaN/±Inf/decimal
		{
			Operation: opencdc.OperationDelete, // tombstone: no payload
			Metadata:  opencdc.Metadata{"a": "1", "b": "2"},
			Key:       opencdc.StructuredData{"id": int64(7)},
		},
		{
			Operation: opencdc.OperationSnapshot,
			Payload: opencdc.Change{
				Before: opencdc.StructuredData{"items": []any{int64(1), int64(2), int64(3)}},
				After:  opencdc.StructuredData{"nested": map[string]any{"ok": true, "n": float64(1.5)}},
			},
		},
	}
}

// TestSourceRunResponseBytesByteIdentical proves the AppendV2 fast path emits
// exactly the bytes the ToProtoV2 + deterministic proto.Marshal path emits, for
// every representative record and for a multi-record batch.
func TestSourceRunResponseBytesByteIdentical(t *testing.T) {
	is := is.New(t)

	for _, rec := range representativeRecords() {
		resp := pconnector.SourceRunResponse{Records: []opencdc.Record{rec}}

		wantTree, err := toproto.SourceRunResponse(resp)
		is.NoErr(err)
		want, err := proto.MarshalOptions{Deterministic: true}.Marshal(wantTree)
		is.NoErr(err)

		got, err := toproto.SourceRunResponseBytes(nil, resp)
		is.NoErr(err)

		is.Equal(string(want), string(got)) // string() so a mismatch renders as text, not hex dump
	}

	// Multi-record batch: the framing around each record must match too.
	batch := pconnector.SourceRunResponse{Records: representativeRecords()}
	wantTree, err := toproto.SourceRunResponse(batch)
	is.NoErr(err)
	want, err := proto.MarshalOptions{Deterministic: true}.Marshal(wantTree)
	is.NoErr(err)
	got, err := toproto.SourceRunResponseBytes(nil, batch)
	is.NoErr(err)
	is.Equal(string(want), string(got))
}

// TestDestinationRunRequestBytesByteIdentical is the destination-side twin of
// TestSourceRunResponseBytesByteIdentical.
func TestDestinationRunRequestBytesByteIdentical(t *testing.T) {
	is := is.New(t)

	for _, rec := range representativeRecords() {
		req := pconnector.DestinationRunRequest{Records: []opencdc.Record{rec}}

		wantTree, err := toproto.DestinationRunRequest(req)
		is.NoErr(err)
		want, err := proto.MarshalOptions{Deterministic: true}.Marshal(wantTree)
		is.NoErr(err)

		got, err := toproto.DestinationRunRequestBytes(nil, req)
		is.NoErr(err)

		is.Equal(string(want), string(got))
	}
}

// TestSourceRunResponseBytesAppends proves the fast path appends to a non-empty
// destination buffer instead of overwriting it, matching AppendV2's contract.
func TestSourceRunResponseBytesAppends(t *testing.T) {
	is := is.New(t)

	resp := pconnector.SourceRunResponse{Records: []opencdc.Record{lossyRecord()}}
	prefix := []byte("prefix:")

	got, err := toproto.SourceRunResponseBytes(prefix, resp)
	is.NoErr(err)

	var bare []byte
	bare, err = toproto.SourceRunResponseBytes(bare, resp)
	is.NoErr(err)

	is.True(bytes.HasPrefix(got, prefix))
	is.Equal(string(bare), string(got[len(prefix):]))
}

// TestSourceRunResponseBytesRoundTrip proves a record encoded with the
// AppendV2 fast path still decodes (unmarshal + fromproto) losslessly, so the
// destination decode side is unaffected by the emit-only change.
func TestSourceRunResponseBytesRoundTrip(t *testing.T) {
	is := is.New(t)

	resp := pconnector.SourceRunResponse{Records: representativeRecords()}
	data, err := toproto.SourceRunResponseBytes(nil, resp)
	is.NoErr(err)

	var msg connectorv3.Source_Run_Response
	err = proto.Unmarshal(data, &msg)
	is.NoErr(err)

	got, err := fromproto.SourceRunResponse(&msg)
	is.NoErr(err)

	is.Equal(len(resp.Records), len(got.Records))
	// The lossless record is index 2 in representativeRecords.
	assertLosslessAfter(t, got.Records[2].Payload.After)
}

// TestSourceRunResponseBytesErrorParity proves the fast path returns an error
// (and not silently different bytes) for a record that ToProtoV2 also rejects.
func TestSourceRunResponseBytesErrorParity(t *testing.T) {
	is := is.New(t)

	// A data type the wire format cannot carry is rejected by both paths.
	rec := opencdc.Record{
		Operation: opencdc.OperationCreate,
		Payload:   opencdc.Change{After: opencdc.StructuredData{"bad": complex(1, 2)}},
	}
	resp := pconnector.SourceRunResponse{Records: []opencdc.Record{rec}}

	_, wantErr := toproto.SourceRunResponse(resp)
	_, gotErr := toproto.SourceRunResponseBytes(nil, resp)

	is.True(wantErr != nil)
	is.True(gotErr != nil)
}

// TestRawMessageMarshal proves the rawMessage wrapper round-trips through the
// exact marshal sequence the gRPC proto codec uses (proto.Size then
// proto.MarshalOptions{UseCachedSize: true}.Marshal) and emits the original
// bytes unchanged, with no aliasing.
func TestRawMessageMarshal(t *testing.T) {
	is := is.New(t)

	resp := pconnector.SourceRunResponse{Records: representativeRecords()}
	data, err := toproto.SourceRunResponseBytes(nil, resp)
	is.NoErr(err)

	msg := toproto.RawMessage(data)

	is.Equal(len(data), proto.Size(msg))

	out, err := (proto.MarshalOptions{UseCachedSize: true}).Marshal(msg)
	is.NoErr(err)
	is.Equal(string(data), string(out))

	// Marshal must copy, not alias: mutating the source buffer must not change
	// the marshaled output.
	out[0] ^= 0xff
	is.True(!bytes.Equal(data, out))
}

// --- Benchmarks ------------------------------------------------------------

// benchSink is where benchmark results are written so the compiler can't
// optimize away the work being measured.
var benchSink []byte

// BenchmarkSourceRunResponseBytes reports the fast-path emit cost with a reused
// buffer. Run it side by side with BenchmarkSourceRunResponseToProtoMarshal for
// the before/after numbers.
func BenchmarkSourceRunResponseBytes(b *testing.B) {
	resp := pconnector.SourceRunResponse{Records: []opencdc.Record{lossyRecord()}}
	buf := make([]byte, 0, 1024)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var err error
		buf, err = toproto.SourceRunResponseBytes(buf[:0], resp)
		if err != nil {
			b.Fatal(err)
		}
		benchSink = buf
	}
}

// BenchmarkSourceRunResponseToProtoMarshal is the "before": build the
// opencdc.v2 message tree with ToProtoV2, then marshal it with proto reflection.
func BenchmarkSourceRunResponseToProtoMarshal(b *testing.B) {
	resp := pconnector.SourceRunResponse{Records: []opencdc.Record{lossyRecord()}}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree, err := toproto.SourceRunResponse(resp)
		if err != nil {
			b.Fatal(err)
		}
		data, err := proto.Marshal(tree)
		if err != nil {
			b.Fatal(err)
		}
		benchSink = data
	}
}
