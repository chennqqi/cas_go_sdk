package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func init() {
	subcommands.Register(&fetchJobOutput{}, "")
}

type fetchJobOutput struct {
	vaultName string
	jobid     string
	localFile string
	force     bool
	start     int64
	size      int64
}

func (*fetchJobOutput) Name() string     { return "fetch_job_output" }
func (*fetchJobOutput) Synopsis() string { return "fetch job output." }
func (*fetchJobOutput) Usage() string {
	return `fetch_job_output [-capitalize] <some text>:
  fetch job output.
`
}

func (p *fetchJobOutput) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name`)
	f.StringVar(&p.jobid, "jobid", "", "jobId createjob returned")
	f.StringVar(&p.localFile, "local_file", "", "local file output written to'")
	f.BoolVar(&p.force, "-f", false, "force overwrite if file exists")
	f.Int64Var(&p.start, "--start", 0, "start position to download output retrieved, default to be 0")
	f.Int64Var(&p.size, "--size", 0, "size to download, default to be (totalsize - start)")
}

func (p *fetchJobOutput) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	client := openapi.NewAPIClient(nil)
	job := client.JobApi
	var opt openapi.UIDVaultsVaultNameJobsGetOpts

	desc, _, err := job.UIDVaultsVaultNameJobsJobIDGet(ctx, "-", p.vaultName, p.jobid, &opt)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	switch desc.(type) {
	case openapi.JobArchiveSearchInfo:
		info := (desc.(openapi.JobArchiveSearchInfo))

	case openapi.JobArchiveListSearchInfo:
		info := (desc.(openapi.JobArchiveListSearchInfo))

	case openapi.JobArchiveImportInfo:
		info := (desc.(openapi.JobArchiveImportInfo))

	case openapi.JobArchiveExportInfo:
		info := (desc.(openapi.JobArchiveExportInfo))
	}

	//	desc, _, err := job.UIDVaultsVaultNameJobsJobIDOutputGet(ctx, "-", p.vaultName, p.jobid, &opt)
	//	if err != nil {
	//		fmt.Println("ERROR:", err)
	//	}

	fmt.Println()
	return subcommands.ExitSuccess
}
