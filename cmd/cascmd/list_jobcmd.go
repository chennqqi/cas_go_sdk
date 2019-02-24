package main

import (
	"context"
	"flag"
	"fmt"

	"strings"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&listJobCmd{}, "")
}

type listJobCmd struct {
	capitalize bool
}

func (*listJobCmd) Name() string     { return "print" }
func (*listJobCmd) Synopsis() string { return "Print args to stdout." }
func (*listJobCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *listJobCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *listJobCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if p.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
