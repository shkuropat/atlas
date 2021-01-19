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
	"time"

	_ "github.com/mailru/go-clickhouse"
)

const (
	defaultTimeout = 10 * time.Second
)

// ConnParams
type ConnParams struct {
	// username which to connect with
	username string
	// password which to connect with
	password string

	// hostname where to connect to
	hostname string
	// port where to connect to
	port int

	// Ready-to-use DSN string
	dsn string
	// DSN string with hiddent credentials. Can be used in logs, etc
	dsnHiddenCredentials string

	// Timeout to be used with connection
	timeout time.Duration
}

// NewConnParams
func NewConnParams(username, password string, hostname string, port int) *ConnParams {
	return &ConnParams{
		username:             username,
		password:             password,
		hostname:             hostname,
		port:                 port,
		timeout:              defaultTimeout,
		dsn:                  makeDSN(username, password, hostname, port, false),
		dsnHiddenCredentials: makeDSN(username, password, hostname, port, true),
	}
}

// GetDSN
func (c *ConnParams) GetDSN() string {
	return c.dsn
}

// GetDSNWithHiddenCredentials
func (c *ConnParams) GetDSNWithHiddenCredentials() string {
	return c.dsnHiddenCredentials
}
