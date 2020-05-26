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

package minio

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v6"
	log "github.com/sirupsen/logrus"
)

var errorNotConnected = fmt.Errorf("minio is not connected")

type MinIO struct {
	Endpoint        string
	Secure          bool
	AccessKeyID     string
	SecretAccessKey string

	client *minio.Client
}

// NewMinIO
func NewMinIO(endpoint string, secure bool, accessKeyID, secretAccessKey string) (*MinIO, error) {
	var err error
	min := &MinIO{
		Endpoint:        endpoint,
		Secure:          secure,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	}
	min.client, err = minio.New(endpoint, accessKeyID, secretAccessKey, secure)

	if err != nil {
		log.Errorf("ERROR call minio.New() %v", err.Error())
	}

	return min, err
}

// Put
func (m *MinIO) Put(bucketName, objectName string, reader io.Reader) (int64, error) {
	if m.client == nil {
		return 0, errorNotConnected
	}

	ctx := context.Background()
	// Specify -1 in case object size in unknown in advance
	size := int64(-1)
	options := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	}

	return m.client.PutObjectWithContext(ctx, bucketName, objectName, reader, size, options)
}

// FPut
func (m *MinIO) FPut(bucketName, objectName, fileName string) (int64, error) {
	if m.client == nil {
		return 0, errorNotConnected
	}

	ctx := context.Background()
	options := minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	}

	return m.client.FPutObjectWithContext(ctx, bucketName, objectName, fileName, options)
}

// Get
func (m *MinIO) Get(bucketName, objectName string) (io.Reader, error) {
	if m.client == nil {
		return nil, errorNotConnected
	}

	ctx := context.Background()
	opts := minio.GetObjectOptions{}
	//opts.SetModified(time.Now().Round(10 * time.Minute)) // get object if was modified within the last 10 minutes
	return m.client.GetObjectWithContext(ctx, bucketName, objectName, opts)
}

// FGet
func (m *MinIO) FGet(bucketName, objectName, fileName string) error {
	if m.client == nil {
		return errorNotConnected
	}

	ctx := context.Background()
	opts := minio.GetObjectOptions{}
	//opts.SetModified(time.Now().Round(10 * time.Minute)) // get object if was modified within the last 10 minutes
	return m.client.FGetObjectWithContext(ctx, bucketName, objectName, fileName, opts)
}

// Remove
func (m *MinIO) Remove(bucketName, objectName string) error {
	if m.client == nil {
		return errorNotConnected
	}

	return m.client.RemoveObject(bucketName, objectName)
}

// Copy
func (m *MinIO) Copy(srcBucketName, srcObjectName, dstBucketName, dstObjectName string) error {
	if m.client == nil {
		return errorNotConnected
	}

	src := minio.NewSourceInfo(srcBucketName, srcObjectName, nil)
	dst, err := minio.NewDestinationInfo(dstBucketName, dstObjectName, nil, nil)

	if err != nil {
		return err
	}

	return m.client.CopyObject(dst, src)
}

// ListBuckets
func (m *MinIO) ListBuckets() ([]string, error) {
	if m.client == nil {
		return nil, errorNotConnected
	}

	buckets, err := m.client.ListBuckets()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	var res []string
	for _, bucket := range buckets {
		res = append(res, bucket.Name)
	}

	return res, nil
}

// CreateBucket
func (m *MinIO) CreateBucket(bucketName string) error {
	if m.client == nil {
		return errorNotConnected
	}

	location := "us-east-1"
	return m.client.MakeBucket(bucketName, location)
}

// RemoveBucket
func (m *MinIO) RemoveBucket(bucketName string) error {
	if m.client == nil {
		return errorNotConnected
	}
	return m.client.RemoveBucket(bucketName)
}
