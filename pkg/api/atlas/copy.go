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

package atlas

import (
	"io"

	log "github.com/sirupsen/logrus"
)

// cp copies from src into dst. Have to use `cp` because `copy` is a non-exported built-in function
func cp(dst io.Writer, src io.Reader, sizes ...int) (int64, error) {
	log.Tracef("cp() - start")
	defer log.Tracef("cp() - end")

	// Allocate relay buffer
	size := 32 * 1024
	if len(sizes) > 0 {
		size = sizes[0]
	}
	buf := make([]byte, size)

	// Copy data
	var written int64
	var err error
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
