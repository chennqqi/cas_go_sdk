package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/antihax/optional"
	"github.com/google/subcommands"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func init() {
	subcommands.Register(&listMultriPartCmd{}, "")
}

type listMultriPartCmd struct {
	vaultName string
	marker    string
	limit     int64
}

func (*listMultriPartCmd) Name() string     { return "list_multipart_upload" }
func (*listMultriPartCmd) Synopsis() string { return "list all multipart uploads in a vault." }
func (*listMultriPartCmd) Usage() string {
	return `print [-capitalize] <some text>:
  list all multipart uploads in a vault.
`
}

func (p *listMultriPartCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "format cas://vault-name")
	f.StringVar(&p.marker, "marker", "", "list start multiupload position marker")
	f.Int64Var(&p.limit, "limit", 0, "number to be listed, max 1000")
}

func (p *listMultriPartCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	client := openapi.NewAPIClient(openapi.NewConfiguration())
	archive := client.ArchiveApi
	var opt openapi.UIDVaultsVaultNameMultipartUploadsGetOpts
	if p.marker != "" {
		opt.Marker = optional.NewString(p.vaultName)
		opt.Limit = optional.NewInt64(p.limit)
	}
	archive.UIDVaultsVaultNameMultipartUploadsGet(ctx, "-", p.vaultName, &opt)
	//TODO: fmt.Println result

	fmt.Println()
	return subcommands.ExitSuccess
}
