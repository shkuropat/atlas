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

package cmd

import (
	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
	cmd "github.com/spf13/cobra"

	"github.com/binarly-io/atlas/pkg/config/client"
)

var configCmd = &cmd.Command{
	Use:   "config",
	Short: "Print config",
	Long: heredoc.Docf(`
			Print software config and exit
			`,
	),
	Run: func(cmd *cmd.Command, args []string) {
		log.Infof("Config:\n%v", config_client.Config.String())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
