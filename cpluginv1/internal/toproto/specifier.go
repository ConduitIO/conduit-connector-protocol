// Copyright © 2021 Meroxa Inc
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
	"fmt"

	"github.com/conduitio/conduit-plugin/cpluginv1"
	"github.com/conduitio/conduit-plugin/cpluginv1/internal/cproto"
)

func SpecifierSpecifyRequest(in cpluginv1.SpecifierSpecifyRequest) (*cproto.Specifier_Specify_Request, error) {
	return &cproto.Specifier_Specify_Request{}, nil
}

func SpecifierSpecifyResponse(in cpluginv1.SpecifierSpecifyResponse) (*cproto.Specifier_Specify_Response, error) {
	specMap := func(in map[string]cpluginv1.SpecifierParameter) (map[string]*cproto.Specifier_Parameter, error) {
		out := make(map[string]*cproto.Specifier_Parameter, len(in))
		var err error
		for k, v := range in {
			out[k], err = SpecifierParameter(v)
			if err != nil {
				return nil, fmt.Errorf("error converting SpecifierParameter %q: %w", k, err)
			}
		}
		return out, nil
	}

	sourceSpec, err := specMap(in.SourceSpec)
	if err != nil {
		return nil, fmt.Errorf("error converting SourceSpec: %w", err)
	}

	destinationSpec, err := specMap(in.DestinationSpec)
	if err != nil {
		return nil, fmt.Errorf("error converting DestinationSpec: %w", err)
	}

	out := cproto.Specifier_Specify_Response{
		Summary:         in.Summary,
		Description:     in.Description,
		Version:         in.Version,
		Author:          in.Author,
		DestinationSpec: destinationSpec,
		SourceSpec:      sourceSpec,
	}
	return &out, nil
}

func SpecifierParameter(in cpluginv1.SpecifierParameter) (*cproto.Specifier_Parameter, error) {
	out := cproto.Specifier_Parameter{
		Default:     in.Default,
		Required:    in.Required,
		Description: in.Description,
	}
	return &out, nil
}
