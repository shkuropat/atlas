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
	"github.com/binarly-io/atlas/pkg/journal"
)

// RPCContext
type RPCContext struct {
	metadata *atlas.Metadata
	claims   jwt.Claims
	journal  journal.Journaller
}

// New
func New() *RPCContext {
	return &RPCContext{
		metadata: atlas.NewMetadata().SetRandomUUID().CreateTimestamp(),
	}
}

// GetClaims
func (c *RPCContext) GetClaims() jwt.Claims {
	if c == nil {
		return nil
	}
	return c.claims
}

// SetClaims
func (c *RPCContext) SetClaims(claims jwt.Claims) *RPCContext {
	if c == nil {
		return nil
	}
	c.claims = claims
	return c
}

// GetJournal
func (c *RPCContext) GetJournal() journal.Journaller {
	if c == nil {
		return nil
	}
	return c.journal
}

// SetJournal
func (c *RPCContext) SetJournal(j journal.Journaller) *RPCContext {
	if c == nil {
		return nil
	}
	c.journal = j
	return c
}

// GetType
func (c *RPCContext) GetType() int32 {
	if c == nil {
		return 0
	}
	return c.metadata.GetType()
}

// SetType
func (c *RPCContext) SetType(_type int32) *RPCContext {
	if c == nil {
		return nil
	}
	c.metadata.SetType(_type)
	return c
}

// GetName
func (c *RPCContext) GetName() string {
	if c == nil {
		return ""
	}
	return c.metadata.GetName()
}

// SetName
func (c *RPCContext) SetName(name string) *RPCContext {
	if c == nil {
		return nil
	}
	c.metadata.SetName(name)
	return c
}

// GetUUID
func (c *RPCContext) GetUUID() *atlas.UUID {
	if c == nil {
		return nil
	}
	return c.metadata.GetUUID()
}

// SetUUID
func (c *RPCContext) SetUUID(id *atlas.UUID) *RPCContext {
	if c == nil {
		return nil
	}
	c.metadata.SetUUID(id)
	return c
}

// GetUUIDAsString
func (c *RPCContext) GetUUIDAsString() string {
	if c == nil {
		return ""
	}
	return c.metadata.GetUUID().String()
}

// SetCallUUIDFromString
func (c *RPCContext) SetCallUUIDFromString(id string) *RPCContext {
	if c == nil {
		return nil
	}
	c.metadata.SetUUIDFromString(id)
	return c
}
