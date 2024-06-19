// Copyright © 2024 Meroxa, Inc.
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

package pconnector

import (
	"bytes"
	"maps"

	"github.com/conduitio/conduit-commons/opencdc"
)

type cloner[T any] interface {
	Clone() T
}

var (
	_ cloner[SourceConfigureRequest]           = SourceConfigureRequest{}
	_ cloner[SourceConfigureResponse]          = SourceConfigureResponse{}
	_ cloner[SourceOpenRequest]                = SourceOpenRequest{}
	_ cloner[SourceOpenResponse]               = SourceOpenResponse{}
	_ cloner[SourceRunRequest]                 = SourceRunRequest{}
	_ cloner[SourceRunResponse]                = SourceRunResponse{}
	_ cloner[SourceStopRequest]                = SourceStopRequest{}
	_ cloner[SourceStopResponse]               = SourceStopResponse{}
	_ cloner[SourceTeardownRequest]            = SourceTeardownRequest{}
	_ cloner[SourceTeardownResponse]           = SourceTeardownResponse{}
	_ cloner[SourceLifecycleOnCreatedRequest]  = SourceLifecycleOnCreatedRequest{}
	_ cloner[SourceLifecycleOnCreatedResponse] = SourceLifecycleOnCreatedResponse{}
	_ cloner[SourceLifecycleOnUpdatedRequest]  = SourceLifecycleOnUpdatedRequest{}
	_ cloner[SourceLifecycleOnUpdatedResponse] = SourceLifecycleOnUpdatedResponse{}
	_ cloner[SourceLifecycleOnDeletedRequest]  = SourceLifecycleOnDeletedRequest{}
	_ cloner[SourceLifecycleOnDeletedResponse] = SourceLifecycleOnDeletedResponse{}

	_ cloner[DestinationConfigureRequest]           = DestinationConfigureRequest{}
	_ cloner[DestinationConfigureResponse]          = DestinationConfigureResponse{}
	_ cloner[DestinationOpenRequest]                = DestinationOpenRequest{}
	_ cloner[DestinationOpenResponse]               = DestinationOpenResponse{}
	_ cloner[DestinationRunRequest]                 = DestinationRunRequest{}
	_ cloner[DestinationRunResponse]                = DestinationRunResponse{}
	_ cloner[DestinationRunResponseAck]             = DestinationRunResponseAck{}
	_ cloner[DestinationStopRequest]                = DestinationStopRequest{}
	_ cloner[DestinationStopResponse]               = DestinationStopResponse{}
	_ cloner[DestinationTeardownRequest]            = DestinationTeardownRequest{}
	_ cloner[DestinationTeardownResponse]           = DestinationTeardownResponse{}
	_ cloner[DestinationLifecycleOnCreatedRequest]  = DestinationLifecycleOnCreatedRequest{}
	_ cloner[DestinationLifecycleOnCreatedResponse] = DestinationLifecycleOnCreatedResponse{}
	_ cloner[DestinationLifecycleOnUpdatedRequest]  = DestinationLifecycleOnUpdatedRequest{}
	_ cloner[DestinationLifecycleOnUpdatedResponse] = DestinationLifecycleOnUpdatedResponse{}
	_ cloner[DestinationLifecycleOnDeletedRequest]  = DestinationLifecycleOnDeletedRequest{}
	_ cloner[DestinationLifecycleOnDeletedResponse] = DestinationLifecycleOnDeletedResponse{}
)

func (r SourceConfigureRequest) Clone() SourceConfigureRequest {
	return SourceConfigureRequest{
		Config: maps.Clone(r.Config),
	}
}

func (r SourceConfigureResponse) Clone() SourceConfigureResponse {
	return SourceConfigureResponse{}
}

func (r SourceOpenRequest) Clone() SourceOpenRequest {
	return SourceOpenRequest{
		Position: bytes.Clone(r.Position),
	}
}

func (r SourceOpenResponse) Clone() SourceOpenResponse {
	return SourceOpenResponse{}
}

func (r SourceRunRequest) Clone() SourceRunRequest {
	return SourceRunRequest{
		AckPositions: cloneSlice(r.AckPositions, func(p opencdc.Position) opencdc.Position {
			return bytes.Clone(p)
		}),
	}
}

func (r SourceRunResponse) Clone() SourceRunResponse {
	return SourceRunResponse{
		Records: cloneSlice(r.Records, func(r opencdc.Record) opencdc.Record {
			return r.Clone()
		}),
	}
}

func (r SourceStopRequest) Clone() SourceStopRequest {
	return SourceStopRequest{}
}

func (r SourceStopResponse) Clone() SourceStopResponse {
	return SourceStopResponse{
		LastPosition: bytes.Clone(r.LastPosition),
	}
}

func (r SourceTeardownRequest) Clone() SourceTeardownRequest {
	return SourceTeardownRequest{}
}

func (r SourceTeardownResponse) Clone() SourceTeardownResponse {
	return SourceTeardownResponse{}
}

func (r SourceLifecycleOnCreatedRequest) Clone() SourceLifecycleOnCreatedRequest {
	return SourceLifecycleOnCreatedRequest{
		Config: maps.Clone(r.Config),
	}
}

func (r SourceLifecycleOnCreatedResponse) Clone() SourceLifecycleOnCreatedResponse {
	return SourceLifecycleOnCreatedResponse{}
}

func (r SourceLifecycleOnUpdatedRequest) Clone() SourceLifecycleOnUpdatedRequest {
	return SourceLifecycleOnUpdatedRequest{
		ConfigBefore: maps.Clone(r.ConfigBefore),
		ConfigAfter:  maps.Clone(r.ConfigAfter),
	}
}

func (r SourceLifecycleOnUpdatedResponse) Clone() SourceLifecycleOnUpdatedResponse {
	return SourceLifecycleOnUpdatedResponse{}
}

func (r SourceLifecycleOnDeletedRequest) Clone() SourceLifecycleOnDeletedRequest {
	return SourceLifecycleOnDeletedRequest{
		Config: maps.Clone(r.Config),
	}
}

func (r SourceLifecycleOnDeletedResponse) Clone() SourceLifecycleOnDeletedResponse {
	return SourceLifecycleOnDeletedResponse{}
}

func (r DestinationConfigureRequest) Clone() DestinationConfigureRequest {
	return DestinationConfigureRequest{
		Config: maps.Clone(r.Config),
	}
}

func (r DestinationConfigureResponse) Clone() DestinationConfigureResponse {
	return DestinationConfigureResponse{}
}

func (r DestinationOpenRequest) Clone() DestinationOpenRequest {
	return DestinationOpenRequest{}
}

func (r DestinationOpenResponse) Clone() DestinationOpenResponse {
	return DestinationOpenResponse{}
}

func (r DestinationRunRequest) Clone() DestinationRunRequest {
	return DestinationRunRequest{
		Records: cloneSlice(r.Records, func(r opencdc.Record) opencdc.Record {
			return r.Clone()
		}),
	}
}

func (r DestinationRunResponse) Clone() DestinationRunResponse {
	return DestinationRunResponse{
		Acks: cloneSlice(r.Acks, func(a DestinationRunResponseAck) DestinationRunResponseAck {
			return a.Clone()
		}),
	}
}

func (r DestinationRunResponseAck) Clone() DestinationRunResponseAck {
	return DestinationRunResponseAck{
		Position: bytes.Clone(r.Position),
		Error:    r.Error,
	}
}

func (r DestinationStopRequest) Clone() DestinationStopRequest {
	return DestinationStopRequest{
		LastPosition: bytes.Clone(r.LastPosition),
	}
}

func (r DestinationStopResponse) Clone() DestinationStopResponse {
	return DestinationStopResponse{}
}

func (r DestinationTeardownRequest) Clone() DestinationTeardownRequest {
	return DestinationTeardownRequest{}
}

func (r DestinationTeardownResponse) Clone() DestinationTeardownResponse {
	return DestinationTeardownResponse{}
}

func (r DestinationLifecycleOnCreatedRequest) Clone() DestinationLifecycleOnCreatedRequest {
	return DestinationLifecycleOnCreatedRequest{
		Config: maps.Clone(r.Config),
	}
}

func (r DestinationLifecycleOnCreatedResponse) Clone() DestinationLifecycleOnCreatedResponse {
	return DestinationLifecycleOnCreatedResponse{}
}

func (r DestinationLifecycleOnUpdatedRequest) Clone() DestinationLifecycleOnUpdatedRequest {
	return DestinationLifecycleOnUpdatedRequest{
		ConfigBefore: maps.Clone(r.ConfigBefore),
		ConfigAfter:  maps.Clone(r.ConfigAfter),
	}
}

func (r DestinationLifecycleOnUpdatedResponse) Clone() DestinationLifecycleOnUpdatedResponse {
	return DestinationLifecycleOnUpdatedResponse{}
}

func (r DestinationLifecycleOnDeletedRequest) Clone() DestinationLifecycleOnDeletedRequest {
	return DestinationLifecycleOnDeletedRequest{
		Config: maps.Clone(r.Config),
	}
}

func (r DestinationLifecycleOnDeletedResponse) Clone() DestinationLifecycleOnDeletedResponse {
	return DestinationLifecycleOnDeletedResponse{}
}

func cloneSlice[T any](s []T, clone func(T) T) []T {
	if s == nil {
		return nil
	}
	cloned := make([]T, len(s))
	for i, v := range s {
		cloned[i] = clone(v)
	}
	return cloned
}
