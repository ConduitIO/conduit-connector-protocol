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
	"testing"

	"github.com/conduitio/conduit-commons/config"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/matryer/is"
)

func TestSpecifierSpecifyResponse(t *testing.T) {
	have := pconnector.SpecifierSpecifyResponse{
		Specification: pconnector.Specification{
			Name:        "TestPlugin",
			Summary:     "This is a test plugin",
			Description: "A plugin for testing purposes",
			Version:     "v1.0.0",
			Author:      "Test Author",
			DestinationParams: config.Parameters{
				"param1": config.Parameter{
					Default:     "value1",
					Description: "Description of param1",
					Type:        config.ParameterTypeString,
					Validations: []config.Validation{
						config.ValidationInclusion{List: []string{"value1", "value2"}},
					},
				},
			},
			SourceParams: config.Parameters{
				"param2": config.Parameter{
					Default:     "20",
					Description: "Description of param2",
					Type:        config.ParameterTypeInt,
					Validations: []config.Validation{
						config.ValidationExclusion{List: []string{"10", "11", "12"}},
					},
				},
			},
		},
	}

	want := &connectorv1.Specifier_Specify_Response{
		Name:        "TestPlugin",
		Summary:     "This is a test plugin",
		Description: "A plugin for testing purposes",
		Version:     "v1.0.0",
		Author:      "Test Author",
		DestinationParams: map[string]*connectorv1.Specifier_Parameter{
			"param1": {
				Default:     "value1",
				Description: "Description of param1",
				Type:        connectorv1.Specifier_Parameter_TYPE_STRING,
				Validations: []*connectorv1.Specifier_Parameter_Validation{
					{Type: connectorv1.Specifier_Parameter_Validation_TYPE_INCLUSION, Value: "value1,value2"},
				},
			},
		},
		SourceParams: map[string]*connectorv1.Specifier_Parameter{
			"param2": {
				Default:     "20",
				Description: "Description of param2",
				Type:        connectorv1.Specifier_Parameter_TYPE_INT,
				Validations: []*connectorv1.Specifier_Parameter_Validation{
					{Type: connectorv1.Specifier_Parameter_Validation_TYPE_EXCLUSION, Value: "10,11,12"},
				},
			},
		},
	}

	is := is.New(t)
	got := SpecifierSpecifyResponse(have)
	is.Equal(
		"",
		cmp.Diff(want, got,
			cmpopts.IgnoreUnexported(connectorv1.Specifier_Specify_Response{}),
			cmpopts.IgnoreUnexported(connectorv1.Specifier_Parameter{}),
			cmpopts.IgnoreUnexported(connectorv1.Specifier_Parameter_Validation{}),
		),
	)
}
