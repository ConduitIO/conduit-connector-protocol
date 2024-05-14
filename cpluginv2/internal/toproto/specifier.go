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
	configv1 "github.com/conduitio/conduit-commons/proto/config/v1"
	"github.com/conduitio/conduit-connector-protocol/cpluginv2"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
)

func SpecifierSpecifyResponse(in cpluginv2.SpecifierSpecifyResponse) (*connectorv2.Specifier_Specify_Response, error) {
	sourceParams := make(map[string]*configv1.Parameter, len(in.SourceParams))
	in.SourceParams.ToProto(sourceParams)

	destinationParams := make(map[string]*configv1.Parameter, len(in.DestinationParams))
	in.DestinationParams.ToProto(destinationParams)

	out := connectorv2.Specifier_Specify_Response{
		Name:              in.Name,
		Summary:           in.Summary,
		Description:       in.Description,
		Version:           in.Version,
		Author:            in.Author,
		DestinationParams: destinationParams,
		SourceParams:      sourceParams,
	}
	return &out, nil
}
