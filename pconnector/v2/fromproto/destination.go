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
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
)

// -- Request Conversions -----------------------------------------------------

func DestinationConfigureRequest(in *connectorv2.Destination_Configure_Request) pconnector.DestinationConfigureRequest {
	return pconnector.DestinationConfigureRequest{
		Config: in.Config,
	}
}

func DestinationOpenRequest(_ *connectorv2.Destination_Open_Request) pconnector.DestinationOpenRequest {
	return pconnector.DestinationOpenRequest{}
}

func DestinationRunRequest(in *connectorv2.Destination_Run_Request) (pconnector.DestinationRunRequest, error) {
	records := make([]opencdc.Record, len(in.Records))
	for i, rec := range in.Records {
		err := records[i].FromProto(rec)
		if err != nil {
			return pconnector.DestinationRunRequest{}, err
		}
	}
	return pconnector.DestinationRunRequest{
		Records: records,
	}, nil
}

func DestinationStopRequest(in *connectorv2.Destination_Stop_Request) pconnector.DestinationStopRequest {
	return pconnector.DestinationStopRequest{
		LastPosition: in.LastPosition,
	}
}

func DestinationTeardownRequest(_ *connectorv2.Destination_Teardown_Request) pconnector.DestinationTeardownRequest {
	return pconnector.DestinationTeardownRequest{}
}

func DestinationLifecycleOnCreatedRequest(in *connectorv2.Destination_Lifecycle_OnCreated_Request) pconnector.DestinationLifecycleOnCreatedRequest {
	return pconnector.DestinationLifecycleOnCreatedRequest{
		Config: in.Config,
	}
}

func DestinationLifecycleOnUpdatedRequest(in *connectorv2.Destination_Lifecycle_OnUpdated_Request) pconnector.DestinationLifecycleOnUpdatedRequest {
	return pconnector.DestinationLifecycleOnUpdatedRequest{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}

func DestinationLifecycleOnDeletedRequest(in *connectorv2.Destination_Lifecycle_OnDeleted_Request) pconnector.DestinationLifecycleOnDeletedRequest {
	return pconnector.DestinationLifecycleOnDeletedRequest{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func DestinationConfigureResponse(_ *connectorv2.Destination_Configure_Response) pconnector.DestinationConfigureResponse {
	return pconnector.DestinationConfigureResponse{}
}

func DestinationOpenResponse(_ *connectorv2.Destination_Open_Response) pconnector.DestinationOpenResponse {
	return pconnector.DestinationOpenResponse{}
}

func DestinationRunResponse(in *connectorv2.Destination_Run_Response) pconnector.DestinationRunResponse {
	acks := make([]pconnector.DestinationRunResponseAck, len(in.Acks))
	for i, ack := range in.Acks {
		acks[i] = DestinationRunResponseAck(ack)
	}
	return pconnector.DestinationRunResponse{
		Acks: acks,
	}
}

func DestinationRunResponseAck(in *connectorv2.Destination_Run_Response_Ack) pconnector.DestinationRunResponseAck {
	return pconnector.DestinationRunResponseAck{
		Position: in.Position,
		Error:    in.Error,
	}
}

func DestinationStopResponse(_ *connectorv2.Destination_Stop_Response) pconnector.DestinationStopResponse {
	return pconnector.DestinationStopResponse{}
}

func DestinationTeardownResponse(_ *connectorv2.Destination_Teardown_Response) pconnector.DestinationTeardownResponse {
	return pconnector.DestinationTeardownResponse{}
}

func DestinationLifecycleOnCreatedResponse(_ *connectorv2.Destination_Lifecycle_OnCreated_Response) pconnector.DestinationLifecycleOnCreatedResponse {
	return pconnector.DestinationLifecycleOnCreatedResponse{}
}

func DestinationLifecycleOnUpdatedResponse(_ *connectorv2.Destination_Lifecycle_OnUpdated_Response) pconnector.DestinationLifecycleOnUpdatedResponse {
	return pconnector.DestinationLifecycleOnUpdatedResponse{}
}

func DestinationLifecycleOnDeletedResponse(_ *connectorv2.Destination_Lifecycle_OnDeleted_Response) pconnector.DestinationLifecycleOnDeletedResponse {
	return pconnector.DestinationLifecycleOnDeletedResponse{}
}
