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
)

const (
	// Used to build DSN for logs, etc
	usernameReplacer = "*"
	passwordReplacer = "*"

	// DSN section
	// DSN URL should be like http://user:password@host:8123/
	dsnURLPattern = "http://%s%s:%s/"

	dsnUsernamePasswordPairPattern             = "%s:%s@"
	dsnUsernamePasswordPairUsernameOnlyPattern = "%s@"
)

// makeUserPassPair makes "username:password@" part for DSN
func makeUserPassPair(username, password string, hidden bool) string {

	// In case of hidden username+password pair we'd just return replacement
	if hidden {
		return fmt.Sprintf(dsnUsernamePasswordPairPattern, usernameReplacer, passwordReplacer)
	}

	// We may have neither username nor password
	if username == "" && password == "" {
		return ""
	}

	// Password may be omitted
	if password == "" {
		return fmt.Sprintf(dsnUsernamePasswordPairUsernameOnlyPattern, username)
	}

	// Expecting both username and password to be in place
	return fmt.Sprintf(dsnUsernamePasswordPairPattern, username, password)
}

// makeDSN makes ClickHouse DSN
func makeDSN(username, password string, hostname string, port int, hideCredentials bool) string {
	return fmt.Sprintf(
		dsnURLPattern,
		makeUserPassPair(username, password, hideCredentials),
		hostname,
		strconv.Itoa(port),
	)
}
