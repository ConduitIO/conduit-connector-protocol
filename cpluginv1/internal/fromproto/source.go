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
	"github.com/conduitio/conduit-plugin/cpluginv1/internal/cproto"
)

func SourceConfigureRequest(in *cproto.Source_Configure_Request) (cpluginv1.SourceConfigureRequest, error) {
	out := cpluginv1.SourceConfigureRequest{
		Config: in.Config,
	}
	return out, nil
}

func SourceConfigureResponse(in *cproto.Source_Configure_Response) (cpluginv1.SourceConfigureResponse, error) {
	return cpluginv1.SourceConfigureResponse{}, nil
}

func SourceStartRequest(in *cproto.Source_Start_Request) (cpluginv1.SourceStartRequest, error) {
	out := cpluginv1.SourceStartRequest{
		Position: in.Position,
	}
	return out, nil
}

func SourceStartResponse(in *cproto.Source_Start_Response) (cpluginv1.SourceStartResponse, error) {
	return cpluginv1.SourceStartResponse{}, nil
}

func SourceRunRequest(in *cproto.Source_Run_Request) (cpluginv1.SourceRunRequest, error) {
	out := cpluginv1.SourceRunRequest{
		AckPosition: in.AckPosition,
	}
	return out, nil
}

func SourceRunResponse(in *cproto.Source_Run_Response) (cpluginv1.SourceRunResponse, error) {
	rec, err := Record(in.Record)
	if err != nil {
		return cpluginv1.SourceRunResponse{}, err
	}

	out := cpluginv1.SourceRunResponse{
		Record: rec,
	}
	return out, nil
}

func SourceStopRequest(in *cproto.Source_Stop_Request) (cpluginv1.SourceStopRequest, error) {
	return cpluginv1.SourceStopRequest{}, nil
}

func SourceStopResponse(in *cproto.Source_Stop_Response) (cpluginv1.SourceStopResponse, error) {
	return cpluginv1.SourceStopResponse{}, nil
}
