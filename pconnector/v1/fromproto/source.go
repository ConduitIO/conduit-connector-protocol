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
	"github.com/conduitio/conduit-commons/opencdc"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
)

// -- Request Conversions -----------------------------------------------------

func SourceConfigureRequest(in *connectorv1.Source_Configure_Request) pconnector.SourceConfigureRequest {
	return pconnector.SourceConfigureRequest{
		Config: in.Config,
	}
}

func SourceStartRequest(in *connectorv1.Source_Start_Request) pconnector.SourceOpenRequest {
	return pconnector.SourceOpenRequest{
		Position: in.Position,
	}
}

func SourceRunRequest(in *connectorv1.Source_Run_Request) pconnector.SourceRunRequest {
	return pconnector.SourceRunRequest{
		AckPositions: []opencdc.Position{in.AckPosition},
	}
}

func SourceStopRequest(_ *connectorv1.Source_Stop_Request) pconnector.SourceStopRequest {
	return pconnector.SourceStopRequest{}
}

func SourceTeardownRequest(_ *connectorv1.Source_Teardown_Request) pconnector.SourceTeardownRequest {
	return pconnector.SourceTeardownRequest{}
}

func SourceLifecycleOnCreatedRequest(in *connectorv1.Source_Lifecycle_OnCreated_Request) pconnector.SourceLifecycleOnCreatedRequest {
	return pconnector.SourceLifecycleOnCreatedRequest{
		Config: in.Config,
	}
}

func SourceLifecycleOnUpdatedRequest(in *connectorv1.Source_Lifecycle_OnUpdated_Request) pconnector.SourceLifecycleOnUpdatedRequest {
	return pconnector.SourceLifecycleOnUpdatedRequest{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}

func SourceLifecycleOnDeletedRequest(in *connectorv1.Source_Lifecycle_OnDeleted_Request) pconnector.SourceLifecycleOnDeletedRequest {
	return pconnector.SourceLifecycleOnDeletedRequest{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func SourceConfigureResponse(_ *connectorv1.Source_Configure_Response) pconnector.SourceConfigureResponse {
	return pconnector.SourceConfigureResponse{}
}

func SourceStartResponse(_ *connectorv1.Source_Start_Response) pconnector.SourceOpenResponse {
	return pconnector.SourceOpenResponse{}
}

func SourceRunResponse(in *connectorv1.Source_Run_Response) (pconnector.SourceRunResponse, error) {
	var rec opencdc.Record
	err := rec.FromProto(in.Record)
	if err != nil {
		return pconnector.SourceRunResponse{}, err
	}
	return pconnector.SourceRunResponse{
		Records: []opencdc.Record{rec},
	}, nil
}

func SourceStopResponse(in *connectorv1.Source_Stop_Response) pconnector.SourceStopResponse {
	return pconnector.SourceStopResponse{
		LastPosition: in.LastPosition,
	}
}

func SourceTeardownResponse(_ *connectorv1.Source_Teardown_Response) pconnector.SourceTeardownResponse {
	return pconnector.SourceTeardownResponse{}
}

func SourceLifecycleOnCreatedResponse(_ *connectorv1.Source_Lifecycle_OnCreated_Response) pconnector.SourceLifecycleOnCreatedResponse {
	return pconnector.SourceLifecycleOnCreatedResponse{}
}

func SourceLifecycleOnUpdatedResponse(_ *connectorv1.Source_Lifecycle_OnUpdated_Response) pconnector.SourceLifecycleOnUpdatedResponse {
	return pconnector.SourceLifecycleOnUpdatedResponse{}
}

func SourceLifecycleOnDeletedResponse(_ *connectorv1.Source_Lifecycle_OnDeleted_Response) pconnector.SourceLifecycleOnDeletedResponse {
	return pconnector.SourceLifecycleOnDeletedResponse{}
}
