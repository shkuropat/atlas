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
	"crypto/rsa"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "No metadata provided")
	errMissingToken    = status.Errorf(codes.Unauthenticated, "No authorization token provided")
	errMissingBearer   = status.Errorf(codes.Unauthenticated, "No bearer token provided within authorization token")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "Invalid token")
	errInvalidClaims   = status.Errorf(codes.Unauthenticated, "Invalid claims")

	// jwtVerificationRSAPublicKey specifies public key to be used by server for JWT verification
	jwtVerificationRSAPublicKey *rsa.PublicKey
)

// authorize ensures a valid token exists within a request's metadata and authorizes the token received from Metadata
func authorize(ctx context.Context) error {
	// Fetch Bearer token
	// In case it is provided and is correct, consider auth completed
	_, err := fetchJWTToken(ctx)

	return err
}

// GetClaims fetches authorization claims from a request's metadata out of RPC context
func GetClaims(ctx context.Context) (jwt.MapClaims, error) {
	token, err := fetchJWTToken(ctx)
	if err != nil {
		return nil, err
	}

	return getClaims(token), nil
}

// fetchJWTToken
func fetchJWTToken(ctx context.Context) (*jwt.Token, error) {
	var err error

	// Fetch metadata from request's context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	// Fetch authorization metadata from request's metadata
	authMetadata, ok := md["authorization"]
	if !ok {
		return nil, errMissingToken
	}
	dumpMetadata(md)

	// Fetch authorization token from authorization metadata
	tokenString, err := fetchBearerToken(authMetadata)
	if err != nil {
		return nil, err
	}

	return parseToken(tokenString)
}

// fetchBearerToken
func fetchBearerToken(md []string) (string, error) {
	if len(md) < 1 {
		return "", errMissingBearer
	}

	// Fetch token line "Bearer XXXXXXXXXXXX"
	bearer := md[0]

	// Fetch bearer token itself - trim prefix
	tokenString := strings.TrimPrefix(bearer, "Bearer ")
	if len(tokenString) < 1 {
		return "", errMissingBearer
	}
	log.Infof("Bearer %s", tokenString)

	return tokenString, nil
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
		return jwtVerificationRSAPublicKey, nil
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
		// This method IS NOT SUPPORTED
	}

	if _, ok := token.Method.(*jwt.SigningMethodECDSA); ok {
		// This method IS NOT SUPPORTED
	}

	return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method: %v", token.Header["alg"])
}

// parseToken parses authorization token string into authorization token
func parseToken(tokenString string) (*jwt.Token, error) {
	// Parse takes JWT token (string) and a function which returns key (public) used for JWT verification.
	token, err := jwt.Parse(tokenString, getTokenVerificationKey)
	if err != nil {
		log.Errorf("jwt.Parse() FAILED with error %v", err)
		return nil, errInvalidToken
	}
	if !token.Valid {
		log.Errorf("jwt.Parse() FAILED with !token.Valid")
		return nil, errInvalidToken
	}

	// From now JWT Token is parsed and validated (by typically config-provided public key)

	// Ensure token's payload structure is reasonable
	err = verifyClaims(token)
	if err != nil {
		log.Errorf("JWT claims sanity check FAILED with error %v", err)
		return nil, err
	}

	// Main part - token's payload
	dumpClaims(token)

	return token, nil
}

// verifyClaims verifies token's payload structure is reasonbale
func verifyClaims(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Errorf("no token.Claims available")
		return errInvalidClaims
	}
	if len(claims) == 0 {
		log.Errorf("zero token.Claims available - this is not correct")
		return errInvalidClaims
	}
	return nil
}

// getClaims fetches jwt.MapClaims from the token
func getClaims(token *jwt.Token) jwt.MapClaims {
	return token.Claims.(jwt.MapClaims)
}

// dumpClaims
func dumpClaims(token *jwt.Token) {
	// Main part - token's payload
	claims := getClaims(token)
	log.Infof("Claims:")
	for name, value := range claims {
		log.Infof("%s: %v", name, value)
	}
}

//{
//	"aud":"https://atlas-aud",
//	"iss":"https://iss/",
//	"sub":"sub@clients",
//	"iat":1630265536,
//	"exp":1630351936,
//	"azp":"some azp",
//	"gty":"client-credentials"
//	"scope":"permission-scope-read",
// }
type Claims struct {
	jwt.StandardClaims
	Scope string `json:"scope"`
}

// parseCustomizedToken parses authorization token string into authorization token
func parseCustomizedToken(tokenString string) (*jwt.Token, error) {
	// Parse takes JWT token (string) and a function which returns key (public) used for JWT verification.
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, getTokenVerificationKey)
	if err != nil {
		log.Errorf("jwt.Parse() FAILED with error %v", err)
		return nil, errInvalidToken
	}
	if !token.Valid {
		log.Errorf("jwt.Parse() FAILED with !token.Valid")
		return nil, errInvalidToken
	}

	// From now JWT Token is parsed and validated (by typically config-provided public key)

	// Ensure token's payload structure is reasonable
	err = verifyCustomizedClaims(token)
	if err != nil {
		log.Errorf("JWT claims sanity check FAILED with error %v", err)
		return nil, err
	}

	// Main part - token's payload
	dumpCustomizedClaims(token)

	return token, nil
}

// verifyCustomizedClaims verifies token's payload structure is reasonbale
func verifyCustomizedClaims(token *jwt.Token) error {
	_, ok := token.Claims.(*Claims)
	if !ok {
		log.Errorf("no token.Claims available")
		return errInvalidClaims
	}
	// Perform Claims verification
	return nil
}

// getCustomizedClaims fetches jwt.MapClaims from the token
func getCustomizedClaims(token *jwt.Token) *Claims {
	return token.Claims.(*Claims)
}

// dumpCustomizedClaims
func dumpCustomizedClaims(token *jwt.Token) {
	// Main part - token's payload
	claims := getCustomizedClaims(token)
	log.Infof("Customized Claims:")
	log.Infof("Audience  %s", claims.Audience)
	log.Infof("ExpiresAt %d", claims.ExpiresAt)
	log.Infof("Id        %s", claims.Id)
	log.Infof("IssuedAt  %d", claims.IssuedAt)
	log.Infof("Issuer    %s", claims.Issuer)
	log.Infof("NotBefore %d", claims.NotBefore)
	log.Infof("Subject   %s", claims.Subject)
	log.Infof("Scope     %s", claims.Scope)
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func getVerificationCertFromIssuer(token *jwt.Token) (string, error) {
	claims := getCustomizedClaims(token)
	// "https://issuer.url.com/.well-known/jwks.json"
	url := claims.Issuer + "/.well-known/jwks.json"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	jwks := Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return "", err
	}

	for i := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[i].Kid {
			return "" +
					"-----BEGIN CERTIFICATE-----\n" +
					jwks.Keys[i].X5c[0] +
					"\n-----END CERTIFICATE-----",
				nil
		}
	}

	return "", errors.New("unable to find appropriate key")
}
