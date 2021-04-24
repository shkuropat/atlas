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
	"fmt"
	_ "github.com/mailru/go-clickhouse"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/config/interfaces"
	"github.com/binarly-io/atlas/pkg/trail"
)

// JournalClickHouse
type JournalClickHouse struct {
	trail.JournalBase
}

// Validate interface compatibility
var _ trail.Journaller = &JournalClickHouse{}

// NewJournalClickHouseConfig
func NewJournalClickHouseConfig(cfg interfaces.ClickHouseConfigurator, endpointID int32) (*JournalClickHouse, error) {
	dsn := cfg.GetClickHouseEndpoint()
	return NewJournalClickHouse(dsn, endpointID)
}

// NewJournalClickHouse
func NewJournalClickHouse(dsn string, endpointID int32) (*JournalClickHouse, error) {
	if dsn == "" {
		str := "ClickHouse address in Config is empty"
		log.Errorf(str)
		return nil, fmt.Errorf(str)
	}
	adapter, err := NewAdapterClickHouse(dsn)
	if err != nil {
		return nil, err
	}
	journal, err := trail.NewJournalBase(endpointID, adapter)
	if err != nil {
		return nil, err
	}
	return &JournalClickHouse{
		JournalBase: *journal,
	}, nil
}
