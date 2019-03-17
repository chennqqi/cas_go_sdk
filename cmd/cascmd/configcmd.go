package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
)

func init() {
}

type configCmd struct {
	endPoint   string
	appId      string
	secretId   string
	secretKey  string
	configFile string
}

func (*configCmd) Name() string     { return "config" }
func (*configCmd) Synopsis() string { return "config cascmd." }
func (*configCmd) Usage() string {
	return `config <params>:
  save config to file.
`
}

func (p *configCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.endPoint, "-endpoint", "", "cas endpoint host")
	f.StringVar(&p.appId, "-appid", "", "user appid")
	f.StringVar(&p.secretId, "-secretid", "", "user secretid")
	f.StringVar(&p.secretKey, "-secretkey", "", "user secretkey")
	f.StringVar(&p.configFile, "-config-file", "", "file to save configuration")
}

func (p *configCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf(p.configFile)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("config error:", err)
		return subcommands.ExitFailure
	}

	if p.endPoint != "" {
		conf.Host = p.endPoint
	}
	if p.appId != "" {
		conf.AppId = p.appId
	}
	if p.secretId != "" {
		conf.AccessKey = p.secretId
	}
	if p.secretKey != "" {
		conf.AccessSecret = p.secretKey
	}
	if err := saveConf(p.configFile, conf); err != nil {
		fmt.Println("config ERROR:", err)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
