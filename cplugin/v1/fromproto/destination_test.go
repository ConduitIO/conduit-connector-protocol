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
	opencdcv1 "github.com/conduitio/conduit-commons/proto/opencdc/v1"
	"github.com/conduitio/conduit-connector-protocol/cplugin"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/matryer/is"
	"google.golang.org/protobuf/types/known/structpb"
)

var (
	testRecordGo = opencdc.Record{
		Position:  opencdc.Position("standing"),
		Operation: opencdc.OperationUpdate,
		Metadata:  opencdc.Metadata{"foo": "bar"},
		Key:       opencdc.RawData("padlock-key"),
		Payload: opencdc.Change{
			Before: opencdc.RawData("yellow"),
			After: opencdc.StructuredData{
				"bool": true,

				"int":   float64(1),
				"int32": float64(1),
				"int64": float64(1),

				"float32": float64(1.5),
				"float64": float64(1.5),

				"string": "orange",

				"nested": map[string]any{
					"bool":    true,
					"int":     float64(2),
					"float32": float64(2.3),
					"string":  "blue",
				},
			},
		},
	}
	testRecordProto = opencdcv1.Record{
		Position:  []byte("standing"),
		Operation: opencdcv1.Operation_OPERATION_UPDATE,
		Metadata:  map[string]string{"foo": "bar"},
		Key:       &opencdcv1.Data{Data: &opencdcv1.Data_RawData{RawData: []byte("padlock-key")}},
		Payload: &opencdcv1.Change{
			Before: &opencdcv1.Data{Data: &opencdcv1.Data_RawData{RawData: []byte("yellow")}},
			After: &opencdcv1.Data{
				Data: &opencdcv1.Data_StructuredData{StructuredData: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"bool":    {Kind: &structpb.Value_BoolValue{BoolValue: true}},
						"int":     {Kind: &structpb.Value_NumberValue{NumberValue: 1}},
						"int32":   {Kind: &structpb.Value_NumberValue{NumberValue: 1}},
						"int64":   {Kind: &structpb.Value_NumberValue{NumberValue: 1}},
						"float32": {Kind: &structpb.Value_NumberValue{NumberValue: 1.5}},
						"float64": {Kind: &structpb.Value_NumberValue{NumberValue: 1.5}},
						"string":  {Kind: &structpb.Value_StringValue{StringValue: "orange"}},
						"nested": {Kind: &structpb.Value_StructValue{StructValue: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"bool":    {Kind: &structpb.Value_BoolValue{BoolValue: true}},
								"int":     {Kind: &structpb.Value_NumberValue{NumberValue: 2}},
								"float32": {Kind: &structpb.Value_NumberValue{NumberValue: 2.3}},
								"string":  {Kind: &structpb.Value_StringValue{StringValue: "blue"}},
							},
						}}},
					},
				}},
			},
		},
	}
)

func TestDestinationConfigureRequest(t *testing.T) {
	have := &connectorv1.Destination_Configure_Request{
		Config: map[string]string{
			"foo": "test_config",
		},
	}
	want := cplugin.DestinationConfigureRequest{
		Config: map[string]string{
			"foo": "test_config",
		},
	}

	is := is.New(t)
	got := DestinationConfigureRequest(have)
	is.Equal(
		"",
		cmp.Diff(want, got,
			cmpopts.IgnoreUnexported(cplugin.DestinationConfigureRequest{}),
		),
	)
}

func TestDestinationRunRequest(t *testing.T) {
	have := &connectorv1.Destination_Run_Request{
		Record: &testRecordProto,
	}
	want := cplugin.DestinationRunRequest{
		Records: []opencdc.Record{testRecordGo},
	}

	is := is.New(t)
	got, err := DestinationRunRequest(have)
	is.NoErr(err)
	is.Equal(
		"",
		cmp.Diff(want, got,
			cmpopts.IgnoreUnexported(cplugin.DestinationRunRequest{}),
			cmpopts.IgnoreUnexported(opencdc.Record{}),
		),
	)
}

func TestDestinationStopRequest(t *testing.T) {
	have := &connectorv1.Destination_Stop_Request{
		LastPosition: []byte("test_position"),
	}
	want := cplugin.DestinationStopRequest{
		LastPosition: []byte("test_position"),
	}

	is := is.New(t)
	got := DestinationStopRequest(have)
	is.Equal(
		"",
		cmp.Diff(want, got,
			cmpopts.IgnoreUnexported(cplugin.DestinationStopRequest{}),
		),
	)
}
