package main

import (
	"context"
	"flag"
	"fmt"

	"strings"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&fetchJobOutput{}, "")
}

type fetchJobOutput struct {
	capitalize bool
}

func (*fetchJobOutput) Name() string     { return "print" }
func (*fetchJobOutput) Synopsis() string { return "Print args to stdout." }
func (*fetchJobOutput) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *fetchJobOutput) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *fetchJobOutput) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if p.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
