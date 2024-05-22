// Copyright © 2022 Meroxa, Inc.
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
	"github.com/conduitio/conduit-commons/config"
	"github.com/conduitio/conduit-connector-protocol/cplugin"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	var vTypes [1]struct{}
	_ = vTypes[int(config.ValidationTypeRequired)-int(connectorv1.Specifier_Parameter_Validation_TYPE_REQUIRED)]
	_ = vTypes[int(config.ValidationTypeLessThan)-int(connectorv1.Specifier_Parameter_Validation_TYPE_LESS_THAN)]
	_ = vTypes[int(config.ValidationTypeGreaterThan)-int(connectorv1.Specifier_Parameter_Validation_TYPE_GREATER_THAN)]
	_ = vTypes[int(config.ValidationTypeInclusion)-int(connectorv1.Specifier_Parameter_Validation_TYPE_INCLUSION)]
	_ = vTypes[int(config.ValidationTypeExclusion)-int(connectorv1.Specifier_Parameter_Validation_TYPE_EXCLUSION)]
	_ = vTypes[int(config.ValidationTypeRegex)-int(connectorv1.Specifier_Parameter_Validation_TYPE_REGEX)]
	// parameter types
	_ = vTypes[int(config.ParameterTypeString)-int(connectorv1.Specifier_Parameter_TYPE_STRING)]
	_ = vTypes[int(config.ParameterTypeInt)-int(connectorv1.Specifier_Parameter_TYPE_INT)]
	_ = vTypes[int(config.ParameterTypeFloat)-int(connectorv1.Specifier_Parameter_TYPE_FLOAT)]
	_ = vTypes[int(config.ParameterTypeBool)-int(connectorv1.Specifier_Parameter_TYPE_BOOL)]
	_ = vTypes[int(config.ParameterTypeFile)-int(connectorv1.Specifier_Parameter_TYPE_FILE)]
	_ = vTypes[int(config.ParameterTypeDuration)-int(connectorv1.Specifier_Parameter_TYPE_DURATION)]
}

// -- Request Conversions -----------------------------------------------------

func SpecifierSpecifyRequest(_ cplugin.SpecifierSpecifyRequest) *connectorv1.Specifier_Specify_Request {
	return &connectorv1.Specifier_Specify_Request{}
}

// -- Response Conversions ----------------------------------------------------

func SpecifierSpecifyResponse(in cplugin.SpecifierSpecifyResponse) *connectorv1.Specifier_Specify_Response {
	out := connectorv1.Specifier_Specify_Response{
		Name:              in.Name,
		Summary:           in.Summary,
		Description:       in.Description,
		Version:           in.Version,
		Author:            in.Author,
		DestinationParams: SpecifierParameters(in.DestinationParams),
		SourceParams:      SpecifierParameters(in.SourceParams),
	}
	return &out
}

func SpecifierParameters(in config.Parameters) map[string]*connectorv1.Specifier_Parameter {
	out := make(map[string]*connectorv1.Specifier_Parameter, len(in))
	for k, v := range in {
		out[k] = SpecifierParameter(v)
	}
	return out
}

func SpecifierParameter(in config.Parameter) *connectorv1.Specifier_Parameter {
	return &connectorv1.Specifier_Parameter{
		Default:     in.Default,
		Description: in.Description,
		Type:        connectorv1.Specifier_Parameter_Type(in.Type),
		Validations: SpecifierParameterValidations(in.Validations),
	}
}

func SpecifierParameterValidations(in []config.Validation) []*connectorv1.Specifier_Parameter_Validation {
	out := make([]*connectorv1.Specifier_Parameter_Validation, len(in))
	for i, v := range in {
		out[i] = &connectorv1.Specifier_Parameter_Validation{
			Type:  connectorv1.Specifier_Parameter_Validation_Type(v.Type()),
			Value: v.Value(),
		}
	}
	return out
}
