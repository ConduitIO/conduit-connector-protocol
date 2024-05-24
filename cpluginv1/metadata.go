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

package cpluginv1

const (
	// OpenCDCVersion is a constant that should be used as the value in the
	// metadata field MetadataVersion. It ensures the OpenCDC format version can
	// be easily identified in case the record gets marshaled into a different
	// untyped format (e.g. JSON).
	OpenCDCVersion = "v1"

	// MetadataOpenCDCVersion is a Record.Metadata key for the version of the
	// OpenCDC format (e.g. "v1"). This field exists to ensure the OpenCDC
	// format version can be easily identified in case the record gets marshaled
	// into a different untyped format (e.g. JSON).
	MetadataOpenCDCVersion = "opencdc.version"
	// MetadataCreatedAt is a Record.Metadata key for the time when the record
	// was created in the 3rd party system. The expected format is a unix
	// timestamp in nanoseconds.
	MetadataCreatedAt = "opencdc.createdAt"
	// MetadataReadAt is a Record.Metadata key for the time when the record was
	// read from the 3rd party system. The expected format is a unix timestamp
	// in nanoseconds.
	MetadataReadAt = "opencdc.readAt"
	// MetadataCollection is a Record.Metadata key for the name of the collection
	// where the record originated from and/or where it should be stored.
	MetadataCollection = "opencdc.collection"

	// MetadataConduitSourcePluginName is a Record.Metadata key for the name of
	// the source plugin that created this record.
	MetadataConduitSourcePluginName = "conduit.source.plugin.name"
	// MetadataConduitSourcePluginVersion is a Record.Metadata key for the
	// version of the source plugin that created this record.
	MetadataConduitSourcePluginVersion = "conduit.source.plugin.version"
	// MetadataConduitDestinationPluginName is a Record.Metadata key for the
	// name of the destination plugin that has written this record
	// (only available in records once they are written by a destination).
	MetadataConduitDestinationPluginName = "conduit.destination.plugin.name"
	// MetadataConduitDestinationPluginVersion is a Record.Metadata key for the
	// version of the destination plugin that has written this record
	// (only available in records once they are written by a destination).
	MetadataConduitDestinationPluginVersion = "conduit.destination.plugin.version"

	// MetadataConduitSourceConnectorID is a Record.Metadata key for the ID of
	// the source connector that produced this record.
	MetadataConduitSourceConnectorID = "conduit.source.connector.id"
	// MetadataConduitDLQNackError is a Record.Metadata key for the error that
	// caused a record to be nacked and pushed to the dead-letter queue.
	MetadataConduitDLQNackError = "conduit.dlq.nack.error"
	// MetadataConduitDLQNackNodeID is a Record.Metadata key for the ID of the
	// internal node that nacked the record.
	MetadataConduitDLQNackNodeID = "conduit.dlq.nack.node.id"
)
