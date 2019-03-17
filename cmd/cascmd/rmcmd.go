package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
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
	return `print [-capitalize] <some text>:
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
