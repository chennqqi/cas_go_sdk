package main

import (
	"context"
	"flag"
	"fmt"

	"strings"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&createJobCmd{}, "")
}

type createJobCmd struct {
	capitalize bool
}

func (*createJobCmd) Name() string     { return "print" }
func (*createJobCmd) Synopsis() string { return "Print args to stdout." }
func (*createJobCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *createJobCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *createJobCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if p.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
