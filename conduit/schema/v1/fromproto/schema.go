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
	cschema "github.com/conduitio/conduit-commons/schema"
	"github.com/conduitio/conduit-connector-protocol/conduit/schema"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	var cTypes [1]struct{}
	_ = cTypes[int(cschema.TypeAvro)-int(conduitv1.Schema_TYPE_AVRO)]
}

func CreateRequest(req *conduitv1.CreateSchemaRequest) schema.CreateRequest {
	return schema.CreateRequest{
		Name:  req.Name,
		Bytes: req.Bytes,
	}
}

func CreateResponse(resp *conduitv1.CreateSchemaResponse) schema.CreateResponse {
	return schema.CreateResponse{
		Instance: cschema.Instance{
			ID:      resp.Schema.Id,
			Name:    resp.Schema.Name,
			Version: resp.Schema.Version,
			Type:    cschema.Type(resp.Schema.Type),
			Bytes:   resp.Schema.Bytes,
		},
	}
}
