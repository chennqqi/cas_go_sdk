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
	openapi "gogs.fastapi.org/gitadmin/cas/go"
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
	return `fetch_job_output [-capitalize] <some text>:
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
	if desc.JobArchiveSearchInfo != nil {
		req := desc.JobArchiveSearchInfo
		status = req.StatusCode
	} else if desc.JobArchiveListSearchInfo != nil {
		req := desc.JobArchiveListSearchInfo
		status = req.StatusCode
	} else if desc.JobArchiveExportInfo != nil {
		req := desc.JobArchiveExportInfo
		status = req.StatusCode
	} else if desc.JobArchiveImportInfo != nil {
		req := desc.JobArchiveImportInfo
		status = req.StatusCode
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
	var brange interface{}
	if jtype == "archive-retrieval" {
		// brange = [int(x) for x in job['RetrievalByteRange'].split('-')]
	} else {
		//brange = [0, int(job['InventorySizeInBytes']) - 1]
	}
	var output_range interface{}

	var outputOpt openapi.UIDVaultsVaultNameJobsJobIDOutputGetOpts

	//TODO:
	outputOpt.Range_ = optional.NewString(fmt.Sprintf("%d-%d"))

	output, resp, err := job.UIDVaultsVaultNameJobsJobIDOutputGet(ctx,
		conf.AppId, p.vaultName, p.jobid, &outputOpt)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fp, err := os.Create(p.localFile)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}
	wfp := bufio.NewWriter(fp)
	var totalRead int64
	//	for {
	//		io.CopyN(wfp, resp.Body, )
	//	}
	//TODO:
	io.Copy(wfp, resp.Body)

	fmt.Println()
	return subcommands.ExitSuccess
}
