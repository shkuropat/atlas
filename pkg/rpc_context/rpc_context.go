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
)

// RPCContext
type RPCContext struct {
	Metadata *atlas.Metadata
	Claims   jwt.MapClaims
}

// New
func New() *RPCContext {
	return &RPCContext{
		Metadata: atlas.NewMetadata().CreateID().CreateTimestamp(),
	}
}

// SetClaims
func (c *RPCContext) SetClaims(claims jwt.MapClaims) *RPCContext {
	c.Claims = claims
	return c
}

// SetType
func (c *RPCContext) SetType(_type atlas.MetadataType) *RPCContext {
	c.Metadata.SetType(int32(_type))
	return c
}

// GetType
func (c *RPCContext) GetType() int32 {
	return c.Metadata.GetType()
}

// SetName
func (c *RPCContext) SetName(name string) *RPCContext {
	c.Metadata.SetName(name)
	return c
}

// GetName
func (c *RPCContext) GetName() string {
	return c.Metadata.GetName()
}

// SetID
func (c *RPCContext) SetID(id *atlas.UUID) *RPCContext {
	c.Metadata.SetID(id)
	return c
}

// SetIDFromString
func (c *RPCContext) SetCallIDFromString(id string) *RPCContext {
	c.Metadata.SetIDFromString(id)
	return c
}

// GetID
func (c *RPCContext) GetID() *atlas.UUID {
	return c.Metadata.GetId()
}

// GetIDAsString
func (c *RPCContext) GetIDAsString() string {
	return c.Metadata.GetId().GetString()
}
