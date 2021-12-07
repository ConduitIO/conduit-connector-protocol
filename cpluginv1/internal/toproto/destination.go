// Copyright © 2021 Meroxa Inc
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
	"github.com/conduitio/conduit-plugin/cpluginv1"
	"github.com/conduitio/conduit-plugin/cpluginv1/internal/cproto"
)

func DestinationConfigureRequest(in cpluginv1.DestinationConfigureRequest) (*cproto.Destination_Configure_Request, error) {
	out := cproto.Destination_Configure_Request{
		Config: in.Config,
	}
	return &out, nil
}

func DestinationConfigureResponse(in cpluginv1.DestinationConfigureResponse) (*cproto.Destination_Configure_Response, error) {
	return &cproto.Destination_Configure_Response{}, nil
}

func DestinationStartRequest(in cpluginv1.DestinationStartRequest) (*cproto.Destination_Start_Request, error) {
	return &cproto.Destination_Start_Request{}, nil
}

func DestinationStartResponse(in cpluginv1.DestinationStartResponse) (*cproto.Destination_Start_Response, error) {
	return &cproto.Destination_Start_Response{}, nil
}

func DestinationRunRequest(in cpluginv1.DestinationRunRequest) (*cproto.Destination_Run_Request, error) {
	rec, err := Record(in.Record)
	if err != nil {
		return nil, err
	}
	out := cproto.Destination_Run_Request{
		Record: rec,
	}
	return &out, nil
}

func DestinationRunResponse(in cpluginv1.DestinationRunResponse) (*cproto.Destination_Run_Response, error) {
	out := cproto.Destination_Run_Response{
		AckPosition: in.AckPosition,
		Error:       in.Error,
	}
	return &out, nil
}

func DestinationStopRequest(in cpluginv1.DestinationStopRequest) (*cproto.Destination_Stop_Request, error) {
	return &cproto.Destination_Stop_Request{}, nil
}

func DestinationStopResponse(in cpluginv1.DestinationStopResponse) (*cproto.Destination_Stop_Response, error) {
	return &cproto.Destination_Stop_Response{}, nil
}
