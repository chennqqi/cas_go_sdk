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
	DEFAULT_CONFIG_FILE = `~/.cascmdgo.json`
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(&configCmd{}, "")

	//	subcommands.Register(subcommands.FlagsCommand(), "")
	//	subcommands.Register(subcommands.CommandsCommand(), "")
	flag.Parse()

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
