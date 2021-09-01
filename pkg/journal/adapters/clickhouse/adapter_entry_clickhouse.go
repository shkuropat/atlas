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

package clickhouse

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/journal"
)

// AdapterEntryClickHouse defines journal entry structure
type AdapterEntryClickHouse struct {
	// Call section
	d                  time.Time
	endpointID         int32
	endpointInstanceID string
	sourceID           string
	contextID          string
	taskID             string
	typeID             int32
	duration           int64
	// Object section
	_type   int32
	size    uint64
	address string
	domain  string
	name    string
	digest  string
	data    string
	// Result section
	result string
	status string
	error  string
}

// String
func (ce *AdapterEntryClickHouse) String() string {
	if ce == nil {
		return "this is nil"
	}

	b := &bytes.Buffer{}

	// Call section
	_, _ = fmt.Fprintf(b, "d:%s\n", ce.d)
	_, _ = fmt.Fprintf(b, "endpointID:%d\n", ce.endpointID)
	_, _ = fmt.Fprintf(b, "endpointInstanceID:%d\n", ce.endpointInstanceID)
	_, _ = fmt.Fprintf(b, "sourceID:%s\n", ce.sourceID)
	_, _ = fmt.Fprintf(b, "contextID:%s\n", ce.contextID)
	_, _ = fmt.Fprintf(b, "taskID:%s\n", ce.taskID)
	_, _ = fmt.Fprintf(b, "typeID:%d\n", ce.typeID)
	_, _ = fmt.Fprintf(b, "duration:%d\n", ce.duration)
	// Object section
	_, _ = fmt.Fprintf(b, "_type:%d\n", ce._type)
	_, _ = fmt.Fprintf(b, "size:%d\n", ce.size)
	_, _ = fmt.Fprintf(b, "address:%s\n", ce.address)
	_, _ = fmt.Fprintf(b, "domain:%s\n", ce.domain)
	_, _ = fmt.Fprintf(b, "name:%s\n", ce.name)
	_, _ = fmt.Fprintf(b, "digest:%s\n", ce.digest)
	_, _ = fmt.Fprintf(b, "data:%s\n", ce.data)
	// Result section
	_, _ = fmt.Fprintf(b, "result:%s\n", ce.result)
	_, _ = fmt.Fprintf(b, "status:%s\n", ce.status)
	_, _ = fmt.Fprintf(b, "error:%s\n", ce.error)

	return b.String()
}

// NewAdapterEntryClickHouse
func NewAdapterEntryClickHouse() *AdapterEntryClickHouse {
	return &AdapterEntryClickHouse{}
}

// Import
func (ce *AdapterEntryClickHouse) Import(entry *journal.Entry) *AdapterEntryClickHouse {
	// Call section
	ce.d = entry.Time
	ce.endpointID = entry.EndpointID
	ce.endpointInstanceID = entry.EndpointInstanceID.String()
	ce.sourceID = entry.SourceID.String()
	ce.contextID = entry.ContextID.String()
	ce.taskID = entry.TaskID.String()
	ce.typeID = entry.Type
	ce.duration = ce.d.Sub(entry.StartTime).Nanoseconds()
	// Object section
	ce._type = entry.ObjectType
	ce.size = entry.ObjectSize
	ce.address = entry.ObjectAddress.FullString()
	ce.domain = entry.ObjectMetadata.GetDomain().GetName()
	ce.name = entry.ObjectMetadata.GetFilename()
	ce.digest = string(entry.ObjectMetadata.GetDigest().GetData())
	ce.data = string(entry.ObjectData)
	// Result section
	ce.result = entry.Result
	ce.status = entry.Status
	if entry.Error != nil {
		ce.error = entry.Error.Error()
	}

	return ce
}

// Export
func (ce *AdapterEntryClickHouse) Export() *journal.Entry {
	// Call section
	entry := journal.NewEntry()
	entry.Time = ce.d
	entry.EndpointID = ce.endpointID
	entry.SetEndpointInstanceID(atlas.NewUUIDFromString(ce.endpointInstanceID))
	entry.SetSourceID(atlas.NewUserID().SetString(ce.sourceID))
	entry.SetCtxID(atlas.NewUUIDFromString(ce.contextID))
	entry.SetTaskID(atlas.NewUUIDFromString(ce.taskID))
	entry.Type = ce.typeID
	//ce.duration = ce.d.Sub(entry.StartTime).Nanoseconds()
	// Object section
	entry.ObjectType = ce._type
	entry.ObjectSize = ce.size
	entry.SetObjectAddress(atlas.NewAddressFromString(ce.address))
	entry.EnsureObjectMetadata().SetDomain(atlas.NewDomain().SetName(ce.domain))
	entry.EnsureObjectMetadata().SetFilename(ce.name)
	entry.EnsureObjectMetadata().SetDigest(atlas.NewDigest().SetDataFromString(ce.digest))
	entry.ObjectData = []byte(ce.data)
	// Result section
	entry.Result = ce.result
	entry.Status = ce.status
	if ce.error != "" {
		entry.Error = fmt.Errorf(ce.error)
	}

	return entry
}

// Fields
func (ce *AdapterEntryClickHouse) Fields() string {
	return `
		d, 
		endpoint_id,
		endpoint_instance_id,
		source_id,
		context_id,
		task_id,
		type_id,
		duration,
		type, 
		size,
		address,
		domain,
		name,
		digest,
		data, 
		result, 
		status, 
		error
	`
}

// StmtParamsPlaceholder
func (ce *AdapterEntryClickHouse) StmtParamsPlaceholder() string {
	return `
		/* d */
		?,
		/* endpoint_id */
		?,
		/* endpoint_source_id */
		?,
		/* source_id */
		?,
		/* context_id */
		?,
		/* task_id */
		?,
		/* type_id */
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
		/* result */
		?,
		/* status */
		?,
		/* error */
		?
	`
}

// AsUntypedSlice
func (ce *AdapterEntryClickHouse) AsUntypedSlice() []interface{} {
	return []interface{}{
		ce.d,
		ce.endpointID,
		ce.endpointInstanceID,
		ce.sourceID,
		ce.contextID,
		ce.taskID,
		ce.typeID,
		ce.duration,
		ce._type,
		ce.size,
		ce.address,
		ce.domain,
		ce.name,
		ce.digest,
		ce.data,
		ce.result,
		ce.status,
		ce.error,
	}
}

// Scan
func (ce *AdapterEntryClickHouse) Scan(rows *sql.Rows) error {
	if ce == nil {
		return fmt.Errorf("unable to scan into nil")
	}
	return rows.Scan(
		&ce.d,
		&ce.endpointID,
		&ce.endpointInstanceID,
		&ce.sourceID,
		&ce.contextID,
		&ce.taskID,
		&ce.typeID,
		&ce.duration,
		&ce._type,
		&ce.size,
		&ce.address,
		&ce.domain,
		&ce.name,
		&ce.digest,
		&ce.data,
		&ce.result,
		&ce.status,
		&ce.error,
	)
}

// AdapterEntryClickHouseSearch defines journal entry structure
type AdapterEntryClickHouseSearch struct {
	// Call section
	d                  *time.Time
	endpointID         *int32
	endpointInstanceID *string
	sourceID           *string
	contextID          *string
	taskID             *string
	typeID             *int32
	duration           *int64
	// Object section
	_type   *int32
	size    *uint64
	address *string
	domain  *string
	name    *string
	digest  *string
	data    *string
	// Result section
	result *string
	status *string
	error  *string
}

// NewAdapterEntryClickHouseSearch
func NewAdapterEntryClickHouseSearch() *AdapterEntryClickHouseSearch {
	return &AdapterEntryClickHouseSearch{}
}

// Import
func (ce *AdapterEntryClickHouseSearch) Import(entry *journal.Entry) *AdapterEntryClickHouseSearch {
	ce.d = nil
	ce.endpointID = nil
	ce.endpointInstanceID = nil
	if entry.SourceID.String() != "" {
		sourceID := entry.SourceID.String()
		ce.sourceID = &sourceID
	}
	if entry.ContextID.String() != "" {
		contextID := entry.ContextID.String()
		ce.contextID = &contextID
	}
	if entry.TaskID.String() != "" {
		taskID := entry.TaskID.String()
		ce.taskID = &taskID
	}
	if entry.Type != journal.EntryTypeUnknown {
		typeID := entry.Type
		ce.typeID = &typeID
	}
	ce.duration = nil
	if entry.ObjectType != journal.ObjectTypeUnknown {
		_type := entry.ObjectType
		ce._type = &_type
	}
	if entry.ObjectSize > 0 {
		size := entry.ObjectSize
		ce.size = &size
	}
	if entry.ObjectAddress.String() != "" {
		address := entry.ObjectAddress.String()
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
	ce.result = nil
	ce.status = nil
	if entry.Error != nil {
		e := entry.Error.Error()
		if e != "" {
			ce.error = &e
		}
	}

	return ce
}

// StmtSearchParamsPlaceholderAndArgs
func (ce *AdapterEntryClickHouseSearch) StmtSearchParamsPlaceholderAndArgs() (string, []interface{}) {
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
	if ce.endpointInstanceID != nil {
		params = append(params, "(endpoint_instance_id == ?)")
		args = append(args, *ce.endpointInstanceID)
	}
	if ce.sourceID != nil {
		params = append(params, "(source_id == ?)")
		args = append(args, *ce.sourceID)
	}
	if ce.contextID != nil {
		params = append(params, "(context_id ==?)")
		args = append(args, *ce.contextID)
	}
	if ce.taskID != nil {
		params = append(params, "(task_id ==?)")
		args = append(args, *ce.taskID)
	}
	if ce.typeID != nil {
		params = append(params, "(type_id == ?)")
		args = append(args, *ce.typeID)
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
	if ce.result != nil {
		params = append(params, "(result == ?)")
		args = append(args, *ce.result)
	}
	if ce.status != nil {
		params = append(params, "(status == ?)")
		args = append(args, *ce.status)
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
