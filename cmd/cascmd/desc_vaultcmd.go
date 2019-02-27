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
	client := openapi.NewAPIClient(nil)
	vault := client.VaultApi

	u, e := url.Parse(p.vaultName)
	if e != nil || p.vaultName == "" {
		fmt.Println("ERROR parse vault:", e)
		return subcommands.ExitFailure
	}
	var vaultName = u.Path
	info, _, err := vault.GetVault(ctx, "-", vaultName)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}
	fmt.Println("vault info:")
	fmt.Println(info)

	fmt.Println()
	return subcommands.ExitSuccess
}
