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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	"io"
	"strconv"
	"time"

	"os"

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

	tls                bool
	caFile             string
	serverHostOverride string
)

func runMServiceControlPlaneClient(client pb.MServiceControlPlaneClient) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rpcControl, err := client.Commands(ctx)
	if err != nil {
		log.Fatalf("client.Control() failed %v", err)
		os.Exit(1)
	}
	defer rpcControl.CloseSend()

	log.Infof("rpcControl()")

	waitc := make(chan struct{})
	go func() {
		for {
			log.Infof("before Recv()")
			msg, err := rpcControl.Recv()
			log.Infof("after Recv()")
			if err == io.EOF {
				log.Infof("Recv() received EOF, return from func")
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("failed to rpcControl.Recv() %v", err)
			}
			log.Info("Control() - got message %v", msg)

			log.Infof("before Recv() sleep")
			time.Sleep(time.Second)
			log.Infof("after Recv() sleep")
		}
	}()

	log.Infof("sleep")
	time.Sleep(5 * time.Second)
	log.Infof("continue")

	for i := 0; i < 5; i++ {
		command := pb.NewCommand(
			pb.CommandType_COMMAND_ECHO_REPLY,
			"",
			0,
			"21-43-65-"+strconv.Itoa(i),
			"",
			0,
			0,
			"desc",
		)
		log.Infof("before Send()")

		err := rpcControl.Send(command)
		if err == io.EOF {
			log.Infof("Send() received EOF, return from func")
			return
		}
		if err != nil {
			log.Fatalf("failed to Send() %v", err)
			os.Exit(1)
		}
		log.Infof("after Send()")

		log.Infof("before Send() sleep")
		time.Sleep(3 * time.Second)
		log.Infof("after Send() sleep")
	}

	<-waitc
}

func sendFile(client pb.MServiceControlPlaneClient) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rpcData, err := client.Data(ctx)
	if err != nil {
		log.Fatalf("client.Data() failed %v", err)
		os.Exit(1)
	}
	defer rpcData.CloseSend()

	log.Infof("rpcData()")
	log.Infof("sleep")
	time.Sleep(5 * time.Second)
	log.Infof("continue")

	for i := 0; i < 5; i++ {
		command := pb.NewDataChunk(
			[]byte(fmt.Sprintf("data piece %d", i)),
		)
		log.Infof("before Send()")

		err := rpcData.Send(command)
		if err == io.EOF {
			log.Infof("Send() received EOF, return from func")
			return
		}
		if err != nil {
			log.Fatalf("failed to Send() %v", err)
			os.Exit(1)
		}
		log.Infof("after Send()")

		log.Infof("before Send() sleep")
		time.Sleep(3 * time.Second)
		log.Infof("after Send() sleep")
	}
}

func init() {
	flag.BoolVar(&versionRequest, "version", false, "Display version and exit")
	flag.StringVar(&configFile, "config", "", "Path to config file.")
	flag.StringVar(&serviceAddress, "service-address", "localhost:10000", "The address of service to use in the format host:port, as localhost:10000")
	flag.BoolVar(&tls, "tls", false, "Connection uses TLS if true, else plain TCP")
	flag.StringVar(&caFile, "ca-file", "", "The file containing the CA root cert file")
	flag.StringVar(&serverHostOverride, "server-host-override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")

	flag.Parse()
}

// Run is an entry point of the application
func Run() {
	if versionRequest {
		fmt.Printf("%s\n", version.Version)
		os.Exit(0)
	}

	// Set OS signals and termination context
	//	ctx, cancelFunc := context.WithCancel(context.Background())
	//	stopChan := make(chan os.Signal, 2)
	//	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	//	go func() {
	//		<-stopChan
	//		cancelFunc()
	//		<-stopChan
	//		os.Exit(1)
	//	}()

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

	sendFile(client)
	runMServiceControlPlaneClient(client)
}
