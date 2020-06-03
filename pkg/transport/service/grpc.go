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
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/binarly-io/atlas/pkg/auth/service"
)

// GetGRPCServerOptions builds gRPC server options
func GetGRPCServerOptions(tls, auth bool, tlsCertFile, tlsKeyFile, jwtPublicKeyFile string) []grpc.ServerOption {
	var opts []grpc.ServerOption

	if tls {
		log.Infof("TLS requested")

		if transportOpts, err := setupTLS(tlsCertFile, tlsKeyFile); err == nil {
			opts = append(opts, transportOpts...)
		} else {
			log.Fatalf("%s", err.Error())
		}
	}

	if auth {
		log.Infof("OAuth2 requested")
		if !tls {
			log.Fatalf("Need TLS to be enabled")
		}

		if oAuthOpts, err := service_auth.SetupOAuth(jwtPublicKeyFile); err == nil {
			opts = append(opts, oAuthOpts...)
		} else {
			log.Fatalf("%s", err.Error())
		}
	}

	return opts
}
