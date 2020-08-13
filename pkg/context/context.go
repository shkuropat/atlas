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

package context

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// Context
type Context struct {
	Metadata *atlas.Metadata
	Claims   jwt.MapClaims
}

// NewContext
func NewContext() *Context {
	return &Context{
		Metadata: atlas.NewMetadata().CreateID().CreateTimestamp(),
	}
}

// SetClaims
func (c *Context) SetClaims(claims jwt.MapClaims) *Context {
	c.Claims = claims
	return c
}

// SetType
func (c *Context) SetType(_type atlas.MetadataType) *Context {
	c.Metadata.SetType(_type)
	return c
}

// GetType
func (c *Context) GetType() int32 {
	return c.Metadata.GetType()
}

// SetName
func (c *Context) SetName(name string) *Context {
	c.Metadata.SetName(name)
	return c
}

// GetName
func (c *Context) GetName() string {
	return c.Metadata.GetName()
}

// SetID
func (c *Context) SetID(id *atlas.UUID) *Context {
	c.Metadata.SetID(id)
	return c
}

// SetIDFromString
func (c *Context) SetCallIDFromString(id string) *Context {
	c.Metadata.SetIDFromString(id)
	return c
}

// GetID
func (c *Context) GetID() *atlas.UUID {
	return c.Metadata.GetId()
}

// GetIDAsString
func (c *Context) GetIDAsString() string {
	return c.Metadata.GetId().GetString()
}
