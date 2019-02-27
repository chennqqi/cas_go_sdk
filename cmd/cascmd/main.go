package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/subcommands"
)

func usage() {
	fmt.Println(`
usage: cascmd.py [-h] cmd ...

optional arguments:
  -h, --help            show this help message and exit

Supported actions:
  Commands {ls, cv, rm, upload, create_job, fetch} provide easier ways to
  use CAS by combining commands below them. Generally they will suffice your
  daily use. For advanced operations, use commands like {createvault...}

  cmd
    config              config cascmd
    help                show a detailed help message and exit
    ls                  list all vaults
    cv                  create a vault
    rm                  remove a vault or an archive
    upload              upload a local file
    create_job          create an inventory/archive retrieval job
    fetch               fetch job output
    create_vault        create a vault
    delete_vault        delete a vault
    list_vault          list all vaults
    desc_vault          get detailed vault description
    upload_archive      upload a local file
    delete_archive      delete an archive
    file_tree_etag      calculate tree sha256 hash of a file
    part_tree_etag      calculate tree sha256 hash of a multipart upload part
    init_multipart_upload
                        initiate a multipart upload
    abort_multipart_upload
                        abort a multipart upload
    list_multipart_upload
                        list all multipart uploads in a vault
    upload_part         upload one part
    list_part           list all parts uploaded in one upload
    complete_multipart_upload
                        complete the multipart upload
    desc_job            get job status description
    fetch_job_output    fetch job output
    list_job            list all jobs except expired`)
}

const (
	//	DEFAULT_CONFIG_FILE = os.path.expanduser('~') + `/.cascmd_credentials`
	CONFIG_SECTION = `CASCredentials`
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	//subcommands.Register(&printCmd{}, "")
	//	var endpoint, appid, secretid, secretkey string
	//	flag.StringVar(&endpoint, "endpoint", "", "endpoint")
	//	flag.StringVar(&appid, "endpoint", "", "endpoint")
	//	flag.StringVar(&secretid, "endpoint", "", "endpoint")
	//	flag.StringVar(&secretkey, "endpoint", "", "endpoint")
	//	flag.Parse()

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
