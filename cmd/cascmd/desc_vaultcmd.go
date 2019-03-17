package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"

	openapi "gogs.fastapi.org/gitadmin/cas/go"
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
