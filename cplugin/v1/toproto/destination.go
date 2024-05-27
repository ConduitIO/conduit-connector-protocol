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

func DestinationConfigureRequest(in cplugin.DestinationConfigureRequest) *connectorv1.Destination_Configure_Request {
	return &connectorv1.Destination_Configure_Request{
		Config: in.Config,
	}
}

func DestinationStartRequest(_ cplugin.DestinationOpenRequest) *connectorv1.Destination_Start_Request {
	return &connectorv1.Destination_Start_Request{}
}

func DestinationRunRequest(in cplugin.DestinationRunRequest) ([]*connectorv1.Destination_Run_Request, error) {
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

func DestinationStopRequest(in cplugin.DestinationStopRequest) *connectorv1.Destination_Stop_Request {
	return &connectorv1.Destination_Stop_Request{
		LastPosition: in.LastPosition,
	}
}

func DestinationTeardownRequest(_ cplugin.DestinationTeardownRequest) *connectorv1.Destination_Teardown_Request {
	return &connectorv1.Destination_Teardown_Request{}
}

func DestinationLifecycleOnCreatedRequest(in cplugin.DestinationLifecycleOnCreatedRequest) *connectorv1.Destination_Lifecycle_OnCreated_Request {
	return &connectorv1.Destination_Lifecycle_OnCreated_Request{
		Config: in.Config,
	}
}

func DestinationLifecycleOnUpdatedRequest(in cplugin.DestinationLifecycleOnUpdatedRequest) *connectorv1.Destination_Lifecycle_OnUpdated_Request {
	return &connectorv1.Destination_Lifecycle_OnUpdated_Request{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}

func DestinationLifecycleOnDeletedRequest(in cplugin.DestinationLifecycleOnDeletedRequest) *connectorv1.Destination_Lifecycle_OnDeleted_Request {
	return &connectorv1.Destination_Lifecycle_OnDeleted_Request{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func DestinationConfigureResponse(_ cplugin.DestinationConfigureResponse) *connectorv1.Destination_Configure_Response {
	return &connectorv1.Destination_Configure_Response{}
}

func DestinationStartResponse(_ cplugin.DestinationOpenResponse) *connectorv1.Destination_Start_Response {
	return &connectorv1.Destination_Start_Response{}
}

func DestinationRunResponse(in cplugin.DestinationRunResponse) []*connectorv1.Destination_Run_Response {
	out := make([]*connectorv1.Destination_Run_Response, len(in.Acks))
	for i, ack := range in.Acks {
		out[i] = &connectorv1.Destination_Run_Response{
			AckPosition: ack.Position,
			Error:       ack.Error,
		}
	}
	return out
}

func DestinationStopResponse(_ cplugin.DestinationStopResponse) *connectorv1.Destination_Stop_Response {
	return &connectorv1.Destination_Stop_Response{}
}

func DestinationTeardownResponse(_ cplugin.DestinationTeardownResponse) *connectorv1.Destination_Teardown_Response {
	return &connectorv1.Destination_Teardown_Response{}
}

func DestinationLifecycleOnCreatedResponse(_ cplugin.DestinationLifecycleOnCreatedResponse) *connectorv1.Destination_Lifecycle_OnCreated_Response {
	return &connectorv1.Destination_Lifecycle_OnCreated_Response{}
}
func DestinationLifecycleOnUpdatedResponse(_ cplugin.DestinationLifecycleOnUpdatedResponse) *connectorv1.Destination_Lifecycle_OnUpdated_Response {
	return &connectorv1.Destination_Lifecycle_OnUpdated_Response{}
}
func DestinationLifecycleOnDeletedResponse(_ cplugin.DestinationLifecycleOnDeletedResponse) *connectorv1.Destination_Lifecycle_OnDeleted_Response {
	return &connectorv1.Destination_Lifecycle_OnDeleted_Response{}
}
