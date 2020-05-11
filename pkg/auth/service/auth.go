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
	"strings"

	"github.com/dgrijalva/jwt-go"
	log "github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "No metadata provided")
	errMissingToken    = status.Errorf(codes.Unauthenticated, "No authorization token provided")
	errMissingBearer   = status.Errorf(codes.Unauthenticated, "No bearer token provided within authorization token")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "Invalid token")

	jwtRSAPublicKey *rsa.PublicKey
)

// authorize ensures a valid token exists within a request's metadata and authorizes the token received from Metadata
func authorize(ctx context.Context) error {
	// Fetch Bearer token
	// In case it is provided and is correct, consider auth completed
	_, err := fetchJWTToken(ctx)

	return err
}

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
	tokenStr, err := fetchBearerToken(authMetadata)
	if err != nil {
		return nil, err
	}

	return parseToken(tokenStr)
}

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

// parseToken validates the authorization token
func parseToken(_token string) (*jwt.Token, error) {

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify which key to use,
	// but the parsed token (head and claims) is provided to the callback, providing flexibility.
	token, err := jwt.Parse(_token, func(token *jwt.Token) (interface{}, error) {
		// Parse methods use this callback function to supply
		// the key for verification.  The function receives the parsed,
		// but unverified Token.  This allows you to use properties in the
		// Header of the token (such as `kid`) to identify which key to use.

		// What method is used in this token?

		if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
			// This method is supported

			// Return RSA Public Key
			return jwtRSAPublicKey, nil
		}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			// This method IS NOT SUPPORTED
		}

		if _, ok := token.Method.(*jwt.SigningMethodECDSA); ok {
			// This method IS NOT SUPPORTED
		}

		return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method: %v", token.Header["alg"])
	})
	if err != nil {
		log.Errorf("jwt.Parse() FAILED with error %v", err)
		return nil, errInvalidToken
	}
	if !token.Valid {
		log.Errorf("jwt.Parse() FAILED with !token.Valid")
		return nil, errInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Errorf("no token.Claims available")
		return nil, errInvalidToken
	}
	if len(claims) == 0 {
		log.Errorf("zero token.Claims available - this is not correct")
		return nil, errInvalidToken
	}

	log.Infof("Claims:")
	for key, value := range claims {
		log.Infof("%s: %s", key, value)
	}

	return token, nil
}
