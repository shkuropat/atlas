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

package service_transport

import (
	"github.com/binarly-io/atlas/pkg/config/sections"
	"github.com/binarly-io/atlas/pkg/devcerts"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path/filepath"
)

type TLSPathsConfigurator interface {
	sections.PathsConfigurator
	sections.TLSConfigurator
}

func setupTLS(config TLSPathsConfigurator) ([]grpc.ServerOption, error) {
	certFile := config.GetTLSPublicCertFile()
	if certFile == "" {
		certFile = devcerts.Path("service.pem")
		if _, err := os.Stat(certFile); err != nil {
			path := config.GetPathsOne("tls", sections.PathsOptsRebaseOnCWD)
			certFile = filepath.Join(path, "service.pem")
		}
	} else {
		if _, err := os.Stat(certFile); err != nil {
			path := config.GetPathsOne("tls", sections.PathsOptsRebaseOnCWD)
			certFile = filepath.Join(path, certFile)
		}
	}
	keyFile := config.GetTLSPrivateKeyFile()
	if keyFile == "" {
		keyFile = devcerts.Path("service.key")
		if _, err := os.Stat(keyFile); err != nil {
			path := config.GetPathsOne("tls", sections.PathsOptsRebaseOnCWD)
			keyFile = filepath.Join(path, "service.key")
		}
	} else {
		if _, err := os.Stat(keyFile); err != nil {
			path := config.GetPathsOne("tls", sections.PathsOptsRebaseOnCWD)
			keyFile = filepath.Join(path, keyFile)
		}
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
	}

	log.Infof("enabling TLS with cert=%s", certFile)
	log.Infof("enabling TLS with key =%s", keyFile)

	opts := []grpc.ServerOption{
		// Enable TLS transport for connections
		grpc.Creds(transportCredentials),
	}

	return opts, nil
}
