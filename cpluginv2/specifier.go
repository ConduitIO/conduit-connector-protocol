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

//go:generate mockgen -destination=mock/specifier.go -package=mock -mock_names=SpecifierPlugin=SpecifierPlugin . SpecifierPlugin

package cpluginv2

import (
	"context"

	"github.com/conduitio/conduit-commons/config"
)

type SpecifierPlugin interface {
	Specify(context.Context, SpecifierSpecifyRequest) (SpecifierSpecifyResponse, error)
}

type SpecifierSpecifyRequest struct{}
type SpecifierSpecifyResponse struct {
	Name              string
	Summary           string
	Description       string
	Version           string
	Author            string
	DestinationParams config.Parameters
	SourceParams      config.Parameters
}
