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

	openapi "github.com/chennqqi/cas_go_sdk/go"
	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&abortMultipartUploadCmd{}, "")
}

type abortMultipartUploadCmd struct {
	vaultName string
	uploadId  string
}

func (*abortMultipartUploadCmd) Name() string     { return "abort_multipart_upload" }
func (*abortMultipartUploadCmd) Synopsis() string { return "abort a multipart upload." }
func (*abortMultipartUploadCmd) Usage() string {
	return `abort_multipart_upload <params>:
  Abort a multipart upload..
`
}

func (p *abortMultipartUploadCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name`)
	f.StringVar(&p.uploadId, "upload_id", "", "ID of multipart upload to be aborted")
}

func (p *abortMultipartUploadCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}
	p.vaultName = parseVaultName(p.vaultName)

	client := openapi.NewAPIClient(conf)
	archive := client.ArchiveApi

	_, err = archive.UIDVaultsVaultNameMultipartUploadsUploadIDDelete(ctx,
		conf.AppId, p.vaultName, p.uploadId)
	if err != nil {
		fmt.Println("abort_multipart_upload ERROR:", err)
		return subcommands.ExitFailure
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
