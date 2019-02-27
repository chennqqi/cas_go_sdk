package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	openapi "gogs.fastapi.org/gitadmin/cas/go"

	"github.com/antihax/optional"
)

func init() {
	subcommands.Register(&initMultipartUploadCmd{}, "")
}

type initMultipartUploadCmd struct {
	vaultName string
	partSize  int64
	desc      string
}

func (*initMultipartUploadCmd) Name() string     { return "print" }
func (*initMultipartUploadCmd) Synopsis() string { return "initiate a multipart upload." }
func (*initMultipartUploadCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *initMultipartUploadCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "capitalize output")
	f.Int64Var(&p.partSize, "part_size", 0, "size of each multipart upload")
	f.StringVar(&p.desc, "--desc", "", "capitalize output")
}

func (p *initMultipartUploadCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	client := openapi.NewAPIClient(openapi.NewConfiguration())
	archive := client.ArchiveApi

	var opt openapi.UIDVaultsVaultNameMultipartUploadsPostOpts
	if p.desc != "" {
		opt.XCasArchiveDescription = optional.NewString(p.desc)
	}

	resp, err := archive.UIDVaultsVaultNameMultipartUploadsPost(ctx,
		"-", p.vaultName, fmt.Sprintf("%d", p.partSize), &opt)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	location := resp.Header.Get("Location")
	uploadId := resp.Header.Get("x-cas-multipart-upload-id")
	fmt.Println("Location:", location)
	fmt.Println("uploadId:", uploadId)

	return subcommands.ExitSuccess
}
