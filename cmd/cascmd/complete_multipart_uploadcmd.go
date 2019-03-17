package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func init() {
	subcommands.Register(&completeMultiPartCmd{}, "")
}

type completeMultiPartCmd struct {
	vaultName string
	uploadId  string
	size      int64
	treeTag   string
}

func (*completeMultiPartCmd) Name() string     { return "complete_multipart_upload" }
func (*completeMultiPartCmd) Synopsis() string { return "complete the multipart upload." }
func (*completeMultiPartCmd) Usage() string {
	return `complete_multipart_upload <params>:
  complete the multipart upload.
`
}

func (p *completeMultiPartCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "vault where the upload initiated")
	f.StringVar(&p.uploadId, "upload_id", "", "ID create multipartupload returned")
	f.Int64Var(&p.size, "size", 0, "size of the file")
	f.StringVar(&p.treeTag, "tree_etag", "", "tree sha256 hash vaule of the file")
}

func (p *completeMultiPartCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	p.vaultName = parseVaultName(p.vaultName)

	client := openapi.NewAPIClient(conf)
	archive := client.ArchiveApi

	resp, err := archive.UIDVaultsVaultNameMultipartUploadsUploadIDPost(ctx,
		conf.AppId, p.vaultName, p.uploadId, p.treeTag, fmt.Sprintf("%d", p.size))
	if err != nil {
		fmt.Println("complete_multipart_upload ERROR:", err)
		return subcommands.ExitFailure
	}

	location := resp.Header.Get("Location")
	archiveId := resp.Header.Get("x-cas-archive-id")
	fmt.Println("Location:", location)
	fmt.Println("Archive ID:", archiveId)

	fmt.Println()
	return subcommands.ExitSuccess
}
