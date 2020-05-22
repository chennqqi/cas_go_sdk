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
	"context"
	"flag"
	"fmt"

	openapi "github.com/chennqqi/cas_go_sdk/cas"
	"github.com/google/subcommands"

	"github.com/antihax/optional"
)

func init() {
	subcommands.Register(&initMultipartUploadCmd{}, "")
}

type initMultipartUploadCmd struct {
	vaultName string
	partSize  int64
	desc      string
}

func (*initMultipartUploadCmd) Name() string     { return "init_multipart_upload" }
func (*initMultipartUploadCmd) Synopsis() string { return "initiate a multipart upload." }
func (*initMultipartUploadCmd) Usage() string {
	return `init_multipart_upload <params>:
  initiate a multipart upload.
`
}

func (p *initMultipartUploadCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "capitalize output")
	f.Int64Var(&p.partSize, "part_size", 0, "size of each multipart upload")
	f.StringVar(&p.desc, "--desc", "", "capitalize output")
}

func (p *initMultipartUploadCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	p.vaultName = parseVaultName(p.vaultName)
	client := openapi.NewAPIClient(conf)
	archive := client.ArchiveApi

	var opt openapi.VaultsVaultNameMultipartUploadsPostOpts
	if p.desc != "" {
		opt.XCasArchiveDescription = optional.NewString(p.desc)
	}

	resp, err := archive.VaultsVaultNameMultipartUploadsPost(ctx,
		p.vaultName, fmt.Sprintf("%d", p.partSize), &opt)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	location := resp.Header.Get("Location")
	uploadId := resp.Header.Get("x-cas-multipart-upload-id")
	fmt.Println("Location:", location)
	fmt.Println("uploadId:", uploadId)

	return subcommands.ExitSuccess
}
