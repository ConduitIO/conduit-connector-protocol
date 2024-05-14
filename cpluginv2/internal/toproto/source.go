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
	opencdcv1 "github.com/conduitio/conduit-commons/proto/opencdc/v1"
	"github.com/conduitio/conduit-connector-protocol/cpluginv2"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
)

func SourceConfigureResponse(_ cpluginv2.SourceConfigureResponse) (*connectorv2.Source_Configure_Response, error) {
	return &connectorv2.Source_Configure_Response{}, nil
}

func SourceStartResponse(_ cpluginv2.SourceStartResponse) (*connectorv2.Source_Start_Response, error) {
	return &connectorv2.Source_Start_Response{}, nil
}

func SourceRunResponse(in cpluginv2.SourceRunResponse) (*connectorv2.Source_Run_Response, error) {
	out := connectorv2.Source_Run_Response{
		Record: &opencdcv1.Record{},
	}
	err := in.Record.ToProto(out.Record)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func SourceStopResponse(in cpluginv2.SourceStopResponse) (*connectorv2.Source_Stop_Response, error) {
	out := connectorv2.Source_Stop_Response{
		LastPosition: in.LastPosition,
	}
	return &out, nil
}

func SourceTeardownResponse(_ cpluginv2.SourceTeardownResponse) (*connectorv2.Source_Teardown_Response, error) {
	return &connectorv2.Source_Teardown_Response{}, nil
}

func SourceLifecycleOnCreatedResponse(_ cpluginv2.SourceLifecycleOnCreatedResponse) (*connectorv2.Source_Lifecycle_OnCreated_Response, error) {
	return &connectorv2.Source_Lifecycle_OnCreated_Response{}, nil
}
func SourceLifecycleOnUpdatedResponse(_ cpluginv2.SourceLifecycleOnUpdatedResponse) (*connectorv2.Source_Lifecycle_OnUpdated_Response, error) {
	return &connectorv2.Source_Lifecycle_OnUpdated_Response{}, nil
}
func SourceLifecycleOnDeletedResponse(_ cpluginv2.SourceLifecycleOnDeletedResponse) (*connectorv2.Source_Lifecycle_OnDeleted_Response, error) {
	return &connectorv2.Source_Lifecycle_OnDeleted_Response{}, nil
}
