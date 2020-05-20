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
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/araddon/dateparse"
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
	expireAt   string
}

func (*configCmd) Name() string     { return "config" }
func (*configCmd) Synopsis() string { return "config cascmd." }
func (*configCmd) Usage() string {
	return `config <params>:
  save config to file.
`
}

func (p *configCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.endPoint, "endpoint", "", "cas endpoint host")
	f.StringVar(&p.appId, "appid", "", "user appid")
	f.StringVar(&p.secretId, "secretid", "", "user secretid")
	f.StringVar(&p.secretKey, "secretkey", "", "user secretkey")
	f.StringVar(&p.expireAt, "expireat", "", "set signkey expire")
	f.StringVar(&p.configFile, "config-file", "", "file to save configuration")
}

func (p *configCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	conf, err := loadConf(p.configFile)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("config error:", err)
		return subcommands.ExitFailure
	}

	if p.endPoint != "" {
		if strings.HasPrefix(p.endPoint, "http://") {
			conf.BasePath = p.endPoint
			conf.Host = p.endPoint[7:]
		} else if strings.HasPrefix(p.endPoint, "https://") {
			conf.BasePath = p.endPoint
			conf.Host = p.endPoint[8:]
		} else {
			conf.BasePath = "http://" + p.endPoint
			conf.Host = p.endPoint
		}
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
	if p.expireAt != "" {
		now := time.Now()
		at, err := dateparse.ParseLocal(p.expireAt)
		if err != nil {
			fmt.Println("config parse expire time ERROR:", err)
			return subcommands.ExitFailure
		} else if at.Before(now) {
			fmt.Println("expire time should after now")
			return subcommands.ExitFailure
		}
		start := now.Unix()
		end := at.Unix()
		conf.SignKeyStart = start
		conf.SignKeyExpire = time.Duration(end-start) * time.Second
		h := hmac.New(sha1.New, []byte(conf.AccessSecret))
		h.Write([]byte(fmt.Sprintf("%d;%d", start, end)))
		conf.SignKey = hex.EncodeToString(h.Sum(nil))
		//clear secret, only use signKey
		conf.AccessSecret = ""
	}

	if err := saveConf(p.configFile, conf); err != nil {
		fmt.Println("config ERROR:", err)
		return subcommands.ExitFailure
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
