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

	openapi "github.com/chennqqi/cas/go"
)

func init() {
	subcommands.Register(&descVaultCmd{}, "")
}

type descVaultCmd struct {
	vaultName string
}

func (*descVaultCmd) Name() string     { return "desc_vault" }
func (*descVaultCmd) Synopsis() string { return "describe a vault" }
func (*descVaultCmd) Usage() string {
	return `desc_vault format <cas://vaultname>':
  Describe a vault as 'vaultname'.
`
}

func (p *descVaultCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "format <cas://vaultname>")
}

func (p *descVaultCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}
	p.vaultName = parseVaultName(p.vaultName)
	client := openapi.NewAPIClient(conf)
	vault := client.VaultApi

	info, _, err := vault.GetVault(ctx, conf.AppId, p.vaultName)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}
	fmt.Println("vault info:")
	fmt.Println(info)

	fmt.Println()
	return subcommands.ExitSuccess
}
