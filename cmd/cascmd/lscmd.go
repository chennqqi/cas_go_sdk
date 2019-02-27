package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&lsCmd{new(listVaultCmd)}, "")
}

type lsCmd struct {
	*listVaultCmd
}

func (*lsCmd) Name() string     { return "ls" }
func (*lsCmd) Synopsis() string { return "list all vaults" }
func (c *lsCmd) Usage() string {
	return c.listVaultCmd.Usage()
}

func (p *lsCmd) SetFlags(f *flag.FlagSet) {
	p.listVaultCmd.SetFlags(f)
}

func (p *lsCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return p.listVaultCmd.Execute(ctx, f)
}
