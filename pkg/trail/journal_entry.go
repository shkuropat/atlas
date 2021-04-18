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

package trail

import (
	"github.com/binarly-io/atlas/pkg/api/atlas"
	"time"
)

// JournalEntry defines journal entry structure
type JournalEntry struct {
	// Base info tells about the origin of the journal entry
	Time      time.Time
	Start     time.Time
	Endpoint  int32
	SourceID  *atlas.UserID
	ContextID *atlas.UUID
	Action    int32

	// Object info tells about object, if any
	ObjectType     int32
	ObjectAddress  *atlas.Address
	ObjectSize     uint64
	ObjectMetadata *atlas.Metadata
	ObjectData     []byte

	// Error info tells about error, if any
	Error error
}

// NewJournalEntry
func NewJournalEntry() *JournalEntry {
	return &JournalEntry{}
}

// SetBaseInfo
func (e *JournalEntry) SetBaseInfo(start time.Time, endpoint int32, ctxID *atlas.UUID, action int32) *JournalEntry {
	e.Time = time.Now()
	e.Start = start
	e.SetEndpoint(endpoint)
	e.SetCtxID(ctxID)
	e.SetAction(action)
	return e
}

// SetEndpoint
func (e *JournalEntry) SetEndpoint(endpoint int32) *JournalEntry {
	e.Endpoint = endpoint
	return e
}

// SetSourceID
func (e *JournalEntry) SetSourceID(userID *atlas.UserID) *JournalEntry {
	e.SourceID = userID
	return e
}

// SetCtxID
func (e *JournalEntry) SetCtxID(ctxID *atlas.UUID) *JournalEntry {
	e.ContextID = ctxID
	return e
}

// SetAction
func (e *JournalEntry) SetAction(action int32) *JournalEntry {
	e.Action = action
	return e
}

// SetObject
func (e *JournalEntry) SetObject(
	objectType int32,
	address *atlas.Address,
	size uint64,
	metadata *atlas.Metadata,
	data []byte,
) *JournalEntry {
	e.SetObjectType(objectType)
	e.SetObjectAddress(address)
	e.SetObjectSize(size)
	e.SetObjectMetadata(metadata)
	e.SetObjectData(data)
	return e
}

// SetObjectType
func (e *JournalEntry) SetObjectType(objectType int32) *JournalEntry {
	e.ObjectType = objectType
	return e
}

// SetObjectAddress
func (e *JournalEntry) SetObjectAddress(address *atlas.Address) *JournalEntry {
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

// InsertInto
func (e *JournalEntry) InsertInto(a Adapter) {
	a.Insert(e)
}
