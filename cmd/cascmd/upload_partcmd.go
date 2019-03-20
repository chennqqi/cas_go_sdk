// Copyright 2019 chennqqi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/antihax/optional"
	"github.com/google/subcommands"

	openapi "github.com/chennqqi/cas_go_sdk/go"
)

func init() {
	subcommands.Register(&uploadPartCmd{}, "")
}

type uploadPartCmd struct {
	vaultName string
	uploadId  string
	localFile string
	start     int64
	end       int64
	eTag      string
	treeTag   string
}

func (*uploadPartCmd) Name() string     { return "upload_part" }
func (*uploadPartCmd) Synopsis() string { return "upload one part." }
func (*uploadPartCmd) Usage() string {
	return `upload_part <params>:
  upload one part..
`
}

func (p *uploadPartCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "format cas://vault-name")
	f.StringVar(&p.uploadId, "upload_id", "", "ID createmupload returned")
	f.StringVar(&p.localFile, "local_file", "", "file to read from")
	f.Int64Var(&p.start, "start", 0, "read start position, start must be divided by partsize")
	f.Int64Var(&p.end, "end", 0, `read end position, end+1 must be the size of file or partsize larger than start`)
	f.StringVar(&p.eTag, "etag", "", "sha256 hash value")
	f.StringVar(&p.treeTag, "tree_tag", "", "tree sha256 hash value of part")
}

func (p *uploadPartCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	p.vaultName = parseVaultName(p.vaultName)
	client := openapi.NewAPIClient(conf)
	archive := client.ArchiveApi

	var opt openapi.UIDVaultsVaultNameMultipartUploadsUploadIDPutOpts
	var size = p.end - p.start

	fp, err := os.Open(p.localFile)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}
	defer fp.Close()
	fp.Seek(p.start, os.SEEK_SET)

	buf := bufio.NewReader(fp)
	r, w := io.Pipe()
	defer r.Close()
	errChan := make(chan error, 1)

	go func() {
		defer w.Close()
		if _, err := io.CopyN(w, buf, size); err != nil {
			errChan <- err
			return
		}
		errChan <- nil
		close(errChan)
	}()

	opt.ContentLength = optional.NewString(fmt.Sprintf("%d", size))
	contentRange := fmt.Sprintf("bytes %d-%d", p.start, p.end)

	resp, err1 := archive.UIDVaultsVaultNameMultipartUploadsUploadIDPut(ctx,
		conf.AppId, p.vaultName, p.uploadId, contentRange, p.eTag, p.treeTag, &opt, r,
	)
	err2 := <-errChan
	if err1 != nil {
		fmt.Println("ERROR:", err1)
	}
	if err2 != nil {
		fmt.Println("ERROR:", err2)
	}
	fmt.Println("x-cas-sha256-tree-hash:", resp.Header.Get("x-cas-sha256-tree-hash"))

	fmt.Println()
	return subcommands.ExitSuccess
}
