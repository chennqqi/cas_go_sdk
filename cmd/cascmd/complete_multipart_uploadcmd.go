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

func (*completeMultiPartCmd) Name() string     { return "print" }
func (*completeMultiPartCmd) Synopsis() string { return "Print args to stdout." }
func (*completeMultiPartCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *completeMultiPartCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "vault where the upload initiated")
	f.StringVar(&p.uploadId, "upload_id", "", "ID create multipartupload returned")
	f.Int64Var(&p.size, "size", 0, "size of the file")
	f.StringVar(&p.treeTag, "tree_etag", "", "tree sha256 hash vaule of the file")
}

func (p *completeMultiPartCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	client := openapi.NewAPIClient(openapi.NewConfiguration())
	archive := client.ArchiveApi

	//TODO: empty check
	resp, err := archive.UIDVaultsVaultNameMultipartUploadsUploadIDPost(ctx, "-",
		p.vaultName, p.uploadId, p.treeTag, fmt.Sprintf("%d", p.size))
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}

	location := resp.Header.Get("Location")
	archiveId := resp.Header.Get("x-cas-archive-id")
	fmt.Println("Location:", location)
	fmt.Println("archiveId:", archiveId)

	fmt.Println()
	return subcommands.ExitSuccess
}
