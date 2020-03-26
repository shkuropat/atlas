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
	"io"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	log "github.com/golang/glog"

	pb "github.com/sunsingerus/mservice/pkg/mservice"

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

type mserviceEndpoint struct {
	pb.UnimplementedMServiceControlPlaneServer
}

var (
	waitTransieverStarted  chan bool
	maxIncomingOutstanding int32 = 100
	incomingQueue          chan *pb.Command
	maxOutgoingOutstanding int32 = 100
	outgoingQueue          chan *pb.Command
)

func (s *mserviceEndpoint) Commands(stream pb.MServiceControlPlane_CommandsServer) error {
	log.Info("Commands() called")

	close(waitTransieverStarted)

	waitIncoming := make(chan bool)
	go func() {
		for {
			msg, err := stream.Recv()
			if err == nil {
				// All went well
				log.Infof("Recv() got msg")
				incomingQueue <- msg
			} else if err == io.EOF {
				// Correct EOF
				log.Infof("Recv() get EOF")

				close(waitIncoming)
				return
			} else {
				// Stream broken
				log.Infof("Recv() got err: %v", err)

				close(waitIncoming)
				return
			}

		}
	}()

	waitOutgoing := make(chan bool)
	go func() {
		for {
			command := <-outgoingQueue
			log.Infof("got command to send")
			err := stream.Send(command)
			if err == nil {
				// All went well
				log.Infof("Send() ok")
			} else if err == io.EOF {
				log.Infof("Send() got EOF")

				close(waitOutgoing)
				return
			} else {
				log.Fatalf("Send() got err: %v", err)

				close(waitOutgoing)
				return
			}
		}
	}()

	<-waitIncoming
	<-waitOutgoing

	return nil
}

func (s *mserviceEndpoint) Data(stream pb.MServiceControlPlane_DataServer) error {
	log.Info("Data() called")
	defer log.Info("Data() exited")

	for {
		dataChunk, err := stream.Recv()
		if err == nil {
			// All went well
			log.Infof("Recv() got msg")
			fmt.Printf("%s\n", string(dataChunk.Bytes))
		} else if err == io.EOF {
			// Correct EOF
			log.Infof("Recv() get EOF")

			return nil
		} else {
			// Stream broken
			log.Infof("Recv() got err: %v", err)

			return nil
		}
	}
}

func (s *mserviceEndpoint) Metrics(stream pb.MServiceControlPlane_MetricsServer) error {
	log.Info("Metrics() called")
	return nil
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

	incomingQueue = make(chan *pb.Command, maxIncomingOutstanding)
	outgoingQueue = make(chan *pb.Command, maxOutgoingOutstanding)
	waitTransieverStarted = make(chan bool)

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterMServiceControlPlaneServer(grpcServer, &mserviceEndpoint{})

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to Serve() %v", err)
			os.Exit(1)
		}
	}()

	log.Infof("wait for transiever started")
	<-waitTransieverStarted

	go func() {
		for i := 0; i < 5; i++ {
			command := pb.NewCommand(
				pb.CommandType_COMMAND_ECHO_REQUEST,
				"",
				0,
				"12-34-56-"+strconv.Itoa(i),
				"",
				0,
				0,
				"desc",
			)
			log.Infof("before Transmit")
			outgoingQueue <- command
			log.Infof("after Transmit")

			log.Infof("before Transmit sleep")
			time.Sleep(3 * time.Second)
			log.Infof("after Transmit sleep")
		}
	}()

	go func() {
		for {
			cmd := <-incomingQueue
			log.Infof("Got cmd %v", cmd)
		}
	}()

	<-ctx.Done()
}
