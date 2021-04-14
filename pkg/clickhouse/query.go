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

	log "github.com/sirupsen/logrus"

	_ "github.com/mailru/go-clickhouse"
)

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
