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
	connectorv1 "go.buf.build/grpc/go/conduitio/conduit-connector-protocol/connector/v1"
)

func SourceConfigureRequest(in cpluginv1.SourceConfigureRequest) (*connectorv1.Source_Configure_Request, error) {
	out := connectorv1.Source_Configure_Request{
		Config: in.Config,
	}
	return &out, nil
}
func SourceConfigureResponse(in cpluginv1.SourceConfigureResponse) (*connectorv1.Source_Configure_Response, error) {
	return &connectorv1.Source_Configure_Response{}, nil
}

func SourceStartRequest(in cpluginv1.SourceStartRequest) (*connectorv1.Source_Start_Request, error) {
	out := connectorv1.Source_Start_Request{
		Position: in.Position,
	}
	return &out, nil
}
func SourceStartResponse(in cpluginv1.SourceStartResponse) (*connectorv1.Source_Start_Response, error) {
	return &connectorv1.Source_Start_Response{}, nil
}

func SourceRunRequest(in cpluginv1.SourceRunRequest) (*connectorv1.Source_Run_Request, error) {
	out := connectorv1.Source_Run_Request{
		AckPosition: in.AckPosition,
	}
	return &out, nil
}
func SourceRunResponse(in cpluginv1.SourceRunResponse) (*connectorv1.Source_Run_Response, error) {
	rec, err := Record(in.Record)
	if err != nil {
		return nil, err
	}

	out := connectorv1.Source_Run_Response{
		Record: rec,
	}
	return &out, nil
}

func SourceStopRequest(in cpluginv1.SourceStopRequest) (*connectorv1.Source_Stop_Request, error) {
	return &connectorv1.Source_Stop_Request{}, nil
}
func SourceStopResponse(in cpluginv1.SourceStopResponse) (*connectorv1.Source_Stop_Response, error) {
	out := connectorv1.Source_Stop_Response{
		LastPosition: in.LastPosition,
	}
	return &out, nil
}

func SourceTeardownRequest(in cpluginv1.SourceTeardownRequest) (*connectorv1.Source_Teardown_Request, error) {
	return &connectorv1.Source_Teardown_Request{}, nil
}
func SourceTeardownResponse(in cpluginv1.SourceTeardownResponse) (*connectorv1.Source_Teardown_Response, error) {
	return &connectorv1.Source_Teardown_Response{}, nil
}
