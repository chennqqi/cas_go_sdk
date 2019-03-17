package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&fetchCmd{new(fetchJobOutputCmd)}, "")
}

type fetchCmd struct {
	*fetchJobOutputCmd
}

func (*fetchCmd) Name() string     { return "fetch" }
func (*fetchCmd) Synopsis() string { return "fetch job output." }
func (*fetchCmd) Usage() string {
	return `fetch <params>:
 fetch job output.
`
}

func (p *fetchCmd) SetFlags(f *flag.FlagSet) {
	p.fetchJobOutputCmd.SetFlags(f)
}

func (p *fetchCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return p.fetchJobOutputCmd.Execute(ctx, f)
}
