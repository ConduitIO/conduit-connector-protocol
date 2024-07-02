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
	"github.com/conduitio/conduit-connector-protocol/pconduit"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
)

// -- Request Conversions -----------------------------------------------------

func CreateSchemaRequest(in *conduitv1.CreateSchemaRequest) pconduit.CreateSchemaRequest {
	return pconduit.CreateSchemaRequest{
		Subject: in.Subject,
		Type:    schema.Type(in.Type),
		Bytes:   in.Bytes,
	}
}

func GetSchemaRequest(in *conduitv1.GetSchemaRequest) pconduit.GetSchemaRequest {
	return pconduit.GetSchemaRequest{
		Subject: in.Subject,
		Version: int(in.Version),
	}
}

// -- Response Conversions ----------------------------------------------------

func CreateSchemaResponse(in *conduitv1.CreateSchemaResponse) (pconduit.CreateSchemaResponse, error) {
	var outSchema schema.Schema
	err := outSchema.FromProto(in.Schema)
	if err != nil {
		return pconduit.CreateSchemaResponse{}, err
	}
	return pconduit.CreateSchemaResponse{Schema: outSchema}, nil
}

func GetSchemaResponse(in *conduitv1.GetSchemaResponse) (pconduit.GetSchemaResponse, error) {
	var outSchema schema.Schema
	err := outSchema.FromProto(in.Schema)
	if err != nil {
		return pconduit.GetSchemaResponse{}, err
	}
	return pconduit.GetSchemaResponse{Schema: outSchema}, nil
}
