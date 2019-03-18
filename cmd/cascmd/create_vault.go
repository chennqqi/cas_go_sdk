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
	subcommands.Register(&createVaultCmd{}, "")
}

type createVaultCmd struct {
	vaultName string
}

func (*createVaultCmd) Name() string     { return "create_vault" }
func (*createVaultCmd) Synopsis() string { return "create a vault" }
func (*createVaultCmd) Usage() string {
	return `create_vault format <cas://vaultname>':
  Create a vault as 'vaultname'.
`
}

func (p *createVaultCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name'`)
}

func (p *createVaultCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	p.vaultName = parseVaultName(p.vaultName)
	client := openapi.NewAPIClient(conf)
	vault := client.VaultApi

	resp, err := vault.CreateVault(ctx, conf.AppId, p.vaultName)
	if err != nil {
		if goe, ok := err.(openapi.GenericOpenAPIError); ok {
			fmt.Println("ERROR:", goe.Model())
		} else {
			fmt.Println("ERROR:", err)
		}
		return subcommands.ExitFailure
	}
	location := resp.Header.Get("Location")
	fmt.Printf(`Vault Location: %s\n`, location)
	fmt.Println()
	return subcommands.ExitSuccess
}
