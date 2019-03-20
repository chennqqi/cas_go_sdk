// Copyright 2019 chennqqi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/chennqqi/cas_go_sdk/treehash"
	"github.com/google/subcommands"
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
