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
	"fmt"

	"github.com/conduitio/conduit-commons/schema"
	conduitv1 "github.com/conduitio/conduit-connector-protocol/proto/conduit/v1"
)

func SchemaInstance(req *conduitv1.CreateRequest) (schema.Instance, error) {
	typ, err := SchemaType(req.Type)
	if err != nil {
		return schema.Instance{}, fmt.Errorf("invalid schema type: %w", err)
	}

	return schema.Instance{
		Name:  req.Name,
		Type:  typ,
		Bytes: req.Bytes,
	}, nil
}

func SchemaType(typ conduitv1.Schema_Type) (schema.Type, error) {
	switch typ {
	case conduitv1.Schema_TYPE_AVRO:
		return schema.TypeAvro, nil
	default:
		return 0, fmt.Errorf("unsupported %q", typ)
	}
}
