package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"

	//	"github.com/antihax/optional"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func init() {
	subcommands.Register(&descJobCmd{}, "")
}

type descJobCmd struct {
	vaultName string
	jobId     string
}

func (*descJobCmd) Name() string     { return "desc_job" }
func (*descJobCmd) Synopsis() string { return "get job status description" }
func (*descJobCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
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
	desc, _, err := job.UIDVaultsVaultNameJobsJobIDGet(ctx, conf.AppId, p.vaultName, p.jobId)
	if err != nil {
		fmt.Println("ERROR:", err)
		return subcommands.ExitFailure
	}
	fmt.Println(desc)

	fmt.Println()
	return subcommands.ExitSuccess
}
