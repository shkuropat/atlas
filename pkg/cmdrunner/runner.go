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

package cmdrunner

import (
	"bytes"
	"io"
	"time"

	exe "github.com/go-cmd/cmd"
	log "github.com/sirupsen/logrus"
)

// Runner
type Runner struct {
	name   string
	args   []string
	cmd    *exe.Cmd
	status exe.Status
}

// New
func New(name string, args ...string) *Runner {
	return &Runner{
		name: name,
		args: args,
	}
}

// Run
func (r *Runner) Run(options *Options) exe.Status {
	log.Infof("Run() - start")
	defer log.Infof("Run() - end")

	// Start a long-running process, capture stdout and stderr
	r.cmd = exe.NewCmdOptions(
		options.GetOptions(),
		r.name,
		r.args...,
	)
	quitTick := r.startTicker(options)
	quitTimeout := r.startTimeout(options)
	log.Infof("wait for cmd to complete")

	// Start and wait for command to complete
	r.cmd.Start()
	<-r.cmd.Done()

	r.stopTicker(quitTick)
	r.stopTimeout(quitTimeout)

	r.status = r.cmd.Status()

	r.WriteOutput(options.GetStdoutWriter(), options.GetStderrWriter())

	return r.status
}

// WriteOutput
func (r *Runner) WriteOutput(stdout, stderr io.Writer) {
	log.Infof("WriteOutput() - start")
	defer log.Infof("WriteOutput() - end")

	if stdout != nil {
		n, err := io.Copy(stdout, r.GetStdoutReader())
		log.Infof("copied to stdout %d bytes. err: %v", n, err)
	}
	if stderr != nil {
		n, err := io.Copy(stderr, r.GetStderrReader())
		log.Infof("copied to stderr %d bytes. err: %v", n, err)
	}
}

// startTicker
func (r *Runner) startTicker(options *Options) chan bool {
	if options.HasTick() {
		quit := make(chan bool)
		// Print last line of stdout every `tick`
		log.Infof("ticker start")
		ticker := time.NewTicker(options.GetTick())
		go func() {
			for {
				select {
				case <-quit:
					log.Infof("ticker stop")
					ticker.Stop()
					return // func
				case <-ticker.C:
					log.Infof("ticker tick")
					status := r.cmd.Status()
					n := len(status.Stdout)
					if n > 0 {
						log.Infof("runtime:%f:%s", status.Runtime, status.Stdout[n-1])
					}
				}
			}
		}()
		return quit
	}

	return nil
}

// stopTicker
func (r *Runner) stopTicker(quit chan bool) {
	if quit == nil {
		return
	}

	quit <- true
}

// startTimeout
func (r *Runner) startTimeout(options *Options) chan bool {
	if options.HasTimeout() {
		quit := make(chan bool)
		// Stop command after specified `timeout`
		log.Infof("timeout start")
		go func() {
			select {
			case <-quit:
				log.Infof("timeout stop")
				return // func
			case <-time.After(options.GetTimeout()):
				log.Warnf("timout trigger")
				_ = r.cmd.Stop()
				return // func
			}
		}()
		return quit
	}

	return nil
}

// stopTimeout
func (r *Runner) stopTimeout(quit chan bool) {
	if quit == nil {
		return
	}

	quit <- true
}

// GetStdoutReader
func (r *Runner) GetStdoutReader() io.Reader {
	buf := &bytes.Buffer{}
	for i := range r.status.Stdout {
		buf.WriteString(r.status.Stdout[i])
		buf.WriteString("\n")
	}

	return buf
}

// GetStderrReader
func (r *Runner) GetStderrReader() io.Reader {
	buf := &bytes.Buffer{}
	for i := range r.status.Stderr {
		buf.WriteString(r.status.Stderr[i])
		buf.WriteString("\n")
	}

	return buf
}
