package main

import (
	"context"
	"flag"
	"fmt"

	//	"github.com/antihax/optional"
	"github.com/google/subcommands"

	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func init() {
	subcommands.Register(&uploadPartCmd{}, "")
}

type uploadPartCmd struct {
	vaultName string
	uploadId  string
	localFile string
	start     int64
	end       int64
	eTag      string
	treeTag   string
}

func (*uploadPartCmd) Name() string     { return "upload_part" }
func (*uploadPartCmd) Synopsis() string { return "upload one part." }
func (*uploadPartCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *uploadPartCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "format cas://vault-name")
	f.StringVar(&p.uploadId, "upload_id", "", "ID createmupload returned")
	f.StringVar(&p.localFile, "local_file", "", "file to read from")
	f.Int64Var(&p.start, "start", 0, "read start position, start must be divided by partsize")
	f.Int64Var(&p.end, "end", 0, `read end position, end+1 must be the size of file or partsize larger than start`)
	f.StringVar(&p.eTag, "etag", "", "sha256 hash value")
	f.StringVar(&p.treeTag, "tree_tag", "", "tree sha256 hash value of part")
}

func (p *uploadPartCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	client := openapi.NewAPIClient(openapi.NewConfiguration())
	archive := client.ArchiveApi

	resp, err := archive.UIDVaultsVaultNameMultipartUploadsUploadIDPost(ctx,
		"-", p.vaultName, p.uploadId, p.eTag, p.treeTag,
	)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println()
	return subcommands.ExitSuccess
}
