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

package fromproto

import (
	"errors"
	"fmt"

	"github.com/conduitio/conduit-plugin/cpluginv1"
	"github.com/conduitio/conduit-plugin/cpluginv1/internal/cproto"
)

func Record(record *cproto.Record) (cpluginv1.Record, error) {
	key, err := Data(record.Key)
	if err != nil {
		return cpluginv1.Record{}, fmt.Errorf("error converting key: %w", err)
	}

	payload, err := Data(record.Payload)
	if err != nil {
		return cpluginv1.Record{}, fmt.Errorf("error converting payload: %w", err)
	}

	out := cpluginv1.Record{
		Position:  record.Position,
		Metadata:  record.Metadata,
		CreatedAt: record.CreatedAt.AsTime(),
		Key:       key,
		Payload:   payload,
	}
	return out, nil
}

func Data(in *cproto.Data) (cpluginv1.Data, error) {
	d := in.GetData()
	if d == nil {
		return nil, nil
	}

	switch v := d.(type) {
	case *cproto.Data_RawData:
		return cpluginv1.RawData(v.RawData.Raw), nil
	case *cproto.Data_StructuredData:
		return cpluginv1.StructuredData(v.StructuredData.AsMap()), nil
	default:
		return nil, errors.New("invalid Data type")
	}
}
