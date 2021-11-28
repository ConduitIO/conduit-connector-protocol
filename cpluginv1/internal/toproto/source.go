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

func SourceConfigureRequest(in cpluginv1.SourceConfigureRequest) (*cproto.Source_Configure_Request, error) {
	out := cproto.Source_Configure_Request{
		Config: in.Config,
	}
	return &out, nil
}

func SourceConfigureResponse(in cpluginv1.SourceConfigureResponse) (*cproto.Source_Configure_Response, error) {
	return &cproto.Source_Configure_Response{}, nil
}

func SourceStartRequest(in cpluginv1.SourceStartRequest) (*cproto.Source_Start_Request, error) {
	out := cproto.Source_Start_Request{
		Position: in.Position,
	}
	return &out, nil
}

func SourceStartResponse(in cpluginv1.SourceStartResponse) (*cproto.Source_Start_Response, error) {
	return &cproto.Source_Start_Response{}, nil
}

func SourceRunRequest(in cpluginv1.SourceRunRequest) (*cproto.Source_Run_Request, error) {
	out := cproto.Source_Run_Request{
		AckPosition: in.AckPosition,
	}
	return &out, nil
}

func SourceRunResponse(in cpluginv1.SourceRunResponse) (*cproto.Source_Run_Response, error) {
	rec, err := Record(in.Record)
	if err != nil {
		return nil, err
	}

	out := cproto.Source_Run_Response{
		Record: rec,
	}
	return &out, nil
}

func SourceStopRequest(in cpluginv1.SourceStopRequest) (*cproto.Source_Stop_Request, error) {
	return &cproto.Source_Stop_Request{}, nil
}

func SourceStopResponse(in cpluginv1.SourceStopResponse) (*cproto.Source_Stop_Response, error) {
	return &cproto.Source_Stop_Response{}, nil
}
