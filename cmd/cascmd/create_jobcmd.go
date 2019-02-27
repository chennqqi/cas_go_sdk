package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"

	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func init() {
	subcommands.Register(&createJobCmd{}, "")
}

type createJobCmd struct {
	vaultName string
	archiveId string
	start     int64
	size      int64
	desc      string
	limit     int64
	marker    string
	startDate string
	endDate   string
	tier      string
}

func (*createJobCmd) Name() string     { return "print" }
func (*createJobCmd) Synopsis() string { return "Print args to stdout." }
func (*createJobCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *createJobCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.vaultName, "vault", "", `format cas://vault-name`)
	f.StringVar(&p.archiveId, "archive_id", "", `ID of archive to be downloaded. If not provided, an inventory-retrieval job will be created`)
	f.Int64Var(&p.start, "--start", 0, "start position of archive to retrieve, default to be 0")
	f.Int64Var(&p.size, "--size", 0, "size to retrieve, default to be (totalsize - start)")
	f.StringVar(&p.desc, "--desc", "", "description of the job")
	f.Int64Var(&p.limit, "--limit", 0, "number of archives to be listed, default to be 10000")
	f.StringVar(&p.marker, "--marker", "", "list start position marker")
	f.StringVar(&p.startDate, "--start_date", "", "the start date of archive created, format: YYYY-MM-DDThh:mm:ssZ")
	f.StringVar(&p.endDate, "--end_date", "", "the end date of archive created, format: YYYY-MM-DDThh:mm:ssZ")
	f.StringVar(&p.tier, "--tier", "", `The retrieval option to use for the archive retrieval. Standard is the default value used.`)
}

func (p *createJobCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	client := openapi.NewAPIClient(openapi.NewConfiguration())
	job := client.JobApi
	var opt openapi.UIDVaultsVaultNameJobsPostOpts

	//TODO:
	job.UIDVaultsVaultNameJobsPost(ctx, "-", p.vaultName, &opt)

	fmt.Println()
	return subcommands.ExitSuccess
}
