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
	"errors"
	"fmt"

	"github.com/conduitio/conduit-connector-protocol/pconnector"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// knownErrors contains known error messages that are mapped to internal error
// types. gRPC does not retain error types, so we have to resort to relying on
// the error message itself.
var knownErrors = map[string]error{
	"context canceled":          context.Canceled,
	"context deadline exceeded": context.DeadlineExceeded,
}

// UnwrapGRPCError removes the gRPC wrapper from the error and returns a known
// error if possible, otherwise creates an internal error.
func UnwrapGRPCError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	if st.Code() == codes.Unimplemented {
		return fmt.Errorf("%s: %w", st.Message(), pconnector.ErrUnimplemented)
	}
	if knownErr, ok := knownErrors[st.Message()]; ok {
		return knownErr
	}
	return errors.New(st.Message())
}
