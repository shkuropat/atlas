// Copyright 2020 The Atlas Authors. All rights reserved.
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

package rpc_context

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/journal/base"
)

// RPCContext
type RPCContext struct {
	metadata *atlas.Metadata
	claims   jwt.MapClaims
	journal  base.Journaller
}

// New
func New() *RPCContext {
	return &RPCContext{
		metadata: atlas.NewMetadata().SetRandomUUID().CreateTimestamp(),
	}
}

// SetClaims
func (c *RPCContext) SetClaims(claims jwt.MapClaims) *RPCContext {
	if c == nil {
		return nil
	}
	c.claims = claims
	return c
}

// GetJournal
func (c *RPCContext) GetJournal() base.Journaller {
	if c == nil {
		return nil
	}
	return c.journal
}

// SetJournal
func (c *RPCContext) SetJournal(j base.Journaller) *RPCContext {
	if c == nil {
		return nil
	}
	c.journal = j
	return c
}

// SetType
func (c *RPCContext) SetType(_type int32) *RPCContext {
	if c == nil {
		return nil
	}
	c.metadata.SetType(_type)
	return c
}

// GetType
func (c *RPCContext) GetType() int32 {
	if c == nil {
		return 0
	}
	return c.metadata.GetType()
}

// SetName
func (c *RPCContext) SetName(name string) *RPCContext {
	if c == nil {
		return nil
	}
	c.metadata.SetName(name)
	return c
}

// GetName
func (c *RPCContext) GetName() string {
	if c == nil {
		return ""
	}
	return c.metadata.GetName()
}

// SetUUID
func (c *RPCContext) SetUUID(id *atlas.UUID) *RPCContext {
	if c == nil {
		return nil
	}
	c.metadata.SetUUID(id)
	return c
}

// SetCallUUIDFromString
func (c *RPCContext) SetCallUUIDFromString(id string) *RPCContext {
	if c == nil {
		return nil
	}
	c.metadata.SetUUIDFromString(id)
	return c
}

// GetUUID
func (c *RPCContext) GetUUID() *atlas.UUID {
	if c == nil {
		return nil
	}
	return c.metadata.GetUUID()
}

// GetUUIDAsString
func (c *RPCContext) GetUUIDAsString() string {
	if c == nil {
		return ""
	}
	return c.metadata.GetUUID().String()
}
