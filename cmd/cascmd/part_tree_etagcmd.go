package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"gogs.fastapi.org/gitadmin/cas/treehash"
)

func init() {
	subcommands.Register(&partTreeEtagCmd{}, "")
}

type partTreeEtagCmd struct {
	localFile string
	start     int64
	end       int64
}

func (*partTreeEtagCmd) Name() string { return "part_tree_etag" }
func (*partTreeEtagCmd) Synopsis() string {
	return "calculate tree sha256 hash of a multipart upload part."
}
func (*partTreeEtagCmd) Usage() string {
	return `part_tree_etag <params>:
  calculate tree sha256 hash of a multipart upload part.
`
}

func (p *partTreeEtagCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.localFile, "local_file", "", "file to be read from")
	f.Int64Var(&p.start, "start", 0, "start position to read")
	f.Int64Var(&p.end, "end", 0, "end position to read")
}

func (p *partTreeEtagCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.end%(1024*1024) == 0 {
		p.end = p.end - 1
	}
	var size = p.end - p.start + 1
	ch, th, e := treehash.ComputeHashFromFile(p.localFile, p.start, size, 1024*1024)
	if e != nil {
		fmt.Println("ERROR:", e)
		return subcommands.ExitFailure
	}
	fmt.Println("content-hash:", ch)
	fmt.Println("tree-hash:", th)

	fmt.Println()
	return subcommands.ExitSuccess
}
