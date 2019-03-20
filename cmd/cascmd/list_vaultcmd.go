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

	"github.com/antihax/optional"
	"github.com/google/subcommands"

	openapi "github.com/chennqqi/cas_go_sdk/go"
)

func init() {
	subcommands.Register(&listVaultCmd{}, "")
}

type listVaultCmd struct {
	marker string
	limit  int /*1-1000*/
}

func (*listVaultCmd) Name() string     { return "list_vault" }
func (*listVaultCmd) Synopsis() string { return "list vaults" }
func (*listVaultCmd) Usage() string {
	return `list_vault <params>:
  list vaults.
`
}

func (p *listVaultCmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&p.limit, "limit", 0, "number of vaults to be listed, max 1000")
	f.StringVar(&p.marker, "marker", "", "list start position marker")
}

func (p *listVaultCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	client := openapi.NewAPIClient(conf)
	vault := client.VaultApi

	var opt openapi.UIDVaultsGetOpts
	if p.limit != 0 {
		opt.Limit = optional.NewInt64(int64(p.limit))
	}
	if p.marker != "" {
		opt.Marker = optional.NewString(p.marker)
	}

	sm, resp, err := vault.UIDVaultsGet(ctx, conf.AppId, &opt)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return subcommands.ExitFailure
	}
	if resp.StatusCode == 200 {
		fmt.Println("Marker:", sm.Marker)
		fmt.Println("Vault count:", len(sm.VaultList))
	}
	for i := 0; i < len(sm.VaultList); i++ {
		fmt.Println(sm.VaultList[i])
	}

	fmt.Println()
	return subcommands.ExitSuccess
}
