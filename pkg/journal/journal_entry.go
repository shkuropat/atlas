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

// Entry defines journal entry structure
type Entry struct {
	// Base info tells about the origin of the action
	Endpoint  EndpointIDType
	SourceID  *atlas.UserID
	ContextID *atlas.UUID
	Action    ActionType

	// Object info tells about object, if any
	ObjectType     ObjectType
	ObjectAddress  *atlas.S3Address
	ObjectSize     uint64
	ObjectMetadata *atlas.Metadata
	ObjectData     []byte

	// Error info tells about error, if any
	Error error
}

// NewEntry
func NewEntry() *Entry {
	return &Entry{}
}

// SetBaseInfo
func (e *Entry) SetBaseInfo(ctxID *atlas.UUID, action ActionType) *Entry {
	e.SetCtxID(ctxID)
	e.SetAction(action)
	return e
}

// SetSourceID
func (e *Entry) SetSourceID(userID *atlas.UserID) *Entry {
	e.SourceID = userID
	return e
}

// SetCtxID
func (e *Entry) SetCtxID(ctxID *atlas.UUID) *Entry {
	e.ContextID = ctxID
	return e
}

// SetAction
func (e *Entry) SetAction(action ActionType) *Entry {
	e.Action = action
	return e
}

// SetObject
func (e *Entry) SetObject(
	objectType ObjectType,
	address *atlas.S3Address,
	size uint64,
	metadata *atlas.Metadata,
	data []byte,
) *Entry {
	e.SetObjectType(objectType)
	e.SetObjectAddress(address)
	e.SetObjectSize(size)
	e.SetObjectMetadata(metadata)
	e.SetObjectData(data)
	return e
}

// SetObjectType
func (e *Entry) SetObjectType(objectType ObjectType) *Entry {
	e.ObjectType = objectType
	return e
}

// SetObjectAddress
func (e *Entry) SetObjectAddress(address *atlas.S3Address) *Entry {
	e.ObjectAddress = address
	return e
}

// SetObjectSize
func (e *Entry) SetObjectSize(size uint64) *Entry {
	e.ObjectSize = size
	return e
}

// SetObjectMetadata
func (e *Entry) SetObjectMetadata(metadata *atlas.Metadata) *Entry {
	e.ObjectMetadata = metadata
	return e
}

// SetObjectData
func (e *Entry) SetObjectData(data []byte) *Entry {
	e.ObjectData = data
	return e
}

// SetError
func (e *Entry) SetError(err error) *Entry {
	e.Error = err
	return e
}

// InsertInto
func (e *Entry) InsertInto(j Journaller) {
	j.Insert(e)
}
