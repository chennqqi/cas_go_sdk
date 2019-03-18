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
	"net/url"

	"github.com/google/subcommands"

	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func init() {
	subcommands.Register(&deleteVaultCmd{}, "")
}

type deleteVaultCmd struct {
	vaultName string
}

func (*deleteVaultCmd) Name() string     { return "delete_vault" }
func (*deleteVaultCmd) Synopsis() string { return "delete a vault" }
func (*deleteVaultCmd) Usage() string {
	return `delete_vault -vault <cas://vault-name>:
 	Delete 'vault-name' vault.
`
}

func (p *deleteVaultCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "format cas://vault-name")
}

func (p *deleteVaultCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}
	p.vaultName = parseVaultName(p.vaultName)

	client := openapi.NewAPIClient(conf)
	vault := client.VaultApi

	u, e := url.Parse(p.vaultName)
	if e != nil || p.vaultName == "" {
		fmt.Println("ERROR parse vault:", e)
		return subcommands.ExitFailure
	}
	var vaultName = u.Path
	_, err = vault.UIDVaultsVaultNameDelete(ctx, conf.AppId, vaultName)

	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}

	fmt.Println()
	return subcommands.ExitSuccess
}
