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

package internal

import (
	"context"

	"github.com/conduitio/conduit-connector-protocol/pconduit"
	"google.golang.org/grpc/metadata"
)

const MetadataConnectorTokenKey = "conduit-connector-token"

// RepackConnectorTokenOutgoingContext takes the connector token from ctx using
// pconduit.ConnectorTokenFromContext and repacks it into gRPC metadata so it can
// be sent to the server.
func RepackConnectorTokenOutgoingContext(ctx context.Context) context.Context {
	token := pconduit.ConnectorTokenFromContext(ctx)
	if token == "" {
		return ctx // no connector token attached
	}
	return metadata.AppendToOutgoingContext(ctx, MetadataConnectorTokenKey, token)
}

// RepackConnectorTokenIncomingContext takes the connector token from the gRPC
// metadata and repacks it into the context so it can be retrieved using
// pconduit.ConnectorTokenFromContext.
func RepackConnectorTokenIncomingContext(ctx context.Context) context.Context {
	token := metadata.ValueFromIncomingContext(ctx, MetadataConnectorTokenKey)
	if len(token) == 0 {
		return ctx // no connector token attached
	}
	return pconduit.ContextWithConnectorToken(ctx, token[0])
}
