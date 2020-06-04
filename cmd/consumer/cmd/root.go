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
	"fmt"

	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
	cmd "github.com/spf13/cobra"
	conf "github.com/spf13/viper"

	"github.com/binarly-io/atlas/pkg/ainit"
	"github.com/binarly-io/atlas/pkg/config"
	"github.com/binarly-io/atlas/pkg/config/consumer"
	"github.com/binarly-io/atlas/pkg/logger"
	"github.com/binarly-io/atlas/pkg/softwareid"
)

const (
	etcConfigFireDir       = "/etc/atlas"
	homedirConfigFileDir   = ".atlas"
	defaultConfigFileNoExt = "consumer"
	defaultConfigFile      = defaultConfigFileNoExt + ".yaml"
)

// CLI parameter variables
var (
	// brokers specifies list of Kafka brokers
	brokers string

	// topic specifies Kafka topic work with
	topic string

	// rootCmd represents the base command when called without any sub-commands
	rootCmd = &cmd.Command{
		Use:   "atlas kafka consumer [COMMAND]",
		Short: "Atlas kafka consumer.",
		Long: heredoc.Docf(`
			Atlas kafka consumer
			`,
		),
		PersistentPreRun: func(cmd *cmd.Command, args []string) {
		},
	}
)

func init() {
	cmd.OnInitialize(func() {
		ainit.Init([]string{etcConfigFireDir}, []string{homedirConfigFileDir}, defaultConfigFileNoExt, softwareid.Name)
		config_consumer.ReadIn()
	})

	// Common section
	rootCmd.PersistentFlags().StringVar(&config.ConfigFile, "config", "", fmt.Sprintf("config file (default: %s)", config.PrintConfigFilePaths([]string{etcConfigFireDir, "$HOME/" + homedirConfigFileDir}, defaultConfigFile)))
	rootCmd.PersistentFlags().BoolVarP(&logger.Verbose, "verbose", "v", false, "verbose output")

	// Kafka section
	rootCmd.PersistentFlags().StringVar(&brokers, "brokers", "", "List of Kafka brokers")
	rootCmd.PersistentFlags().StringVar(&topic, "topic", "", "Kafka topic to work with")

	// Bind full flag set to the configuration
	if err := conf.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
