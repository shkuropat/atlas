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
	// name specifies name of the executable to launch. Ex.: "/path/to/binary"
	name string
	// args specifies arguments of the executable. Ex.: ["-f=filename", "-v"]
	args []string

	cmd    *exe.Cmd
	status exe.Status

	// stopTickerChan specifies chan to stop periodical ticker
	stopTickerChan chan bool
	// stopTimeoutChan specifies chan to stop command timeout
	stopTimeoutChan chan bool
}

// New creates new runner having all executable command specified as args
func New(args ...string) *Runner {
	if len(args) == 0 {
		return nil
	}
	name := args[0]
	args = args[1:]
	return NewWithName(name, args...)
}

// NewWithName creates new cmdrunner having name explicitly specified
func NewWithName(name string, args ...string) *Runner {
	return &Runner{
		name: name,
		args: args,
	}
}

// Run
func (r *Runner) Run(options *Options) exe.Status {
	log.Infof("Run() - start")
	defer log.Infof("Run() - end")

	// Create new command. It is not launched here yet
	r.cmd = exe.NewCmdOptions(
		options.GetOptions(),
		r.name,
		r.args...,
	)
	r.startTicker(options)
	r.startTimeout(options)
	log.Infof("wait for cmd to complete")

	// Start command and wait for it to complete
	log.Infof("Starting command:\n%s %s", r.name, strings.Join(r.args, " "))
	r.cmd.Start()
	<-r.cmd.Done()

	r.stopTicker()
	r.stopTimeout()
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
func (r *Runner) startTicker(options *Options) {
	if !options.HasTick() {
		// No tick specified, unable to start ticker
		return
	}

	log.Infof("ticker start")

	// Chan to receive quit request
	r.stopTickerChan = make(chan bool)
	ticker := time.NewTicker(options.GetTick())
	go func() {
		for {
			select {
			case <-ticker.C:
				// Ticker's tick arrived. Time to log last line from stdout
				log.Infof("ticker tick")
				status := r.cmd.Status()
				n := len(status.Stdout)
				if n > 0 {
					log.Infof("runtime:%f:%s", status.Runtime, status.Stdout[n-1])
				}
			case <-r.stopTickerChan:
				// Quit request arrived
				log.Infof("ticker stop")
				ticker.Stop()
				return
			}
		}
	}()
}

// stopTicker sends quit request to specified chan
func (r *Runner) stopTicker() {
	r.stop(r.stopTickerChan)
}

// stop sends quit request to specified chan
func (r *Runner) stop(quit chan bool) {
	if quit == nil {
		return
	}

	close(quit)
}

// startTimeout starts goroutine which stops command after specified `timeout`
func (r *Runner) startTimeout(options *Options) {
	if !options.HasTimeout() {
		return
	}

	log.Infof("timeout start")

	// Chan to receive quit request
	r.stopTimeoutChan = make(chan bool)
	go func() {
		select {
		case <-time.After(options.GetTimeout()):
			// Time to stop the command
			log.Warnf("timout trigger")
			_ = r.cmd.Stop()
			return
		case <-r.stopTimeoutChan:
			// Quit request arrived
			log.Infof("timeout stop")
			return
		}
	}()
}

// stopTimeout sends quit request to specified chan
func (r *Runner) stopTimeout() {
	r.stop(r.stopTimeoutChan)
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
