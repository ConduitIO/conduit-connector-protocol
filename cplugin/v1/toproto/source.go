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

func SourceConfigureRequest(in cplugin.SourceConfigureRequest) *connectorv1.Source_Configure_Request {
	return &connectorv1.Source_Configure_Request{
		Config: in.Config,
	}
}

func SourceStartRequest(in cplugin.SourceStartRequest) *connectorv1.Source_Start_Request {
	return &connectorv1.Source_Start_Request{
		Position: in.Position,
	}
}

func SourceRunRequest(in cplugin.SourceRunRequest) []*connectorv1.Source_Run_Request {
	out := make([]*connectorv1.Source_Run_Request, len(in.AckPositions))
	for i, pos := range in.AckPositions {
		out[i] = &connectorv1.Source_Run_Request{
			AckPosition: pos,
		}
	}
	return out
}

func SourceStopRequest(_ cplugin.SourceStopRequest) *connectorv1.Source_Stop_Request {
	return &connectorv1.Source_Stop_Request{}
}

func SourceTeardownRequest(_ cplugin.SourceTeardownRequest) *connectorv1.Source_Teardown_Request {
	return &connectorv1.Source_Teardown_Request{}
}

func SourceLifecycleOnCreatedRequest(in cplugin.SourceLifecycleOnCreatedRequest) *connectorv1.Source_Lifecycle_OnCreated_Request {
	return &connectorv1.Source_Lifecycle_OnCreated_Request{
		Config: in.Config,
	}
}

func SourceLifecycleOnUpdatedRequest(in cplugin.SourceLifecycleOnUpdatedRequest) *connectorv1.Source_Lifecycle_OnUpdated_Request {
	return &connectorv1.Source_Lifecycle_OnUpdated_Request{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}

func SourceLifecycleOnDeletedRequest(in cplugin.SourceLifecycleOnDeletedRequest) *connectorv1.Source_Lifecycle_OnDeleted_Request {
	return &connectorv1.Source_Lifecycle_OnDeleted_Request{
		Config: in.Config,
	}
}

// -- Response Conversions ----------------------------------------------------

func SourceConfigureResponse(_ cplugin.SourceConfigureResponse) *connectorv1.Source_Configure_Response {
	return &connectorv1.Source_Configure_Response{}
}

func SourceStartResponse(_ cplugin.SourceStartResponse) *connectorv1.Source_Start_Response {
	return &connectorv1.Source_Start_Response{}
}

func SourceRunResponse(in cplugin.SourceRunResponse) ([]*connectorv1.Source_Run_Response, error) {
	out := make([]*connectorv1.Source_Run_Response, len(in.Records))
	for i, rec := range in.Records {
		outResp := connectorv1.Source_Run_Response{
			Record: &opencdcv1.Record{},
		}
		err := rec.ToProto(outResp.Record)
		if err != nil {
			return nil, err
		}
		out[i] = &outResp
	}
	return out, nil
}

func SourceStopResponse(in cplugin.SourceStopResponse) *connectorv1.Source_Stop_Response {
	return &connectorv1.Source_Stop_Response{
		LastPosition: in.LastPosition,
	}
}

func SourceTeardownResponse(_ cplugin.SourceTeardownResponse) *connectorv1.Source_Teardown_Response {
	return &connectorv1.Source_Teardown_Response{}
}

func SourceLifecycleOnCreatedResponse(_ cplugin.SourceLifecycleOnCreatedResponse) *connectorv1.Source_Lifecycle_OnCreated_Response {
	return &connectorv1.Source_Lifecycle_OnCreated_Response{}
}
func SourceLifecycleOnUpdatedResponse(_ cplugin.SourceLifecycleOnUpdatedResponse) *connectorv1.Source_Lifecycle_OnUpdated_Response {
	return &connectorv1.Source_Lifecycle_OnUpdated_Response{}
}
func SourceLifecycleOnDeletedResponse(_ cplugin.SourceLifecycleOnDeletedResponse) *connectorv1.Source_Lifecycle_OnDeleted_Response {
	return &connectorv1.Source_Lifecycle_OnDeleted_Response{}
}
