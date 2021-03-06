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

	"github.com/google/subcommands"

	"github.com/antihax/optional"
	openapi "github.com/chennqqi/cas_go_sdk/cas"
)

func init() {
	subcommands.Register(&listJobCmd{}, "")
}

type listJobCmd struct {
	vaultName string
	marker    string
	limit     int
}

func (*listJobCmd) Name() string     { return "list_job" }
func (*listJobCmd) Synopsis() string { return "list all jobs except expired" }
func (*listJobCmd) Usage() string {
	return `list_job <params>:
  list all jobs except expired.
`
}

func (p *listJobCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name`)
	f.StringVar(&p.marker, "-marker", "", `start list job position marker`)
	f.IntVar(&p.limit, "-limit", 10, `number to be listed`)
}

func (p *listJobCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}
	client := openapi.NewAPIClient(conf)
	job := client.JobApi

	var opt openapi.VaultsVaultNameJobsGetOpts
	if p.marker != "" {
		opt.Marker = optional.NewString(p.marker)
	}
	if p.limit != 0 {
		opt.Limit = optional.NewInt64(int64(p.limit))
	}

	jobs, _, err := job.VaultsVaultNameJobsGet(ctx, p.vaultName, &opt)
	if err != nil {
		ge := err.(openapi.GenericOpenAPIError)
		fmt.Println("ERROR:", ge.Model(), err)
		return subcommands.ExitFailure
	}
	fmt.Println("Marker:", jobs.Marker)
	fmt.Println("Job count:", len(jobs.JobList))
	if len(jobs.JobList) == 0 {
		fmt.Println()
		return subcommands.ExitSuccess
	}
	//TODO:
	fmt.Println(jobs.JobList)

	fmt.Println()
	return subcommands.ExitSuccess
}
