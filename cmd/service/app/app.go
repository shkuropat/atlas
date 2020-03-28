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
	controller "github.com/sunsingerus/mservice/pkg/controller/service"
	"github.com/sunsingerus/mservice/pkg/transiever/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/golang/glog"

	pb "github.com/sunsingerus/mservice/pkg/api/mservice"

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

	// tls defines whether to use TLS for connection
	tls bool

	// certFile defines path to certificate file. To be used with TLS
	certFile string

	// keyFile defines path to key file. To be used with TLS
	keyFile string

	// port
	port int
)

func init() {
	flag.BoolVar(&versionRequest, "version", false, "Display version and exit")
	flag.StringVar(&configFile, "config", "", "Path to config file.")
	flag.StringVar(&serviceAddress, "service-address", ":10000", "The address of service to use in the format host:port, as localhost:10000")
	flag.BoolVar(&tls, "tls", false, "Whether to use TLS or plain TCP")
	flag.StringVar(&certFile, "cert_file", "", "The TLS cert file. To be used with TLS")
	flag.StringVar(&keyFile, "key_file", "", "The TLS key file. To be used with TLS")
	flag.IntVar(&port, "port", 10000, "The server port")

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

	log.Infof("Starting service. Version:%s GitSHA:%s BuiltAt:%s\n", version.Version, version.GitSHA, version.BuiltAt)

	log.Infof("Listening on %s", serviceAddress)
	listener, err := net.Listen("tcp", serviceAddress)
	if err != nil {
		log.Fatalf("failed to Listen() %v", err)
		os.Exit(1)
	}

	var opts []grpc.ServerOption
	if tls {
		if certFile == "" {
			certFile = testdata.Path("server1.pem")
		}
		if keyFile == "" {
			keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to generate credentials %v", err)
			os.Exit(1)
		}
		opts = []grpc.ServerOption{
			grpc.Creds(creds),
		}
		log.Infof("enabling TLS with cert=%s key=%s", certFile, keyFile)
	}

	transiever_service.Init()

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMServiceControlPlaneServer(grpcServer, &transiever_service.MServiceControlPlaneEndpoint{})
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to Serve() %v", err)
			os.Exit(1)
		}
	}()

	log.Infof("wait for transiever started")

	transiever_service.WaitTransieverStarted()
	go controller.DispatchEchoRequest(transiever_service.GetOutgoingQueue())
	go controller.HandleIncomingCommands(transiever_service.GetIncomingQueue())

	<-ctx.Done()
}
