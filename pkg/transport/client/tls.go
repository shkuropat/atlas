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

package client_transport

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

func setupTLS(config TLSOAuthConfigurator) ([]grpc.DialOption, error) {
	//caFile := config.GetTLSCAFile()
	//if caFile == "" {
	//	caFile = testdata.Path("ca.pem")
	//}
	certFile := config.GetTLSPublicCertFile()
	if certFile == "" {
		certFile = testdata.Path("server1.pem")
	}

	//transportCredentials, err := credentials.NewClientTLSFromFile(caFile, config.GetTLSServerHostOverride())
	transportCredentials, err := credentials.NewClientTLSFromFile(certFile, config.GetTLSServerHostOverride())
	if err != nil {
		log.Fatalf("failed to create TLS credentials %v", err)
	}

	//log.Infof("enabling TLS with ca=%s", caFile)
	log.Infof("enabling TLS with cert=%s", certFile)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(transportCredentials),
	}

	return opts, nil
}
