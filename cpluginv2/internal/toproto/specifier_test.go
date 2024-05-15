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
	configv1 "github.com/conduitio/conduit-commons/proto/config/v1"
	"github.com/conduitio/conduit-connector-protocol/cpluginv2"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/matryer/is"
)

func TestSpecifierSpecifyResponse(t *testing.T) {
	have := cpluginv2.SpecifierSpecifyResponse{
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
	}

	want := &connectorv2.Specifier_Specify_Response{
		Name:        "TestPlugin",
		Summary:     "This is a test plugin",
		Description: "A plugin for testing purposes",
		Version:     "v1.0.0",
		Author:      "Test Author",
		DestinationParams: map[string]*configv1.Parameter{
			"param1": {
				Default:     "value1",
				Description: "Description of param1",
				Type:        configv1.Parameter_TYPE_STRING,
				Validations: []*configv1.Validation{
					{Type: configv1.Validation_TYPE_INCLUSION, Value: "value1,value2"},
				},
			},
		},
		SourceParams: map[string]*configv1.Parameter{
			"param2": {
				Default:     "20",
				Description: "Description of param2",
				Type:        configv1.Parameter_TYPE_INT,
				Validations: []*configv1.Validation{
					{Type: configv1.Validation_TYPE_EXCLUSION, Value: "10,11,12"},
				},
			},
		},
	}

	is := is.New(t)
	got := SpecifierSpecifyResponse(have)
	is.Equal(
		"",
		cmp.Diff(want, got,
			cmpopts.IgnoreUnexported(connectorv2.Specifier_Specify_Response{}),
			cmpopts.IgnoreUnexported(configv1.Parameter{}),
			cmpopts.IgnoreUnexported(configv1.Validation{}),
		),
	)
}
