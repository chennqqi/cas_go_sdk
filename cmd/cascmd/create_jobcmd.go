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
	openapi "github.com/chennqqi/cas_go_sdk/cas"
	"github.com/google/subcommands"
)

//Initiate Job
//https://cloud.tencent.com/document/product/572/8828

func init() {
	subcommands.Register(&createJobCmd{}, "")
}

type createJobCmd struct {
	vaultName string
	archiveId string
	start     int64
	size      int64
	desc      string
	tier      string
}

func (*createJobCmd) Name() string     { return "create_job" }
func (*createJobCmd) Synopsis() string { return "create a job." }
func (*createJobCmd) Usage() string {
	return `create_job [-options]:
  create a job.
`
}

func (p *createJobCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name`)
	f.StringVar(&p.archiveId, "archive_id", "", `ID of archive to be downloaded. If not provided, an inventory-retrieval job will be created`)
	f.Int64Var(&p.start, "--start", 0, "start position of archive to retrieve, default to be 0")
	f.Int64Var(&p.size, "--size", 0, "size to retrieve, default to be (totalsize - start)")
	f.StringVar(&p.desc, "--desc", "", "description of the job")
	f.StringVar(&p.tier, "--tier", "Expedited", `The retrieval option to use for the archive retrieval.[Expedited/Standard/Bulk]`)
}

func (p *createJobCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}
	p.vaultName = parseVaultName(p.vaultName)

	client := openapi.NewAPIClient(conf)
	job := client.JobApi

	var jtype = "archive-retrieval"
	if p.size != 0 || p.start != 0 {
		fmt.Println("Tip: Inventory-retrieval does NOT support range, ignored")
		p.start, p.size = 0, 0
	}

	var byteRange string
	if p.start != 0 && p.size != 0 {
		byteRange = fmt.Sprintf("%d-%d", p.start, p.start+p.size-1)
	} else if p.start != 0 && p.size == 0 {
		byteRange = fmt.Sprintf("%d-", p.start)
	} else if p.start == 0 && p.size != 0 {
		byteRange = fmt.Sprintf("0-%d", p.size-1)
	}
	if byteRange != "" {
		fmt.Println("Archive retrieval range:", byteRange)
	}

	var opt openapi.VaultsVaultNameJobsPostOpts
	opt.UNKNOWNBASETYPE = optional.NewInterface(openapi.JobArchiveSearchReq{
		Type:               jtype,
		ArchiveId:          p.archiveId,
		CallBackUrl:        "",
		Description:        p.desc,
		RetrievalByteRange: byteRange,
		Tier:               p.tier,
	})

	resp, err := job.VaultsVaultNameJobsPost(ctx, p.vaultName, &opt)
	if err != nil {
		fmt.Println("Error:", err)
		return subcommands.ExitFailure
	}
	jobId := resp.Header.Get("x-cas-job-id")
	fmt.Printf(`%s job created, job ID: %s\n`, jtype, jobId)
	fmt.Printf(`Use\n\n    cascmd fetch %s %s <local_file>\n\nto check job progress and download the data when job finished`,
		p.vaultName, job)
	fmt.Println(`NOTICE: Jobs usually take about 4 HOURS to complete.`)

	fmt.Println()
	return subcommands.ExitSuccess
}
