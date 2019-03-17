package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/antihax/optional"
	"github.com/google/subcommands"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
	"gogs.fastapi.org/gitadmin/cas/treehash"
)

type postArchiveCmd struct {
	vaultName, localFile, desc string
}

func (*postArchiveCmd) Name() string     { return "post_archive" }
func (*postArchiveCmd) Synopsis() string { return "upload a local file" }
func (*postArchiveCmd) Usage() string {
	return `post_archive <params>:
  Upload a local file.
`
}

func (p *postArchiveCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name`)
	f.StringVar(&p.localFile, "local_file", "", "file to be uploaded")
	f.StringVar(&p.desc, "-desc", "", "description of the file")
}

func (p *postArchiveCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.localFile == "" {
		fmt.Println("local file is must parameter")
		return subcommands.ExitFailure
	}
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	p.vaultName = parseVaultName(p.vaultName)
	client := openapi.NewAPIClient(conf)
	archive := client.ArchiveApi

	st, err := os.Stat(p.localFile)
	if os.IsNotExist(err) || err != nil {
		fmt.Println("ERROR local file:", err)
		return subcommands.ExitFailure
	}
	size := st.Size()
	var desc = p.desc
	if desc == "" {
		desc = filepath.Base(p.localFile)
	}
	tag, treeEtag, err := treehash.ComputeHashFromFile(p.localFile, 0, 0, 0)

	fp, err := os.Open(p.localFile)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}
	defer fp.Close()

	buf := bufio.NewReader(fp)
	r, w := io.Pipe()
	defer r.Close()
	errChan := make(chan error, 1)

	go func() {
		defer w.Close()
		if _, err := io.CopyN(w, buf, size); err != nil {
			errChan <- err
			return
		}
		errChan <- nil
		close(errChan)
	}()

	var opt openapi.UIDVaultsVaultNameArchivesPostOpts
	opt.XCasArchiveDescription = optional.NewString(desc)

	contentLength := fmt.Sprintf("%d", size)
	resp, err1 := archive.UIDVaultsVaultNameArchivesPost(ctx,
		conf.AppId, p.vaultName, tag, contentLength, treeEtag, r, &opt,
	)
	err2 := <-errChan
	if err1 != nil {
		fmt.Println("ERROR:", err1)
		return subcommands.ExitFailure
	}
	if err2 != nil {
		fmt.Println("ERROR:", err2)
		return subcommands.ExitFailure
	}
	fmt.Println("Archive ID: %s", resp.Header.Get("x-cas-archive-id"))
	fmt.Println()
	return subcommands.ExitSuccess
}
