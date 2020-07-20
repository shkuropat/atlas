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

package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

var (
	// Level specifies verbosity level in string form
	// Available levels are:
	// "panic"
	// "fatal"
	// "error"
	// "warn", "warning"
	// "info"
	// "debug"
	// "trace"
	Level string

	// Formatter specifies log format
	// Available formatters are:
	// "json"
	// "text", "txt"
	Formatter string
)

// InitLog sets logging options
func InitLog() {
	formatter, err := parseFormatter(Formatter)
	if err != nil {
		// Use default formatter
		formatter = &log.TextFormatter{}
	}
	log.SetFormatter(formatter)

	level, err := log.ParseLevel(Level)
	if err != nil {
		// Set default level
		level = log.InfoLevel
	}
	log.SetLevel(level)
}

// parseFormatter makes Formatter out of its string name
func parseFormatter(str string) (log.Formatter, error) {
	switch strings.ToLower(str) {
	case "json":
		return &log.JSONFormatter{}, nil
	case "txt", "text":
		return &log.TextFormatter{}, nil
	}

	return nil, fmt.Errorf("not a valid logrus formatter: %q", str)
}
