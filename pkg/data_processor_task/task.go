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
	"strings"

	"gopkg.in/yaml.v2"
)

// Status specifies status of the DataProcessorTask
type Status struct {
	// Status represents status code
	Status int32
	// Errors is a list of errors, if any
	Errors []string
}

// Format specifies DataProcessorTask serialization formats
type Format string

const (
	// DataProcessorTask serialization formats
	Empty   Format = ""
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
	InputTables  = "input_tables"
	OutputDirs   = "output_dirs"
	OutputFiles  = "output_files"
	OutputTables = "output_tables"
	ReportLevel  = "report_level"
	SummaryLevel = "summary_level"
	TraceLevel   = "trace_level"
)

const (
	// Names of the directories
	ConfigDirName = "config"
	InputDirName  = "input"
	OutputDirName = "output"
)

// DataProcessorTask specifies task to be launched as external process for data processing
type DataProcessorTask struct {
	// Items contains named sections of various parameters. Each section is represented as a slice.
	Items map[string][]string `json:"items,omitempty" yaml:"items,omitempty"`
	// Status represents status of the data processor task
	Status Status `json:"status,omitempty" yaml:"status,omitempty"`

	// RootDir specifies name of the root directory in case dirs have nested structure
	RootDir string `json:"root,omitempty" yaml:"root,omitempty"`
	// TaskFile specifies path/name where to/from serialize/un-serialize a task
	TaskFile string `json:"task,omitempty" yaml:"task,omitempty"`
	// Format specifies DataProcessorTask serialization formats
	Format Format `json:"-" yaml:"-"`
}

// DataProcessorTaskFile defines what DataProcessorTask file should be used.
// It has to be exported var in order to be used in cases such as:
// rootCmd.PersistentFlags().StringVar(&data_processor_task.DataProcessorTaskFile, "task", "", "DataProcessorTask file")
var DataProcessorTaskFile string

// Task is the DataProcessorTask un-serialized from DataProcessorTaskFile
// It has to be exported var in order to be used in external modules to access the Task specification
var Task *DataProcessorTask

// ReadIn reads/un-serializes the DataProcessorTask from DataProcessorTaskFile
func ReadIn() {
	if DataProcessorTaskFile == "" {
		// No task file specified
		return
	}

	Task = New()
	if err := Task.ReadFrom(DataProcessorTaskFile); err != nil {
		// Unable to read task file, need to clear task
		Task = nil
		return
	}

	// Task read successfully
}

// New creates new task
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
		return t
	}
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return t
	}

	// Create sub-folders
	config := filepath.Join(root, ConfigDirName)
	input := filepath.Join(root, InputDirName)
	output := filepath.Join(root, OutputDirName)
	if err := os.Mkdir(config, 0700); err != nil {
		return t
	}
	if err := os.Mkdir(input, 0700); err != nil {
		return t
	}
	if err := os.Mkdir(output, 0700); err != nil {
		return t
	}

	if _, err := os.Stat(config); os.IsNotExist(err) {
		return t
	}
	if _, err := os.Stat(input); os.IsNotExist(err) {
		return t
	}
	if _, err := os.Stat(output); os.IsNotExist(err) {
		return t
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

// Exists checks whether specified section exists within DataProcessorTask.
// Returns tru if section exists. Section may have 0 items in it and return true
func (t *DataProcessorTask) Exists(section string) bool {
	if t == nil {
		return false
	}
	if t.Items == nil {
		return false
	}
	_, ok := t.Items[section]
	return ok
}

// Has checks whether DataProcessorTask has something in specified section.
// Returns true only in case section has > 0 items in it.
func (t *DataProcessorTask) Has(section string) bool {
	return t.Len(section) > 0
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
func (t *DataProcessorTask) Walk(f func(section string, items []string) bool) *DataProcessorTask {
	for _, section := range t.Sections() {
		if !f(section, t.GetAll(section)) {
			// Stop walking
			break
		}
	}
	return t
}

// WalkSection walk over all items in a section with a function
func (t *DataProcessorTask) WalkSection(section string, f func(item string) bool) *DataProcessorTask {
	return t.Walk(func(s string, items []string) bool {
		if s == section {
			// This is our section, walk all items
			for _, itm := range items {
				if !f(itm) {
					// Stop walking
					break
				}
			}
			// Section found, no need to continue
			return false
		}
		return true
	})
}

// GetAll gets all entities of a section
func (t *DataProcessorTask) GetAll(section string) []string {
	if t.Exists(section) {
		return t.Items[section]
	}
	return nil
}

// Len gets number of items within a section
func (t *DataProcessorTask) Len(section string) int {
	return len(t.GetAll(section))
}

// Get first item from a section or default value. Default value can be provided explicitly or "" used otherwise
func (t *DataProcessorTask) Get(section string, defaultValue ...string) string {
	// Prepare default value
	_default := ""
	if len(defaultValue) > 0 {
		_default = defaultValue[0]
	}
	if t == nil {
		return _default
	}
	if t.Len(section) == 0 {
		return _default
	}
	return t.GetAll(section)[0]
}

// Delete deletes a section
func (t *DataProcessorTask) Delete(section string) *DataProcessorTask {
	if t.Exists(section) {
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

// IsFormatKnown checks whether specified format is known to parser
func (t *DataProcessorTask) IsFormatKnown() bool {
	switch t.GetFormat() {
	case
		YAML,
		JSON:
		return true
	}
	return false
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
	return fmt.Errorf("unspecified format to unmarshal from")
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

// ReadFrom reads DataProcessorTask from specified file and tries to understand what is the format of the specified file
func (t *DataProcessorTask) ReadFrom(file string) error {
	if t == nil {
		return nil
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// Let's check, whether format is specified, and in case it is not, let's guess it from filename
	if !t.IsFormatKnown() {
		switch strings.ToLower(filepath.Ext(file)) {
		case
			".yaml",
			".yml":
			t.SetFormat(YAML)
		case
			".json":
			t.SetFormat(JSON)
		default:
			return fmt.Errorf("unable to understand what format task file shoud be read")
		}
	}

	return t.Unmarshal(b)
}

// String
func (t *DataProcessorTask) String() string {
	if t == nil {
		return ""
	}
	res := ""
	t.Walk(func(section string, items []string) bool {
		res += fmt.Sprintln(section+":", strings.Join(items, ":"))
		return true
	})
	res += fmt.Sprintln("root:", t.GetRootDir())
	res += fmt.Sprintln("task:", t.GetTaskFile())
	res += fmt.Sprintln("status:", t.GetStatus())
	res += fmt.Sprintln("errors:", strings.Join(t.GetErrors(), ","))
	return res
}
