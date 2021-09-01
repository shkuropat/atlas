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
	"github.com/binarly-io/atlas/pkg/config/sections"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/binarly-io/atlas/pkg/auth/client"
)

type PathsTLSOAuthConfigurator interface {
	sections.PathsConfigurator
	sections.TLSConfigurator
	sections.OAuthConfigurator
}

// GetGRPCClientOptions  builds gRPC dial options
func GetGRPCClientOptions(config PathsTLSOAuthConfigurator) []grpc.DialOption {
	var opts []grpc.DialOption

	if config.GetTLSEnabled() {
		log.Infof("TLS requested")
		if transportOpts, err := setupTLS(config); err == nil {
			opts = append(opts, transportOpts...)
		} else {
			log.Fatalf("%s", err.Error())
		}
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	if config.GetOAuthEnabled() {
		log.Infof("OAuth2 requested")
		if !config.GetTLSEnabled() {
			log.Fatalf("Need TLS to be enabled")
		}

		if oAuthOpts, err := client_auth.SetupOAuth(config); err == nil {
			opts = append(opts, oAuthOpts...)
		} else {
			log.Fatalf("%s", err.Error())
		}
	}

	opts = append(opts, grpc.WithBlock())

	return opts
}
