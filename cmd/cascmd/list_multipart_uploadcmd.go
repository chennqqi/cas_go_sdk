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

	"github.com/antihax/optional"
	openapi "github.com/chennqqi/cas_go_sdk/go"
	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&listMultriPartCmd{}, "")
}

type listMultriPartCmd struct {
	vaultName string
	marker    string
	limit     int64
}

func (*listMultriPartCmd) Name() string     { return "list_multipart_upload" }
func (*listMultriPartCmd) Synopsis() string { return "list all multipart uploads in a vault." }
func (*listMultriPartCmd) Usage() string {
	return `list_multipart_upload <params>:
  list all multipart uploads in a vault.
`
}

func (p *listMultriPartCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", "format cas://vault-name")
	f.StringVar(&p.marker, "marker", "", "list start multiupload position marker")
	f.Int64Var(&p.limit, "limit", 0, "number to be listed, max 1000")
}

func (p *listMultriPartCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	p.vaultName = parseVaultName(p.vaultName)
	client := openapi.NewAPIClient(conf)
	archive := client.ArchiveApi
	var opt openapi.UIDVaultsVaultNameMultipartUploadsGetOpts
	if p.marker != "" {
		opt.Marker = optional.NewString(p.vaultName)
		opt.Limit = optional.NewInt64(p.limit)
	}
	archive.UIDVaultsVaultNameMultipartUploadsGet(ctx, conf.AppId, p.vaultName, &opt)
	//TODO: fmt.Println result

	fmt.Println()
	return subcommands.ExitSuccess
}
