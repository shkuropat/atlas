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
	controller "github.com/sunsingerus/mservice/pkg/controller/service"
	"github.com/sunsingerus/mservice/pkg/transiever/health"
	"github.com/sunsingerus/mservice/pkg/transiever/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/testdata"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	pbHealth "github.com/sunsingerus/mservice/pkg/api/health"
	pbMService "github.com/sunsingerus/mservice/pkg/api/mservice"
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

	transiever_service.Init()

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

		// TransportCredentials can be created by two ways
		// 1. Directly from files via NewServerTLSFromFile()
		// 2. Or through intermediate Certificate

		// Create TransportCredentials directly from files
		transportCredentials, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		// Create TransportCredentials through intermediate Certificate
		// needs "crypto/tls"
		// cert, err := tls.LoadX509KeyPair(testdata.Path("server1.pem"), testdata.Path("server1.key"))
		// transportCredentials := credentials.NewServerTLSFromCert(&cert)

		if err != nil {
			log.Fatalf("failed to generate credentials %v", err)
			os.Exit(1)
		}
		opts = []grpc.ServerOption{
			// Enable TLS transport for connections
			grpc.Creds(transportCredentials),
		}

		log.Infof("enabling TLS with cert=%s key=%s", certFile, keyFile)
	}

	opts = append(opts, []grpc.ServerOption{
		// Add an interceptor for all unary RPCs.
		grpc.UnaryInterceptor(unaryInterceptor),

		// Add an interceptor for all stream RPCs.
		grpc.StreamInterceptor(streamInterceptor),
	}...,
	)


	grpcServer := grpc.NewServer(opts...)
	pbMService.RegisterMServiceControlPlaneServer(grpcServer, &transiever_service.MServiceControlPlaneEndpoint{})
	pbHealth.RegisterHealthServer(grpcServer, &transiever_health.HealthEndpoint{})

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to Serve() %v", err)
			os.Exit(1)
		}
	}()

	go controller.IncomingCommandsHandler(transiever_service.GetIncomingQueue(), transiever_service.GetOutgoingQueue())

	<-ctx.Done()
}

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "No metadata provided")
	errMissingToken    = status.Errorf(codes.Unauthenticated, "No authorization token provided")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "Invalid token")
)

// In case of failed authorization, the interceptor blocks execution of the handler and returns an error.
// type grpc.StreamServerInterceptor
func streamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Infof("streamInterceptor %s %s %s", info.FullMethod, info.IsClientStream, info.IsServerStream)

	ctx := ss.Context()
	if err := authorize(ctx); err != nil {
		log.Infof("AUTH FAILED streamInterceptor %s %s %s", info.FullMethod, info.IsClientStream, info.IsServerStream)
		return err
	}

	log.Infof("AUTH OK streamInterceptor %s %s %s", info.FullMethod, info.IsClientStream, info.IsServerStream)

	// Continue execution of handler
	return handler(srv, ss)
}

// In case of failed authorization, the interceptor blocks execution of the handler and returns an error.
// type grpc.StreamClientInterceptor
func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Infof("unaryInterceptor %s", info.FullMethod)

	// Skip authorize when GetJWT is requested
	//if info.FullMethod != "/proto.EventStoreService/GetJWT" {
	//	if err := authorize(ctx); err != nil {
	//		return nil, err
	//	}
	//}

	if err := authorize(ctx); err != nil {
		return nil, err
	}

	// Continue execution of handler
	return handler(ctx, req)
}

// authorize ensures a valid token exists within a request's metadata and authorizes the token received from Metadata
func authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errMissingMetadata
	}

	authMetadata, ok := md["authorization"]
	if !ok {
		return errMissingToken
	}

	token := strings.TrimPrefix(authMetadata[0], "Bearer ")
	log.Infof("Bearer %s", token)
	err := validateToken(token)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}

	return nil
}

// valid validates the authorization.
func validateToken(token string) error {
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	if token == "my-secret-token" {
		return nil
	} else {
		return errInvalidToken
	}
}
