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

package base

import (
	"github.com/binarly-io/atlas/pkg/api/atlas"
)

// JournalNOP
type JournalNOP struct {
	ctx  Contexter
	task Tasker
}

// Validate interface compatibility
var _ Journaller = &JournalNOP{}

// NewJournalNOP
func NewJournalNOP() (*JournalNOP, error) {
	return &JournalNOP{}, nil
}

// SetContext
func (j *JournalNOP) SetContext(ctx Contexter) Journaller {
	return nil
}

// SetTask
func (j *JournalNOP) SetTask(task Tasker) Journaller {
	return nil
}

// WithContext
func (j *JournalNOP) WithContext(ctx Contexter) Journaller {
	return nil
}

// WithTask
func (j *JournalNOP) WithTask(task Tasker) Journaller {
	return nil
}

// NewEntry
func (j *JournalNOP) NewEntry(action int32) *Entry {
	return nil
}

// Insert
func (j *JournalNOP) Insert(entry *Entry) error {
	return nil
}

// FindAll
func (j *JournalNOP) FindAll(entry *Entry) ([]*Entry, error) {
	return nil, nil
}

// RequestStart journals beginning of the request processing
func (j *JournalNOP) RequestStart() {

}

// RequestCompleted journals request completed successfully
func (j *JournalNOP) RequestEnd() {

}

// RequestError journals request error
func (j *JournalNOP) RequestError(callErr error) {

}

// SaveData journals data saved successfully
func (j *JournalNOP) SaveData(
	dataAddress *atlas.Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
	data []byte,
) {

}

// SaveDataError journals data not saved due to an error
func (j *JournalNOP) SaveDataError(callErr error) {

}

// ProcessData journals data processed successfully
func (j *JournalNOP) ProcessData(
	dataAddress *atlas.Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
) {

}

// ProcessDataError journals data not processed due to an error
func (j *JournalNOP) ProcessDataError(callErr error) {

}

// SaveTask journals task saved successfully
func (j *JournalNOP) SaveTask(task *atlas.Task) {

}

// SaveTaskError journals task not saved due to an error
func (j *JournalNOP) SaveTaskError(task *atlas.Task, callErr error) {

}

// ProcessTask journals task processed successfully
func (j *JournalNOP) ProcessTask(task *atlas.Task) {

}

// ProcessTaskError journals task not processed due to an error
func (j *JournalNOP) ProcessTaskError(task *atlas.Task, callErr error) {

}
