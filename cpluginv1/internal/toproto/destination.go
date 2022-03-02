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
	"github.com/conduitio/conduit-connector-protocol/cpluginv1"
	connectorv1 "go.buf.build/library/go-grpc/conduitio/conduit-connector-protocol/connector/v1"
)

func DestinationConfigureRequest(in cpluginv1.DestinationConfigureRequest) (*connectorv1.Destination_Configure_Request, error) {
	out := connectorv1.Destination_Configure_Request{
		Config: in.Config,
	}
	return &out, nil
}
func DestinationConfigureResponse(in cpluginv1.DestinationConfigureResponse) (*connectorv1.Destination_Configure_Response, error) {
	return &connectorv1.Destination_Configure_Response{}, nil
}

func DestinationStartRequest(in cpluginv1.DestinationStartRequest) (*connectorv1.Destination_Start_Request, error) {
	return &connectorv1.Destination_Start_Request{}, nil
}
func DestinationStartResponse(in cpluginv1.DestinationStartResponse) (*connectorv1.Destination_Start_Response, error) {
	return &connectorv1.Destination_Start_Response{}, nil
}

func DestinationRunRequest(in cpluginv1.DestinationRunRequest) (*connectorv1.Destination_Run_Request, error) {
	rec, err := Record(in.Record)
	if err != nil {
		return nil, err
	}
	out := connectorv1.Destination_Run_Request{
		Record: rec,
	}
	return &out, nil
}
func DestinationRunResponse(in cpluginv1.DestinationRunResponse) (*connectorv1.Destination_Run_Response, error) {
	out := connectorv1.Destination_Run_Response{
		AckPosition: in.AckPosition,
		Error:       in.Error,
	}
	return &out, nil
}

func DestinationStopRequest(in cpluginv1.DestinationStopRequest) (*connectorv1.Destination_Stop_Request, error) {
	return &connectorv1.Destination_Stop_Request{}, nil
}
func DestinationStopResponse(in cpluginv1.DestinationStopResponse) (*connectorv1.Destination_Stop_Response, error) {
	return &connectorv1.Destination_Stop_Response{}, nil
}

func DestinationTeardownRequest(in cpluginv1.DestinationTeardownRequest) (*connectorv1.Destination_Teardown_Request, error) {
	return &connectorv1.Destination_Teardown_Request{}, nil
}
func DestinationTeardownResponse(in cpluginv1.DestinationTeardownResponse) (*connectorv1.Destination_Teardown_Response, error) {
	return &connectorv1.Destination_Teardown_Response{}, nil
}
