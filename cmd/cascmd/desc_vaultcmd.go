package main

import (
	"context"
	"flag"
	"fmt"

	"strings"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&descVaultCmd{}, "")
}

type descVaultCmd struct {
	capitalize bool
}

func (*descVaultCmd) Name() string     { return "print" }
func (*descVaultCmd) Synopsis() string { return "Print args to stdout." }
func (*descVaultCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *descVaultCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *descVaultCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if p.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
