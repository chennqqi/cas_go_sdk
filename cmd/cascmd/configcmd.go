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
	"os"
	"time"

	"github.com/google/subcommands"
)

func init() {
}

type configCmd struct {
	region string
	appId  string
	key    string
	secret string
	expire string

	sign  string
	start int64
	end   int64

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
	f.StringVar(&p.region, "region", "beijing", "cas region [beijing/shanghai/guangzhou/chengdu]")
	f.StringVar(&p.appId, "appid", "-", "user appid")
	f.StringVar(&p.key, "key", "", "user api key, required")

	f.StringVar(&p.secret, "secret", "", "user api secret, using secret mode")
	f.StringVar(&p.expire, "expire", "86400s", "set access secret expire")

	f.StringVar(&p.sign, "sign", "", "set signkey, using signkey mode")
	f.Int64Var(&p.start, "start", 0, "set signkey start, if 'sign' set, this opition is required")
	f.Int64Var(&p.end, "end", 0, "set signkey expire,if 'sign' set, this opition is required")

	f.StringVar(&p.configFile, "config-file", "", "file to save configuration")
}

func (p *configCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf(p.configFile)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("config error:", err)
		return subcommands.ExitFailure
	}

	if p.key == "" {
		fmt.Println("ERROR: api key is required")
		return subcommands.ExitFailure
	}
	conf.AppId = p.appId
	conf.AccessKey = p.key
	conf.Region = p.region
	du, _ := time.ParseDuration(p.expire)

	if p.secret != "" {
		conf.AccessSecret = p.secret
		conf.SecretExpire = du
	} else {
		conf.SignKey = p.sign
		conf.SignKeyStart = p.start
		conf.SignKeyEnd = p.end
	}

	if err := saveConf(p.configFile, conf); err != nil {
		fmt.Println("config ERROR:", err)
		return subcommands.ExitFailure
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
