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

package adapter

import (
	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/config/sections"
	"github.com/binarly-io/atlas/pkg/minio"
	minio_go "github.com/minio/minio-go/v6"
)

// TaskMinIO
type TaskMinIO struct {
	Config sections.MinIOConfigurator
	Task   *atlas.Task
}

// NewTaskMinIOAdapter
func NewTaskMinIOAdapter(config sections.MinIOConfigurator, task *atlas.Task) *TaskMinIO {
	return &TaskMinIO{
		Config: config,
		Task:   task,
	}
}

const (
	task   = "task"
	in     = "in"
	out    = "out"
	result = "result"
)

// GetBucket
func (a *TaskMinIO) GetBucket() string {
	return a.Config.GetMinIOBucket()
}

// getPath
func (a *TaskMinIO) getPath(what string) string {
	switch what {
	case task:
		return a.Task.GetUUID().String()
	case in:
		return minio.PathJoin(a.GetTaskPath(), "in")
	case out:
		return minio.PathJoin(a.GetTaskPath(), "out")
	case result:
		return minio.PathJoin(a.GetTaskPath(), "result")
	}
	return ""
}

// getPrefix
func (a *TaskMinIO) getPrefix(what string) string {
	switch what {
	case task:
		return a.GetTaskPath() + "/"
	case in:
		return a.GetInPath() + "/"
	case out:
		return a.GetOutPath() + "/"
	case result:
		return a.GetResultPath() + "/"
	}
	return ""
}

// getFile
func (a *TaskMinIO) getFile(what, file string) string {
	switch what {
	case task:
		return minio.PathJoin(a.GetTaskPath(), file)
	case in:
		return minio.PathJoin(a.GetInPath(), file)
	case out:
		return minio.PathJoin(a.GetOutPath(), file)
	case result:
		return minio.PathJoin(a.GetResultPath(), file)
	}
	return ""
}

// walkFiles
func (a *TaskMinIO) walkFiles(what string, f func(*minio.MinIO, *minio_go.ObjectInfo, *atlas.S3Address) bool) error {
	mi, err := minio.NewMinIOFromConfig(a.Config)
	if err != nil {
		return err
	}
	bucket := a.GetBucket()
	list, err := mi.List(bucket, a.getPrefix(what), -1)
	if err != nil {
		return err
	}

	for i := range list {
		info := list[i]
		s3address := atlas.NewS3Address(bucket, info.Key)
		if !f(mi, &info, s3address) {
			break
		}
	}
	return nil
}

/***********************/
/*     Wrappers        */
/***********************/

// GetTaskPath
func (a *TaskMinIO) GetTaskPath() string {
	return a.getPath(task)
}

// GetInPath
func (a *TaskMinIO) GetInPath() string {
	return a.getPath(in)
}

// GetInPrefix
func (a *TaskMinIO) GetInPrefix() string {
	return a.getPrefix(in)
}

// GetInFile
func (a *TaskMinIO) GetInFile(file string) string {
	return a.getFile(in, file)
}

// WalkInFiles
func (a *TaskMinIO) WalkInFiles(f func(*minio.MinIO, *minio_go.ObjectInfo, *atlas.S3Address) bool) error {
	return a.walkFiles(in, f)
}

// GetOutPath
func (a *TaskMinIO) GetOutPath() string {
	return a.getPath(out)
}

// GetOutPrefix
func (a *TaskMinIO) GetOutPrefix() string {
	return a.getPrefix(out)
}

// GetOutFile
func (a *TaskMinIO) GetOutFile(file string) string {
	return a.getFile(out, file)
}

// WalkOutFiles
func (a *TaskMinIO) WalkOutFiles(f func(*minio.MinIO, *minio_go.ObjectInfo, *atlas.S3Address) bool) error {
	return a.walkFiles(out, f)
}

// GetResultPath
func (a *TaskMinIO) GetResultPath() string {
	return a.getPath(result)
}

// GetResultPrefix
func (a *TaskMinIO) GetResultPrefix() string {
	return a.getPrefix(result)
}

// GetResultFile
func (a *TaskMinIO) GetResultFile(file string) string {
	return a.getFile(result, file)
}

// WalkResultFiles
func (a *TaskMinIO) WalkResultFiles(f func(*minio.MinIO, *minio_go.ObjectInfo, *atlas.S3Address) bool) error {
	return a.walkFiles(result, f)
}
