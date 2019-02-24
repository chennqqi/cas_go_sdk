package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&abortMultipartUploadCmd{}, "")
}

type abortMultipartUploadCmd struct {
	capitalize bool
}

func (*abortMultipartUploadCmd) Name() string     { return "print" }
func (*abortMultipartUploadCmd) Synopsis() string { return "Print args to stdout." }
func (*abortMultipartUploadCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *abortMultipartUploadCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *abortMultipartUploadCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	for _, arg := range f.Args() {
		if p.capitalize {
			arg = strings.ToUpper(arg)
		}
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
