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

package app

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/golang/glog"
	"golang.org/x/oauth2/dcrp"
	"google.golang.org/grpc"

	pb "github.com/binarly-io/binarly-atlas/pkg/api/mservice"
	"github.com/binarly-io/binarly-atlas/pkg/auth/client"
	"github.com/binarly-io/binarly-atlas/pkg/controller/client"
	controller "github.com/binarly-io/binarly-atlas/pkg/controller/client"
	"github.com/binarly-io/binarly-atlas/pkg/transiever/client"
	"github.com/binarly-io/binarly-atlas/pkg/transiever/service"
	"github.com/binarly-io/binarly-atlas/pkg/transport/client"
	"github.com/binarly-io/binarly-atlas/pkg/version"
)

// CLI parameter variables
var (
	// versionRequest defines request for version report
	versionRequest bool

	// configFile defines path to config file to be used
	configFile string

	// serviceAddr specifies address of service to use
	serviceAddress string

	tls                bool
	caFile             string
	serverHostOverride string

	readFilename string
	readStdin    bool

	auth         bool
	clientID     string
	clientSecret string
	tokenURL     string

	registerClient     bool
	registerClientURL  string
	initialAccessToken string
)

func init() {
	flag.BoolVar(&versionRequest, "version", false, "Display version and exit")
	flag.StringVar(&configFile, "config", "", "Path to config file.")
	flag.StringVar(&serviceAddress, "service-address", "localhost:10000", "The address of service to use in the format host:port, as localhost:10000")
	flag.BoolVar(&tls, "tls", false, "Connection uses TLS if true, else plain TCP")
	flag.StringVar(&caFile, "ca-file", "", "The file containing the CA root cert file")
	flag.BoolVar(&auth, "oauth", false, "Whether to use OAuth2 for authentication")
	flag.StringVar(&clientID, "client-id", "", "ClientID used for Identity server access")
	flag.StringVar(&clientSecret, "client-secret", "", "ClientSecret used for Identity server access")
	flag.StringVar(&tokenURL, "token-url", "", "URL of Identity server's token service")

	flag.StringVar(&serverHostOverride, "server-host-override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	flag.StringVar(&readFilename, "read-filename", "", "Read file and send it")
	flag.BoolVar(&readStdin, "read-stdin", false, "Read data from STDIN and send it")

	flag.BoolVar(&registerClient, "register", false, "Register new client")
	flag.StringVar(&registerClientURL, "register-url", "", "Register client URL")
	flag.StringVar(&initialAccessToken, "initial-access-token", "", "Initial access token")

	flag.Parse()
}

// Run is an entry point of the application
func Run() {
	if versionRequest {
		fmt.Printf("%s\n", version.Version)
		os.Exit(0)
	}

	if registerClient {
		if registerClientURL == "" {
			fmt.Printf("Need to specify --register-url=URL and possibly --initial-access-token=TOKEN\n")
			os.Exit(0)
		}
		config := dcrp.Config{
			InitialAccessToken:            initialAccessToken,
			ClientRegistrationEndpointURL: registerClientURL,
			Metadata: dcrp.Metadata{
				ClientName:              "new fluffy client",
				TokenEndpointAuthMethod: "client_secret_basic",
				GrantTypes:              []string{"client_credentials"},
				SoftwareID:              "atlas",
				SoftwareVersion:         "0.0.1",
			},
		}
		if cl, err := config.Register(); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			fmt.Printf("Registered:\nclient_id:%s\nclient_secret:%s\n", cl.ClientID, cl.ClientSecret)
		}
		os.Exit(0)
	}

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
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewMServiceControlPlaneClient(conn)

	transiever_client.Init()

	log.Infof("About to cal RunMServiceControlPlaneClient()")
	time.Sleep(5 * time.Second)
	go transiever_client.RunMServiceControlPlaneClient(client)
	log.Infof("Wait...")
	time.Sleep(5 * time.Second)
	go controller.IncomingCommandsHandler(transiever_service.GetIncomingQueue(), transiever_service.GetOutgoingQueue())
	log.Infof("Wait...")
	time.Sleep(5 * time.Second)
	go controller.SendEchoRequest(transiever_service.GetOutgoingQueue())

	if readFilename != "" {
		controller_client.SendFile(client, readFilename)
	}

	if readStdin {
		controller_client.SendStdin(client)
	}

	log.Infof("Press Ctrl+C to exit...")
	<-ctx.Done()
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
			os.Exit(1)
		}
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	if auth {
		log.Infof("OAuth2 requested")
		if !tls {
			log.Fatalf("Need TLS to be enabled")
			os.Exit(1)
		}

		if oAuthOpts, err := client_auth.SetupOAuth(clientID, clientSecret, tokenURL); err == nil {
			opts = append(opts, oAuthOpts...)
		} else {
			log.Fatalf("%s", err.Error())
			os.Exit(1)
		}
	}

	opts = append(opts, grpc.WithBlock())

	return opts
}
