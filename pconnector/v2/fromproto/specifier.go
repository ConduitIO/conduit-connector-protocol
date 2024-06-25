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
	"github.com/conduitio/conduit-commons/config"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
)

// -- Request Conversions -----------------------------------------------------

func SpecifierSpecifyRequest(_ *connectorv2.Specifier_Specify_Request) pconnector.SpecifierSpecifyRequest {
	return pconnector.SpecifierSpecifyRequest{}
}

// -- Response Conversions ----------------------------------------------------

func SpecifierSpecifyResponse(in *connectorv2.Specifier_Specify_Response) (pconnector.SpecifierSpecifyResponse, error) {
	spec, err := Specification(in.Specification)
	if err != nil {
		return pconnector.SpecifierSpecifyResponse{}, err
	}

	return pconnector.SpecifierSpecifyResponse{
		Specification: spec,
	}, nil
}

func Specification(in *connectorv2.Specification) (pconnector.Specification, error) {
	sourceParams := make(config.Parameters)
	err := sourceParams.FromProto(in.SourceParams)
	if err != nil {
		return pconnector.Specification{}, err
	}

	destinationParams := make(config.Parameters)
	err = destinationParams.FromProto(in.DestinationParams)
	if err != nil {
		return pconnector.Specification{}, err
	}

	return pconnector.Specification{
		Name:              in.Name,
		Summary:           in.Summary,
		Description:       in.Description,
		Version:           in.Version,
		Author:            in.Author,
		DestinationParams: destinationParams,
		SourceParams:      sourceParams,
	}, nil
}
