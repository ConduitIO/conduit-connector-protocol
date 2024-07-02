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

package fromproto

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/conduitio/conduit-commons/config"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
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

func SpecifierSpecifyRequest(_ *connectorv1.Specifier_Specify_Request) pconnector.SpecifierSpecifyRequest {
	return pconnector.SpecifierSpecifyRequest{}
}

// -- Response Conversions ----------------------------------------------------

func SpecifierSpecifyResponse(in *connectorv1.Specifier_Specify_Response) (pconnector.SpecifierSpecifyResponse, error) {
	sourceParams, err := Parameters(in.SourceParams)
	if err != nil {
		return pconnector.SpecifierSpecifyResponse{}, err
	}

	destinationParams, err := Parameters(in.DestinationParams)
	if err != nil {
		return pconnector.SpecifierSpecifyResponse{}, err
	}

	return pconnector.SpecifierSpecifyResponse{
		Specification: pconnector.Specification{
			Name:              in.Name,
			Summary:           in.Summary,
			Description:       in.Description,
			Version:           in.Version,
			Author:            in.Author,
			DestinationParams: destinationParams,
			SourceParams:      sourceParams,
		},
	}, nil
}

func Parameters(proto map[string]*connectorv1.Specifier_Parameter) (config.Parameters, error) {
	if proto == nil {
		return nil, nil //nolint:nilnil // This is the expected behavior.
	}

	out := make(config.Parameters, len(proto))
	for k, v := range proto {
		param, err := Parameter(v)
		if err != nil {
			return nil, fmt.Errorf("error converting parameter: %w", err)
		}
		out[k] = param
	}
	return out, nil
}

func Parameter(proto *connectorv1.Specifier_Parameter) (config.Parameter, error) {
	if proto == nil {
		return config.Parameter{}, nil
	}

	validations, err := Validations(proto.Validations)
	if err != nil {
		return config.Parameter{}, err
	}

	if proto.Required { //nolint:staticcheck // backwards compatibility
		// needed for backward compatibility, proto.Required is converted to a
		// validation of type config.ValidationTypeRequired, making sure not to
		// duplicate the required validation
		isRequired := false
		for _, v := range validations {
			if v.Type() == config.ValidationTypeRequired {
				isRequired = true
				break
			}
		}
		if !isRequired {
			validations = append(validations, config.ValidationRequired{})
		}
	}

	return config.Parameter{
		Default:     proto.Default,
		Description: proto.Description,
		Type:        config.ParameterType(proto.Type),
		Validations: validations,
	}, nil
}

func Validations(proto []*connectorv1.Specifier_Parameter_Validation) ([]config.Validation, error) {
	if proto == nil {
		return nil, nil
	}

	validations := make([]config.Validation, len(proto))
	for i, v := range proto {
		var err error
		validations[i], err = Validation(v)
		if err != nil {
			return nil, fmt.Errorf("error converting validation: %w", err)
		}
	}
	return validations, nil
}

func Validation(proto *connectorv1.Specifier_Parameter_Validation) (config.Validation, error) {
	if proto == nil {
		return nil, nil //nolint:nilnil // This is the expected behavior.
	}

	switch proto.Type {
	case connectorv1.Specifier_Parameter_Validation_TYPE_REQUIRED:
		return config.ValidationRequired{}, nil
	case connectorv1.Specifier_Parameter_Validation_TYPE_GREATER_THAN:
		v, err := strconv.ParseFloat(proto.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing greater than value: %w", err)
		}
		return config.ValidationGreaterThan{V: v}, nil
	case connectorv1.Specifier_Parameter_Validation_TYPE_LESS_THAN:
		v, err := strconv.ParseFloat(proto.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing less than value: %w", err)
		}
		return config.ValidationLessThan{V: v}, nil
	case connectorv1.Specifier_Parameter_Validation_TYPE_INCLUSION:
		return config.ValidationInclusion{List: strings.Split(proto.Value, ",")}, nil
	case connectorv1.Specifier_Parameter_Validation_TYPE_EXCLUSION:
		return config.ValidationExclusion{List: strings.Split(proto.Value, ",")}, nil
	case connectorv1.Specifier_Parameter_Validation_TYPE_REGEX:
		regex, err := regexp.Compile(proto.Value)
		if err != nil {
			return nil, fmt.Errorf("error compiling regex: %w", err)
		}
		return config.ValidationRegex{Regex: regex}, nil
	case connectorv1.Specifier_Parameter_Validation_TYPE_UNSPECIFIED:
		fallthrough
	default:
		return nil, fmt.Errorf("%v: %w", proto.Type, config.ErrInvalidValidationType)
	}
}
