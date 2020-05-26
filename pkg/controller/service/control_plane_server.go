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
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/binarly-atlas/pkg/api/atlas"
	"github.com/binarly-io/binarly-atlas/pkg/auth/service"
	"github.com/binarly-io/binarly-atlas/pkg/config/service"
	"github.com/binarly-io/binarly-atlas/pkg/controller"
	"github.com/binarly-io/binarly-atlas/pkg/minio"
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
func (s *ControlPlaneServer) Commands(server atlas.ControlPlane_CommandsServer) error {
	log.Info("Commands() - start")
	defer log.Info("Commands() - end")

	controller.CommandsExchangeEndlessLoop(server)
	return nil
}

// DataChunks gRPC call
func (s *ControlPlaneServer) DataChunks(DataChunksServer atlas.ControlPlane_DataChunksServer) error {
	log.Info("DataChunks() - start")
	defer log.Info("DataChunks() - end")

	claims, err := service_auth.GetClaims(DataChunksServer.Context())
	log.Infof("Claims:")
	for name, value := range claims {
		log.Infof("%s: %v", name, value)
	}

	_, _, err = relayIntoMinIO(DataChunksServer)
	if err != nil {
		log.Errorf("error: %v", err.Error())
	}

	//	// Send back
	//	var buf2 = &bytes.Buffer{}
	//	buf2.WriteString(strings.ToUpper(buf.String()))
	//
	//	_, err = pb.SendDataChunkFile(DataChunksServer, pb.NewMetadata("returnback.file"), buf2)
	//
	return err
}

// relayIntoMinIO
func relayIntoMinIO(DataChunksServer atlas.ControlPlane_DataChunksServer) (int64, *atlas.Metadata, error) {
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

	bucketName := "bucket1"
	objectName := atlas.CreateNewUUID()

	return atlas.RelayDataChunkFileIntoMinIO(DataChunksServer, mi, bucketName, objectName)
}

func relayIntoKafka() {
	//	_, buf, metadata, err := pb.RecvDataChunkFile(DataChunksServer)
	//
	//	log.Infof("Data() Got file len: %d name: %v", buf.Len(), metadata.GetFilename())
	//
	//	producer := kafka.NewProducer(config_service.Config.Brokers, config_service.Config.Topic)
	//	_ = producer.Send(buf.Bytes())
	//
}
