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
	"time"

	"github.com/MakeNowJust/heredoc"
	log "github.com/sirupsen/logrus"
	cmd "github.com/spf13/cobra"
	conf "github.com/spf13/viper"
	"google.golang.org/grpc"

	pb "github.com/binarly-io/binarly-atlas/pkg/api/mservice"
	"github.com/binarly-io/binarly-atlas/pkg/auth/client"
	"github.com/binarly-io/binarly-atlas/pkg/controller"
	"github.com/binarly-io/binarly-atlas/pkg/controller/client"
	"github.com/binarly-io/binarly-atlas/pkg/transport/client"
	"github.com/binarly-io/binarly-atlas/pkg/version"
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
			Send file or STDIN to service
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

		log.Infof("Starting client. Version:%s GitSHA:%s BuiltAt:%s\n", version.Version, version.GitSHA, version.BuiltAt)

		log.Infof("Dial() to %s", serviceAddress)
		conn, err := grpc.Dial(serviceAddress, getDialOptions()...)
		if err != nil {
			log.Fatalf("fail to dial %v", err)
		}
		defer conn.Close()

		ControlPlaneClient := pb.NewMServiceControlPlaneClient(conn)

		controller.Init()

		log.Infof("About to cal CommandsExchange()")
		time.Sleep(5 * time.Second)
		go controller_client.CommandsExchange(ControlPlaneClient)
		log.Infof("Wait...")
		time.Sleep(5 * time.Second)
		go controller_client.IncomingCommandsHandler(controller.GetIncoming(), controller.GetOutgoing())
		log.Infof("Wait...")
		time.Sleep(5 * time.Second)
		go controller_client.SendEchoRequest(controller.GetOutgoing())

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

// getDialOptions  builds gRPC dial options from flags
func getDialOptions() []grpc.DialOption {
	var opts []grpc.DialOption

	if tls {
		log.Infof("TLS requested")
		if transportOpts, err := client_transport.SetupTransport(caFile, serverHostOverride); err == nil {
			opts = append(opts, transportOpts...)
		} else {
			log.Fatalf("%s", err.Error())
		}
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	if auth {
		log.Infof("OAuth2 requested")
		if !tls {
			log.Fatalf("Need TLS to be enabled")
		}

		if oAuthOpts, err := client_auth.SetupOAuth(clientID, clientSecret, tokenURL); err == nil {
			opts = append(opts, oAuthOpts...)
		} else {
			log.Fatalf("%s", err.Error())
		}
	}

	opts = append(opts, grpc.WithBlock())

	return opts
}
