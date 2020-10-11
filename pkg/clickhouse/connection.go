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
	"context"
	databasesql "database/sql"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	_ "github.com/mailru/go-clickhouse"
)

type Connection struct {
	params *ConnParams
	conn   *databasesql.DB
}

func NewConnection(params *ConnParams) *Connection {
	// DO not perform connection immediately, do it in lazy manner
	return &Connection{
		params: params,
	}
}

func (c *Connection) connect() {
	log.Infof("Establishing connection: %s", c.params.GetDSNWithHiddenCredentials())
	dbConnection, err := databasesql.Open("clickhouse", c.params.GetDSN())
	if err != nil {
		log.Warnf("FAILED Open(%s) %v", c.params.GetDSNWithHiddenCredentials(), err)
		return
	}

	// Ping should be deadlined
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(defaultTimeout))
	defer cancel()

	if err := dbConnection.PingContext(ctx); err != nil {
		log.Warnf("FAILED Ping(%s) %v", c.params.GetDSNWithHiddenCredentials(), err)
		_ = dbConnection.Close()
		return
	}

	c.conn = dbConnection
}

func (c *Connection) ensureConnected() bool {
	if c.conn != nil {
		log.Infof("Already connected: %s", c.params.GetDSNWithHiddenCredentials())
		return true
	}

	c.connect()

	return c.conn != nil
}

// Query
type Query struct {
	ctx        context.Context
	cancelFunc context.CancelFunc

	Rows *databasesql.Rows
}

// Close
func (q *Query) Close() {
	if q == nil {
		return
	}

	if q.Rows != nil {
		err := q.Rows.Close()
		q.Rows = nil
		if err != nil {
			log.Warnf("UNABLE to close rows. err: %v", err)
		}
	}

	if q.cancelFunc != nil {
		q.cancelFunc()
		q.cancelFunc = nil
	}
}

// Query runs given sql query
func (c *Connection) Query(sql string) (*Query, error) {
	if len(sql) == 0 {
		return nil, nil
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(defaultTimeout))

	if !c.ensureConnected() {
		cancel()
		s := fmt.Sprintf("FAILED connect(%s) for SQL: %s", c.params.GetDSNWithHiddenCredentials(), sql)
		log.Warn(s)
		return nil, fmt.Errorf(s)
	}

	rows, err := c.conn.QueryContext(ctx, sql)
	if err != nil {
		cancel()
		s := fmt.Sprintf("FAILED Query(%s) %v for SQL: %s", c.params.GetDSNWithHiddenCredentials(), err, sql)
		log.Warn(s)
		return nil, err
	}

	log.Infof("clickhouse.QueryContext():'%s'", sql)

	return &Query{
		ctx:        ctx,
		cancelFunc: cancel,
		Rows:       rows,
	}, nil
}

// Exec runs given sql query
func (c *Connection) Exec(sql string) error {
	if len(sql) == 0 {
		return nil
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(defaultTimeout))
	defer cancel()

	if !c.ensureConnected() {
		s := fmt.Sprintf("FAILED connect(%s) for SQL: %s", c.params.GetDSNWithHiddenCredentials(), sql)
		log.Warn(s)
		return fmt.Errorf(s)
	}

	_, err := c.conn.ExecContext(ctx, sql)

	if err != nil {
		log.Warnf("FAILED Exec(%s) %v for SQL: %s", c.params.GetDSNWithHiddenCredentials(), err, sql)
		return err
	}

	log.Infof("clickhouse.Exec():\n", sql)

	return nil
}
