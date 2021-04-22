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
	"time"

	_ "github.com/mailru/go-clickhouse"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// JournalBase
type JournalBase struct {
	start      time.Time
	endpointID int32
	adapter    Adapter

	JournalDefault
}

// Validate interface compatibility
var _ Journaller = &JournalBase{}

// NewJournalBase
func NewJournalBase(endpointID int32, adapter Adapter) (*JournalBase, error) {
	return &JournalBase{
		start:      time.Now(),
		endpointID: endpointID,
		adapter:    adapter,
	}, nil
}

// SetContext
func (j *JournalBase) SetContext(ctx Contexter) Journaller {
	j.ctx = ctx
	return j
}

// GetContext
func (j *JournalBase) GetContext() Contexter {
	if j == nil {
		return nil
	}
	return j.ctx
}

// GetContextUUID
func (j *JournalBase) GetContextUUID() *atlas.UUID {
	if j.GetContext() == nil {
		return nil
	}
	return j.GetContext().GetUUID()
}

// NewEntry
func (j *JournalBase) NewEntry(entryType int32) *JournalEntry {
	return NewJournalEntry().SetBaseInfo(j.start, j.endpointID, j.GetContextUUID(), entryType)
}

// Insert
func (j *JournalBase) Insert(entry *JournalEntry) error {
	return j.adapter.Insert(entry)
}

// FindAll
func (j *JournalBase) FindAll(entry *JournalEntry) ([]*JournalEntry, error) {
	return j.adapter.FindAll(entry)
}

// RequestStart journals beginning of the request processing
func (j *JournalBase) RequestStart() {
	e := j.NewEntry(EntryTypeRequestStart)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// RequestCompleted journals request completed successfully
func (j *JournalBase) RequestEnd() {
	e := j.NewEntry(EntryTypeRequestCompleted)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// RequestError journals request error
func (j *JournalBase) RequestError(callErr error) {
	e := j.NewEntry(EntryTypeRequestError).SetError(callErr)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// SaveData journals data saved successfully
func (j *JournalBase) SaveData(
	dataAddress *atlas.Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
	data []byte,
) {
	e := j.NewEntry(EntryTypeSaveData).
		SetSourceID(dataMetadata.GetUserID()).
		SetObject(dataMetadata.GetType(), dataAddress, uint64(dataSize), dataMetadata, data)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// SaveDataError journals data not saved due to an error
func (j *JournalBase) SaveDataError(callErr error) {
	e := j.NewEntry(EntryTypeSaveDataError).SetError(callErr)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// ProcessData journals data processed successfully
func (j *JournalBase) ProcessData(
	dataAddress *atlas.Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
) {
	e := j.NewEntry(EntryTypeProcessData).
		SetSourceID(dataMetadata.GetUserID()).
		SetObject(dataMetadata.GetType(), dataAddress, uint64(dataSize), dataMetadata, nil)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// ProcessDataError journals data not processed due to an error
func (j *JournalBase) ProcessDataError(callErr error) {
	e := j.NewEntry(EntryTypeProcessDataError).SetError(callErr)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// SaveTask journals task saved successfully
func (j *JournalBase) SaveTask(task *atlas.Task) {
	e := j.NewEntry(EntryTypeSaveTask).SetTaskID(task.GetUUID())
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// SaveTaskError journals task not saved due to an error
func (j *JournalBase) SaveTaskError(task *atlas.Task, callErr error) {
	e := j.NewEntry(EntryTypeSaveTaskError).SetError(callErr).SetTaskID(task.GetUUID())
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// ProcessTask journals task processed successfully
func (j *JournalBase) ProcessTask(task *atlas.Task) {
	e := j.NewEntry(EntryTypeProcessTask).SetTaskID(task.GetUUID())
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// ProcessTaskError journals task not processed due to an error
func (j *JournalBase) ProcessTaskError(task *atlas.Task, callErr error) {
	e := j.NewEntry(EntryTypeProcessTaskError).SetError(callErr).SetTaskID(task.GetUUID())
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}
