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

// NOPJournal
type NOPJournal struct {
	ctx  Contexter
	task Tasker
}

// Validate interface compatibility
var _ Journaller = &NOPJournal{}

// NewNOPJournal
func NewNOPJournal() (*NOPJournal, error) {
	return &NOPJournal{}, nil
}

// SetContext
func (j *NOPJournal) SetContext(ctx Contexter) Journaller {
	return nil
}

// SetTask
func (j *NOPJournal) SetTask(task Tasker) Journaller {
	return nil
}

// WithContext
func (j *NOPJournal) WithContext(ctx Contexter) Journaller {
	return nil
}

// WithTask
func (j *NOPJournal) WithTask(task Tasker) Journaller {
	return nil
}

// NewEntry
func (j *NOPJournal) NewEntry(action int32) *Entry {
	return nil
}

// Insert
func (j *NOPJournal) Insert(entry *Entry) error {
	return nil
}

// FindAll
func (j *NOPJournal) FindAll(entry *Entry) ([]*Entry, error) {
	return nil, nil
}

// RequestStart journals beginning of the request processing
func (j *NOPJournal) RequestStart() {
}

// RequestCompleted journals request completed successfully
func (j *NOPJournal) RequestEnd() {
}

// RequestError journals request error
func (j *NOPJournal) RequestError(callErr error) {
}

// SaveData journals data saved successfully
func (j *NOPJournal) SaveData(
	address *atlas.Address,
	size int64,
	metadata *atlas.Metadata,
	data []byte,
) {
}

// SaveDataError journals data not saved due to an error
func (j *NOPJournal) SaveDataError(callErr error) {
}

// ProcessData journals data processed successfully
func (j *NOPJournal) ProcessData(
	address *atlas.Address,
	size int64,
	metadata *atlas.Metadata,
) {
}

// ProcessDataError journals data not processed due to an error
func (j *NOPJournal) ProcessDataError(callErr error) {
}

// Result journals result
func (j *NOPJournal) Result(
	address *atlas.Address,
	size int64,
	metadata *atlas.Metadata,
) {
}

// SaveTask journals task saved successfully
func (j *NOPJournal) SaveTask(task *atlas.Task) {
}

// SaveTaskError journals task not saved due to an error
func (j *NOPJournal) SaveTaskError(task *atlas.Task, callErr error) {
}

// ProcessTask journals task processed successfully
func (j *NOPJournal) ProcessTask(task *atlas.Task) {
}

// ProcessTaskError journals task not processed due to an error
func (j *NOPJournal) ProcessTaskError(task *atlas.Task, callErr error) {
}

// Lookup
func (j *NOPJournal) Lookup(address *atlas.Address) {
}

// LookupError
func (j *NOPJournal) LookupError(address *atlas.Address, callErr error) {
}
