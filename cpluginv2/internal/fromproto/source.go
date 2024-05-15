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

package fromproto

import (
	"github.com/conduitio/conduit-connector-protocol/cpluginv2"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
)

func SourceConfigureRequest(in *connectorv2.Source_Configure_Request) cpluginv2.SourceConfigureRequest {
	return cpluginv2.SourceConfigureRequest{
		Config: in.Config,
	}
}

func SourceStartRequest(in *connectorv2.Source_Start_Request) cpluginv2.SourceStartRequest {
	return cpluginv2.SourceStartRequest{
		Position: in.Position,
	}
}

func SourceRunRequest(in *connectorv2.Source_Run_Request) cpluginv2.SourceRunRequest {
	return cpluginv2.SourceRunRequest{
		AckPosition: in.AckPosition,
	}
}

func SourceStopRequest(_ *connectorv2.Source_Stop_Request) cpluginv2.SourceStopRequest {
	return cpluginv2.SourceStopRequest{}
}

func SourceTeardownRequest(_ *connectorv2.Source_Teardown_Request) cpluginv2.SourceTeardownRequest {
	return cpluginv2.SourceTeardownRequest{}
}

func SourceLifecycleOnCreatedRequest(in *connectorv2.Source_Lifecycle_OnCreated_Request) cpluginv2.SourceLifecycleOnCreatedRequest {
	return cpluginv2.SourceLifecycleOnCreatedRequest{
		Config: in.Config,
	}
}
func SourceLifecycleOnUpdatedRequest(in *connectorv2.Source_Lifecycle_OnUpdated_Request) cpluginv2.SourceLifecycleOnUpdatedRequest {
	return cpluginv2.SourceLifecycleOnUpdatedRequest{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}
func SourceLifecycleOnDeletedRequest(in *connectorv2.Source_Lifecycle_OnDeleted_Request) cpluginv2.SourceLifecycleOnDeletedRequest {
	return cpluginv2.SourceLifecycleOnDeletedRequest{
		Config: in.Config,
	}
}
