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
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func SetupOAuth(jwtRSAPublicKeyFile string) ([]grpc.ServerOption, error) {

	opts := []grpc.ServerOption{
		// Add an interceptor for all unary RPCs.
		grpc.UnaryInterceptor(unaryInterceptor),

		// Add an interceptor for all stream RPCs.
		grpc.StreamInterceptor(streamInterceptor),
	}

	pem, err := ioutil.ReadFile(jwtRSAPublicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("unable to access Public Key file '%s'", pem)
	}

	jwtRSAPublicKey, err = parseRSAPublicKey(pem)
	if err != nil {
		return nil, fmt.Errorf("file '%s' pase error %v", jwtRSAPublicKeyFile, err)
	}

	return opts, nil
}

// In case of failed authorization, the interceptor blocks execution of the handler and returns an error.
// type grpc.StreamServerInterceptor
func streamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Infof("streamInterceptor %s %t %t", info.FullMethod, info.IsClientStream, info.IsServerStream)

	ctx := ss.Context()
	if err := authorize(ctx); err != nil {
		log.Infof("AUTH FAILED streamInterceptor %s %v", info.FullMethod, err.Error())
		return err
	}

	log.Infof("AUTH OK streamInterceptor %s %t %t", info.FullMethod, info.IsClientStream, info.IsServerStream)

	// Continue execution of handler
	return handler(srv, ss)
}

// In case of failed authorization, the interceptor blocks execution of the handler and returns an error.
// type grpc.StreamClientInterceptor
func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Infof("unaryInterceptor %s", info.FullMethod)

	// Skip authorize when GetJWT is requested
	//if info.FullMethod != "/proto.EventStoreService/GetJWT" {
	//	if err := authorize(ctx); err != nil {
	//		return nil, err
	//	}
	//}

	if err := authorize(ctx); err != nil {
		log.Infof("AUTH FAILED unaryInterceptor %s %v", info.FullMethod, err.Error())
		return nil, err
	}

	log.Infof("AUTH OK unaryInterceptor %s", info.FullMethod)

	// Continue execution of handler
	return handler(ctx, req)
}
