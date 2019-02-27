package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&uploadArchiveCmd{}, "")
}

type uploadArchiveCmd struct {
	localFile string
	vaultName string
	uploadId  string
	desc      string
	partSize  int
}

func (*uploadArchiveCmd) Name() string     { return "upload_archive" }
func (*uploadArchiveCmd) Synopsis() string { return "upload a local file" }
func (*uploadArchiveCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Upload a local file.
`
}

func (p *uploadArchiveCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name`)
	f.StringVar(&p.localFile, "local_file", "", "file to be uploaded")
	f.StringVar(&p.desc, "--desc", "", "description of the file")
	f.IntVar(&p.partSize, "--part-size", 0, "multipart upload part size")
}

func (p *uploadArchiveCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.localFile == "" {
		fmt.Println("local file is must parameter")
		return subcommands.ExitFailure
	}
	st, err := os.Stat(p.localFile)
	if os.IsNotExist(err) || err != nil {
		fmt.Println("ERROR local file:", err)
		return subcommands.ExitFailure
	}
	size := st.Size()
	var desc = p.desc
	if desc == "" {
		desc = filepath.Base(p.localFile)
	}

	fmt.Println()
	return subcommands.ExitSuccess
}
