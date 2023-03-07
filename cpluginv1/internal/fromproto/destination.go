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
	connectorv1 "buf.build/gen/go/conduitio/conduit-connector-protocol/protocolbuffers/go/connector/v1"
	"github.com/conduitio/conduit-connector-protocol/cpluginv1"
)

func DestinationConfigureRequest(in *connectorv1.Destination_Configure_Request) (cpluginv1.DestinationConfigureRequest, error) {
	out := cpluginv1.DestinationConfigureRequest{
		Config: in.Config,
	}
	return out, nil
}
func DestinationConfigureResponse(in *connectorv1.Destination_Configure_Response) (cpluginv1.DestinationConfigureResponse, error) {
	return cpluginv1.DestinationConfigureResponse{}, nil
}

func DestinationStartRequest(in *connectorv1.Destination_Start_Request) (cpluginv1.DestinationStartRequest, error) {
	return cpluginv1.DestinationStartRequest{}, nil
}
func DestinationStartResponse(in *connectorv1.Destination_Start_Response) (cpluginv1.DestinationStartResponse, error) {
	return cpluginv1.DestinationStartResponse{}, nil
}

func DestinationRunRequest(in *connectorv1.Destination_Run_Request) (cpluginv1.DestinationRunRequest, error) {
	rec, err := Record(in.Record)
	if err != nil {
		return cpluginv1.DestinationRunRequest{}, err
	}
	out := cpluginv1.DestinationRunRequest{
		Record: rec,
	}
	return out, nil
}
func DestinationRunResponse(in *connectorv1.Destination_Run_Response) (cpluginv1.DestinationRunResponse, error) {
	out := cpluginv1.DestinationRunResponse{
		AckPosition: in.AckPosition,
		Error:       in.Error,
	}
	return out, nil
}

func DestinationStopRequest(in *connectorv1.Destination_Stop_Request) (cpluginv1.DestinationStopRequest, error) {
	out := cpluginv1.DestinationStopRequest{
		LastPosition: in.LastPosition,
	}
	return out, nil
}
func DestinationStopResponse(in *connectorv1.Destination_Stop_Response) (cpluginv1.DestinationStopResponse, error) {
	return cpluginv1.DestinationStopResponse{}, nil
}

func DestinationTeardownRequest(in *connectorv1.Destination_Teardown_Request) (cpluginv1.DestinationTeardownRequest, error) {
	return cpluginv1.DestinationTeardownRequest{}, nil
}
func DestinationTeardownResponse(in *connectorv1.Destination_Teardown_Response) (cpluginv1.DestinationTeardownResponse, error) {
	return cpluginv1.DestinationTeardownResponse{}, nil
}
