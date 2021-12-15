// Copyright © 2021 Meroxa Inc
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
	"errors"
	"fmt"

	"github.com/conduitio/conduit-plugin/cpluginv1"
	connectorv1 "go.buf.build/library/go-grpc/conduitio/conduit-plugin/connector/v1"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Record(record cpluginv1.Record) (*connectorv1.Record, error) {
	key, err := Data(record.Key)
	if err != nil {
		return nil, fmt.Errorf("error converting key: %w", err)
	}

	payload, err := Data(record.Payload)
	if err != nil {
		return nil, fmt.Errorf("error converting payload: %w", err)
	}

	out := connectorv1.Record{
		Position:  record.Position,
		Metadata:  record.Metadata,
		CreatedAt: timestamppb.New(record.CreatedAt),
		Key:       key,
		Payload:   payload,
	}
	return &out, nil
}

func Data(in cpluginv1.Data) (*connectorv1.Data, error) {
	var out connectorv1.Data

	switch v := in.(type) {
	case cpluginv1.RawData:
		out.Data = &connectorv1.Data_RawData{
			RawData: v,
		}
	case cpluginv1.StructuredData:
		content, err := structpb.NewStruct(v)
		if err != nil {
			return nil, err
		}
		out.Data = &connectorv1.Data_StructuredData{
			StructuredData: content,
		}
	default:
		return nil, errors.New("invalid Data type")
	}

	return &out, nil
}
