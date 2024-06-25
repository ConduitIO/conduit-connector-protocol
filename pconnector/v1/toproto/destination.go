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
	"github.com/conduitio/conduit-connector-protocol/pconnector"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
)

// -- Request Conversions -----------------------------------------------------

func DestinationConfigureRequest(in pconnector.DestinationConfigureRequest) *connectorv1.Destination_Configure_Request {
	return &connectorv1.Destination_Configure_Request{
		Config: in.Config,
	}
}

func DestinationStartRequest(_ pconnector.DestinationOpenRequest) *connectorv1.Destination_Start_Request {
	return &connectorv1.Destination_Start_Request{}
}

func DestinationRunRequest(in pconnector.DestinationRunRequest) ([]*connectorv1.Destination_Run_Request, error) {
	out := make([]*connectorv1.Destination_Run_Request, len(in.Records))
	for i, rec := range in.Records {
		outReq := connectorv1.Destination_Run_Request{
			Record: &opencdcv1.Record{},
		}
		err := rec.ToProto(outReq.Record)
		if err != nil {
			return nil, err
		}
		out[i] = &outReq
	}
	return out, nil
}

func DestinationStopRequest(in pconnector.DestinationStopRequest) *connectorv1.Destination_Stop_Request {
	return &connectorv1.Destination_Stop_Request{
		LastPosition: in.LastPosition,
	}
}

func DestinationTeardownRequest(_ pconnector.DestinationTeardownRequest) *connectorv1.Destination_Teardown_Request {
	return &connectorv1.Destination_Teardown_Request{}
}

func DestinationLifecycleOnCreatedRequest(in pconnector.DestinationLifecycleOnCreatedRequest) *connectorv1.Destination_Lifecycle_OnCreated_Request {
	return &connectorv1.Destination_Lifecycle_OnCreated_Request{
		Config: in.Config,
	}
}

func DestinationLifecycleOnUpdatedRequest(in pconnector.DestinationLifecycleOnUpdatedRequest) *connectorv1.Destination_Lifecycle_OnUpdated_Request {
	return &connectorv1.Destination_Lifecycle_OnUpdated_Request{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}

func DestinationLifecycleOnDeletedRequest(in pconnector.DestinationLifecycleOnDeletedRequest) *connectorv1.Destination_Lifecycle_OnDeleted_Request {
	return &connectorv1.Destination_Lifecycle_OnDeleted_Request{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func DestinationConfigureResponse(_ pconnector.DestinationConfigureResponse) *connectorv1.Destination_Configure_Response {
	return &connectorv1.Destination_Configure_Response{}
}

func DestinationStartResponse(_ pconnector.DestinationOpenResponse) *connectorv1.Destination_Start_Response {
	return &connectorv1.Destination_Start_Response{}
}

func DestinationRunResponse(in pconnector.DestinationRunResponse) []*connectorv1.Destination_Run_Response {
	out := make([]*connectorv1.Destination_Run_Response, len(in.Acks))
	for i, ack := range in.Acks {
		out[i] = &connectorv1.Destination_Run_Response{
			AckPosition: ack.Position,
			Error:       ack.Error,
		}
	}
	return out
}

func DestinationStopResponse(_ pconnector.DestinationStopResponse) *connectorv1.Destination_Stop_Response {
	return &connectorv1.Destination_Stop_Response{}
}

func DestinationTeardownResponse(_ pconnector.DestinationTeardownResponse) *connectorv1.Destination_Teardown_Response {
	return &connectorv1.Destination_Teardown_Response{}
}

func DestinationLifecycleOnCreatedResponse(_ pconnector.DestinationLifecycleOnCreatedResponse) *connectorv1.Destination_Lifecycle_OnCreated_Response {
	return &connectorv1.Destination_Lifecycle_OnCreated_Response{}
}

func DestinationLifecycleOnUpdatedResponse(_ pconnector.DestinationLifecycleOnUpdatedResponse) *connectorv1.Destination_Lifecycle_OnUpdated_Response {
	return &connectorv1.Destination_Lifecycle_OnUpdated_Response{}
}

func DestinationLifecycleOnDeletedResponse(_ pconnector.DestinationLifecycleOnDeletedResponse) *connectorv1.Destination_Lifecycle_OnDeleted_Response {
	return &connectorv1.Destination_Lifecycle_OnDeleted_Response{}
}
