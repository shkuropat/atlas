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
	log "github.com/golang/glog"
	"github.com/sunsingerus/mservice/pkg/controller/client"
	"github.com/sunsingerus/mservice/pkg/transiever/client"
	"github.com/sunsingerus/mservice/pkg/transiever/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/sunsingerus/mservice/pkg/api/mservice"
	controller "github.com/sunsingerus/mservice/pkg/controller/client"
	"github.com/sunsingerus/mservice/pkg/version"
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
)

func init() {
	flag.BoolVar(&versionRequest, "version", false, "Display version and exit")
	flag.StringVar(&configFile, "config", "", "Path to config file.")
	flag.StringVar(&serviceAddress, "service-address", "localhost:10000", "The address of service to use in the format host:port, as localhost:10000")
	flag.BoolVar(&tls, "tls", false, "Connection uses TLS if true, else plain TCP")
	flag.StringVar(&caFile, "ca-file", "", "The file containing the CA root cert file")
	flag.StringVar(&serverHostOverride, "server-host-override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	flag.StringVar(&readFilename, "read-filename", "", "Read file and send it")
	flag.BoolVar(&readStdin, "read-stdin", false, "Read data from STDIN and send it")

	flag.Parse()
}

// Run is an entry point of the application
func Run() {
	if versionRequest {
		fmt.Printf("%s\n", version.Version)
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

	//	<-ctx.Done()

	var opts []grpc.DialOption

	if tls {
		if caFile == "" {
			caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(caFile, serverHostOverride)
		if err != nil {
			log.Fatalf("failed to create TLS credentials %v", err)
			os.Exit(1)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
		log.Infof("enabling TLS with ca=%s", caFile)
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	log.Infof("Dial() to %s", serviceAddress)
	conn, err := grpc.Dial(serviceAddress, opts...)
	if err != nil {
		log.Fatalf("fail to dial %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewMServiceControlPlaneClient(conn)

	transiever_client.Init()

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

	<-ctx.Done()
}
