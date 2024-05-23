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
	"github.com/conduitio/conduit-commons/opencdc"
	"github.com/conduitio/conduit-connector-protocol/cplugin"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
)

// -- Request Conversions -----------------------------------------------------

func SourceConfigureRequest(in *connectorv2.Source_Configure_Request) cplugin.SourceConfigureRequest {
	return cplugin.SourceConfigureRequest{
		Config: in.Config,
	}
}

func SourceStartRequest(in *connectorv2.Source_Start_Request) cplugin.SourceStartRequest {
	return cplugin.SourceStartRequest{
		Position: in.Position,
	}
}

func SourceRunRequest(in *connectorv2.Source_Run_Request) cplugin.SourceRunRequest {
	ackPositions := make([]opencdc.Position, len(in.AckPositions))
	for i, pos := range in.AckPositions {
		ackPositions[i] = pos
	}

	return cplugin.SourceRunRequest{
		AckPositions: ackPositions,
	}
}

func SourceStopRequest(_ *connectorv2.Source_Stop_Request) cplugin.SourceStopRequest {
	return cplugin.SourceStopRequest{}
}

func SourceTeardownRequest(_ *connectorv2.Source_Teardown_Request) cplugin.SourceTeardownRequest {
	return cplugin.SourceTeardownRequest{}
}

func SourceLifecycleOnCreatedRequest(in *connectorv2.Source_Lifecycle_OnCreated_Request) cplugin.SourceLifecycleOnCreatedRequest {
	return cplugin.SourceLifecycleOnCreatedRequest{
		Config: in.Config,
	}
}
func SourceLifecycleOnUpdatedRequest(in *connectorv2.Source_Lifecycle_OnUpdated_Request) cplugin.SourceLifecycleOnUpdatedRequest {
	return cplugin.SourceLifecycleOnUpdatedRequest{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}
func SourceLifecycleOnDeletedRequest(in *connectorv2.Source_Lifecycle_OnDeleted_Request) cplugin.SourceLifecycleOnDeletedRequest {
	return cplugin.SourceLifecycleOnDeletedRequest{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func SourceConfigureResponse(_ *connectorv2.Source_Configure_Response) cplugin.SourceConfigureResponse {
	return cplugin.SourceConfigureResponse{}
}

func SourceStartResponse(_ *connectorv2.Source_Start_Response) cplugin.SourceStartResponse {
	return cplugin.SourceStartResponse{}
}

func SourceRunResponse(in *connectorv2.Source_Run_Response) (cplugin.SourceRunResponse, error) {
	records := make([]opencdc.Record, len(in.Records))
	for i, rec := range in.Records {
		err := records[i].FromProto(rec)
		if err != nil {
			return cplugin.SourceRunResponse{}, err
		}
	}
	return cplugin.SourceRunResponse{
		Records: records,
	}, nil
}

func SourceStopResponse(in *connectorv2.Source_Stop_Response) cplugin.SourceStopResponse {
	return cplugin.SourceStopResponse{
		LastPosition: in.LastPosition,
	}
}

func SourceTeardownResponse(_ *connectorv2.Source_Teardown_Response) cplugin.SourceTeardownResponse {
	return cplugin.SourceTeardownResponse{}
}

func SourceLifecycleOnCreatedResponse(_ *connectorv2.Source_Lifecycle_OnCreated_Response) cplugin.SourceLifecycleOnCreatedResponse {
	return cplugin.SourceLifecycleOnCreatedResponse{}
}
func SourceLifecycleOnUpdatedResponse(_ *connectorv2.Source_Lifecycle_OnUpdated_Response) cplugin.SourceLifecycleOnUpdatedResponse {
	return cplugin.SourceLifecycleOnUpdatedResponse{}
}
func SourceLifecycleOnDeletedResponse(_ *connectorv2.Source_Lifecycle_OnDeleted_Response) cplugin.SourceLifecycleOnDeletedResponse {
	return cplugin.SourceLifecycleOnDeletedResponse{}
}
