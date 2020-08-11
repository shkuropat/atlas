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

package journal

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// CallMetadata
type CallMetadata struct {
	Internal *atlas.Metadata
	External jwt.MapClaims
}

// NewCallMetadata
func NewCallMetadata() *CallMetadata {
	return &CallMetadata{
		Internal: atlas.NewMetadata().CreateID().CreateTimestamp(),
	}
}

// SetExternal
func (c *CallMetadata) SetExternal(external jwt.MapClaims) *CallMetadata {
	c.External = external
	return c
}

// SetType
func (c *CallMetadata) SetType(_type atlas.MetadataType) *CallMetadata {
	c.Internal.SetType(_type)
	return c
}

// GetType
func (c *CallMetadata) GetType() int32 {
	return c.Internal.GetType()
}

// SetName
func (c *CallMetadata) SetName(name string) *CallMetadata {
	c.Internal.SetName(name)
	return c
}

// GetName
func (c *CallMetadata) GetName() string {
	return c.Internal.GetName()
}

// SetCallID
func (c *CallMetadata) SetCallID(callID *atlas.UUID) *CallMetadata {
	c.Internal.SetID(callID)
	return c
}

// SetCallIDFromString
func (c *CallMetadata) SetCallIDFromString(callID string) *CallMetadata {
	c.Internal.SetIDFromString(callID)
	return c
}

// GetCallID
func (c *CallMetadata) GetCallID() *atlas.UUID {
	return c.Internal.GetId()
}

// GetCallIDAsString
func (c *CallMetadata) GetCallIDAsString() string {
	return c.Internal.GetId().GetStringValue()
}
