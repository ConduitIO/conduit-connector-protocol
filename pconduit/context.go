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

package pconduit

import (
	"context"
)

// -- Connector token ----------------------------------------------------------

type connectorTokenCtxKey struct{}

// ContextWithConnectorToken wraps ctx and returns a context that contains the connector token.
func ContextWithConnectorToken(ctx context.Context, connectorToken string) context.Context {
	return context.WithValue(ctx, connectorTokenCtxKey{}, connectorToken)
}

// ConnectorTokenFromContext fetches the connector token from the context. If the
// context does not contain a connector token it returns an empty string.
func ConnectorTokenFromContext(ctx context.Context) string {
	connectorToken := ctx.Value(connectorTokenCtxKey{})
	if connectorToken != nil {
		return connectorToken.(string) //nolint:forcetypeassert // only this package can set the value, it has to be a string
	}
	return ""
}
