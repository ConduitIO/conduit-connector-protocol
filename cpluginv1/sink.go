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

package cpluginv1

import (
	"context"
)

type SinkPluginServer interface {
	Configure(context.Context, SinkConfigureRequest) (SinkConfigureResponse, error)
	Start(context.Context, SinkStartRequest) (SinkStartResponse, error)
	Run(context.Context, SinkRunRequest) (SinkRunResponse, error)
	Stop(context.Context, SinkStopRequest) (SinkStopResponse, error)
}

type SinkConfigureRequest struct {
	Config map[string]string
}
type SinkConfigureResponse struct{}

type SinkStartRequest struct{}
type SinkStartResponse struct{}

type SinkRunRequest struct{}
type SinkRunResponse struct{}

type SinkStopRequest struct{}
type SinkStopResponse struct{}
