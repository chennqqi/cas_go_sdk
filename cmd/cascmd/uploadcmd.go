package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&uploadCmd{new(uploadArchiveCmd)}, "")
}

type uploadCmd struct {
	*uploadArchiveCmd
}

func (*uploadCmd) Name() string     { return "upload" }
func (*uploadCmd) Synopsis() string { return "upload a local file." }
func (*uploadCmd) Usage() string {
	return `upload <params>:
  upload a local file.
`
}

func (p *uploadCmd) SetFlags(f *flag.FlagSet) {
	p.uploadArchiveCmd.SetFlags(f)
}

func (p *uploadCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return p.uploadArchiveCmd.Execute(ctx, f)
}
