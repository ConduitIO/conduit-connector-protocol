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
	"github.com/conduitio/conduit-connector-protocol/cpluginv1"
	connectorv1 "github.com/conduitio/conduit-connector-protocol/proto/connector/v1"
)

func SourceConfigureResponse(_ cpluginv1.SourceConfigureResponse) (*connectorv1.Source_Configure_Response, error) {
	return &connectorv1.Source_Configure_Response{}, nil
}

func SourceStartResponse(_ cpluginv1.SourceStartResponse) (*connectorv1.Source_Start_Response, error) {
	return &connectorv1.Source_Start_Response{}, nil
}

func SourceRunResponse(in cpluginv1.SourceRunResponse) (*connectorv1.Source_Run_Response, error) {
	rec, err := Record(in.Record)
	if err != nil {
		return nil, err
	}

	out := connectorv1.Source_Run_Response{
		Record: rec,
	}
	return &out, nil
}

func SourceStopResponse(in cpluginv1.SourceStopResponse) (*connectorv1.Source_Stop_Response, error) {
	out := connectorv1.Source_Stop_Response{
		LastPosition: in.LastPosition,
	}
	return &out, nil
}

func SourceTeardownResponse(_ cpluginv1.SourceTeardownResponse) (*connectorv1.Source_Teardown_Response, error) {
	return &connectorv1.Source_Teardown_Response{}, nil
}

func SourceLifecycleOnCreatedResponse(_ cpluginv1.SourceLifecycleOnCreatedResponse) (*connectorv1.Source_Lifecycle_OnCreated_Response, error) {
	return &connectorv1.Source_Lifecycle_OnCreated_Response{}, nil
}
func SourceLifecycleOnUpdatedResponse(_ cpluginv1.SourceLifecycleOnUpdatedResponse) (*connectorv1.Source_Lifecycle_OnUpdated_Response, error) {
	return &connectorv1.Source_Lifecycle_OnUpdated_Response{}, nil
}
func SourceLifecycleOnDeletedResponse(_ cpluginv1.SourceLifecycleOnDeletedResponse) (*connectorv1.Source_Lifecycle_OnDeleted_Response, error) {
	return &connectorv1.Source_Lifecycle_OnDeleted_Response{}, nil
}
