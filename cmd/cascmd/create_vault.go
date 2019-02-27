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
	client := openapi.NewAPIClient(nil)
	vault := client.VaultApi

	u, e := url.Parse(p.vaultName)
	if e != nil || p.vaultName == "" {
		fmt.Println("ERROR parse vault:", e)
		return subcommands.ExitFailure
	}
	var vaultName = u.Path

	resp, err := vault.CreateVault(ctx, "-", vaultName)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}
	location := resp.Header.Get("Location")
	fmt.Printf(`Vault Location: %s\n`, location)
	fmt.Println()
	return subcommands.ExitSuccess
}
