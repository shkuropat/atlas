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

package client_auth

import (
	"context"

	"github.com/coreos/go-oidc"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	cc "golang.org/x/oauth2/clientcredentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/binarly-io/atlas/pkg/config/sections"
)

// requestOAuthToken
func requestOAuthToken(config sections.OAuthConfigurator) (*oauth2.Token, error) {
	log.Infof("Setup OAuth params:\nClientID:%s\nTokenURL:%s\n", config.GetOAuthClientID(), config.GetOAuthTokenURL())

	// Client credential access
	request := &cc.Config{
		ClientID:       config.GetOAuthClientID(),
		ClientSecret:   config.GetOAuthClientSecret(),
		TokenURL:       config.GetOAuthTokenURL(),
		EndpointParams: config.GetOAuthEndpointParams(),
	}

	// Client can be modified (for example ignoring bad certs or otherwise)
	// by modifying the context
	ctx := context.Background()
	// Fetch token from token URL
	token, err := request.Token(ctx)
	if err != nil {
		log.Infof("Error token request %v", err)
		return nil, err
	}
	log.Infof("Token received============\nAccessToken:\n%s\nTokenType:\n%s\nRefreshToken:\n%s\nExpiry:\n%s",
		token.AccessToken,
		token.TokenType,
		token.RefreshToken,
		token.Expiry,
	)
	return token, nil
}

// prepareToken prepares oauth2.Token
func prepareToken(config sections.OAuthConfigurator) *oauth2.Token {
	// Token can be fetched from external party - such as Identity Server or other token provider.
	// External party is charge of token provision is expected to be accessible by tokenURL
	token, _ := requestOAuthToken(config)
	return token

	// Token can be prepared locally (for example for test purposes)
	//	return &oauth2.Token{
	//		AccessToken: "my-secret-token",
	//	}
}

// SetupOAuth
func SetupOAuth(config sections.OAuthConfigurator) ([]grpc.DialOption, error) {
	perRPCCredentials := oauth.NewOauthAccess(prepareToken(config))

	// Set token once per connection
	// It will be sent by gRPC on each call, without need to do it manually
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPCCredentials),
	}

	return opts, nil
}

func qwe() {
	_, err := oidc.NewProvider(context.Background(), "providerURI")
	if err != nil {
		log.Fatal(err)
	}
}
