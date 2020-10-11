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

// DefaultJournal provides empty implementations of all interface functions
type DefaultJournal struct {
}

// RequestStart journals beginning of the request processing
func (j *DefaultJournal) RequestStart(
	ctx *rpc_context.RPCContext,
) {
}

// RequestCompleted journals request completed successfully
func (j *DefaultJournal) RequestCompleted(
	ctx *rpc_context.RPCContext,
) {
}

// RequestError journals request error
func (j *DefaultJournal) RequestError(
	ctx *rpc_context.RPCContext,
	callErr error,
) {
}

// NewEntry
func (j *DefaultJournal) NewEntry(ctxID *atlas.UUID, action ActionType) *Entry {
	return nil
}

// Insert
func (j *DefaultJournal) Insert(entry *Entry) error {
	return nil
}

// SaveData journals data saved successfully
func (j *DefaultJournal) SaveData(
	ctx *rpc_context.RPCContext,

	dataS3Address *atlas.S3Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
	data []byte,
) {
}

// SaveDataError journals data not saved due to an error
func (j *DefaultJournal) SaveDataError(
	ctx *rpc_context.RPCContext,
	callErr error,
) {
}

// ProcessData journals data processed successfully
func (j *DefaultJournal) ProcessData(
	ctx *rpc_context.RPCContext,

	dataS3Address *atlas.S3Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
) {
}

// ProcessDataError journals data not processed due to an error
func (j *DefaultJournal) ProcessDataError(
	ctx *rpc_context.RPCContext,
	callErr error,
) {
}
