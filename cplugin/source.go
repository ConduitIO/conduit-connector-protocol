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

package cplugin

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

// SourceRunStream is the bidirectional stream interface for SourcePlugin.Run.
// It combines the client and server interfaces into a single interface.
type SourceRunStream interface {
	// Client is only allowed to be used by the host (Conduit).
	Client() SourceRunStreamClient
	// Server is only allowed to be used by the plugin (connector).
	Server() SourceRunStreamServer
}

// SourceRunStreamClient is the client-side interface for a bidirectional stream
// of SourceRunRequest and SourceRunResponse messages.
type SourceRunStreamClient interface {
	Send(SourceRunRequest) error
	Recv() (SourceRunResponse, error)
}

// SourceRunStreamServer is the server-side interface for a bidirectional stream
// of SourceRunRequest and SourceRunResponse messages.
type SourceRunStreamServer interface {
	Send(SourceRunResponse) error
	Recv() (SourceRunRequest, error)
}

type SourceConfigureRequest struct {
	Config map[string]string
}
type SourceConfigureResponse struct{}

type SourceStartRequest struct {
	Position opencdc.Position
}
type SourceStartResponse struct{}

type SourceRunRequest struct {
	AckPositions []opencdc.Position
}
type SourceRunResponse struct {
	Records []opencdc.Record
}

type SourceStopRequest struct{}
type SourceStopResponse struct {
	LastPosition opencdc.Position
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
