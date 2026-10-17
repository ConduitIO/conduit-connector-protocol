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

func SourceConfigureRequest(in pconnector.SourceConfigureRequest) *connectorv3.Source_Configure_Request {
	return &connectorv3.Source_Configure_Request{
		Config: in.Config,
	}
}

func SourceOpenRequest(in pconnector.SourceOpenRequest) *connectorv3.Source_Open_Request {
	return &connectorv3.Source_Open_Request{
		Position: in.Position,
	}
}

func SourceRunRequest(in pconnector.SourceRunRequest) *connectorv3.Source_Run_Request {
	ackPositions := make([][]byte, len(in.AckPositions))
	for i, pos := range in.AckPositions {
		ackPositions[i] = pos
	}

	return &connectorv3.Source_Run_Request{
		AckPositions: ackPositions,
	}
}

func SourceStopRequest(_ pconnector.SourceStopRequest) *connectorv3.Source_Stop_Request {
	return &connectorv3.Source_Stop_Request{}
}

func SourceTeardownRequest(_ pconnector.SourceTeardownRequest) *connectorv3.Source_Teardown_Request {
	return &connectorv3.Source_Teardown_Request{}
}

func SourceLifecycleOnCreatedRequest(in pconnector.SourceLifecycleOnCreatedRequest) *connectorv3.Source_Lifecycle_OnCreated_Request {
	return &connectorv3.Source_Lifecycle_OnCreated_Request{
		Config: in.Config,
	}
}

func SourceLifecycleOnUpdatedRequest(in pconnector.SourceLifecycleOnUpdatedRequest) *connectorv3.Source_Lifecycle_OnUpdated_Request {
	return &connectorv3.Source_Lifecycle_OnUpdated_Request{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}

func SourceLifecycleOnDeletedRequest(in pconnector.SourceLifecycleOnDeletedRequest) *connectorv3.Source_Lifecycle_OnDeleted_Request {
	return &connectorv3.Source_Lifecycle_OnDeleted_Request{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func SourceConfigureResponse(_ pconnector.SourceConfigureResponse) *connectorv3.Source_Configure_Response {
	return &connectorv3.Source_Configure_Response{}
}

func SourceOpenResponse(_ pconnector.SourceOpenResponse) *connectorv3.Source_Open_Response {
	return &connectorv3.Source_Open_Response{}
}

func SourceRunResponse(in pconnector.SourceRunResponse) (*connectorv3.Source_Run_Response, error) {
	out := connectorv3.Source_Run_Response{
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

func SourceStopResponse(in pconnector.SourceStopResponse) *connectorv3.Source_Stop_Response {
	return &connectorv3.Source_Stop_Response{
		LastPosition: in.LastPosition,
	}
}

func SourceTeardownResponse(_ pconnector.SourceTeardownResponse) *connectorv3.Source_Teardown_Response {
	return &connectorv3.Source_Teardown_Response{}
}

func SourceLifecycleOnCreatedResponse(_ pconnector.SourceLifecycleOnCreatedResponse) *connectorv3.Source_Lifecycle_OnCreated_Response {
	return &connectorv3.Source_Lifecycle_OnCreated_Response{}
}

func SourceLifecycleOnUpdatedResponse(_ pconnector.SourceLifecycleOnUpdatedResponse) *connectorv3.Source_Lifecycle_OnUpdated_Response {
	return &connectorv3.Source_Lifecycle_OnUpdated_Response{}
}

func SourceLifecycleOnDeletedResponse(_ pconnector.SourceLifecycleOnDeletedResponse) *connectorv3.Source_Lifecycle_OnDeleted_Response {
	return &connectorv3.Source_Lifecycle_OnDeleted_Response{}
}
