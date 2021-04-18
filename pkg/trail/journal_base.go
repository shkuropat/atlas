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
	"github.com/binarly-io/atlas/pkg/rpc_context"
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

// NewJournalEntry
func (j *JournalBase) NewEntry(ctxID *atlas.UUID, action int32) *JournalEntry {
	return NewJournalEntry().SetBaseInfo(j.start, j.endpointID, ctxID, action)
}

// RequestStart journals beginning of the request processing
func (j *JournalBase) RequestStart(ctx *rpc_context.RPCContext) {
	e := NewJournalEntry().SetBaseInfo(j.start, j.endpointID, ctx.GetUUID(), ActionRequestStart)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// RequestCompleted journals request completed successfully
func (j *JournalBase) RequestEnd(ctx *rpc_context.RPCContext) {
	e := NewJournalEntry().SetBaseInfo(j.start, j.endpointID, ctx.GetUUID(), ActionRequestCompleted)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// RequestError journals request error
func (j *JournalBase) RequestError(ctx *rpc_context.RPCContext, callErr error) {
	e := NewJournalEntry().SetBaseInfo(j.start, j.endpointID, ctx.GetUUID(), ActionRequestError).
		SetError(callErr)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// SaveData journals data saved successfully
func (j *JournalBase) SaveData(
	ctx *rpc_context.RPCContext,

	dataAddress *atlas.Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
	data []byte,
) {
	e := NewJournalEntry().
		SetBaseInfo(j.start, j.endpointID, ctx.GetUUID(), ActionSaveData).
		SetSourceID(dataMetadata.GetUserID()).
		SetObject(dataMetadata.GetType(), dataAddress, uint64(dataSize), dataMetadata, data)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// SaveDataError journals data not saved due to an error
func (j *JournalBase) SaveDataError(
	ctx *rpc_context.RPCContext,
	callErr error,
) {
	e := NewJournalEntry().
		SetBaseInfo(j.start, j.endpointID, ctx.GetUUID(), ActionSaveDataError).
		SetError(callErr)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// ProcessData journals data processed successfully
func (j *JournalBase) ProcessData(
	ctx *rpc_context.RPCContext,

	dataAddress *atlas.Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
) {
	e := NewJournalEntry().
		SetBaseInfo(j.start, j.endpointID, ctx.GetUUID(), ActionProcessData).
		SetSourceID(dataMetadata.GetUserID()).
		SetObject(dataMetadata.GetType(), dataAddress, uint64(dataSize), dataMetadata, nil)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// ProcessDataError journals data not processed due to an error
func (j *JournalBase) ProcessDataError(
	ctx *rpc_context.RPCContext,
	callErr error,
) {
	e := NewJournalEntry().
		SetBaseInfo(j.start, j.endpointID, ctx.GetUUID(), ActionProcessDataError).
		SetError(callErr)
	if err := j.adapter.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}
