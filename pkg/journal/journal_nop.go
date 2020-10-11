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
	"github.com/binarly-io/atlas/pkg/rpc_context"
)

// NopJournal
type NopJournal struct {
	DefaultJournal
}

// NewJournalNOP
func NewJournalNOP() (*NopJournal, error) {
	return &NopJournal{}, nil
}

// RequestStart journals beginning of the request processing
func (j *NopJournal) RequestStart(
	ctx *rpc_context.RPCContext,
) {

}

// RequestCompleted journals request completed successfully
func (j *NopJournal) RequestCompleted(
	ctx *rpc_context.RPCContext,
) {

}

// RequestError journals request error
func (j *NopJournal) RequestError(
	ctx *rpc_context.RPCContext,
	callErr error,
) {

}

// NewEntry
func (j *NopJournal) NewEntry(ctxID *atlas.UUID, action ActionType) *Entry {
	return nil
}

// Insert
func (j *NopJournal) Insert(entry *Entry) error {
	return nil
}

// FindAll
func (j *NopJournal) FindAll(entry *Entry) ([]ClickHouseEntry, error) {
	return nil, nil
}

// SaveData journals data saved successfully
func (j *NopJournal) SaveData(
	ctx *rpc_context.RPCContext,

	dataS3Address *atlas.S3Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
	data []byte,
) {

}

// SaveDataError journals data not saved due to an error
func (j *NopJournal) SaveDataError(
	ctx *rpc_context.RPCContext,
	callErr error,
) {

}

// ProcessData journals data processed successfully
func (j *NopJournal) ProcessData(
	ctx *rpc_context.RPCContext,

	dataS3Address *atlas.S3Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
) {

}

// ProcessDataError journals data not processed due to an error
func (j *NopJournal) ProcessDataError(
	ctx *rpc_context.RPCContext,
	callErr error,
) {

}
