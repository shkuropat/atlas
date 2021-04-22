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
	//
	// Base info tells about the origin of the journal entry
	//

	// Time of the entry/event
	Time time.Time
	// StartTime specifies start time of the action sequence or execution context
	StartTime time.Time

	// EndpointID [MANDATORY] specifies ID of the endpoint (API call handler/Task processor/etc) which produces the entry
	// See EndpointTypeEnum for available options.
	EndpointID int32
	// SourceID [OPTIONAL] specifies ID of the source (possibly external) of the entry
	SourceID *atlas.UserID
	// ContextID [OPTIONAL] specifies ID of the execution/rpc context associated with the entry
	ContextID *atlas.UUID
	// TaskID [OPTIONAL] specifies ID of the task associated with the entry
	TaskID *atlas.UUID
	// Type [MANDATORY] specifies type of the entry - what this entry is about.
	// See EntryTypeEnum for available options.
	Type int32

	// Object info tells about object, if any
	// ObjectType specified object type
	// See ObjectTypeEnum for available options.
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
func (e *JournalEntry) SetBaseInfo(start time.Time, endpoint int32, ctxID *atlas.UUID, _type int32) *JournalEntry {
	e.Time = time.Now()
	e.StartTime = start
	e.SetEndpointID(endpoint)
	e.SetCtxID(ctxID)
	e.SetType(_type)
	return e
}

// SetEndpointID
func (e *JournalEntry) SetEndpointID(endpoint int32) *JournalEntry {
	e.EndpointID = endpoint
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

// SetTaskID
func (e *JournalEntry) SetTaskID(taskID *atlas.UUID) *JournalEntry {
	e.TaskID = taskID
	return e
}

// SetType
func (e *JournalEntry) SetType(_type int32) *JournalEntry {
	e.Type = _type
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

// EnsureObjectMetadata
func (e *JournalEntry) EnsureObjectMetadata() *atlas.Metadata {
	if e.ObjectMetadata == nil {
		e.ObjectMetadata = atlas.NewMetadata()
	}
	return e.ObjectMetadata
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

// InsertInto inserts entry into a journal
func (e *JournalEntry) InsertInto(j Journaller) {
	j.Insert(e)
}
