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
	"context"

	"github.com/conduitio/conduit-commons/config"
	"github.com/conduitio/conduit-commons/opencdc"
)

//go:generate mockgen -destination=mock/destination.go -package=mock -mock_names=DestinationPlugin=DestinationPlugin . DestinationPlugin

type DestinationPlugin interface {
	Configure(context.Context, DestinationConfigureRequest) (DestinationConfigureResponse, error)
	Open(context.Context, DestinationOpenRequest) (DestinationOpenResponse, error)
	Run(context.Context, DestinationRunStream) error
	Stop(context.Context, DestinationStopRequest) (DestinationStopResponse, error)
	Teardown(context.Context, DestinationTeardownRequest) (DestinationTeardownResponse, error)

	LifecycleOnCreated(context.Context, DestinationLifecycleOnCreatedRequest) (DestinationLifecycleOnCreatedResponse, error)
	LifecycleOnUpdated(context.Context, DestinationLifecycleOnUpdatedRequest) (DestinationLifecycleOnUpdatedResponse, error)
	LifecycleOnDeleted(context.Context, DestinationLifecycleOnDeletedRequest) (DestinationLifecycleOnDeletedResponse, error)
}

// DestinationRunStream is the bidirectional stream interface for DestinationPlugin.Run.
// It combines the client and server interfaces into a single interface.
type DestinationRunStream interface {
	// Client is only allowed to be used by the host (Conduit).
	Client() DestinationRunStreamClient
	// Server is only allowed to be used by the plugin (connector).
	Server() DestinationRunStreamServer
}

// DestinationRunStreamClient is the client-side interface for a bidirectional stream
// of DestinationRunRequest and DestinationRunResponse messages.
type DestinationRunStreamClient interface {
	Send(DestinationRunRequest) error
	Recv() (DestinationRunResponse, error)
}

// DestinationRunStreamServer is the server-side interface for a bidirectional stream
// of DestinationRunRequest and DestinationRunResponse messages.
type DestinationRunStreamServer interface {
	Send(DestinationRunResponse) error
	Recv() (DestinationRunRequest, error)
}

type DestinationConfigureRequest struct {
	Config config.Config
}
type DestinationConfigureResponse struct{}

type (
	DestinationOpenRequest  struct{}
	DestinationOpenResponse struct{}
)

type DestinationRunRequest struct {
	Records []opencdc.Record
}
type DestinationRunResponse struct {
	Acks []DestinationRunResponseAck
}

type DestinationRunResponseAck struct {
	Position opencdc.Position
	Error    string
}

type DestinationStopRequest struct {
	LastPosition opencdc.Position
}
type DestinationStopResponse struct{}

type (
	DestinationTeardownRequest  struct{}
	DestinationTeardownResponse struct{}
)

type DestinationLifecycleOnCreatedRequest struct {
	Config config.Config
}
type DestinationLifecycleOnCreatedResponse struct{}

type DestinationLifecycleOnUpdatedRequest struct {
	ConfigBefore config.Config
	ConfigAfter  config.Config
}
type DestinationLifecycleOnUpdatedResponse struct{}

type DestinationLifecycleOnDeletedRequest struct {
	Config config.Config
}
type DestinationLifecycleOnDeletedResponse struct{}
