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

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&fetchCmd{new(fetchJobOutputCmd)}, "")
}

type fetchCmd struct {
	*fetchJobOutputCmd
}

func (*fetchCmd) Name() string     { return "fetch" }
func (*fetchCmd) Synopsis() string { return "fetch job output." }
func (*fetchCmd) Usage() string {
	return `fetch <params>:
 fetch job output.
`
}

func (p *fetchCmd) SetFlags(f *flag.FlagSet) {
	p.fetchJobOutputCmd.SetFlags(f)
}

func (p *fetchCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return p.fetchJobOutputCmd.Execute(ctx, f)
}
