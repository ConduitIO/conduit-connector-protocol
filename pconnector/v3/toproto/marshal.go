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

package toproto

import (
	"encoding/binary"
	"fmt"
	"sync"

	"github.com/conduitio/conduit-commons/opencdc"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

// recordBytesPool reuses the buffer that the run-stream Send path builds the
// record envelope into. Reusing it means the emit path allocates nothing per
// record: the buffer grows to the largest batch seen and is then retained for
// the next batch.
//
// The pool is goroutine-safe, so concurrent streams never share a buffer. A
// buffer is returned to the pool only after the gRPC stream has copied the
// bytes into its own frame (see RawMessage), so there is no aliasing between a
// pooled buffer and bytes still on the wire.
var recordBytesPool = sync.Pool{
	New: func() any {
		buf := make([]byte, 0, 1024)
		return &buf
	},
}

// GetRecordBytes returns a reusable byte buffer for building a record envelope.
// The returned slice has length zero; its backing array is retained between
// calls so a batch that fits the previous capacity allocates nothing.
func GetRecordBytes() []byte {
	buf := recordBytesPool.Get().(*[]byte) //nolint:forcetypeassert // the pool only ever holds *[]byte
	return (*buf)[:0]
}

// PutRecordBytes returns a record buffer to the pool. It keeps its backing
// array so the next batch reuses the capacity.
func PutRecordBytes(b []byte) {
	recordBytesPool.Put(&b)
}

// SourceRunResponseBytes appends the wire encoding of the v3 Source_Run_Response
// (a single repeated opencdc.v2.Record field, number 1) to dst and returns the
// extended slice. Each record is encoded directly from the Go record with
// Record.AppendV2, bypassing the ToProtoV2 message tree and proto reflection.
//
// The output is byte-identical to
// proto.MarshalOptions{Deterministic: true}.Marshal(toproto.SourceRunResponse(in)).
func SourceRunResponseBytes(dst []byte, in pconnector.SourceRunResponse) ([]byte, error) {
	return marshalRecordsV2(dst, in.Records)
}

// DestinationRunRequestBytes appends the wire encoding of the v3
// Destination_Run_Request (a single repeated opencdc.v2.Record field, number 1)
// to dst and returns the extended slice. See SourceRunResponseBytes.
func DestinationRunRequestBytes(dst []byte, in pconnector.DestinationRunRequest) ([]byte, error) {
	return marshalRecordsV2(dst, in.Records)
}

// marshalRecordsV2 appends the wire encoding of a repeated opencdc.v2.Record
// field (field number 1) to dst: per record a field-1 tag, a varint length
// prefix, then the record's own wire encoding from Record.AppendV2.
func marshalRecordsV2(dst []byte, records []opencdc.Record) ([]byte, error) {
	for i := range records {
		dst = protowire.AppendTag(dst, 1, protowire.BytesType)

		// Reserve 10 bytes (the maximum varint width) for the length prefix,
		// append the record body after it, then shift the body left and write
		// the real length. This mirrors the protobuf marshaler and allocates
		// nothing beyond growing dst.
		lenPos := len(dst)
		dst = append(dst, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
		bodyStart := len(dst)

		var err error
		dst, err = records[i].AppendV2(dst)
		if err != nil {
			return dst, err
		}

		bodyLen := len(dst) - bodyStart
		var tmp [binary.MaxVarintLen64]byte
		prefix := protowire.AppendVarint(tmp[:0], uint64(bodyLen)) //nolint:gosec // bodyLen is a non-negative length
		m := len(prefix)
		copy(dst[lenPos+m:], dst[bodyStart:]) // shift the body left by (10 - m)
		dst = dst[:lenPos+m+bodyLen]
		copy(dst[lenPos:], prefix)
	}
	return dst, nil
}

// RawMessage wraps pre-marshaled wire bytes as a proto.Message so a gRPC stream
// can send them with SendMsg without proto reflection re-marshaling a message
// tree. Marshal copies data into the caller's buffer, so the returned message
// does not alias data after Marshal returns; data may be reused as soon as the
// gRPC send returns.
func RawMessage(data []byte) proto.Message {
	return &rawMessage{data: data}
}

type rawMessage struct {
	data []byte
}

func (m *rawMessage) Reset() { m.data = nil }

func (m *rawMessage) String() string {
	return fmt.Sprintf("raw message (%d bytes)", len(m.data))
}

func (m *rawMessage) ProtoReflect() protoreflect.Message {
	return &rawMessageReflect{m: m}
}

// rawMessageReflect is the minimal protoreflect.Message needed to marshal a
// rawMessage. It embeds the protoreflect.Message interface so the ~20 methods
// that Marshal and Size never use need no implementation; only IsValid and
// ProtoMethods are overridden. Calling any other method panics, which is the
// intended behavior: a rawMessage is only ever marshaled.
type rawMessageReflect struct {
	protoreflect.Message
	m *rawMessage
}

func (r *rawMessageReflect) IsValid() bool { return true }

func (r *rawMessageReflect) ProtoMethods() *protoiface.Methods {
	return &rawMessageMethods
}

// rawMessageMethods is the proto fast path shared by every rawMessage: Size
// reports the length of the pre-marshaled bytes and Marshal appends them to the
// caller's buffer (a copy, so the pre-marshaled bytes are never aliased).
var rawMessageMethods = protoiface.Methods{
	Flags: protoiface.SupportMarshalDeterministic,
	Size: func(in protoiface.SizeInput) protoiface.SizeOutput {
		return protoiface.SizeOutput{Size: len(rawMessageOf(in.Message).data)}
	},
	Marshal: func(in protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		return protoiface.MarshalOutput{Buf: append(in.Buf, rawMessageOf(in.Message).data...)}, nil
	},
	// CheckInitialized is implemented so proto.Marshal's required-field check
	// (AllowPartial is false for the gRPC codec) takes the fast path instead of
	// reflecting over the message, which a raw message cannot do. The bytes were
	// already produced by a valid encode, so there are no required fields to
	// check.
	CheckInitialized: func(protoiface.CheckInitializedInput) (protoiface.CheckInitializedOutput, error) {
		return protoiface.CheckInitializedOutput{}, nil
	},
}

func rawMessageOf(m protoreflect.Message) *rawMessage {
	return m.(*rawMessageReflect).m //nolint:forcetypeassert // ProtoReflect always returns *rawMessageReflect
}
