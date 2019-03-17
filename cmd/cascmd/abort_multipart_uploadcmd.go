package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func init() {
	subcommands.Register(&abortMultipartUploadCmd{}, "")
}

type abortMultipartUploadCmd struct {
	vaultName string
	uploadId  string
}

func (*abortMultipartUploadCmd) Name() string     { return "abort_multipart_upload" }
func (*abortMultipartUploadCmd) Synopsis() string { return "abort a multipart upload." }
func (*abortMultipartUploadCmd) Usage() string {
	return `abort_multipart_upload <params>:
  Abort a multipart upload..
`
}

func (p *abortMultipartUploadCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name`)
	f.StringVar(&p.uploadId, "upload_id", "", "ID of multipart upload to be aborted")
}

func (p *abortMultipartUploadCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}
	p.vaultName = parseVaultName(p.vaultName)

	client := openapi.NewAPIClient(conf)
	archive := client.ArchiveApi

	_, err = archive.UIDVaultsVaultNameMultipartUploadsUploadIDDelete(ctx,
		conf.AppId, p.vaultName, p.uploadId)
	if err != nil {
		fmt.Println("abort_multipart_upload ERROR:", err)
		return subcommands.ExitFailure
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
