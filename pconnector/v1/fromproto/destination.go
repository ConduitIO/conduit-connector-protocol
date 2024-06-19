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

func DestinationConfigureRequest(in *connectorv1.Destination_Configure_Request) pconnector.DestinationConfigureRequest {
	return pconnector.DestinationConfigureRequest{
		Config: in.Config,
	}
}

func DestinationStartRequest(_ *connectorv1.Destination_Start_Request) pconnector.DestinationOpenRequest {
	return pconnector.DestinationOpenRequest{}
}

func DestinationRunRequest(in *connectorv1.Destination_Run_Request) (pconnector.DestinationRunRequest, error) {
	var rec opencdc.Record
	err := rec.FromProto(in.Record)
	if err != nil {
		return pconnector.DestinationRunRequest{}, err
	}
	return pconnector.DestinationRunRequest{
		Records: []opencdc.Record{rec},
	}, nil
}

func DestinationStopRequest(in *connectorv1.Destination_Stop_Request) pconnector.DestinationStopRequest {
	return pconnector.DestinationStopRequest{
		LastPosition: in.LastPosition,
	}
}

func DestinationTeardownRequest(_ *connectorv1.Destination_Teardown_Request) pconnector.DestinationTeardownRequest {
	return pconnector.DestinationTeardownRequest{}
}

func DestinationLifecycleOnCreatedRequest(in *connectorv1.Destination_Lifecycle_OnCreated_Request) pconnector.DestinationLifecycleOnCreatedRequest {
	return pconnector.DestinationLifecycleOnCreatedRequest{
		Config: in.Config,
	}
}

func DestinationLifecycleOnUpdatedRequest(in *connectorv1.Destination_Lifecycle_OnUpdated_Request) pconnector.DestinationLifecycleOnUpdatedRequest {
	return pconnector.DestinationLifecycleOnUpdatedRequest{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}

func DestinationLifecycleOnDeletedRequest(in *connectorv1.Destination_Lifecycle_OnDeleted_Request) pconnector.DestinationLifecycleOnDeletedRequest {
	return pconnector.DestinationLifecycleOnDeletedRequest{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func DestinationConfigureResponse(_ *connectorv1.Destination_Configure_Response) pconnector.DestinationConfigureResponse {
	return pconnector.DestinationConfigureResponse{}
}

func DestinationStartResponse(_ *connectorv1.Destination_Start_Response) pconnector.DestinationOpenResponse {
	return pconnector.DestinationOpenResponse{}
}

func DestinationRunResponse(in *connectorv1.Destination_Run_Response) pconnector.DestinationRunResponse {
	return pconnector.DestinationRunResponse{
		Acks: []pconnector.DestinationRunResponseAck{{
			Position: in.AckPosition,
			Error:    in.Error,
		}},
	}
}

func DestinationStopResponse(_ *connectorv1.Destination_Stop_Response) pconnector.DestinationStopResponse {
	return pconnector.DestinationStopResponse{}
}

func DestinationTeardownResponse(_ *connectorv1.Destination_Teardown_Response) pconnector.DestinationTeardownResponse {
	return pconnector.DestinationTeardownResponse{}
}

func DestinationLifecycleOnCreatedResponse(_ *connectorv1.Destination_Lifecycle_OnCreated_Response) pconnector.DestinationLifecycleOnCreatedResponse {
	return pconnector.DestinationLifecycleOnCreatedResponse{}
}

func DestinationLifecycleOnUpdatedResponse(_ *connectorv1.Destination_Lifecycle_OnUpdated_Response) pconnector.DestinationLifecycleOnUpdatedResponse {
	return pconnector.DestinationLifecycleOnUpdatedResponse{}
}

func DestinationLifecycleOnDeletedResponse(_ *connectorv1.Destination_Lifecycle_OnDeleted_Response) pconnector.DestinationLifecycleOnDeletedResponse {
	return pconnector.DestinationLifecycleOnDeletedResponse{}
}
