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
	"strings"
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
	stopTickerChan := r.startTicker(options)
	stopTimeoutChan := r.startTimeout(options)
	log.Infof("wait for cmd to complete")

	// Start command and wait for it to complete
	log.Infof("Starting command:\n%s %s", r.name, strings.Join(r.args, " "))
	r.cmd.Start()
	<-r.cmd.Done()

	r.stopTicker(stopTickerChan)
	r.stopTimeout(stopTimeoutChan)
	r.status = r.cmd.Status()

	r.WriteOutput(options.GetStdoutWriter(), options.GetStderrWriter())

	return r.status
}

// WriteOutput writes output into provided stdout and stderr writers from run app's stdout and stderr
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

// startTicker starts goroutine which prints last line of stdout every `tick`
// Returns chan where to send quit/stop request
func (r *Runner) startTicker(options *Options) chan bool {
	if !options.HasTick() {
		// No tick specified, unable to start ticker
		return nil
	}

	log.Infof("ticker start")

	// Chan to receive quit request
	quit := make(chan bool)
	ticker := time.NewTicker(options.GetTick())
	go func() {
		for {
			select {
			case <-ticker.C:
				// Time to log last line from stdout
				log.Infof("ticker tick")
				status := r.cmd.Status()
				n := len(status.Stdout)
				if n > 0 {
					log.Infof("runtime:%f:%s", status.Runtime, status.Stdout[n-1])
				}
			case <-quit:
				// Quit request arrived
				log.Infof("ticker stop")
				ticker.Stop()
				return
			}
		}
	}()
	return quit
}

// stop sends quit request to specified chan
func (r *Runner) stop(quit chan bool) {
	if quit == nil {
		return
	}

	close(quit)
}

// stopTicker sends quit request to specified chan
func (r *Runner) stopTicker(quit chan bool) {
	r.stop(quit)
}

// startTimeout starts goroutine which stops command after specified `timeout`
func (r *Runner) startTimeout(options *Options) chan bool {
	if !options.HasTimeout() {
		return nil
	}

	log.Infof("timeout start")

	// Chan to receive quit request
	quit := make(chan bool)
	go func() {
		select {
		case <-time.After(options.GetTimeout()):
			// Time to stop the command
			log.Warnf("timout trigger")
			_ = r.cmd.Stop()
			return // func
		case <-quit:
			// Quit request arrived
			log.Infof("timeout stop")
			return
		}
	}()
	return quit
}

// stopTimeout sends quit request to specified chan
func (r *Runner) stopTimeout(quit chan bool) {
	r.stop(quit)
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
