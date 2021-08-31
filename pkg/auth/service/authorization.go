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

package service_auth

import (
	"context"
	"strings"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	jwks2 "github.com/binarly-io/atlas/pkg/auth/jwks"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "No metadata provided")
	errMissingToken    = status.Errorf(codes.Unauthenticated, "No authorization token provided")
	errMissingBearer   = status.Errorf(codes.Unauthenticated, "No bearer token provided within authorization token")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "Invalid token")
	errInvalidClaims   = status.Errorf(codes.Unauthenticated, "Invalid claims")

	// jwks specifies public key to be used by server for JWT verification
	jwks *jwks2.JWKS
)

// authorize ensures a valid token exists within a request's metadata and authorizes the token received from Metadata
func authorize(ctx context.Context) error {
	// Fetch Bearer token
	// In case it is provided and is correct, consider auth completed
	_, err := GetClaims(ctx, jwt.MapClaims{})

	return err
}

// GetClaims fetches authorization claims from a request's metadata out of RPC context
// claims jwt.Claims is typically a pointer to a struct. Ex.: &ScopeClaims{}
func GetClaims(ctx context.Context, claims jwt.Claims) (jwt.Claims, error) {
	token, err := fetchJWTToken(ctx, claims)
	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}

// fetchJWTToken
func fetchJWTToken(ctx context.Context, claims jwt.Claims) (*jwt.Token, error) {
	var err error

	// Fetch metadata from request's context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	dumpMetadata(md)

	// Fetch authorization metadata from request's metadata
	authorization, ok := md["authorization"]
	if !ok {
		return nil, errMissingToken
	}

	// Fetch authorization token as a string from authorization metadata
	token, err := fetchBearerToken(authorization)
	if err != nil {
		return nil, err
	}

	return parseToken(token, claims)
}

// fetchBearerToken fetches token (as a string) from authorization metadata
func fetchBearerToken(md []string) (string, error) {
	if len(md) < 1 {
		return "", errMissingBearer
	}

	// Fetch token line "Bearer XXXXXXXXXXXX"
	bearer := md[0]

	// Fetch bearer token itself - trim prefix
	token := strings.TrimPrefix(bearer, "Bearer ")
	if len(token) < 1 {
		return "", errMissingBearer
	}
	log.Infof("Bearer %s", token)

	return token, nil
}

// dumpMetadata
func dumpMetadata(md metadata.MD) {
	log.Infof("Dump Metadata ---")
	// Metadata is a map[string][]string
	for key, value := range md {
		log.Infof("[%s]=", key)
		for _, str := range value {
			log.Infof("    %s", str)
		}
	}
	log.Infof("End Dump Metadata ---")
}

// getTokenVerificationKey gets key/cert used for JWT verification
func getTokenVerificationKey(token *jwt.Token) (interface{}, error) {
	// token *jwt.Token is parsed, but unverified Token.
	// Function must return the key for verification of the specified token.
	// Receiving parsed token, allows us to use properties in the Header (such as `kid`) to identify which key to use.
	// For examples, to access 'kid' use token.Header["kid"]
	// This function is especially useful if you have multiple keys (say for various signing methods - RSA, HMAC, etc).
	// The standard is to use 'kid' from the token's Header to identify which key to use.
	// However the parsed token (header and claims) is provided to the callback, thus extending flexibility.
	// JWT Header example:
	// {"alg":"RS256", "typ":"JWT", "kid":"M_GO0JNz4iRvra7NEFI-n"}
	// 'kid' specifies key id/name of the key which should be used for verification
	// Verification keys (with their name/kid) must be provided by OAuth identity server, which issued the token.

	// What signing method is used in this token?
	if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
		// This method is supported

		// Return RSA Public Key (typically provided to server via config) to be used for JWT verification
		return jwks.GetVerificationPublicKey(token, true)
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
		// This method IS NOT SUPPORTED
	}

	if _, ok := token.Method.(*jwt.SigningMethodECDSA); ok {
		// This method IS NOT SUPPORTED
	}

	return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method: %v", token.Header["alg"])
}

// parseToken accepts Claims to fill and parses authorization token string into authorization token
func parseToken(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	// Parse takes JWT token (string) and a function which returns key (public) used for JWT verification.
	token, err := jwt.ParseWithClaims(tokenString, claims, getTokenVerificationKey)
	if err != nil {
		log.Errorf("jwt.Parse() FAILED with error %v", err)
		return nil, errInvalidToken
	}
	if !token.Valid {
		log.Errorf("jwt.Parse() FAILED with !token.Valid")
		return nil, errInvalidToken
	}

	return token, nil
}

/*
func getVerificationCertFromIssuer(token *jwt.Token) (string, error) {
	claims := getCustomizedClaims(token)
	// "https://issuer.url.com/.well-known/jwks.json"
	url := claims.Issuer + "/.well-known/jwks.json"
	resp, err := http.Get(url)
}
*/
