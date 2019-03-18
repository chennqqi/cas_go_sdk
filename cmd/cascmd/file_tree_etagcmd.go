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
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/google/subcommands"
	"gogs.fastapi.org/gitadmin/cas/treehash"
)

func init() {
	subcommands.Register(&fileTreeCmd{}, "")
}

type fileTreeCmd struct {
	localFile string
}

func (*fileTreeCmd) Name() string     { return "file_tree_etag" }
func (*fileTreeCmd) Synopsis() string { return "calculate tree sha256 hash of a file" }
func (*fileTreeCmd) Usage() string {
	return `file_tree_etag <params>:
  calculate tree sha256 hash of a file.
`
}

func (p *fileTreeCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.localFile, "local_file", "", "file to be calculated")
}

func (p *fileTreeCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fp, err := os.Open(p.localFile)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}
	defer fp.Close()

	cacheReader := bufio.NewReader(fp)

	h := sha256.New()
	tree := treehash.New(1024*1024, sha256.New())
	mw := io.MultiWriter(h, tree)
	io.Copy(mw, cacheReader)

	contentHash := hex.EncodeToString(h.Sum(nil))
	treeHash := hex.EncodeToString(tree.Sum(nil))

	fmt.Println("content-hash:", contentHash)
	fmt.Println("tree-hash,", treeHash)

	fmt.Println()
	return subcommands.ExitSuccess
}
