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

// Level represents a log level.
type Level int32

const (
	// NoLevel is a special level used to indicate that no level has been
	// set and allow for a default to be used.
	NoLevel Level = 0

	// Trace is the most verbose level. Intended to be used for the tracing
	// of actions in code, such as function enters/exits, etc.
	Trace Level = 1

	// Debug information for programmer low-level analysis.
	Debug Level = 2

	// Info information about steady state operations.
	Info Level = 3

	// Warn information about rare but handled events.
	Warn Level = 4

	// Error information about unrecoverable events.
	Error Level = 5

	// Off disables all logging output.
	Off Level = 6
)

type PluginConfig struct {
	Token       string
	ConnectorID string
	Level       Level
}
