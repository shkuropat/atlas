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
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/minio/minio-go/v6"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/config/interfaces"
)

var errorNotConnected = fmt.Errorf("minio is not connected")

// MinIO
type MinIO struct {
	Endpoint        string
	Secure          bool
	AccessKeyID     string
	SecretAccessKey string

	client *minio.Client
}

// NewMinIO
func NewMinIO(
	endpoint string,
	secure bool,
	insecureSkipVerify bool,
	accessKeyID,
	secretAccessKey string,
) (*MinIO, error) {
	var err error
	min := &MinIO{
		Endpoint:        endpoint,
		Secure:          secure,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
	}
	min.client, err = minio.New(endpoint, accessKeyID, secretAccessKey, secure)
	if secure && insecureSkipVerify {
		// All this dance is for TLSClientConfig - set InsecureSkipVerify: true
		transport := &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		min.client.SetCustomTransport(transport)
	}
	if err != nil {
		log.Errorf("ERROR call minio.New() %v", err.Error())
	}

	return min, err
}

// NewMinIOFromConfig
func NewMinIOFromConfig(cfg interfaces.MinIOEndpointConfig) (*MinIO, error) {
	return NewMinIO(
		cfg.GetMinIOEndpoint(),
		cfg.GetMinIOSecure(),
		cfg.GetMinIOInsecureSkipVerify(),
		cfg.GetMinIOAccessKeyID(),
		cfg.GetMinIOSecretAccessKey(),
	)
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

// PutA
func (m *MinIO) PutA(addr *atlas.S3Address, reader io.Reader) (int64, error) {
	return m.Put(addr.Bucket, addr.Object, reader)
}

// PutUUID
func (m *MinIO) PutUUID(bucketName string, reader io.Reader) (string, int64, error) {
	target, n, err := m.PutUUIDA(bucketName, reader)
	return target.Object, n, err
}

// PutUUIDA
func (m *MinIO) PutUUIDA(bucketName string, reader io.Reader) (*atlas.S3Address, int64, error) {
	target := &atlas.S3Address{
		Bucket: bucketName,
		Object: atlas.NewUUIDRandom().String(),
	}
	n, err := m.PutA(target, reader)
	return target, n, err
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

// FPutA
func (m *MinIO) FPutA(addr *atlas.S3Address, fileName string) (int64, error) {
	return m.FPut(addr.Bucket, addr.Object, fileName)
}

// FPutUUID
func (m *MinIO) FPutUUID(bucketName, fileName string) (string, int64, error) {
	target, n, err := m.FPutUUIDA(bucketName, fileName)
	return target.Object, n, err
}

// FPutUUIDA
func (m *MinIO) FPutUUIDA(bucketName, fileName string) (*atlas.S3Address, int64, error) {
	target := &atlas.S3Address{
		Bucket: bucketName,
		Object: atlas.NewUUIDRandom().String(),
	}
	n, err := m.FPutA(target, fileName)
	return target, n, err
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

// GetA
func (m *MinIO) GetA(addr *atlas.S3Address) (io.Reader, error) {
	return m.Get(addr.Bucket, addr.Object)
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

// FGetA
func (m *MinIO) FGetA(addr *atlas.S3Address, fileName string) error {
	return m.FGet(addr.Bucket, addr.Object, fileName)
}

// FGetTempFile
func (m *MinIO) FGetTempFile(bucketName, objectName, dir, pattern string) (string, error) {
	r, err := m.Get(bucketName, objectName)
	if err != nil {
		log.Errorf("unable to get MinIO object %s/%s err: %v", bucketName, objectName, err)
		return "", err
	}

	f, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		log.Errorf("unable to create tmp file dir: %s pattern: %s err: %v", dir, pattern, err)
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, r)
	if err != nil {
		log.Errorf("unable to get copy MinIO object %s/%s err: %v", bucketName, objectName, err)
		os.Remove(f.Name())
		return "", err
	}

	return f.Name(), nil
}

// FGetTempFileA
func (m *MinIO) FGetTempFileA(addr *atlas.S3Address, dir, pattern string) (string, error) {
	return m.FGetTempFile(addr.Bucket, addr.Object, dir, pattern)
}

// Remove
func (m *MinIO) Remove(bucketName, objectName string) error {
	if m.client == nil {
		return errorNotConnected
	}

	return m.client.RemoveObject(bucketName, objectName)
}

// RemoveA
func (m *MinIO) RemoveA(addr *atlas.S3Address) error {
	return m.Remove(addr.Bucket, addr.Object)
}

// Copy
func (m *MinIO) Copy(dstBucketName, dstObjectName, srcBucketName, srcObjectName string) error {
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

// Digest
func (m *MinIO) Digest(bucketName, objectName string, _type atlas.DigestType) (*atlas.Digest, error) {
	reader, err := m.Get(bucketName, objectName)
	if err != nil {
		return nil, err
	}

	switch _type {
	case atlas.DigestType_DIGEST_SHA256:
		break
	default:
		return nil, fmt.Errorf("unable to calc digest - unknown digest type %v", _type)
	}

	h := sha256.New()
	_, err = io.Copy(h, reader)
	if err != nil {
		return nil, err
	}
	digest := h.Sum(nil)

	res := &atlas.Digest{}
	res.Type = _type
	res.Data = digest

	return res, nil
}

// DigestA
func (m *MinIO) DigestA(addr *atlas.S3Address, _type atlas.DigestType) (*atlas.Digest, error) {
	return m.Digest(addr.Bucket, addr.Object, _type)
}

// CopyA
func (m *MinIO) CopyA(dst, src *atlas.S3Address) error {
	return m.Copy(dst.Bucket, dst.Object, src.Bucket, src.Object)
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

// List n objects from a bucket with a prefix.
func (m *MinIO) List(bucket, prefix string, n int) ([]minio.ObjectInfo, error) {
	var res []minio.ObjectInfo

	// Done channel controls 'ListObjects' go routine
	doneCh := make(chan struct{}, 1)
	defer close(doneCh)

	i := 0
	for object := range m.client.ListObjects(bucket, prefix, false, doneCh) {
		if object.Err != nil {
			continue
		}
		res = append(res, object)
		i++

		if (n > 0) && (i >= n) {
			// Number of objects to list is limited and limit reached
			// Indicate ListObjects go-routine to exit and stop feeding the objectInfo channel.
			doneCh <- struct{}{}
		}
	}

	return res, nil
}
