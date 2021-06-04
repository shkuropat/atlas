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
	"bytes"
	"fmt"
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

// String
func (e *JournalEntry) String() string {
	if e == nil {
		return "this JE is nil"
	}

	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Time:%s\n", e.Time)
	_, _ = fmt.Fprintf(b, "StartTime:%s\n", e.StartTime)

	_, _ = fmt.Fprintf(b, "EndpointID:%d", e.EndpointID)
	_, _ = fmt.Fprintf(b, "SourceID:%s\n", e.SourceID)
	_, _ = fmt.Fprintf(b, "ContextID:%s\n", e.ContextID)
	_, _ = fmt.Fprintf(b, "TaskID:%s\n", e.TaskID)
	_, _ = fmt.Fprintf(b, "Type:%d\n", e.Type)

	_, _ = fmt.Fprintf(b, "ObjectType:%d\n", e.ObjectType)
	_, _ = fmt.Fprintf(b, "ObjectAddress:%s\n", e.ObjectAddress)
	_, _ = fmt.Fprintf(b, "ObjectSize:%d\n", e.ObjectSize)
	_, _ = fmt.Fprintf(b, "ObjectMetadata:%s\n", e.ObjectMetadata)
	_, _ = fmt.Fprintf(b, "ObjectData:%s\n", e.ObjectData)

	_, _ = fmt.Fprintf(b, "Error:%s\n", e.Error)

	return b.String()
}

// NewJournalEntry
func NewJournalEntry() *JournalEntry {
	return &JournalEntry{}
}

// SetBaseInfo
func (e *JournalEntry) SetBaseInfo(start time.Time, endpoint int32, ctxID *atlas.UUID, _type int32) *JournalEntry {
	if e == nil {
		return nil
	}
	e.Time = time.Now()
	e.StartTime = start
	e.SetEndpointID(endpoint)
	e.SetCtxID(ctxID)
	e.SetType(_type)
	return e
}

// SetEndpointID
func (e *JournalEntry) SetEndpointID(endpoint int32) *JournalEntry {
	if e == nil {
		return nil
	}
	e.EndpointID = endpoint
	return e
}

// SetSourceID
func (e *JournalEntry) SetSourceID(userID *atlas.UserID) *JournalEntry {
	if e == nil {
		return nil
	}
	e.SourceID = userID
	return e
}

// SetCtxID
func (e *JournalEntry) SetCtxID(ctxID *atlas.UUID) *JournalEntry {
	if e == nil {
		return nil
	}
	e.ContextID = ctxID
	return e
}

// SetTaskID
func (e *JournalEntry) SetTaskID(taskID *atlas.UUID) *JournalEntry {
	if e == nil {
		return nil
	}
	e.TaskID = taskID
	return e
}

// SetType
func (e *JournalEntry) SetType(_type int32) *JournalEntry {
	if e == nil {
		return nil
	}
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
	if e == nil {
		return nil
	}
	e.SetObjectType(objectType)
	e.SetObjectAddress(address)
	e.SetObjectSize(size)
	e.SetObjectMetadata(metadata)
	e.SetObjectData(data)
	return e
}

// SetObjectType
func (e *JournalEntry) SetObjectType(objectType int32) *JournalEntry {
	if e == nil {
		return nil
	}
	e.ObjectType = objectType
	return e
}

// SetObjectAddress
func (e *JournalEntry) SetObjectAddress(address *atlas.Address) *JournalEntry {
	if e == nil {
		return nil
	}
	e.ObjectAddress = address
	return e
}

// SetObjectSize
func (e *JournalEntry) SetObjectSize(size uint64) *JournalEntry {
	if e == nil {
		return nil
	}
	e.ObjectSize = size
	return e
}

// SetObjectMetadata
func (e *JournalEntry) SetObjectMetadata(metadata *atlas.Metadata) *JournalEntry {
	if e == nil {
		return nil
	}
	e.ObjectMetadata = metadata
	return e
}

// EnsureObjectMetadata
func (e *JournalEntry) EnsureObjectMetadata() *atlas.Metadata {
	if e == nil {
		return nil
	}
	if e.ObjectMetadata == nil {
		e.ObjectMetadata = atlas.NewMetadata()
	}
	return e.ObjectMetadata
}

// SetObjectData
func (e *JournalEntry) SetObjectData(data []byte) *JournalEntry {
	if e == nil {
		return nil
	}
	e.ObjectData = data
	return e
}

// SetError
func (e *JournalEntry) SetError(err error) *JournalEntry {
	if e == nil {
		return nil
	}
	e.Error = err
	return e
}

// InsertInto inserts entry into a journal
func (e *JournalEntry) InsertInto(j Journaller) {
	j.Insert(e)
}
