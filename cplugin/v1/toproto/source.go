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
	opencdcv1 "github.com/conduitio/conduit-commons/proto/opencdc/v1"
	"github.com/conduitio/conduit-connector-protocol/cplugin"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
)

// -- Request Conversions -----------------------------------------------------

// TODO

// -- Response Conversions ----------------------------------------------------

func SourceConfigureResponse(_ cplugin.SourceConfigureResponse) *connectorv1.Source_Configure_Response {
	return &connectorv1.Source_Configure_Response{}
}

func SourceStartResponse(_ cplugin.SourceStartResponse) *connectorv1.Source_Start_Response {
	return &connectorv1.Source_Start_Response{}
}

func SourceRunResponse(in cplugin.SourceRunResponse) (*connectorv1.Source_Run_Response, error) {
	out := connectorv1.Source_Run_Response{
		Record: &opencdcv1.Record{},
	}
	err := in.Record.ToProto(out.Record)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func SourceStopResponse(in cplugin.SourceStopResponse) *connectorv1.Source_Stop_Response {
	return &connectorv1.Source_Stop_Response{
		LastPosition: in.LastPosition,
	}
}

func SourceTeardownResponse(_ cplugin.SourceTeardownResponse) *connectorv1.Source_Teardown_Response {
	return &connectorv1.Source_Teardown_Response{}
}

func SourceLifecycleOnCreatedResponse(_ cplugin.SourceLifecycleOnCreatedResponse) *connectorv1.Source_Lifecycle_OnCreated_Response {
	return &connectorv1.Source_Lifecycle_OnCreated_Response{}
}
func SourceLifecycleOnUpdatedResponse(_ cplugin.SourceLifecycleOnUpdatedResponse) *connectorv1.Source_Lifecycle_OnUpdated_Response {
	return &connectorv1.Source_Lifecycle_OnUpdated_Response{}
}
func SourceLifecycleOnDeletedResponse(_ cplugin.SourceLifecycleOnDeletedResponse) *connectorv1.Source_Lifecycle_OnDeleted_Response {
	return &connectorv1.Source_Lifecycle_OnDeleted_Response{}
}
