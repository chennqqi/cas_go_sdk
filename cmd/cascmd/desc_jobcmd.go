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

	openapi "github.com/chennqqi/cas_go_sdk/cas"
	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&descJobCmd{}, "")
}

type descJobCmd struct {
	vaultName string
	jobId     string
}

func (*descJobCmd) Name() string     { return "desc_job" }
func (*descJobCmd) Synopsis() string { return "describe a job" }
func (*descJobCmd) Usage() string {
	return `desc_job [-options]:
  describe a job.
`
}

func (p *descJobCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name`)
	f.StringVar(&p.jobId, "jobid", "", "the id of createjob returned")
}

func (p *descJobCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}
	p.vaultName = parseVaultName(p.vaultName)
	client := openapi.NewAPIClient(conf)

	job := client.JobApi
	desc, _, err := job.VaultsVaultNameJobsJobIDGet(ctx, p.vaultName, p.jobId)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}
	fmt.Println(desc)

	fmt.Println()
	return subcommands.ExitSuccess
}
