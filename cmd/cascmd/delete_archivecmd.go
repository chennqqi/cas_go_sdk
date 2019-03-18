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

	"github.com/google/subcommands"

	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func init() {
	subcommands.Register(&deleteArchiveCmd{}, "")
}

type deleteArchiveCmd struct {
	vaultName string
	archiveId string
}

func (*deleteArchiveCmd) Name() string     { return "delete_archive" }
func (*deleteArchiveCmd) Synopsis() string { return "Delete an archive." }
func (*deleteArchiveCmd) Usage() string {
	return `delete_archive -vault <cas://vault-name> -achive_id <achiveid>:
  Delete an archive.
`
}

func (p *deleteArchiveCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "format cas://vault-name")
	f.StringVar(&p.archiveId, "archive_id", "", "ID of archive to be deleted")
}

func (p *deleteArchiveCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}
	p.vaultName = parseVaultName(p.vaultName)

	client := openapi.NewAPIClient(conf)
	archive := client.ArchiveApi

	_, err = archive.UIDVaultsVaultNameMultipartUploadsUploadIDDelete(ctx,
		conf.AppId, p.vaultName, p.archiveId)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}

	fmt.Println()
	return subcommands.ExitSuccess
}
