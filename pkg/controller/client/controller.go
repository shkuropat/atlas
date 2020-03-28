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

package controller_client

import (
	log "github.com/golang/glog"
	pb "github.com/sunsingerus/mservice/pkg/api/mservice"
	"github.com/sunsingerus/mservice/pkg/transiever/client"
	"io/ioutil"
	"os"
)

type FileGetter struct {
	filename string
}

func NewFileGetter(filename string) *FileGetter {
	return &FileGetter{
		filename: filename,
	}
}

func (f *FileGetter) Get() ([]byte, bool) {
	data, _ := ioutil.ReadFile(f.filename)
	return data, true
}

func SendFile(client pb.MServiceControlPlaneClient, filename string) {
	if _, err := os.Stat(filename); err == nil {
		log.Infof("Has file %s", filename)
		if f, err := os.Open(filename); err == nil {
			log.Infof("START send file %s", filename)
			n, err := transiever_client.StreamDataChunks(client, f)
			log.Infof("DONE send file %s size %d err %v", filename, n, err)
		} else {
			log.Infof("ERROR open file %s", filename)
		}
	}
}

func SendStdin(client pb.MServiceControlPlaneClient) {
	n, err := transiever_client.StreamDataChunks(client, os.Stdin)
	log.Infof("DONE send %s size %d err %v", os.Stdin.Name(), n, err)
}
