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

package fromproto

import (
	"github.com/conduitio/conduit-plugin/cpluginv1"
	cproto "github.com/conduitio/conduit-plugin/proto/gen/go/conduitio/cplugin/v1"
)

func DestinationConfigureRequest(in *cproto.Destination_Configure_Request) (cpluginv1.DestinationConfigureRequest, error) {
	out := cpluginv1.DestinationConfigureRequest{
		Config: in.Config,
	}
	return out, nil
}

func DestinationConfigureResponse(in *cproto.Destination_Configure_Response) (cpluginv1.DestinationConfigureResponse, error) {
	return cpluginv1.DestinationConfigureResponse{}, nil
}

func DestinationStartRequest(in *cproto.Destination_Start_Request) (cpluginv1.DestinationStartRequest, error) {
	return cpluginv1.DestinationStartRequest{}, nil
}

func DestinationStartResponse(in *cproto.Destination_Start_Response) (cpluginv1.DestinationStartResponse, error) {
	return cpluginv1.DestinationStartResponse{}, nil
}

func DestinationRunRequest(in *cproto.Destination_Run_Request) (cpluginv1.DestinationRunRequest, error) {
	rec, err := Record(in.Record)
	if err != nil {
		return cpluginv1.DestinationRunRequest{}, err
	}
	out := cpluginv1.DestinationRunRequest{
		Record: rec,
	}
	return out, nil
}

func DestinationRunResponse(in *cproto.Destination_Run_Response) (cpluginv1.DestinationRunResponse, error) {
	out := cpluginv1.DestinationRunResponse{
		AckPosition: in.AckPosition,
		Error:       in.Error,
	}
	return out, nil
}

func DestinationStopRequest(in *cproto.Destination_Stop_Request) (cpluginv1.DestinationStopRequest, error) {
	return cpluginv1.DestinationStopRequest{}, nil
}

func DestinationStopResponse(in *cproto.Destination_Stop_Response) (cpluginv1.DestinationStopResponse, error) {
	return cpluginv1.DestinationStopResponse{}, nil
}
