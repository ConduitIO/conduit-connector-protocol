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

package v1

import (
	"fmt"
	"github.com/conduitio/conduit-commons/schema"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
)

func GetResponse(inst schema.Instance) *conduitv1.GetResponse {
	return &conduitv1.GetResponse{
		Schema: &conduitv1.Schema{
			Id:      inst.ID,
			Name:    inst.Name,
			Version: inst.Version,
			Type:    ProtoType(inst.Type),
			Bytes:   inst.Bytes,
		},
	}
}

func ProtoType(t schema.Type) conduitv1.Schema_Type {
	switch t {
	case schema.TypeAvro:
		return conduitv1.Schema_TYPE_AVRO
	default:
		panic(fmt.Errorf("unsupported schema type %q", t))
	}
}

func CreateResponse(inst schema.Instance) *conduitv1.CreateResponse {
	return &conduitv1.CreateResponse{
		Schema: &conduitv1.Schema{
			Id:      inst.ID,
			Name:    inst.Name,
			Version: inst.Version,
			Type:    ProtoType(inst.Type),
			Bytes:   inst.Bytes,
		},
	}
}
