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
	"github.com/conduitio/conduit-commons/opencdc"
	"github.com/conduitio/conduit-connector-protocol/cplugin"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
)

func DestinationConfigureRequest(in *connectorv2.Destination_Configure_Request) cplugin.DestinationConfigureRequest {
	return cplugin.DestinationConfigureRequest{
		Config: in.Config,
	}
}

func DestinationStartRequest(_ *connectorv2.Destination_Start_Request) cplugin.DestinationStartRequest {
	return cplugin.DestinationStartRequest{}
}

func DestinationRunRequest(in *connectorv2.Destination_Run_Request) (cplugin.DestinationRunRequest, error) {
	var rec opencdc.Record
	err := rec.FromProto(in.Record)
	if err != nil {
		return cplugin.DestinationRunRequest{}, err
	}
	return cplugin.DestinationRunRequest{
		Record: rec,
	}, nil
}

func DestinationStopRequest(in *connectorv2.Destination_Stop_Request) cplugin.DestinationStopRequest {
	return cplugin.DestinationStopRequest{
		LastPosition: in.LastPosition,
	}
}

func DestinationTeardownRequest(_ *connectorv2.Destination_Teardown_Request) cplugin.DestinationTeardownRequest {
	return cplugin.DestinationTeardownRequest{}
}

func DestinationLifecycleOnCreatedRequest(in *connectorv2.Destination_Lifecycle_OnCreated_Request) cplugin.DestinationLifecycleOnCreatedRequest {
	return cplugin.DestinationLifecycleOnCreatedRequest{
		Config: in.Config,
	}
}
func DestinationLifecycleOnUpdatedRequest(in *connectorv2.Destination_Lifecycle_OnUpdated_Request) cplugin.DestinationLifecycleOnUpdatedRequest {
	return cplugin.DestinationLifecycleOnUpdatedRequest{
		ConfigBefore: in.ConfigBefore,
		ConfigAfter:  in.ConfigAfter,
	}
}
func DestinationLifecycleOnDeletedRequest(in *connectorv2.Destination_Lifecycle_OnDeleted_Request) cplugin.DestinationLifecycleOnDeletedRequest {
	return cplugin.DestinationLifecycleOnDeletedRequest{
		Config: in.Config,
	}
}
