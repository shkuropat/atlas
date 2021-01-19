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
	databasesql "database/sql"
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc"
	_ "github.com/mailru/go-clickhouse"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/config"
	"github.com/binarly-io/atlas/pkg/rpc_context"
)

// JournalClickHouse
type JournalClickHouse struct {
	start      time.Time
	endpointID EndpointIDType
	connect    *databasesql.DB

	DefaultJournal
}

// Validate interface compatibility
var _ Journal = &JournalClickHouse{}

const (
	defaultObjectType ObjectType = ObjectTypeUnknown
)

// NewJournalClickHouseConfig
func NewJournalClickHouseConfig(cfg config.ClickHouseEndpointConfig, endpointID EndpointIDType) (*JournalClickHouse, error) {
	dsn := cfg.GetClickHouseEndpoint()
	if dsn == "" {
		str := "ClickHouse address in Config is empty"
		log.Errorf(str)
		return nil, fmt.Errorf(str)
	}

	return NewJournalClickHouse(dsn, endpointID)
}

// NewJournalClickHouse
func NewJournalClickHouse(dsn string, endpointID EndpointIDType) (*JournalClickHouse, error) {
	now := time.Now()

	log.Infof("connect to ClickHouse %s", dsn)

	connect, err := databasesql.Open("clickhouse", dsn)
	if err != nil {
		log.Errorf("unable to open ClickHouse err: %v", err)
		return nil, err
	}

	if err := connect.Ping(); err != nil {
		log.Errorf("unable to ping ClickHouse. err: %v", err)
		return nil, err
	}

	return &JournalClickHouse{
		start:      now,
		endpointID: endpointID,
		connect:    connect,
	}, nil
}

// NewEntry
func (j *JournalClickHouse) NewEntry(ctxID *atlas.UUID, action ActionType) *Entry {
	return NewEntry().SetBaseInfo(ctxID, action)
}

// RequestStart journals beginning of the request processing
func (j *JournalClickHouse) RequestStart(ctx *rpc_context.RPCContext) {
	e := NewEntry().SetBaseInfo(ctx.GetID(), ActionRequestStart)
	if err := j.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// RequestCompleted journals request completed successfully
func (j *JournalClickHouse) RequestEnd(
	ctx *rpc_context.RPCContext,
) {
	e := NewEntry().SetBaseInfo(ctx.GetID(), ActionRequestCompleted)
	if err := j.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// RequestError journals request error
func (j *JournalClickHouse) RequestError(
	ctx *rpc_context.RPCContext,
	callErr error,
) {
	e := NewEntry().SetBaseInfo(ctx.GetID(), ActionRequestError).
		SetError(callErr)
	if err := j.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// SaveData journals data saved successfully
func (j *JournalClickHouse) SaveData(
	ctx *rpc_context.RPCContext,

	dataS3Address *atlas.S3Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
	data []byte,
) {
	e := NewEntry().
		SetBaseInfo(ctx.GetID(), ActionSaveData).
		SetSourceID(dataMetadata.GetUserId()).
		SetObject(defaultObjectType, dataS3Address, uint64(dataSize), dataMetadata, data)
	if err := j.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// SaveDataError journals data not saved due to an error
func (j *JournalClickHouse) SaveDataError(
	ctx *rpc_context.RPCContext,
	callErr error,
) {
	e := NewEntry().
		SetBaseInfo(ctx.GetID(), ActionSaveDataError).
		SetError(callErr)
	if err := j.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// ProcessData journals data processed successfully
func (j *JournalClickHouse) ProcessData(
	ctx *rpc_context.RPCContext,

	dataS3Address *atlas.S3Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
) {
	e := NewEntry().
		SetBaseInfo(ctx.GetID(), ActionProcessData).
		SetSourceID(dataMetadata.GetUserId()).
		SetObject(defaultObjectType, dataS3Address, uint64(dataSize), dataMetadata, nil)
	if err := j.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// ProcessDataError journals data not processed due to an error
func (j *JournalClickHouse) ProcessDataError(
	ctx *rpc_context.RPCContext,
	callErr error,
) {
	e := NewEntry().
		SetBaseInfo(ctx.GetID(), ActionProcessDataError).
		SetError(callErr)
	if err := j.Insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// Insert
func (j *JournalClickHouse) Insert(entry *Entry) error {
	e := NewClickHouseEntry().Accept(j, entry)

	sql := heredoc.Doc(
		fmt.Sprintf(`
			INSERT INTO api_journal (
				%s
			) VALUES (
				%s
			)
			`,
			e.Fields(),
			e.StmtParamsPlaceholder(),
		),
	)

	tx, err := j.connect.Begin()
	if err != nil {
		log.Errorf("unable to begin tx. err: %v", err)
		return err
	}

	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Errorf("unable to prepare stmt. err: %v", err)
		return err
	}

	if _, err := stmt.Exec(e.AsUntypedSlice()...); err != nil {
		log.Errorf("exec failed. err: %v", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("commit failed. err %v", err)
		return err
	}

	return nil
}

// FindAll
func (j *JournalClickHouse) FindAll(entry *Entry) ([]ClickHouseEntry, error) {
	e := NewClickHouseEntrySearch().Accept(entry)
	placeholder, args := e.StmtSearchParamsPlaceholderAndArgs()
	sql := heredoc.Doc(
		fmt.Sprintf(`
			SELECT * FROM api_journal WHERE (1 == 1) %s 
			`,
			placeholder,
		),
	)

	stmt, err := j.connect.Prepare(sql)
	if err != nil {
		log.Errorf("unable to prepare stmt. err: %v", err)
		return nil, err
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		log.Errorf("unable to query stmt. err: %v", err)
		return nil, err
	}
	defer rows.Close()

	var res []ClickHouseEntry
	for rows.Next() {
		var ce ClickHouseEntry
		if err := rows.Scan(
			&ce.d,
			&ce.endpointID,
			&ce.sourceID,
			&ce.contextID,
			&ce.actionID,
			&ce.duration,
			&ce._type,
			&ce.size,
			&ce.address,
			&ce.name,
			&ce.digest,
			&ce.data,
			&ce.error,
		); err == nil {
			res = append(res, ce)
		}
	}

	return res, nil
}
