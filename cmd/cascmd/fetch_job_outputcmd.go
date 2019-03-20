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
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/antihax/optional"
	"github.com/google/subcommands"
	openapi "github.com/chennqqi/cas/go"
	"gopkg.in/cheggaaa/pb.v1"
)

func init() {
	subcommands.Register(&fetchJobOutputCmd{}, "")
}

type fetchJobOutputCmd struct {
	vaultName string
	jobid     string
	localFile string
	force     bool
	start     int64
	size      int64
}

func (*fetchJobOutputCmd) Name() string     { return "fetch_job_output" }
func (*fetchJobOutputCmd) Synopsis() string { return "fetch job output." }
func (*fetchJobOutputCmd) Usage() string {
	return `fetch_job_output <params>:
  fetch job output.
`
}

func (p *fetchJobOutputCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name`)
	f.StringVar(&p.jobid, "jobid", "", "jobId createjob returned")
	f.StringVar(&p.localFile, "local_file", "", "local file output written to'")
	f.BoolVar(&p.force, "-f", false, "force overwrite if file exists")
	f.Int64Var(&p.start, "--start", 0, "start position to download output retrieved, default to be 0")
	f.Int64Var(&p.size, "--size", 0, "size to download, default to be (totalsize - start)")
}

func (p *fetchJobOutputCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	p.vaultName = parseVaultName(p.vaultName)
	client := openapi.NewAPIClient(conf)
	job := client.JobApi

	desc, _, err := job.UIDVaultsVaultNameJobsJobIDGet(ctx, conf.AppId, p.vaultName, p.jobid)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	var status string
	var reqRetrievalByteRange string
	if desc.JobArchiveSearchInfo != nil {
		req := desc.JobArchiveSearchInfo
		status = req.StatusCode
		reqRetrievalByteRange = req.RetrievalByteRange
	} else if desc.JobArchiveListSearchInfo != nil {
		req := desc.JobArchiveListSearchInfo
		status = req.StatusCode
	} else if desc.JobArchiveExportInfo != nil {
		req := desc.JobArchiveExportInfo
		status = req.StatusCode
		fmt.Println("Unsupport PullFromCOS")
		return subcommands.ExitFailure
	} else if desc.JobArchiveImportInfo != nil {
		req := desc.JobArchiveImportInfo
		status = req.StatusCode
		fmt.Println("Unsupport PullFromCOS")
		return subcommands.ExitFailure
	}
	var jtype = "archive-retrieval"

	status = strings.ToLower(status)
	if status == "inprogress" {
		fmt.Println(`%s job still in progress. Repeat this 'command later'`, jtype)
		return subcommands.ExitSuccess
	} else if status == "failed" {
		fmt.Println("%s job failed.", jtype)
		return subcommands.ExitSuccess
	}

	var rangeA, rangeB int64
	if jtype == "archive-retrieval" {
		rangeValues := strings.Split(reqRetrievalByteRange, "-")
		if len(rangeValues) > 0 {
			fmt.Sscanf(rangeValues[0], "%d", &rangeA)
		}
		if len(rangeValues) > 1 {
			fmt.Sscanf(rangeValues[1], "%d", &rangeB)
		}
	} else {
		fmt.Sscanf(desc.InventorySizeInBytes, "%d", &rangeB)
		rangeB -= 1
	}
	if p.start == 0 && p.size == 0 {
		p.size = rangeB - rangeA - p.start + 1
	} else if p.start == 0 && p.size != 0 {
		p.start = 0
	}

	var outputOpt openapi.UIDVaultsVaultNameJobsJobIDOutputGetOpts
	if p.start != 0 && p.size != 0 {
		outputOpt.Range_ = optional.NewString(fmt.Sprintf("bytes=%d-%d", p.start, p.size-1))
	}
	_, err = os.Stat(p.localFile)
	if err == nil {
		fmt.Println("Output file %s existed. Do you wish to overwrite it? (y/n):")
		var c byte
		fmt.Scanf("%c", &c)
		if c != 'y' {
			fmt.Println(`Answer is no. Quit now.`)
			return subcommands.ExitSuccess
		}
	}
	os.Remove(p.localFile)

	//get job output
	output, resp, err := job.UIDVaultsVaultNameJobsJobIDOutputGet(ctx,
		conf.AppId, p.vaultName, p.jobid, &outputOpt)

	if resp.Header.Get("Content-Type") != "application/octet-stream" {
		fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
		fmt.Println("output:", output)
		return subcommands.ExitSuccess
	}
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fp, err := os.Create(p.localFile)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}
	defer fp.Close()

	var totalSize int64
	if p.size != 0 {
		totalSize = p.size
	} else {
		totalSize = rangeB - rangeA + 1
	}
	bar := pb.StartNew(int(totalSize))

	wfp := bufio.NewWriter(fp)
	var totalRead int64

	for {
		l, err := io.CopyN(wfp, resp.Body, 1024*1024)
		if l > 0 {
			totalRead += l
			if jtype == "inventory-retrieval" {
				continue
			}
			bar.Add64(l)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("COPY ERROR:", err)
			return subcommands.ExitFailure
		}
	}

	bar.FinishPrint("Download job output success")
	if jtype == "inventory-retrieval" && desc.InventoryRetrievalParameters.Marker != "" {
		fmt.Println(`NOTICE: Want more archive list? Create a new job with  --marker `, desc.InventoryRetrievalParameters.Marker)
	}

	fmt.Println()
	return subcommands.ExitSuccess
}
