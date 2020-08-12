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
	"database/sql"
	"fmt"
	"github.com/binarly-io/atlas/pkg/context"
	"time"

	"github.com/MakeNowJust/heredoc"
	_ "github.com/mailru/go-clickhouse"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/config"
)

// JournalClickHouse
type JournalClickHouse struct {
	start      time.Time
	endpointID EndpointIDType
	connect    *sql.DB
}

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

	connect, err := sql.Open("clickhouse", dsn)
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

// RequestStart journals beginning of the request processing
func (j *JournalClickHouse) RequestStart(ctx *context.Context) {
	e := NewEntry().SetCtxIDAction(ctx.GetID(), ActionRequestStart)
	if err := j.insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// SaveData journals data saved successfully
func (j *JournalClickHouse) SaveData(
	ctx *context.Context,

	dataS3Address *atlas.S3Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
	data []byte,
) {
	e := NewEntry().
		SetCtxIDAction(ctx.GetID(), ActionSaveData).
		SetSourceID(dataMetadata.GetUserId()).
		SetObject(1, dataS3Address, uint64(dataSize), dataMetadata, data)
	if err := j.insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// SaveDataError journals data not saved due to an error
func (j *JournalClickHouse) SaveDataError(
	ctx *context.Context,
	callErr error,
) {
	e := NewEntry().SetCtxIDAction(ctx.GetID(), ActionSaveDataError).SetError(callErr)
	if err := j.insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// ProcessData journals data processed successfully
func (j *JournalClickHouse) ProcessData(
	ctx *context.Context,

	dataS3Address *atlas.S3Address,
	dataSize int64,
	dataMetadata *atlas.Metadata,
) {
	e := NewEntry().
		SetCtxIDAction(ctx.GetID(), ActionProcessData).
		SetSourceID(dataMetadata.GetUserId()).
		SetObject(1, dataS3Address, uint64(dataSize), dataMetadata, nil)
	if err := j.insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// ProcessDataError journals data not processed due to an error
func (j *JournalClickHouse) ProcessDataError(
	ctx *context.Context,
	callErr error,
) {
	e := NewEntry().
		SetCtxIDAction(ctx.GetID(), ActionProcessDataError).SetError(callErr)
	if err := j.insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// RequestCompleted journals request completed successfully
func (j *JournalClickHouse) RequestCompleted(
	ctx *context.Context,
) {
	e := NewEntry().SetCtxIDAction(ctx.GetID(), ActionRequestCompleted)
	if err := j.insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// RequestError journals request error
func (j *JournalClickHouse) RequestError(
	ctx *context.Context,
	callErr error,
) {
	e := NewEntry().SetCtxIDAction(ctx.GetID(), ActionRequestError).SetError(callErr)
	if err := j.insert(e); err != nil {
		log.Warnf("unable to insert journal entry")
	}
}

// insert
func (j *JournalClickHouse) insert(entry *Entry) error {
	sql := heredoc.Doc(`
		INSERT INTO api_journal (
			d, 
			endpoint_id,
			source_id,
			context_id,
			action_id,
			duration,
			type, 
			size,
			address,
			name,
			data, 
			error
		) VALUES (
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
			/* name */
			?,
			/* data */
			?,
			/* error */
			?
		)
		`,
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

	d := time.Now()
	sourceID := entry.SourceID.GetStringValue()
	contextID := entry.ContextID.GetStringValue()
	actionID := entry.Action
	duration := d.Sub(j.start).Nanoseconds()
	_type := entry.ObjectType
	size := entry.ObjectSize
	address := entry.ObjectAddress.Printable()
	name := entry.ObjectMetadata.GetFilename()
	data := string(entry.ObjectData)
	var e string
	if entry.Error != nil {
		e = entry.Error.Error()
	}
	if _, err := stmt.Exec(
		d,
		j.endpointID,
		sourceID,
		contextID,
		actionID,
		duration,
		_type,
		size,
		address,
		name,
		data,
		e,
	); err != nil {
		log.Errorf("exec failed. err: %v", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("commit failed. err %v", err)
		return err
	}

	return nil
}
