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

// GetBucket
func (a *TaskMinIO) GetBucket() string {
	return a.Config.GetMinIOBucket()
}

// GetTaskPath
func (a *TaskMinIO) GetTaskPath() string {
	return a.Task.GetUUID().String()
}

// GetInPath
func (a *TaskMinIO) GetInPath() string {
	return minio.PathJoin(a.GetTaskPath(), "in")
}

// GetInPrefix
func (a *TaskMinIO) GetInPrefix() string {
	return a.GetInPath() + "/"
}

// GetInFile
func (a *TaskMinIO) GetInFile(file string) string {
	return minio.PathJoin(a.GetInPath(), file)
}

// GetOutPath
func (a *TaskMinIO) GetOutPath() string {
	return minio.PathJoin(a.GetTaskPath(), "out")
}

// GetOutPrefix
func (a *TaskMinIO) GetOutPrefix() string {
	return a.GetOutPath() + "/"
}

// GetOutFile
func (a *TaskMinIO) GetOutFile(file string) string {
	return minio.PathJoin(a.GetOutPath(), file)
}

// GetResultPath
func (a *TaskMinIO) GetResultPath() string {
	return minio.PathJoin(a.GetTaskPath(), "result")
}

// GetResultPrefix
func (a *TaskMinIO) GetResultPrefix() string {
	return a.GetResultPath() + "/"
}

// GetResultFile
func (a *TaskMinIO) GetResultFile(file string) string {
	return minio.PathJoin(a.GetResultPath(), file)
}
