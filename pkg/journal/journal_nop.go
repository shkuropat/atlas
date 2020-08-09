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

// NopJournal
type NopJournal struct {
}

// NewJournalNOP
func NewJournalNOP() (*NopJournal, error) {
	return &NopJournal{}, nil
}

// RequestStart journals beginning of the request processing
func (j *NopJournal) RequestStart(
	callMetadata *CallMetadata,
) {

}

// SaveData journals data saved successfully
func (j *NopJournal) SaveData(
	callMetadata *CallMetadata,

	dataS3Address *atlas.S3Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
) {

}

// SaveDataError journals data not saved due to an error
func (j *NopJournal) SaveDataError(
	callMetadata *CallMetadata,
	callErr error,
) {

}

// ProcessData journals data processed successfully
func (j *NopJournal) ProcessData(
	callMetadata *CallMetadata,

	dataS3Address *atlas.S3Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
) {

}

// ProcessDataError journals data not processed due to an error
func (j *NopJournal) ProcessDataError(
	callMetadata *CallMetadata,
	callErr error,
) {

}

// RequestCompleted journals request completed successfully
func (j *NopJournal) RequestCompleted(
	callMetadata *CallMetadata,
) {

}

// RequestError journals request error
func (j *NopJournal) RequestError(
	callMetadata *CallMetadata,
	callErr error,
) {

}
