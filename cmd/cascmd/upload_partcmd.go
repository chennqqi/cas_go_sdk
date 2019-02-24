package main

import (
	"context"
	"flag"
	"fmt"

	"strings"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&uploadPartCmd{}, "")
}

type uploadPartCmd struct {
	capitalize bool
}

func (*uploadPartCmd) Name() string     { return "print" }
func (*uploadPartCmd) Synopsis() string { return "Print args to stdout." }
func (*uploadPartCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *uploadPartCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *uploadPartCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if p.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
