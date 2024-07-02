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
	"github.com/conduitio/conduit-commons/schema"
	"github.com/conduitio/conduit-connector-protocol/conduit/pschema"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
)

// -- Request Conversions -----------------------------------------------------

func CreateSchemaRequest(in *conduitv1.CreateSchemaRequest) pschema.CreateSchemaRequest {
	return pschema.CreateSchemaRequest{
		Subject: in.Subject,
		Type:    schema.Type(in.Type),
		Bytes:   in.Bytes,
	}
}

func GetSchemaRequest(in *conduitv1.GetSchemaRequest) pschema.GetSchemaRequest {
	return pschema.GetSchemaRequest{
		Subject: in.Subject,
		Version: int(in.Version),
	}
}

// -- Response Conversions ----------------------------------------------------

func CreateSchemaResponse(in *conduitv1.CreateSchemaResponse) (pschema.CreateSchemaResponse, error) {
	var schema schema.Schema
	err := schema.FromProto(in.Schema)
	if err != nil {
		return pschema.CreateSchemaResponse{}, err
	}
	return pschema.CreateSchemaResponse{Schema: schema}, nil
}

func GetSchemaResponse(in *conduitv1.GetSchemaResponse) (pschema.GetSchemaResponse, error) {
	var schema schema.Schema
	err := schema.FromProto(in.Schema)
	if err != nil {
		return pschema.GetSchemaResponse{}, err
	}
	return pschema.GetSchemaResponse{Schema: schema}, nil
}
