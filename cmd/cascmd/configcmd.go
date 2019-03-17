package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/chennqqi/goutils/jsonconfig"
	"github.com/google/subcommands"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func init() {
	subcommands.Register(&configCmd{}, "")
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
	f.StringVar(&p.endPoint, "-endpoint", "", "cas endpoint output")
	f.StringVar(&p.appId, "-appid", "", "user appid")
	f.StringVar(&p.secretId, "-secretid", "", "user secretid")
	f.StringVar(&p.secretKey, "-secretkey", "", "user secretkey")
	f.StringVar(&p.configFile, "-config-file", "", "file to save configuration")
}

func (p *configCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	var conf string
	if p.configFile != "" {
		conf = DEFAULT_CONFIG_FILE
	}

	var cfg openapi.Configuration

	//load old config
	//TODO: use loadConf
	err := jsonconfig.Load(&cfg, conf)
	if err != nil {
		fmt.Println("config error:", err)
		return subcommands.ExitFailure
	}

	if p.endPoint != "" {
		cfg.Host = p.endPoint
	}
	if p.appId != "" {
		cfg.AppId = p.appId
	}
	if p.secretId != "" {
		cfg.AccessKey = p.secretId
	}
	if p.secretKey != "" {
		cfg.AccessSecret = p.secretKey
	}
	if err := jsonconfig.Save(cfg, conf); err != nil {
		fmt.Println("config ERROR:", err)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
