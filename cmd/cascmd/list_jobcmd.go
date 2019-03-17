package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"

	"github.com/antihax/optional"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
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
	f.StringVar(&p.marker, "--marker", "", `start list job position marker`)
	f.IntVar(&p.limit, "--limit", 0, `number to be listed`)
}

func (p *listJobCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}

	client := openapi.NewAPIClient(conf)
	job := client.JobApi

	var opt openapi.UIDVaultsVaultNameJobsGetOpts
	if p.marker != "" {
		opt.Marker = optional.NewString(p.marker)
	}
	if p.limit != 0 {
		opt.Limit = optional.NewInt64(int64(p.limit))
	}

	jobs, _, err := job.UIDVaultsVaultNameJobsGet(ctx, "-", p.vaultName, &opt)
	if err != nil {
		fmt.Println("ERROR:", err)
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
