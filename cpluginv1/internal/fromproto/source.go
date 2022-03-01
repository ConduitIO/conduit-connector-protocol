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
	"github.com/conduitio/conduit-plugin-protocol/cpluginv1"
	connectorv1 "go.buf.build/library/go-grpc/conduitio/conduit-plugin-protocol/connector/v1"
)

func SourceConfigureRequest(in *connectorv1.Source_Configure_Request) (cpluginv1.SourceConfigureRequest, error) {
	out := cpluginv1.SourceConfigureRequest{
		Config: in.Config,
	}
	return out, nil
}
func SourceConfigureResponse(in *connectorv1.Source_Configure_Response) (cpluginv1.SourceConfigureResponse, error) {
	return cpluginv1.SourceConfigureResponse{}, nil
}

func SourceStartRequest(in *connectorv1.Source_Start_Request) (cpluginv1.SourceStartRequest, error) {
	out := cpluginv1.SourceStartRequest{
		Position: in.Position,
	}
	return out, nil
}
func SourceStartResponse(in *connectorv1.Source_Start_Response) (cpluginv1.SourceStartResponse, error) {
	return cpluginv1.SourceStartResponse{}, nil
}

func SourceRunRequest(in *connectorv1.Source_Run_Request) (cpluginv1.SourceRunRequest, error) {
	out := cpluginv1.SourceRunRequest{
		AckPosition: in.AckPosition,
	}
	return out, nil
}
func SourceRunResponse(in *connectorv1.Source_Run_Response) (cpluginv1.SourceRunResponse, error) {
	rec, err := Record(in.Record)
	if err != nil {
		return cpluginv1.SourceRunResponse{}, err
	}

	out := cpluginv1.SourceRunResponse{
		Record: rec,
	}
	return out, nil
}

func SourceStopRequest(in *connectorv1.Source_Stop_Request) (cpluginv1.SourceStopRequest, error) {
	return cpluginv1.SourceStopRequest{}, nil
}
func SourceStopResponse(in *connectorv1.Source_Stop_Response) (cpluginv1.SourceStopResponse, error) {
	return cpluginv1.SourceStopResponse{}, nil
}

func SourceTeardownRequest(in *connectorv1.Source_Teardown_Request) (cpluginv1.SourceTeardownRequest, error) {
	return cpluginv1.SourceTeardownRequest{}, nil
}
func SourceTeardownResponse(in *connectorv1.Source_Teardown_Response) (cpluginv1.SourceTeardownResponse, error) {
	return cpluginv1.SourceTeardownResponse{}, nil
}
