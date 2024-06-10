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

func DestinationConfigureRequest(in *connectorv2.Destination_Configure_Request) cplugin.DestinationConfigureRequest {
	return cplugin.DestinationConfigureRequest{
		Config: in.Config,
	}
}

func DestinationOpenRequest(_ *connectorv2.Destination_Open_Request) cplugin.DestinationOpenRequest {
	return cplugin.DestinationOpenRequest{}
}

func DestinationRunRequest(in *connectorv2.Destination_Run_Request) (cplugin.DestinationRunRequest, error) {
	records := make([]opencdc.Record, len(in.Records))
	for i, rec := range in.Records {
		err := records[i].FromProto(rec)
		if err != nil {
			return cplugin.DestinationRunRequest{}, err
		}
	}
	return cplugin.DestinationRunRequest{
		Records: records,
	}, nil
}

func DestinationStopRequest(in *connectorv2.Destination_Stop_Request) cplugin.DestinationStopRequest {
	return cplugin.DestinationStopRequest{
		LastPosition: in.LastPosition,
	}
}

func DestinationTeardownRequest(_ *connectorv2.Destination_Teardown_Request) cplugin.DestinationTeardownRequest {
	return cplugin.DestinationTeardownRequest{}
}

func DestinationLifecycleOnCreatedRequest(in *connectorv2.Destination_Lifecycle_OnCreated_Request) cplugin.DestinationLifecycleOnCreatedRequest {
	return cplugin.DestinationLifecycleOnCreatedRequest{
		Config: in.Config,
	}
}
func DestinationLifecycleOnUpdatedRequest(in *connectorv2.Destination_Lifecycle_OnUpdated_Request) cplugin.DestinationLifecycleOnUpdatedRequest {
	return cplugin.DestinationLifecycleOnUpdatedRequest{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}
func DestinationLifecycleOnDeletedRequest(in *connectorv2.Destination_Lifecycle_OnDeleted_Request) cplugin.DestinationLifecycleOnDeletedRequest {
	return cplugin.DestinationLifecycleOnDeletedRequest{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func DestinationConfigureResponse(_ *connectorv2.Destination_Configure_Response) cplugin.DestinationConfigureResponse {
	return cplugin.DestinationConfigureResponse{}
}

func DestinationOpenResponse(_ *connectorv2.Destination_Open_Response) cplugin.DestinationOpenResponse {
	return cplugin.DestinationOpenResponse{}
}

func DestinationRunResponse(in *connectorv2.Destination_Run_Response) cplugin.DestinationRunResponse {
	acks := make([]cplugin.DestinationRunResponseAck, len(in.Acks))
	for i, ack := range in.Acks {
		acks[i] = DestinationRunResponseAck(ack)
	}
	return cplugin.DestinationRunResponse{
		Acks: acks,
	}
}

func DestinationRunResponseAck(in *connectorv2.Destination_Run_Response_Ack) cplugin.DestinationRunResponseAck {
	return cplugin.DestinationRunResponseAck{
		Position: in.Position,
		Error:    in.Error,
	}
}

func DestinationStopResponse(_ *connectorv2.Destination_Stop_Response) cplugin.DestinationStopResponse {
	return cplugin.DestinationStopResponse{}
}

func DestinationTeardownResponse(_ *connectorv2.Destination_Teardown_Response) cplugin.DestinationTeardownResponse {
	return cplugin.DestinationTeardownResponse{}
}

func DestinationLifecycleOnCreatedResponse(_ *connectorv2.Destination_Lifecycle_OnCreated_Response) cplugin.DestinationLifecycleOnCreatedResponse {
	return cplugin.DestinationLifecycleOnCreatedResponse{}
}
func DestinationLifecycleOnUpdatedResponse(_ *connectorv2.Destination_Lifecycle_OnUpdated_Response) cplugin.DestinationLifecycleOnUpdatedResponse {
	return cplugin.DestinationLifecycleOnUpdatedResponse{}
}
func DestinationLifecycleOnDeletedResponse(_ *connectorv2.Destination_Lifecycle_OnDeleted_Response) cplugin.DestinationLifecycleOnDeletedResponse {
	return cplugin.DestinationLifecycleOnDeletedResponse{}
}
