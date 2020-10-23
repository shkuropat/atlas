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
	"strings"
	"time"
)

// ClickHouseEntry defines journal entry structure
type ClickHouseEntry struct {
	// Call section
	d          time.Time
	endpointID uint16
	sourceID   string
	contextID  string
	actionID   uint8
	duration   uint64
	// Object section
	_type   uint8
	size    uint64
	address string
	domain  string
	name    string
	digest  string
	data    string
	// Error section
	error string
}

// NewClickHouseEntry
func NewClickHouseEntry() *ClickHouseEntry {
	return &ClickHouseEntry{}
}

// Accept
func (ce *ClickHouseEntry) Accept(j *JournalClickHouse, entry *Entry) *ClickHouseEntry {
	ce.d = time.Now()
	ce.endpointID = uint16(j.endpointID)
	ce.sourceID = entry.SourceID.GetString()
	ce.contextID = entry.ContextID.GetString()
	ce.actionID = uint8(entry.Action)
	ce.duration = uint64(ce.d.Sub(j.start).Nanoseconds())
	ce._type = uint8(entry.ObjectType)
	ce.size = entry.ObjectSize
	ce.address = entry.ObjectAddress.Printable()
	ce.domain = entry.ObjectMetadata.GetDomain().GetName()
	ce.name = entry.ObjectMetadata.GetFilename()
	ce.digest = string(entry.ObjectMetadata.GetDigest().GetData())
	ce.data = string(entry.ObjectData)
	if entry.Error != nil {
		ce.error = entry.Error.Error()
	}

	return ce
}

// Fields
func (ce *ClickHouseEntry) Fields() string {
	return `
		d, 
		endpoint_id,
		source_id,
		context_id,
		action_id,
		duration,
		type, 
		size,
		address,
		domain,
		name,
		digest,
		data, 
		error
	`
}

// StmtParamsPlaceholder
func (ce *ClickHouseEntry) StmtParamsPlaceholder() string {
	return `
		/* d */
		?,
		/* endpoint_id */
		?,
		/* source_id */
		?,
		/* context_id */
		?,
		/* action_id */
		?,
		/* duration */
		?,
		/* type */
		?,
		/* size */
		?,
		/* address */
		?,
		/* domain */
		?,
		/* name */
		?,
		/* digest */
		?,
		/* data */
		?,
		/* error */
		?
	`
}

// AsUntypedSlice
func (ce *ClickHouseEntry) AsUntypedSlice() []interface{} {
	return []interface{}{
		ce.d,
		ce.endpointID,
		ce.sourceID,
		ce.contextID,
		ce.actionID,
		ce.duration,
		ce._type,
		ce.size,
		ce.address,
		ce.domain,
		ce.name,
		ce.digest,
		ce.data,
		ce.error,
	}
}

// ClickHouseEntrySearch defines journal entry structure
type ClickHouseEntrySearch struct {
	// Call section
	d          *time.Time
	endpointID *uint16
	sourceID   *string
	contextID  *string
	actionID   *uint8
	duration   *uint64
	// Object section
	_type   *uint8
	size    *uint64
	address *string
	domain  *string
	name    *string
	digest  *string
	data    *string
	// Error section
	error *string
}

// NewClickHouseEntrySearch
func NewClickHouseEntrySearch() *ClickHouseEntrySearch {
	return &ClickHouseEntrySearch{}
}

// Accept
func (ce *ClickHouseEntrySearch) Accept(entry *Entry) *ClickHouseEntrySearch {
	ce.d = nil
	ce.endpointID = nil
	if entry.SourceID.GetString() != "" {
		sourceID := entry.SourceID.GetString()
		ce.sourceID = &sourceID
	}
	if entry.ContextID.GetString() != "" {
		contextID := entry.ContextID.GetString()
		ce.contextID = &contextID
	}
	if entry.Action != ActionUnknown {
		actionID := uint8(entry.Action)
		ce.actionID = &actionID
	}
	ce.duration = nil
	if entry.ObjectType != ObjectTypeUnknown {
		_type := uint8(entry.ObjectType)
		ce._type = &_type
	}
	if entry.ObjectSize > 0 {
		size := entry.ObjectSize
		ce.size = &size
	}
	if entry.ObjectAddress.Printable() != "" {
		address := entry.ObjectAddress.Printable()
		ce.address = &address
	}
	if entry.ObjectMetadata.GetDomain().GetName() != "" {
		domain := entry.ObjectMetadata.GetDomain().GetName()
		ce.domain = &domain
	}
	if entry.ObjectMetadata.GetFilename() != "" {
		name := entry.ObjectMetadata.GetFilename()
		ce.name = &name
	}
	digest := string(entry.ObjectMetadata.GetDigest().GetData())
	if digest != "" {
		ce.digest = &digest
	}
	data := string(entry.ObjectData)
	if data != "" {
		ce.data = &data
	}
	if entry.Error != nil {
		e := entry.Error.Error()
		if e != "" {
			ce.error = &e
		}
	}

	return ce
}

func (ce *ClickHouseEntrySearch) StmtSearchParamsPlaceholderAndArgs() (string, []interface{}) {
	params := []string{}
	str := ""
	args := []interface{}{}

	if ce.d != nil {
		params = append(params, "(d == ?)")
		args = append(args, *ce.d)
	}
	if ce.endpointID != nil {
		params = append(params, "(endpoint_id == ?)")
		args = append(args, *ce.endpointID)
	}
	if ce.sourceID != nil {
		params = append(params, "(source_id == ?)")
		args = append(args, *ce.sourceID)
	}
	if ce.contextID != nil {
		params = append(params, "(context_id ==?)")
		args = append(args, *ce.contextID)
	}
	if ce.actionID != nil {
		params = append(params, "(action_id == ?)")
		args = append(args, *ce.actionID)
	}
	if ce.duration != nil {
		params = append(params, "(duration == ?)")
		args = append(args, *ce.duration)
	}
	if ce._type != nil {
		params = append(params, "(type == ?)")
		args = append(args, *ce._type)
	}
	if ce.size != nil {
		params = append(params, "(size == ?)")
		args = append(args, *ce.size)
	}
	if ce.address != nil {
		params = append(params, "(address == ?)")
		args = append(args, *ce.address)
	}
	if ce.domain != nil {
		params = append(params, "(domain == ?)")
		args = append(args, *ce.domain)
	}
	if ce.name != nil {
		params = append(params, "(name == ?)")
		args = append(args, *ce.name)
	}
	if ce.digest != nil {
		params = append(params, "(digest == ?)")
		args = append(args, *ce.digest)
	}
	if ce.data != nil {
		params = append(params, "(data == ?)")
		args = append(args, *ce.data)
	}
	if ce.error != nil {
		params = append(params, "(error == ?)")
		args = append(args, *ce.error)
	}

	if len(params) > 0 {
		str = strings.Join(params, " AND ")
	}

	if len(str) > 0 {
		str = " AND " + str
	}

	return str, args
}
