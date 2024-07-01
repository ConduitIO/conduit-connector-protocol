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

package pschema

import (
	"context"
	"fmt"
	"sync"

	"github.com/conduitio/conduit-commons/schema"
)

type inMemoryService struct {
	// schemas is a map of schema subjects to all the versions of that schema
	// versioning starts at 1, newer versions are appended to the end of the versions slice.
	schemas map[string][]schema.Schema
	// m guards access to schemas
	m sync.Mutex
}

func NewInMemoryService() Service {
	return &inMemoryService{
		schemas: make(map[string][]schema.Schema),
	}
}

func (s *inMemoryService) Create(_ context.Context, request CreateRequest) (CreateResponse, error) {
	s.m.Lock()
	defer s.m.Unlock()

	inst := schema.Schema{
		Subject: request.Subject,
		Version: len(s.schemas[request.Subject]) + 1,
		Type:    schema.Type(request.Type),
		Bytes:   request.Bytes,
	}
	s.schemas[request.Subject] = append(s.schemas[request.Subject], inst)

	return CreateResponse{Schema: inst}, nil
}

func (s *inMemoryService) Get(_ context.Context, request GetRequest) (GetResponse, error) {
	s.m.Lock()
	defer s.m.Unlock()

	versions, ok := s.schemas[request.Subject]
	if !ok {
		return GetResponse{}, fmt.Errorf("subject %v: %w", request.Subject, ErrSchemaNotFound)
	}

	if len(versions) < request.Version {
		return GetResponse{}, fmt.Errorf("version %v: %w", request.Version, ErrSchemaNotFound)
	}

	return GetResponse{Schema: versions[request.Version-1]}, nil
}
