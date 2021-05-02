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

package data_processor_task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// Status
type Status struct {
	Status int32
	Errors []string
}

// Format specifies DataProcessorTask serialization formats
type Format string

const (
	// DataProcessorTask serialization formats
	Unknown Format = "unknown"
	YAML    Format = "yaml"
	JSON    Format = "json"
)

const (
	// DataProcessorTask predefined sections
	ConfigDirs   = "config_dirs"
	ConfigFiles  = "config_files"
	InputDirs    = "input_dirs"
	InputFiles   = "input_files"
	OutputDirs   = "output_dirs"
	OutputFiles  = "output_files"
	ReportLevel  = "report_level"
	SummaryLevel = "summary_level"
	TraceLevel   = "trace_level"
)

// DataProcessorTask
type DataProcessorTask struct {
	// Items contains named slices of parameters
	Items map[string][]string `json:"items,omitempty"  yaml:"items,omitempty"`
	// Status represents status of the data processor task
	Status Status `json:"status,omitempty" yaml:"status,omitempty"`

	// RootDir specifies name of the root directory in case dirs have nested structure
	RootDir string `json:"root,omitempty"   yaml:"root,omitempty"`
	// TaskFile specifies path/name where to serialize/un-serialize from a task
	TaskFile string `json:"task,omitempty"   yaml:"task,omitempty"`
	// Format specifies DataProcessorTask serialization formats
	Format Format `json:"-" yaml:"-"`
}

// New
func New() *DataProcessorTask {
	return &DataProcessorTask{
		Format: Unknown,
	}
}

// ensureItems ensures Item are created
func (t *DataProcessorTask) ensureItems() map[string][]string {
	if t == nil {
		return nil
	}
	if t.Items == nil {
		t.Items = make(map[string][]string)
	}
	return t.Items
}

// CreateTempDir
func (t *DataProcessorTask) CreateTempDir(dir, pattern string) *DataProcessorTask {
	// Create root folder
	root, err := ioutil.TempDir(dir, pattern)
	if err != nil {
	}
	if _, err := os.Stat(root); os.IsNotExist(err) {
	}

	// Create sub-folders
	config := filepath.Join(root, "config")
	input := filepath.Join(root, "input")
	output := filepath.Join(root, "output")

	_ = os.Mkdir(config, 0700)
	_ = os.Mkdir(input, 0700)
	_ = os.Mkdir(output, 0700)

	if _, err := os.Stat(config); os.IsNotExist(err) {
	}
	if _, err := os.Stat(input); os.IsNotExist(err) {
	}
	if _, err := os.Stat(output); os.IsNotExist(err) {
	}

	// Setup folders in the task
	t.SetRootDir(root)
	t.AddConfigDir(config)
	t.AddInputDir(input)
	t.AddOutputDir(output)

	return t
}

// GetRootDir
func (t *DataProcessorTask) GetRootDir() string {
	if t == nil {
		return ""
	}
	return t.RootDir
}

// SetRootDir
func (t *DataProcessorTask) SetRootDir(dir string) *DataProcessorTask {
	if t == nil {
		return nil
	}
	t.RootDir = dir
	return t
}

// GetTaskFile
func (t *DataProcessorTask) GetTaskFile() string {
	if t == nil {
		return ""
	}
	return t.TaskFile
}

// SetTaskFile
func (t *DataProcessorTask) SetTaskFile(file string) *DataProcessorTask {
	if t == nil {
		return nil
	}
	t.TaskFile = file
	return t
}

// Has checks whether DataProcessorTask has specified section
func (t *DataProcessorTask) Has(section string) bool {
	if t == nil {
		return false
	}
	if t.Items == nil {
		return false
	}
	_, ok := t.Items[section]
	return ok
}

// Sections lists all sections
func (t *DataProcessorTask) Sections() []string {
	if t == nil {
		return nil
	}
	if t.Items == nil {
		return nil
	}

	var sections []string
	for section := range t.Items {
		sections = append(sections, section)
	}
	return sections
}

// Walk walk over sections with a function
func (t *DataProcessorTask) Walk(f func(section string, items []string) error) *DataProcessorTask {
	for _, section := range t.Sections() {
		_ = f(section, t.GetAll(section))
	}
	return t
}

// GetAll gets all entities of a section
func (t *DataProcessorTask) GetAll(section string) []string {
	if t.Has(section) {
		return t.Items[section]
	}
	return nil
}

// Len gets number of items wihtin a section
func (t *DataProcessorTask) Len(section string) int {
	return len(t.GetAll(section))
}

// Get first item from a section
func (t *DataProcessorTask) Get(section string) string {
	if t == nil {
		return ""
	}
	if t.Len(section) == 0 {
		return ""
	}
	return t.GetAll(section)[0]
}

// Delete deletes a section
func (t *DataProcessorTask) Delete(section string) *DataProcessorTask {
	if t.Has(section) {
		delete(t.Items, section)
	}
	return t
}

// Add adds item(s) to a section
func (t *DataProcessorTask) Add(section string, items ...string) *DataProcessorTask {
	if t == nil {
		return nil
	}
	if len(items) == 0 {
		return t
	}
	t.ensureItems()
	t.Items[section] = append(t.Items[section], items...)
	return t
}

// Set replaces section with specified items
func (t *DataProcessorTask) Set(section string, items ...string) *DataProcessorTask {
	t.Delete(section)
	t.Add(section, items...)
	return t
}

// GetConfigFiles gets all config files
func (t *DataProcessorTask) GetConfigFiles() []string {
	return t.GetAll(ConfigFiles)
}

// GetConfigFile gets the first config file
func (t *DataProcessorTask) GetConfigFile() string {
	return t.Get(ConfigFiles)
}

// GetConfigDirs gets all config dirs
func (t *DataProcessorTask) GetConfigDirs() []string {
	return t.GetAll(ConfigDirs)
}

// GetConfigDir gets the first config dir
func (t *DataProcessorTask) GetConfigDir() string {
	return t.Get(ConfigDirs)
}

// GetInputFiles gets all input files
func (t *DataProcessorTask) GetInputFiles() []string {
	return t.GetAll(InputFiles)
}

// GetInputFile gets the first input file
func (t *DataProcessorTask) GetInputFile() string {
	return t.Get(InputFiles)
}

// GetInputDirs gets all input dirs
func (t *DataProcessorTask) GetInputDirs() []string {
	return t.GetAll(InputDirs)
}

// GetInputDir gets the first input dir
func (t *DataProcessorTask) GetInputDir() string {
	return t.Get(InputDirs)
}

// GetOutputFiles gets all output files
func (t *DataProcessorTask) GetOutputFiles() []string {
	return t.GetAll(OutputFiles)
}

// GetOutputFile gets the first output file
func (t *DataProcessorTask) GetOutputFile() string {
	return t.Get(OutputFiles)
}

// GetOutputDirs gets all output dirs
func (t *DataProcessorTask) GetOutputDirs() []string {
	return t.GetAll(OutputDirs)
}

// GetOutputDir gets the first output dir
func (t *DataProcessorTask) GetOutputDir() string {
	return t.Get(OutputDirs)
}

// AddConfigFile adds config file(s)
func (t *DataProcessorTask) AddConfigFile(file ...string) *DataProcessorTask {
	return t.Add(ConfigFiles, file...)
}

// AddConfigDir adds config dir(s)
func (t *DataProcessorTask) AddConfigDir(dir ...string) *DataProcessorTask {
	return t.Add(ConfigDirs, dir...)
}

// AddInputFile adds input file(s)
func (t *DataProcessorTask) AddInputFile(file ...string) *DataProcessorTask {
	return t.Add(InputFiles, file...)
}

// AddInputDir adds input dir(s)
func (t *DataProcessorTask) AddInputDir(dir ...string) *DataProcessorTask {
	return t.Add(InputDirs, dir...)
}

// AddOutputFile adds output file(s)
func (t *DataProcessorTask) AddOutputFile(file ...string) *DataProcessorTask {
	return t.Add(OutputFiles, file...)
}

// AddOutputDir adds output dir(s)
func (t *DataProcessorTask) AddOutputDir(dir ...string) *DataProcessorTask {
	return t.Add(OutputDirs, dir...)
}

// GetStatus gets status of the task
func (t *DataProcessorTask) GetStatus() int32 {
	return t.Status.Status
}

// GetErrors gets slice of errors reported by the task
func (t *DataProcessorTask) GetErrors() []string {
	return t.Status.Errors
}

// GetFormat gets format to serialize DataProcessorTask to
func (t *DataProcessorTask) GetFormat() Format {
	if t == nil {
		return Unknown
	}
	return t.Format
}

// SetFormat sets format to serialize DataProcessorTask to
func (t *DataProcessorTask) SetFormat(format Format) *DataProcessorTask {
	if t == nil {
		return nil
	}
	t.Format = format
	return t
}

// Marshal marshals DataProcessorTask according with specified format
func (t *DataProcessorTask) Marshal() (out []byte, err error) {
	if t == nil {
		return nil, fmt.Errorf("unable to marshal nil")
	}
	switch t.Format {
	case YAML:
		return yaml.Marshal(t)
	case JSON:
		return json.Marshal(t)
	}
	return nil, fmt.Errorf("unspecified format to marshal")
}

// Unmarshal unmarshalls DataProcessorTask according with specified format
func (t *DataProcessorTask) Unmarshal(in []byte) (err error) {
	if t == nil {
		return fmt.Errorf("unable to unmarshal into nil")
	}
	switch t.Format {
	case YAML:
		return yaml.Unmarshal(in, t)
	case JSON:
		return json.Unmarshal(in, t)
	}
	return fmt.Errorf("unspecified format to unmarshal")
}

// SaveAs saves DataProcessorTask into specified file
func (t *DataProcessorTask) SaveAs(file string) error {
	if t == nil {
		return nil
	}

	b, err := t.Marshal()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, b, 0600)
}

// SaveTempFile saves DataProcessorTask as temp file with pattern-randomized name into specified dir
func (t *DataProcessorTask) SaveTempFile(dir, pattern string) (string, error) {
	if t == nil {
		return "", nil
	}

	f, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := t.Marshal()
	if err != nil {
		return "", err
	}

	_, err = io.Copy(f, bytes.NewBuffer(b))
	if err != nil {
		_ = os.Remove(f.Name())
		return "", err
	}

	return f.Name(), nil
}

// SaveTempTaskFile saves task as temp file and sets TaskFile to produced temp file name
func (t *DataProcessorTask) SaveTempTaskFile(dir, pattern string) error {
	taskFilename, err := t.SaveTempFile(dir, pattern)
	if err != nil {
		return err
	}
	t.SetTaskFile(taskFilename)
	return nil
}

// ReadFrom reads DataProcessorTask from specified file
func (t *DataProcessorTask) ReadFrom(file string) error {
	if t == nil {
		return nil
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return t.Unmarshal(b)
}

// HasReportLevel
func (t *DataProcessorTask) HasReportLevel() bool {
	return t.Len(ReportLevel) > 0
}

// SetReportLevel
func (t *DataProcessorTask) SetReportLevel(level int) *DataProcessorTask {
	return t.Set(ReportLevel, strconv.Itoa(level))
}

// HasSummaryLevel
func (t *DataProcessorTask) HasSummaryLevel() bool {
	return t.Len(SummaryLevel) > 0
}

// SetSummaryLevel
func (t *DataProcessorTask) SetSummaryLevel(level int) *DataProcessorTask {
	return t.Set(SummaryLevel, strconv.Itoa(level))
}

// HasTraceLevel
func (t *DataProcessorTask) HasTraceLevel() bool {
	return t.Len(TraceLevel) > 0
}

// SetTraceLevel
func (t *DataProcessorTask) SetTraceLevel(level int) *DataProcessorTask {
	return t.Set(TraceLevel, strconv.Itoa(level))
}

// String
func (t *DataProcessorTask) String() string {
	if t == nil {
		return ""
	}
	res := ""
	t.Walk(func(section string, items []string) error {
		res += fmt.Sprintln(section+":", strings.Join(items, ":"))
		return nil
	})
	res += fmt.Sprintln("root:", t.GetRootDir())
	res += fmt.Sprintln("task:", t.GetTaskFile())
	res += fmt.Sprintln("status:", t.GetStatus())
	res += fmt.Sprintln("errors:", strings.Join(t.GetErrors(), ","))
	return res
}

// WalkOutputFiles runs a function over each output-specified files
func (t *DataProcessorTask) WalkOutputFiles(f func(string) error) []error {
	if t == nil {
		return nil
	}
	var res []error
	for _, file := range t.GetOutputFiles() {
		err := f(file)
		if err != nil {
			res = append(res, err)
		}
	}
	return res
}
