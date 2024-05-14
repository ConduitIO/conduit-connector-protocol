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

package toproto

import (
	"github.com/conduitio/conduit-connector-protocol/cpluginv2"
	connectorv2 "github.com/conduitio/conduit-connector-protocol/proto/connector/v2"
)

func DestinationConfigureResponse(_ cpluginv2.DestinationConfigureResponse) (*connectorv2.Destination_Configure_Response, error) {
	return &connectorv2.Destination_Configure_Response{}, nil
}

func DestinationStartResponse(_ cpluginv2.DestinationStartResponse) (*connectorv2.Destination_Start_Response, error) {
	return &connectorv2.Destination_Start_Response{}, nil
}

func DestinationRunResponse(in cpluginv2.DestinationRunResponse) (*connectorv2.Destination_Run_Response, error) {
	out := connectorv2.Destination_Run_Response{
		AckPosition: in.AckPosition,
		Error:       in.Error,
	}
	return &out, nil
}

func DestinationStopResponse(_ cpluginv2.DestinationStopResponse) (*connectorv2.Destination_Stop_Response, error) {
	return &connectorv2.Destination_Stop_Response{}, nil
}

func DestinationTeardownResponse(_ cpluginv2.DestinationTeardownResponse) (*connectorv2.Destination_Teardown_Response, error) {
	return &connectorv2.Destination_Teardown_Response{}, nil
}

func DestinationLifecycleOnCreatedResponse(_ cpluginv2.DestinationLifecycleOnCreatedResponse) (*connectorv2.Destination_Lifecycle_OnCreated_Response, error) {
	return &connectorv2.Destination_Lifecycle_OnCreated_Response{}, nil
}
func DestinationLifecycleOnUpdatedResponse(_ cpluginv2.DestinationLifecycleOnUpdatedResponse) (*connectorv2.Destination_Lifecycle_OnUpdated_Response, error) {
	return &connectorv2.Destination_Lifecycle_OnUpdated_Response{}, nil
}
func DestinationLifecycleOnDeletedResponse(_ cpluginv2.DestinationLifecycleOnDeletedResponse) (*connectorv2.Destination_Lifecycle_OnDeleted_Response, error) {
	return &connectorv2.Destination_Lifecycle_OnDeleted_Response{}, nil
}
