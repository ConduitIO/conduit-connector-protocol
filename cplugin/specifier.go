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

	"github.com/conduitio/conduit-commons/config"
)

type SpecifierPlugin interface {
	Specify(context.Context, SpecifierSpecifyRequest) (SpecifierSpecifyResponse, error)
}

type SpecifierSpecifyRequest struct{}
type SpecifierSpecifyResponse struct {
	Specification Specification
}

// Specification is returned by a plugin when Specify is called.
// It contains information about the configuration parameters for plugins
// and allows them to describe their parameters.
type Specification struct {
	// Name is the name of the plugin.
	Name string
	// Summary is a brief description of the plugin and what it does.
	Summary string
	// Description is a more long form area appropriate for README-like text
	// that the author can provide for documentation about the specified
	// Parameters.
	Description string
	// Version string. Should be a semver prepended with `v`, e.g. `v1.54.3`.
	Version string
	// Author declares the entity that created or maintains this plugin.
	Author string
	// SourceParams and DestinationParams are maps of named Parameters that
	// describe how to configure the plugins Destination or Source.
	SourceParams      config.Parameters
	DestinationParams config.Parameters
}
