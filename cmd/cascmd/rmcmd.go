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
	subcommands.Register(&rmCmd{}, "")
}

type rmCmd struct {
	vaultName string
	uploadId  string
}

func (*rmCmd) Name() string     { return "rm" }
func (*rmCmd) Synopsis() string { return "remove a vault or an archive." }
func (*rmCmd) Usage() string {
	return `rm <params>:
  remove a vault or an archive.
`
}

func (p *rmCmd) SetFlags(f *flag.FlagSet) {
	flag.StringVar(&p.vaultName, "vault", "", "format cas://vault-name")
	flag.StringVar(&p.uploadId, "upload_id", "", "ID of archive to be deleted")
}

func (p *rmCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	p.vaultName = parseVaultName(p.vaultName)
	client := openapi.NewAPIClient(conf)
	vault := client.VaultApi
	archive := client.ArchiveApi

	if p.uploadId == "" {
		_, err = vault.UIDVaultsVaultNameDelete(ctx,
			conf.AppId, p.vaultName)
	} else {
		_, err = archive.UIDVaultsVaultNameArchivesArchiveIDDelete(ctx,
			conf.AppId, p.vaultName, p.uploadId)
	}
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println()
	return subcommands.ExitSuccess
}
