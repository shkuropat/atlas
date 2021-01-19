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

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/auth/service"
)

// fetchMetadata fetches claims from incoming gRPC request context
func fetchMetadata(ctx context.Context) jwt.MapClaims {
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
