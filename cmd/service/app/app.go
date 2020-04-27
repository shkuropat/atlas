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
	"github.com/binarly-io/binarly-atlas/pkg/transiever"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/golang/glog"
	"google.golang.org/grpc"

	pbHealth "github.com/binarly-io/binarly-atlas/pkg/api/health"
	pbMService "github.com/binarly-io/binarly-atlas/pkg/api/mservice"
	"github.com/binarly-io/binarly-atlas/pkg/auth/service"
	controller "github.com/binarly-io/binarly-atlas/pkg/controller/service"
	"github.com/binarly-io/binarly-atlas/pkg/transiever/health"
	"github.com/binarly-io/binarly-atlas/pkg/transiever/service"
	"github.com/binarly-io/binarly-atlas/pkg/transport/service"
	"github.com/binarly-io/binarly-atlas/pkg/version"
)

// CLI parameter variables
var (
	// versionRequest specifies request for version report
	versionRequest bool

	// configFile specifies path to config file to be used
	configFile string

	// serviceAddr specifies address of service to use
	serviceAddress string

	// tls specifies whether to use TLS for connection
	tls bool

	// tlsCertFile specifies path to certificate file. To be used with TLS
	tlsCertFile string

	// tlsKeyFile specifies path to key file. To be used with TLS
	tlsKeyFile string

	// port specifies port to listen by gRPC handler
	port int

	// oauth specifies whether to use OAuth2 authentication for clients
	oauth bool

	// jwtPublicKeyFile specifies path to RSA Public Key file to be used for JWT parsing
	jwtPublicKeyFile string

	// brokers specifies list of Kafka brokers
	brokers string
)

func init() {
	flag.BoolVar(&versionRequest, "version", false, "Display version and exit")
	flag.StringVar(&configFile, "config", "", "Path to config file.")
	flag.StringVar(&serviceAddress, "service-address", ":10000", "The address of service to use in the format host:port, as localhost:10000")
	flag.BoolVar(&tls, "tls", false, "Whether to use TLS or plain TCP")
	flag.StringVar(&tlsCertFile, "tls-cert-file", "", "The TLS cert file. To be used with TLS")
	flag.StringVar(&tlsKeyFile, "tls-key-file", "", "The TLS key file. To be used with TLS")
	flag.BoolVar(&oauth, "oauth", false, "Whether to use OAuth2 for authentication")
	flag.StringVar(&jwtPublicKeyFile, "jwt-public-key-file", "", "Public RSA key used for JWT parsing")
	flag.IntVar(&port, "port", 10000, "The server port")
	flag.StringVar(&brokers, "brokers", "", "List of Kafka brokers")

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

	transiever.Init()

	log.Infof("Listening on %s", serviceAddress)
	listener, err := net.Listen("tcp", serviceAddress)
	if err != nil {
		log.Fatalf("failed to Listen() %v", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(getGRPCServerOptions()...)
	pbMService.RegisterMServiceControlPlaneServer(grpcServer, &transiever_service.MServiceControlPlaneServer{})
	pbHealth.RegisterHealthServer(grpcServer, &transiever_health.HealthServer{})

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to Serve() %v", err)
			os.Exit(1)
		}
	}()

	go controller.IncomingCommandsHandler(transiever_service.GetIncomingQueue(), transiever_service.GetOutgoingQueue())

	<-ctx.Done()
}

// getGRPCServerOptions builds gRPC server options from flags
func getGRPCServerOptions() []grpc.ServerOption {
	var opts []grpc.ServerOption
	if tls {
		log.Infof("TLS requested")

		if transportOpts, err := service_transport.SetupTransport(tlsCertFile, tlsKeyFile); err == nil {
			opts = append(opts, transportOpts...)
		} else {
			log.Fatalf("%s", err.Error())
			os.Exit(1)
		}
	}

	if oauth {
		log.Infof("OAuth2 requested")
		if !tls {
			log.Fatalf("Need TLS to be enabled")
			os.Exit(1)
		}

		if oAuthOpts, err := service_auth.SetupOAuth(jwtPublicKeyFile); err == nil {
			opts = append(opts, oAuthOpts...)
		} else {
			log.Fatalf("%s", err.Error())
			os.Exit(1)
		}
	}

	return opts
}
