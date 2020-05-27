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
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
	cmd "github.com/spf13/cobra"
	conf "github.com/spf13/viper"

	"github.com/binarly-io/binarly-atlas/pkg/config/consumer"
	"github.com/binarly-io/binarly-atlas/pkg/kafka/consumer"
	"github.com/binarly-io/binarly-atlas/pkg/softwareid"
)

var (
	// dir specifies dir to write files into
	dir string

	// newest specifies whether to consume starting from newest or oldest available messages
	newest bool

	// ack specifies whether to mark messages as consumed
	ack bool
)

var consumeCmd = &cmd.Command{
	Use:   "consume [OPTION] [FILE]",
	Short: "Consume from Kafka",
	Long: heredoc.Docf(`
			Consume from Kafka
			`,
	),
	Args: func(cmd *cmd.Command, args []string) error {
		//if len(args) < 1 {
		//	return errors.New("requires an filename as argument")
		//}
		return nil
	},
	Run: func(cmd *cmd.Command, args []string) {
		//filename := args[0]

		// Set OS signals and termination context
		ctx, cancelFunc := context.WithCancel(context.Background())
		stopChan := make(chan os.Signal, 2)
		signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-stopChan
			cancelFunc()
			<-stopChan
			os.Exit(1)
		}()

		log.Infof("Starting consumer. Version:%s GitSHA:%s BuiltAt:%s\n", softwareid.Version, softwareid.GitSHA, softwareid.BuiltAt)
		log.Infof("Press Ctrl+C to exit...")

		log.Infof("Config:\n%s", config_consumer.Config.String())

		consumer := kafka.NewConsumer(config_consumer.Config.Brokers, config_consumer.Config.GroupID, config_consumer.Config.Topic)
		consumer.ConsumeLoop(config_consumer.Config.ReadNewest, config_consumer.Config.Ack)

		<-ctx.Done()
	},
}

func init() {
	consumeCmd.PersistentFlags().StringVar(&dir, "dir", "", "Consume into dir")
	consumeCmd.PersistentFlags().BoolVar(&newest, "newest", true, "Consume starting from newest or oldest available messages")
	consumeCmd.PersistentFlags().BoolVar(&ack, "ack", true, "Ack received messages as consumed")

	// Bind full flag set to the configuration
	if err := conf.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(consumeCmd)
}
