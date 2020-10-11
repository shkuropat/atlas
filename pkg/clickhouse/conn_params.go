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
	"strconv"
	"time"

	_ "github.com/mailru/go-clickhouse"
)

const (
	// http://user:password@host:8123/
	dsnURLPattern  = "http://%s%s:%s/"
	defaultTimeout = 10 * time.Second

	usernameReplacer = "*"
	passwordReplacer = "*"

	dsnUsernamePasswordPairPattern             = "%s:%s@"
	dsnUsernamePasswordPairUsernameOnlyPattern = "%s@"
)

// ConnParams
type ConnParams struct {
	hostname string
	username string
	password string
	port     int

	dsn                  string
	dsnHiddenCredentials string

	timeout time.Duration
}

// NewConnParams
func NewConnParams(hostname, username, password string, port int) *ConnParams {
	params := &ConnParams{
		hostname: hostname,
		username: username,
		password: password,
		port:     port,

		timeout: defaultTimeout,
	}

	params.dsn = params.makeDSN(false)
	params.dsnHiddenCredentials = params.makeDSN(true)

	return params
}

// makeUserPassPair makes "username:password" pair for connection
func (c *ConnParams) makeUserPassPair(hidden bool) string {

	// In case of hidden username+password pair we'd just return replacement
	if hidden {
		return fmt.Sprintf(dsnUsernamePasswordPairPattern, usernameReplacer, passwordReplacer)
	}

	// We may have neither username nor password
	if c.username == "" && c.password == "" {
		return ""
	}

	// Password may be omitted
	if c.password == "" {
		return fmt.Sprintf(dsnUsernamePasswordPairUsernameOnlyPattern, c.username)
	}

	// Expecting both username and password to be in place
	return fmt.Sprintf(dsnUsernamePasswordPairPattern, c.username, c.password)
}

// makeDSN makes ClickHouse DSN
func (c *ConnParams) makeDSN(hideCredentials bool) string {
	return fmt.Sprintf(
		dsnURLPattern,
		c.makeUserPassPair(hideCredentials),
		c.hostname,
		strconv.Itoa(c.port),
	)
}

// GetDSN
func (c *ConnParams) GetDSN() string {
	return c.dsn
}

// GetDSNWithHiddenCredentials
func (c *ConnParams) GetDSNWithHiddenCredentials() string {
	return c.dsnHiddenCredentials
}
