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
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}
	location := resp.Header.Get("Location")
	fmt.Printf(`Vault Location: %s\n`, location)
	fmt.Println()
	return subcommands.ExitSuccess
}
