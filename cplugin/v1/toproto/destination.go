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
	"github.com/conduitio/conduit-connector-protocol/cplugin"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
)

// -- Request Conversions -----------------------------------------------------

// TODO

// -- Response Conversions ----------------------------------------------------

func DestinationConfigureResponse(_ cplugin.DestinationConfigureResponse) *connectorv1.Destination_Configure_Response {
	return &connectorv1.Destination_Configure_Response{}
}

func DestinationStartResponse(_ cplugin.DestinationStartResponse) *connectorv1.Destination_Start_Response {
	return &connectorv1.Destination_Start_Response{}
}

func DestinationRunResponse(in cplugin.DestinationRunResponse) *connectorv1.Destination_Run_Response {
	return &connectorv1.Destination_Run_Response{
		AckPosition: in.AckPosition,
		Error:       in.Error,
	}
}

func DestinationStopResponse(_ cplugin.DestinationStopResponse) *connectorv1.Destination_Stop_Response {
	return &connectorv1.Destination_Stop_Response{}
}

func DestinationTeardownResponse(_ cplugin.DestinationTeardownResponse) *connectorv1.Destination_Teardown_Response {
	return &connectorv1.Destination_Teardown_Response{}
}

func DestinationLifecycleOnCreatedResponse(_ cplugin.DestinationLifecycleOnCreatedResponse) *connectorv1.Destination_Lifecycle_OnCreated_Response {
	return &connectorv1.Destination_Lifecycle_OnCreated_Response{}
}
func DestinationLifecycleOnUpdatedResponse(_ cplugin.DestinationLifecycleOnUpdatedResponse) *connectorv1.Destination_Lifecycle_OnUpdated_Response {
	return &connectorv1.Destination_Lifecycle_OnUpdated_Response{}
}
func DestinationLifecycleOnDeletedResponse(_ cplugin.DestinationLifecycleOnDeletedResponse) *connectorv1.Destination_Lifecycle_OnDeleted_Response {
	return &connectorv1.Destination_Lifecycle_OnDeleted_Response{}
}
