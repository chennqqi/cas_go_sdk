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

	"time"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&signCmd{}, "")
}

const (
	TIME_LAYOUT = "2006-01-02 15:03:04"
)

type signCmd struct {
	secret  string
	start   int64
	end     int64
	startAt string
	endAt   string
}

func (*signCmd) Name() string     { return "sign" }
func (*signCmd) Synopsis() string { return "sign cascmd." }
func (*signCmd) Usage() string {
	return `sign <params>:
	sign -secret <secret> -start <start unix> -end <end unix)
	sign -secret <secret> -start-at '2020-01-02 00:00:00' -end-at '2021-12-30 00:00:00'
  calculate sign key.
`
}

func (p *signCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.secret, "secret", "", "user api secret, required")

	f.Int64Var(&p.start, "start", 0, "set signkey start unix timestamp, default now")
	f.Int64Var(&p.end, "end", 0, "set signkey end unix timestamp")

	f.StringVar(&p.startAt, "start-at", "", "set signkey start at time, format: "+TIME_LAYOUT)
	f.StringVar(&p.endAt, "end-at", "", "set signkey end at time, format: "+TIME_LAYOUT)
}

func (p *signCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if p.secret == "" {
		fmt.Println("secret option is required")
		return subcommands.ExitUsageError
	}

	if p.end == 0 && p.endAt == "" {
		fmt.Println("ether end or end-at is required")
		return subcommands.ExitUsageError
	}
	if p.end != 0 {
		if p.start == 0 {
			p.start = time.Now().Unix()
		}
		rangeTime := fmt.Sprintf("%d;%d", p.start, p.end)

		h := hmac.New(sha1.New, []byte(p.secret))
		h.Write([]byte(rangeTime))
		fmt.Println("range:", rangeTime)
		fmt.Println("signkey:", hex.EncodeToString(h.Sum(nil)))
	} else {
		if p.startAt == "" {
			p.startAt = time.Now().Format(TIME_LAYOUT)
		}
		ts, err := time.Parse(TIME_LAYOUT, p.startAt)
		if err != nil {
			fmt.Println("parse start-at", p.startAt, err)
			return subcommands.ExitFailure
		}
		te, err := time.Parse(TIME_LAYOUT, p.endAt)
		if err != nil {
			fmt.Println("parse end-at", p.startAt, err)
			return subcommands.ExitFailure
		}

		rangeTime := fmt.Sprintf("%d;%d", ts.Unix(), te.Unix())
		h := hmac.New(sha1.New, []byte(p.secret))
		h.Write([]byte(rangeTime))
		fmt.Println("range:", rangeTime)
		fmt.Println("signkey:", hex.EncodeToString(h.Sum(nil)))
	}

	fmt.Println()
	return subcommands.ExitSuccess
}
