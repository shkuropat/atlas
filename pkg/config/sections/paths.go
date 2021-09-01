// Copyright 2021 The Atlas Authors. All rights reserved.
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

package sections

import (
	"fmt"
	"github.com/binarly-io/atlas/pkg/config/items"
	"os"
	"path/filepath"
)

type PathsOpts struct {
	// Base specifies the base on top of which to rebase relative paths.
	// In case base == nil no rebase
	// In case *base == "" use CWD as a base
	// Otherwise rebase on top of *base
	Base *string
	// Fallback specifies path which to fallback to in case nothing found
	// In case fallback == nil no fallback
	// In case *fallback == "" use CWD as a fallback
	// Otherwise fallback on *fallback
	Fallback *string
}

var (
	empty                = ""
	PathsOptsNothing     = &PathsOpts{}
	PathsOptsRebaseOnCWD = &PathsOpts{
		Base: &empty,
	}
	PathsOptsRebaseOnCWDFallbackOnCWD = &PathsOpts{
		Base:     &empty,
		Fallback: &empty,
	}
	PathsOptsDefault = PathsOptsNothing
)

// PathsConfigurator
type PathsConfigurator interface {
	GetPathsMap() map[string][]string
	GetPaths(name string, opts *PathsOpts) []string
	GetPathsOne(name string, opts *PathsOpts) string
}

// Interface compatibility
var _ PathsConfigurator = Paths{}

// Paths
type Paths struct {
	Paths *items.Paths `mapstructure:"paths"`
}

// PathsNormalize
func (c Paths) PathsNormalize() Paths {
	if c.Paths == nil {
		c.Paths = items.NewPaths()
	}
	return c
}

// GetPathsMap
func (c Paths) GetPathsMap() map[string][]string {
	if c.Paths == nil {
		return nil
	}
	return c.Paths.GetMap()
}

// GetPaths
func (c Paths) GetPaths(name string, opts *PathsOpts) []string {
	if c.Paths == nil {
		return nil
	}
	if opts == nil {
		opts = PathsOptsDefault
	}
	paths := c.Paths.GetPaths(name)

	// Fallback
	if len(paths) < 1 {
		switch {
		case opts.Fallback == nil:
			// No fallback
		case *opts.Fallback == "":
			// Fallback to CWD
			if cwd, err := os.Getwd(); err == nil {
				// CWD found, all is fine
				paths = []string{cwd}
			} else {
				// Unable to get CWD, fallback to root
				paths = []string{"/"}
			}
		default:
			// Fallback to explicitly specified path
			paths = []string{*opts.Fallback}
		}
	}

	// Rebase
	for i := range paths {
		switch {
		case opts.Base == nil:
			// No rebase
		case *opts.Base == "":
			// Rebase on top CWD
			base, err := os.Getwd()
			if err != nil {
				base = "/"
			}
			paths[i] = filepath.Clean(filepath.Join(base, paths[i]))
		default:
			// Rebase on top of explicitly specified path
			base := *opts.Base
			paths[i] = filepath.Clean(filepath.Join(base, paths[i]))
		}
	}

	return paths
}

// GetPathsOne
func (c Paths) GetPathsOne(name string, opts *PathsOpts) string {
	paths := c.GetPaths(name, opts)
	if len(paths) > 0 {
		return paths[0]
	}
	return ""
}

// String
func (c Paths) String() string {
	return fmt.Sprintf("Paths=%s", c.Paths)
}
