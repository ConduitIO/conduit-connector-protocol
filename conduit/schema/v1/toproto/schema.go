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
	"github.com/conduitio/conduit-commons/schema"
	v1 "github.com/conduitio/conduit-connector-protocol/conduit/schema"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
)

func GetSchemaRequest(request v1.GetRequest) *conduitv1.GetSchemaRequest {
	return &conduitv1.GetSchemaRequest{
		Name:    request.Name,
		Version: int32(request.Version),
	}
}

func GetSchemaResponse(inst schema.Instance) *conduitv1.GetSchemaResponse {
	return &conduitv1.GetSchemaResponse{
		Schema: &conduitv1.Schema{
			Id:      inst.ID,
			Name:    inst.Name,
			Version: inst.Version,
			Type:    conduitv1.Schema_Type(inst.Type),
			Bytes:   inst.Bytes,
		},
	}
}

func CreateSchemaRequest(request v1.CreateRequest) *conduitv1.CreateSchemaRequest {
	return &conduitv1.CreateSchemaRequest{
		Name:  request.Name,
		Type:  conduitv1.Schema_Type(request.Type),
		Bytes: request.Bytes,
	}
}

func CreateSchemaResponse(inst schema.Instance) *conduitv1.CreateSchemaResponse {
	return &conduitv1.CreateSchemaResponse{
		Schema: &conduitv1.Schema{
			Id:      inst.ID,
			Name:    inst.Name,
			Version: inst.Version,
			Type:    conduitv1.Schema_Type(inst.Type),
			Bytes:   inst.Bytes,
		},
	}
}
