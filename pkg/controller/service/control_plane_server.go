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

package controller_service

import (
	"context"
	"fmt"
	"io"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/auth/service"
	"github.com/binarly-io/atlas/pkg/config/service"
	"github.com/binarly-io/atlas/pkg/controller"
	"github.com/binarly-io/atlas/pkg/kafka"
	"github.com/binarly-io/atlas/pkg/minio"
)

func GetOutgoingQueue() chan *atlas.Command {
	return controller.GetOutgoing()
}

func GetIncomingQueue() chan *atlas.Command {
	return controller.GetIncoming()
}

type ControlPlaneServer struct {
	atlas.UnimplementedControlPlaneServer
}

func NewControlPlaneServer() *ControlPlaneServer {
	return &ControlPlaneServer{}
}

// Commands gRPC call
func (s *ControlPlaneServer) Commands(CommandsServer atlas.ControlPlane_CommandsServer) error {
	log.Info("Commands() - start")
	defer log.Info("Commands() - end")

	_ = fetchUserMetadata(CommandsServer.Context())

	controller.CommandsExchangeEndlessLoop(CommandsServer)
	return nil
}

// DataChunks gRPC call
func (s *ControlPlaneServer) DataChunks(DataChunksServer atlas.ControlPlane_DataChunksServer) error {
	log.Info("DataChunks() - start")
	defer log.Info("DataChunks() - end")

	userMetadata := fetchUserMetadata(DataChunksServer.Context())

	s3address, _, dataMetadata, err := relayIntoMinIO(DataChunksServer, userMetadata)
	//_, _, err = relayIntoKafka(DataChunksServer)
	if err != nil {
		log.Errorf("relay error: %v", err.Error())
	}

	log.Infof("DataChunks() saved as %s/%s", s3address.Bucket, s3address.Object)

	// Write event

	err = writeEvent(userMetadata, dataMetadata, s3address)
	if err != nil {
		log.Errorf("event error: %v", err.Error())
	}

	//	// Send back
	//	var buf2 = &bytes.Buffer{}
	//	buf2.WriteString(strings.ToUpper(buf.String()))
	//
	//	_, err = pb.SendDataChunkFile(DataChunksServer, pb.NewMetadata("returnback.file"), buf2)
	//
	return err
}

func writeEvent(userMetadata jwt.MapClaims, dataMetadata *atlas.Metadata, s3address *atlas.S3Address) error {
	kTransport := kafka.NewCommandTransport(kafka.NewProducer(
		&kafka.Endpoint{
			Brokers: config_service.Config.Brokers,
		},
		&atlas.KafkaAddress{
			Topic: config_service.Config.Topic,
		},
	), nil, true)
	if kTransport == nil {
		log.Errorf("no transport")
		return fmt.Errorf("no transport")
	}
	defer kTransport.Close()

	return kTransport.Send(atlas.NewCommand(
		atlas.CommandType_COMMAND_ECHO_REPLY,
		"",
		0,
		atlas.CreateNewUUID(),
		"reference: ",
		0,
		0,
		"desc",
	))
}

func fetchUserMetadata(ctx context.Context) jwt.MapClaims {
	claims, err := service_auth.GetClaims(ctx)
	if err != nil {
		log.Warnf("unable to get claims with err: %v", err)
		return nil
	}

	log.Infof("Claims:")
	for name, value := range claims {
		log.Infof("%s: %v", name, value)
	}

	return claims
}

func getBucketName(metadata jwt.MapClaims) string {
	return "bucket1"
}

// relayIntoMinIO
func relayIntoMinIO(DataChunksServer atlas.ControlPlane_DataChunksServer, userMetadata jwt.MapClaims) (*atlas.S3Address, int64, *atlas.Metadata, error) {
	log.Info("relayIntoMinIO() - start")
	defer log.Info("relayIntoMinIO() - end")

	mi, err := minio.NewMinIO(
		config_service.Config.Endpoint,
		config_service.Config.Secure,
		config_service.Config.AccessKeyID,
		config_service.Config.SecretAccessKey,
	)

	if err != nil {

	}

	s3address := atlas.NewS3Address(getBucketName(userMetadata), atlas.CreateNewUUID())
	n, mt, err := minio.RelayDataChunkFileIntoMinIO(DataChunksServer, mi, s3address)

	return s3address, n, mt, err
}

func relayIntoKafka(DataChunksServer atlas.ControlPlane_DataChunksServer) (int64, *atlas.Metadata, error) {
	// Kafka transport
	kTransport := kafka.NewDataChunkTransport(
		kafka.NewProducer(
			&kafka.Endpoint{
				Brokers: config_service.Config.Brokers,
			},
			&atlas.KafkaAddress{
				Topic: config_service.Config.Topic,
			},
		),
		nil,
		true,
	)
	if kTransport == nil {
		log.Errorf("no transport")
		return 0, nil, fmt.Errorf("no transport")
	}
	defer kTransport.Close()

	// Kafka file
	kFile, err := atlas.OpenDataChunkFile(kTransport)
	if err != nil {
		log.Errorf("got error: %v", err)
		return 0, nil, err
	}
	defer kFile.Close()

	// Server file
	sFile, err := atlas.OpenDataChunkFile(DataChunksServer)
	if err != nil {
		log.Errorf("got error: %v", err)
		return 0, nil, err
	}
	defer sFile.Close()

	n, err := io.Copy(kFile, sFile)
	if err == nil {
		log.Infof("written: %d", n)
		sFile.PayloadMetadata.Log()
		kFile.PayloadMetadata.Log()
	} else {
		log.Errorf("err: %v", err)
	}

	return n, sFile.PayloadMetadata, err
}
