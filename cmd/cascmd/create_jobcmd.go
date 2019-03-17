package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/antihax/optional"
	"github.com/google/subcommands"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
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
	conf, err := loadConf("")
	if err != nil {
		fmt.Println("load conf error:", err)
		return subcommands.ExitFailure
	}
	p.vaultName = parseVaultName(p.vaultName)

	client := openapi.NewAPIClient(conf)
	job := client.JobApi

	var jtype = "archive-retrieval"
	if p.archiveId != "" {
		jtype = "inventory-retrieval"
		if p.size != 0 || p.start != 0 {
			fmt.Println("Tip: Inventory-retrieval does NOT support range, ignored")
			p.start, p.size = 0, 0
		}
		if p.marker != "" && (p.startDate != "" || p.endDate != "") {
			fmt.Println("Tip: Inventory-retrieval does NOT support start_date and end_date when marker is set, ignored")
			p.startDate, p.endDate = "", ""
		}
	} else {
		if p.marker != "" || p.limit != 0 {
			fmt.Println("Tip: Archive-retrieval does NOT support marker and limit, ignored")
			p.marker, p.limit = "", 0
		}
		if p.startDate != "" || p.endDate != "" {
			fmt.Println("Tip: Archive-retrieval does NOT support start_date and end_date, ignored")
			p.startDate, p.endDate = "", ""
		}
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

	var opt openapi.UIDVaultsVaultNameJobsPostOpts

	switch jtype {
	case "archive-retrieval":
		opt.UNKNOWNBASETYPE = optional.NewInterface(openapi.JobArchiveSearchReq{
			Type:               jtype,
			ArchiveId:          p.archiveId,
			CallBackUrl:        "",
			Description:        p.desc,
			RetrievalByteRange: byteRange,
			Tier:               p.tier,
		})
	case "inventory-retrieval":
		opt.UNKNOWNBASETYPE = optional.NewInterface(openapi.JobArchiveListSearchReq{
			// archive-retrieval
			Type:        jtype,
			CallBackUrl: "",
			// 默认JSON
			Format:      "JSON",
			Description: p.desc,
			InventoryRetrievalParameters: openapi.JobInventoryRetrievalParameters{
				StartDate: p.startDate,
				EndDate:   p.endDate,
				Marker:    p.marker,
				Limit:     fmt.Sprint("%d", p.limit),
			},
		})
	case "push-to-cos":
		fmt.Println("Not support posh-to-cols yet!")
		return subcommands.ExitFailure
		//TODO:
		opt.UNKNOWNBASETYPE = optional.NewInterface(openapi.JobArchiveImportReq{})
	case "pull-from-cos":
		//TODO:
		fmt.Println("Not support pull-from-cols yet!")
		return subcommands.ExitFailure
		opt.UNKNOWNBASETYPE = optional.NewInterface(openapi.JobArchiveExportReq{})
	}

	resp, err := job.UIDVaultsVaultNameJobsPost(ctx, conf.AppId, p.vaultName, &opt)
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
