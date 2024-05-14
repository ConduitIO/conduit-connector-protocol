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

//go:generate mockgen -destination=mock/source.go -package=mock -mock_names=SourcePlugin=SourcePlugin,SourceRunStream=SourceRunStream . SourcePlugin,SourceRunStream

package cpluginv2

import (
	"context"

	"github.com/conduitio/conduit-commons/opencdc"
)

type SourcePlugin interface {
	Configure(context.Context, SourceConfigureRequest) (SourceConfigureResponse, error)
	Start(context.Context, SourceStartRequest) (SourceStartResponse, error)
	Run(context.Context, SourceRunStream) error
	Stop(context.Context, SourceStopRequest) (SourceStopResponse, error)
	Teardown(context.Context, SourceTeardownRequest) (SourceTeardownResponse, error)

	LifecycleOnCreated(context.Context, SourceLifecycleOnCreatedRequest) (SourceLifecycleOnCreatedResponse, error)
	LifecycleOnUpdated(context.Context, SourceLifecycleOnUpdatedRequest) (SourceLifecycleOnUpdatedResponse, error)
	LifecycleOnDeleted(context.Context, SourceLifecycleOnDeletedRequest) (SourceLifecycleOnDeletedResponse, error)
}

type SourceConfigureRequest struct {
	Config map[string]string
}
type SourceConfigureResponse struct{}

type SourceStartRequest struct {
	Position []byte
}
type SourceStartResponse struct{}

type SourceRunStream interface {
	Send(SourceRunResponse) error
	Recv() (SourceRunRequest, error)
}
type SourceRunRequest struct {
	AckPosition []byte
}
type SourceRunResponse struct {
	Record opencdc.Record
}

type SourceStopRequest struct{}
type SourceStopResponse struct {
	LastPosition []byte
}

type SourceTeardownRequest struct{}
type SourceTeardownResponse struct{}

type SourceLifecycleOnCreatedRequest struct {
	Config map[string]string
}
type SourceLifecycleOnCreatedResponse struct{}

type SourceLifecycleOnUpdatedRequest struct {
	ConfigBefore map[string]string
	ConfigAfter  map[string]string
}
type SourceLifecycleOnUpdatedResponse struct{}

type SourceLifecycleOnDeletedRequest struct {
	Config map[string]string
}
type SourceLifecycleOnDeletedResponse struct{}
