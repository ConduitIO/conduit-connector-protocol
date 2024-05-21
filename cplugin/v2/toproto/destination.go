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
	"github.com/conduitio/conduit-connector-protocol/cplugin"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
)

// -- Request Conversions -----------------------------------------------------

func DestinationConfigureRequest(in cplugin.DestinationConfigureRequest) *connectorv2.Destination_Configure_Request {
	return &connectorv2.Destination_Configure_Request{
		Config: in.Config,
	}
}

func DestinationStartRequest(_ cplugin.DestinationStartRequest) *connectorv2.Destination_Start_Request {
	return &connectorv2.Destination_Start_Request{}
}

func DestinationRunRequest(in cplugin.DestinationRunRequest) (*connectorv2.Destination_Run_Request, error) {
	out := connectorv2.Destination_Run_Request{
		Records: make([]*opencdcv1.Record, len(in.Records)),
	}

	for i, inRec := range in.Records {
		outRec := &opencdcv1.Record{}
		err := inRec.ToProto(outRec)
		if err != nil {
			return nil, err
		}
		out.Records[i] = outRec
	}

	return &out, nil
}

func DestinationStopRequest(in cplugin.DestinationStopRequest) *connectorv2.Destination_Stop_Request {
	return &connectorv2.Destination_Stop_Request{
		LastPosition: in.LastPosition,
	}
}

func DestinationTeardownRequest(_ cplugin.DestinationTeardownRequest) *connectorv2.Destination_Teardown_Request {
	return &connectorv2.Destination_Teardown_Request{}
}

func DestinationLifecycleOnCreatedRequest(in cplugin.DestinationLifecycleOnCreatedRequest) *connectorv2.Destination_Lifecycle_OnCreated_Request {
	return &connectorv2.Destination_Lifecycle_OnCreated_Request{
		Config: in.Config,
	}
}

func DestinationLifecycleOnUpdatedRequest(in cplugin.DestinationLifecycleOnUpdatedRequest) *connectorv2.Destination_Lifecycle_OnUpdated_Request {
	return &connectorv2.Destination_Lifecycle_OnUpdated_Request{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}

func DestinationLifecycleOnDeletedRequest(in cplugin.DestinationLifecycleOnDeletedRequest) *connectorv2.Destination_Lifecycle_OnDeleted_Request {
	return &connectorv2.Destination_Lifecycle_OnDeleted_Request{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func DestinationConfigureResponse(_ cplugin.DestinationConfigureResponse) *connectorv2.Destination_Configure_Response {
	return &connectorv2.Destination_Configure_Response{}
}

func DestinationStartResponse(_ cplugin.DestinationStartResponse) *connectorv2.Destination_Start_Response {
	return &connectorv2.Destination_Start_Response{}
}

func DestinationRunResponse(in cplugin.DestinationRunResponse) *connectorv2.Destination_Run_Response {
	acks := make([]*connectorv2.Destination_Run_Response_Ack, len(in.Acks))
	for i, inAck := range in.Acks {
		acks[i] = DestinationRunResponseAck(inAck)
	}
	return &connectorv2.Destination_Run_Response{
		Acks: acks,
	}
}

func DestinationRunResponseAck(in cplugin.DestinationRunResponseAck) *connectorv2.Destination_Run_Response_Ack {
	return &connectorv2.Destination_Run_Response_Ack{
		Position: in.Position,
		Error:    in.Error,
	}
}

func DestinationStopResponse(_ cplugin.DestinationStopResponse) *connectorv2.Destination_Stop_Response {
	return &connectorv2.Destination_Stop_Response{}
}

func DestinationTeardownResponse(_ cplugin.DestinationTeardownResponse) *connectorv2.Destination_Teardown_Response {
	return &connectorv2.Destination_Teardown_Response{}
}

func DestinationLifecycleOnCreatedResponse(_ cplugin.DestinationLifecycleOnCreatedResponse) *connectorv2.Destination_Lifecycle_OnCreated_Response {
	return &connectorv2.Destination_Lifecycle_OnCreated_Response{}
}
func DestinationLifecycleOnUpdatedResponse(_ cplugin.DestinationLifecycleOnUpdatedResponse) *connectorv2.Destination_Lifecycle_OnUpdated_Response {
	return &connectorv2.Destination_Lifecycle_OnUpdated_Response{}
}
func DestinationLifecycleOnDeletedResponse(_ cplugin.DestinationLifecycleOnDeletedResponse) *connectorv2.Destination_Lifecycle_OnDeleted_Response {
	return &connectorv2.Destination_Lifecycle_OnDeleted_Response{}
}
