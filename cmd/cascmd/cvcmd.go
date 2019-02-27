package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&cvCmd{new(createVaultCmd)}, "")
}

type cvCmd struct {
	*createVaultCmd
}

func (p *cvCmd) Name() string     { return "cv" }
func (p *cvCmd) Synopsis() string { return p.createVaultCmd.Synopsis() }
func (p *cvCmd) Usage() string {
	return p.createVaultCmd.Usage()
}

func (p *cvCmd) SetFlags(f *flag.FlagSet) {
	p.createVaultCmd.SetFlags(f)
}

func (p *cvCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return p.createVaultCmd.Execute(ctx, f)
}
