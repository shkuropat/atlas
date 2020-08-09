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
	"github.com/binarly-io/atlas/pkg/api/atlas"
)

type JournalEntry struct {
	Endpoint       uint16
	Source         *atlas.UUID
	Call           *atlas.UUID
	Action         ActionType
	ObjectType     uint8
	ObjectAddress  *atlas.S3Address
	ObjectSize     uint64
	ObjectMetadata *atlas.Metadata
	ObjectData     []byte
	Error          error
}

// NewJournalEntry
func NewJournalEntry() *JournalEntry {
	return &JournalEntry{}
}

// SetCallAction
func (e *JournalEntry) SetCallAction(call *atlas.UUID, action ActionType) *JournalEntry {
	e.SetCall(call)
	e.SetAction(action)
	return e
}

// SetCall
func (e *JournalEntry) SetCall(call *atlas.UUID) *JournalEntry {
	e.Call = call
	return e
}

// SetAction
func (e *JournalEntry) SetAction(action ActionType) *JournalEntry {
	e.Action = action
	return e
}

// SetObject
func (e *JournalEntry) SetObject(_type uint8, address *atlas.S3Address, size uint64, metadata *atlas.Metadata) *JournalEntry {
	e.SetObjectType(_type)
	e.SetObjectAddress(address)
	e.SetObjectSize(size)
	e.SetObjectMetadata(metadata)
	return e
}

// SetObjectType
func (e *JournalEntry) SetObjectType(_type uint8) *JournalEntry {
	e.ObjectType = _type
	return e
}

// SetObjectAddress
func (e *JournalEntry) SetObjectAddress(address *atlas.S3Address) *JournalEntry {
	e.ObjectAddress = address
	return e
}

// SetObjectSize
func (e *JournalEntry) SetObjectSize(size uint64) *JournalEntry {
	e.ObjectSize = size
	return e
}

// SetObjectMetadata
func (e *JournalEntry) SetObjectMetadata(metadata *atlas.Metadata) *JournalEntry {
	e.ObjectMetadata = metadata
	return e
}

// SetObjectData
func (e *JournalEntry) SetObjectData(data []byte) *JournalEntry {
	e.ObjectData = data
	return e
}

// SetError
func (e *JournalEntry) SetError(err error) *JournalEntry {
	e.Error = err
	return e
}
