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

package fromproto

import (
	"fmt"

	"github.com/conduitio/conduit-plugin/cpluginv1"
	cproto "github.com/conduitio/conduit-plugin/proto/gen/go/conduitio/cplugin/v1"
)

func SpecifierSpecifyRequest(in *cproto.Specifier_Specify_Request) (cpluginv1.SpecifierSpecifyRequest, error) {
	return cpluginv1.SpecifierSpecifyRequest{}, nil
}

func SpecifierSpecifyResponse(in *cproto.Specifier_Specify_Response) (cpluginv1.SpecifierSpecifyResponse, error) {
	specMap := func(in map[string]*cproto.Specifier_Parameter) (map[string]cpluginv1.SpecifierParameter, error) {
		out := make(map[string]cpluginv1.SpecifierParameter, len(in))
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
		return cpluginv1.SpecifierSpecifyResponse{}, fmt.Errorf("error converting SourceSpec: %w", err)
	}

	destinationSpec, err := specMap(in.DestinationSpec)
	if err != nil {
		return cpluginv1.SpecifierSpecifyResponse{}, fmt.Errorf("error converting DestinationSpec: %w", err)
	}

	out := cpluginv1.SpecifierSpecifyResponse{
		Summary:         in.Summary,
		Description:     in.Description,
		Version:         in.Version,
		Author:          in.Author,
		DestinationSpec: destinationSpec,
		SourceSpec:      sourceSpec,
	}
	return out, nil
}

func SpecifierParameter(in *cproto.Specifier_Parameter) (cpluginv1.SpecifierParameter, error) {
	out := cpluginv1.SpecifierParameter{
		Default:     in.Default,
		Required:    in.Required,
		Description: in.Description,
	}
	return out, nil
}