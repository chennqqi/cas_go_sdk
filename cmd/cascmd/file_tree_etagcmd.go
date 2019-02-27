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
	return `print [-capitalize] <some text>:
  Print args to stdout.
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
