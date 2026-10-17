// Copyright © 2026 Meroxa, Inc.
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
	opencdcv2 "github.com/conduitio/conduit-commons/proto/opencdc/v2"
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	connectorv3 "github.com/conduitio/conduit-connector-protocol/proto/connector/v3"
)

// -- Request Conversions -----------------------------------------------------

func DestinationConfigureRequest(in pconnector.DestinationConfigureRequest) *connectorv3.Destination_Configure_Request {
	return &connectorv3.Destination_Configure_Request{
		Config: in.Config,
	}
}

func DestinationOpenRequest(_ pconnector.DestinationOpenRequest) *connectorv3.Destination_Open_Request {
	return &connectorv3.Destination_Open_Request{}
}

func DestinationRunRequest(in pconnector.DestinationRunRequest) (*connectorv3.Destination_Run_Request, error) {
	out := connectorv3.Destination_Run_Request{
		Records: make([]*opencdcv2.Record, len(in.Records)),
	}

	for i, inRec := range in.Records {
		outRec := &opencdcv2.Record{}
		err := inRec.ToProtoV2(outRec)
		if err != nil {
			return nil, err
		}
		out.Records[i] = outRec
	}

	return &out, nil
}

func DestinationStopRequest(in pconnector.DestinationStopRequest) *connectorv3.Destination_Stop_Request {
	return &connectorv3.Destination_Stop_Request{
		LastPosition: in.LastPosition,
	}
}

func DestinationTeardownRequest(_ pconnector.DestinationTeardownRequest) *connectorv3.Destination_Teardown_Request {
	return &connectorv3.Destination_Teardown_Request{}
}

func DestinationLifecycleOnCreatedRequest(in pconnector.DestinationLifecycleOnCreatedRequest) *connectorv3.Destination_Lifecycle_OnCreated_Request {
	return &connectorv3.Destination_Lifecycle_OnCreated_Request{
		Config: in.Config,
	}
}

func DestinationLifecycleOnUpdatedRequest(in pconnector.DestinationLifecycleOnUpdatedRequest) *connectorv3.Destination_Lifecycle_OnUpdated_Request {
	return &connectorv3.Destination_Lifecycle_OnUpdated_Request{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}

func DestinationLifecycleOnDeletedRequest(in pconnector.DestinationLifecycleOnDeletedRequest) *connectorv3.Destination_Lifecycle_OnDeleted_Request {
	return &connectorv3.Destination_Lifecycle_OnDeleted_Request{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func DestinationConfigureResponse(_ pconnector.DestinationConfigureResponse) *connectorv3.Destination_Configure_Response {
	return &connectorv3.Destination_Configure_Response{}
}

func DestinationOpenResponse(_ pconnector.DestinationOpenResponse) *connectorv3.Destination_Open_Response {
	return &connectorv3.Destination_Open_Response{}
}

func DestinationRunResponse(in pconnector.DestinationRunResponse) *connectorv3.Destination_Run_Response {
	acks := make([]*connectorv3.Destination_Run_Response_Ack, len(in.Acks))
	for i, inAck := range in.Acks {
		acks[i] = DestinationRunResponseAck(inAck)
	}
	return &connectorv3.Destination_Run_Response{
		Acks: acks,
	}
}

func DestinationRunResponseAck(in pconnector.DestinationRunResponseAck) *connectorv3.Destination_Run_Response_Ack {
	return &connectorv3.Destination_Run_Response_Ack{
		Position: in.Position,
		Error:    in.Error,
	}
}

func DestinationStopResponse(_ pconnector.DestinationStopResponse) *connectorv3.Destination_Stop_Response {
	return &connectorv3.Destination_Stop_Response{}
}

func DestinationTeardownResponse(_ pconnector.DestinationTeardownResponse) *connectorv3.Destination_Teardown_Response {
	return &connectorv3.Destination_Teardown_Response{}
}

func DestinationLifecycleOnCreatedResponse(_ pconnector.DestinationLifecycleOnCreatedResponse) *connectorv3.Destination_Lifecycle_OnCreated_Response {
	return &connectorv3.Destination_Lifecycle_OnCreated_Response{}
}

func DestinationLifecycleOnUpdatedResponse(_ pconnector.DestinationLifecycleOnUpdatedResponse) *connectorv3.Destination_Lifecycle_OnUpdated_Response {
	return &connectorv3.Destination_Lifecycle_OnUpdated_Response{}
}

func DestinationLifecycleOnDeletedResponse(_ pconnector.DestinationLifecycleOnDeletedResponse) *connectorv3.Destination_Lifecycle_OnDeleted_Response {
	return &connectorv3.Destination_Lifecycle_OnDeleted_Response{}
}
