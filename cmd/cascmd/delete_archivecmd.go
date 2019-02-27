package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"

	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func init() {
	subcommands.Register(&deleteArchiveCmd{}, "")
}

type deleteArchiveCmd struct {
	vaultName string
	archiveId string
}

func (*deleteArchiveCmd) Name() string     { return "delete_archive" }
func (*deleteArchiveCmd) Synopsis() string { return "Delete an archive." }
func (*deleteArchiveCmd) Usage() string {
	return `delete_archive -vault <cas://vault-name> -achive_id <achiveid>:
  Delete an archive.
`
}

func (p *deleteArchiveCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "format cas://vault-name")
	f.StringVar(&p.archiveId, "archive_id", "", "ID of archive to be deleted")
}

func (p *deleteArchiveCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	client := openapi.NewAPIClient(nil)
	archive := client.ArchiveApi

	_, err := archive.UIDVaultsVaultNameMultipartUploadsUploadIDDelete(ctx, p.vaultName, "-", p.archiveId)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}

	fmt.Println()
	return subcommands.ExitSuccess
}
