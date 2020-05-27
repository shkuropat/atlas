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
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
	cmd "github.com/spf13/cobra"
	conf "github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/binarly-io/binarly-atlas/pkg/api/atlas"
	"github.com/binarly-io/binarly-atlas/pkg/controller"
	"github.com/binarly-io/binarly-atlas/pkg/controller/client"
	"github.com/binarly-io/binarly-atlas/pkg/softwareid"
	"github.com/binarly-io/binarly-atlas/pkg/transport/client"
)

const (
	sendFileFlagName  = "file"
	sendSTDINFlagName = "stdin"
)

var (
	// readFilename specifies file to read and send to service
	sendFilename string

	// readSTDIN specifies whether to read STDIN
	sendSTDIN bool
)

var sendCmd = &cmd.Command{
	Use:   "send [OPTION] [FILE]",
	Short: "Send file or STDIN to service",
	Long: heredoc.Docf(`
			Specify file to send to service or STDIN
	`,
	),
	Args: func(cmd *cmd.Command, args []string) error {
		//if len(args) < 1 {
		//	return errors.New("requires an filename as argument")
		//}
		if (sendFilename == "") && !sendSTDIN {
			return fmt.Errorf("requires either file or STDIN to send - don't know what to send")
		}

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

		log.Infof("Starting client. Version:%s GitSHA:%s BuiltAt:%s\n", softwareid.Version, softwareid.GitSHA, softwareid.BuiltAt)

		log.Infof("Dial() to %s", serviceAddress)
		conn, err := grpc.Dial(serviceAddress, client_transport.GetGRPCClientOptions(tls, auth, caFile, serverHostOverride, clientID, clientSecret, tokenURL)...)
		if err != nil {
			log.Fatalf("fail to dial %v", err)
		}
		defer conn.Close()

		ControlPlaneClient := atlas.NewControlPlaneClient(conn)

		controller.Init()

		//		log.Infof("About to call CommandsExchange()")
		//		time.Sleep(5 * time.Second)
		//		go controller_client.CommandsExchange(ControlPlaneClient)
		//		log.Infof("Wait...")
		//		time.Sleep(5 * time.Second)
		//		go controller_client.IncomingCommandsHandler(controller.GetIncoming(), controller.GetOutgoing())
		//		log.Infof("Wait...")
		//		time.Sleep(5 * time.Second)
		//		go controller_client.SendEchoRequest(controller.GetOutgoing())

		if sendFilename != "" {
			_, _ = controller_client.SendFile(ControlPlaneClient, sendFilename)
		}

		if sendSTDIN {
			_, _ = controller_client.SendStdin(ControlPlaneClient)
		}

		log.Infof("Press Ctrl+C to exit...")
		<-ctx.Done()
	},
}

func init() {
	sendCmd.PersistentFlags().StringVar(&sendFilename, sendFileFlagName, "", "Send file")
	sendCmd.PersistentFlags().BoolVar(&sendSTDIN, sendSTDINFlagName, false, "Read data from STDIN and send it")

	// Bind full flag set to the configuration
	if err := conf.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(sendCmd)
}
