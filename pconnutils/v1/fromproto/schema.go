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
	"github.com/conduitio/conduit-connector-protocol/pconnutils"
	connutilsv1 "github.com/conduitio/conduit-connector-protocol/proto/connutils/v1"
)

// -- Request Conversions -----------------------------------------------------

func CreateSchemaRequest(in *connutilsv1.CreateSchemaRequest) pconnutils.CreateSchemaRequest {
	return pconnutils.CreateSchemaRequest{
		Subject: in.Subject,
		Type:    schema.Type(in.Type),
		Bytes:   in.Bytes,
	}
}

func GetSchemaRequest(in *connutilsv1.GetSchemaRequest) pconnutils.GetSchemaRequest {
	return pconnutils.GetSchemaRequest{
		Subject: in.Subject,
		Version: int(in.Version),
	}
}

// -- Response Conversions ----------------------------------------------------

func CreateSchemaResponse(in *connutilsv1.CreateSchemaResponse) (pconnutils.CreateSchemaResponse, error) {
	var outSchema schema.Schema
	err := outSchema.FromProto(in.Schema)
	if err != nil {
		return pconnutils.CreateSchemaResponse{}, err
	}
	return pconnutils.CreateSchemaResponse{Schema: outSchema}, nil
}

func GetSchemaResponse(in *connutilsv1.GetSchemaResponse) (pconnutils.GetSchemaResponse, error) {
	var outSchema schema.Schema
	err := outSchema.FromProto(in.Schema)
	if err != nil {
		return pconnutils.GetSchemaResponse{}, err
	}
	return pconnutils.GetSchemaResponse{Schema: outSchema}, nil
}
