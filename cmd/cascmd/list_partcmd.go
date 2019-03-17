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
	subcommands.Register(&listPartCmd{}, "")
}

type listPartCmd struct {
	vaultName string
	uploadId  string
	marker    string
	limit     int64
}

func (*listPartCmd) Name() string     { return "list_part" }
func (*listPartCmd) Synopsis() string { return "list all parts uploaded in one upload." }
func (*listPartCmd) Usage() string {
	return `list_part <params>:
  list all parts uploaded in one upload.
`
}

func (p *listPartCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "format cas://vault-name")
	f.StringVar(&p.uploadId, "upload_id", "", "ID of multipart upload")
	f.StringVar(&p.marker, "-marker", "", "list start multiupload position marker")
	f.Int64Var(&p.limit, "-marker", 0, "number to be listed, max 1000")
}

func (p *listPartCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	p.vaultName = parseVaultName(p.vaultName)
	client := openapi.NewAPIClient(conf)
	archive := client.ArchiveApi

	var opt openapi.UIDVaultsVaultNameMultipartUploadsUploadIDGetOpts
	if p.marker != "" {
		opt.Marker = optional.NewString(p.vaultName)
		opt.Limit = optional.NewInt64(p.limit)
	}

	///<UID>/vaults/<VaultName>/multipart-uploads/<uploadID>
	parts, resp, err := archive.UIDVaultsVaultNameMultipartUploadsUploadIDGet(ctx,
		conf.AppId, p.vaultName, p.uploadId, &opt)
	if err != nil {
		fmt.Println("Error:", err)
		return subcommands.ExitFailure
	}
	//TODO: print result

	fmt.Println()
	return subcommands.ExitSuccess
}
