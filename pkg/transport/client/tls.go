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
	"crypto/tls"
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

// setupTLS
func setupTLS(config TLSPathsConfigurator) ([]grpc.DialOption, error) {
	var transportCredentials credentials.TransportCredentials
	var err error
	switch {
	case config.GetTLSPublicCertFile() != "":
		transportCredentials, err = transportCredentialsFromServerCertFile(config)
	case config.GetTLSCAFile() != "":
		transportCredentials, err = transportCredentialsFromServerCAFile(config)
	default:
		transportCredentials, err = transportCredentialsNoValidation()
		//transportCredentials, err = transportCredentialsFromSystemCertPool()
	}

	if err != nil {
		log.Fatalf("failed to create TLS credentials %v", err)
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(transportCredentials),
	}

	return opts, nil
}

// transportCredentialsFromServerCertFile
func transportCredentialsFromServerCertFile(config TLSPathsConfigurator) (credentials.TransportCredentials, error) {
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

	//transportCredentials, err := credentials.NewClientTLSFromFile(caFile, config.GetTLSServerHostOverride())
	transportCredentials, err := credentials.NewClientTLSFromFile(certFile, config.GetTLSServerNameOverride())
	if err != nil {
		log.Fatalf("failed to create TLS credentials %v", err)
		return nil, err
	}

	log.Infof("enabling TLS with cert=%s", certFile)
	return transportCredentials, nil
}

// transportCredentialsFromServerCAFile
func transportCredentialsFromServerCAFile(config TLSPathsConfigurator) (credentials.TransportCredentials, error) {
	caFile := config.GetTLSCAFile()
	if caFile == "" {
		caFile = devcerts.Path("ca.cert")
		if _, err := os.Stat(caFile); err != nil {
			path := config.GetPathsOne("tls", sections.PathsOptsRebaseOnCWD)
			caFile = filepath.Join(path, "ca.cert")
		}
	} else {
		if _, err := os.Stat(caFile); err != nil {
			path := config.GetPathsOne("tls", sections.PathsOptsRebaseOnCWD)
			caFile = filepath.Join(path, caFile)
		}
	}

	/*
		b, err := ioutil.ReadFile(caFile)
		if err != nil {
			return nil, err
		}
		cp := x509.NewCertPool()
		if !cp.AppendCertsFromPEM(b) {
			return nil, errors.New("credentials: failed to append certificates")
		}

		conf := &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            cp,
		}
			log.Infof("enabling TLS with ca=%s", caFile)
			transportCredentials := credentials.NewTLS(conf)
	*/

	transportCredentials, err := credentials.NewClientTLSFromFile(caFile, config.GetTLSServerNameOverride())
	if err != nil {
		log.Fatalf("failed to create TLS credentials %v", err)
		return nil, err
	}
	log.Infof("enabling TLS with ca=%s", caFile)
	return transportCredentials, nil
}

// transportCredentialsFromSystemCertPool
func transportCredentialsFromSystemCertPool(config TLSPathsConfigurator) (credentials.TransportCredentials, error) {
	/*
		we can specify System cert pool explicitly
		certPool, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}

		conf := &tls.Config{
			RootCAs:            certPool,
		}
	*/

	// Empty config leads to loading the system CA certificates and therefore trust well-known CA
	conf := &tls.Config{}

	log.Infof("enabling TLS with system cert pool")
	transportCredentials := credentials.NewTLS(conf)
	return transportCredentials, nil
}

// transportCredentialsNoValidation
func transportCredentialsNoValidation() (credentials.TransportCredentials, error) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	log.Infof("enabling TLS with w/o verification")
	return credentials.NewTLS(config), nil
}
