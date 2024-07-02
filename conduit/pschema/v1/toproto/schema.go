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
	schemav1 "github.com/conduitio/conduit-commons/proto/schema/v1"
	"github.com/conduitio/conduit-connector-protocol/conduit/pschema"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
)

// -- Request Conversions -----------------------------------------------------

func CreateSchemaRequest(in pschema.CreateSchemaRequest) *conduitv1.CreateSchemaRequest {
	return &conduitv1.CreateSchemaRequest{
		Subject: in.Subject,
		Type:    schemav1.Schema_Type(in.Type),
		Bytes:   in.Bytes,
	}
}

func GetSchemaRequest(in pschema.GetSchemaRequest) *conduitv1.GetSchemaRequest {
	return &conduitv1.GetSchemaRequest{
		Subject: in.Subject,
		Version: int32(in.Version),
	}
}

// -- Response Conversions ----------------------------------------------------

func CreateSchemaResponse(in pschema.CreateSchemaResponse) (*conduitv1.CreateSchemaResponse, error) {
	outSchema := &schemav1.Schema{}
	err := in.Schema.ToProto(outSchema)
	if err != nil {
		return &conduitv1.CreateSchemaResponse{}, err
	}

	return &conduitv1.CreateSchemaResponse{
		Schema: outSchema,
	}, nil
}

func GetSchemaResponse(in pschema.GetSchemaResponse) (*conduitv1.GetSchemaResponse, error) {
	outSchema := &schemav1.Schema{}
	err := in.Schema.ToProto(outSchema)
	if err != nil {
		return &conduitv1.GetSchemaResponse{}, err
	}

	return &conduitv1.GetSchemaResponse{
		Schema: outSchema,
	}, nil
}
