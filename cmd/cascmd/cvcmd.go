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
	subcommands.Register(&cvCmd{new(createVaultCmd)}, "")
}

type cvCmd struct {
	*createVaultCmd
}

func (p *cvCmd) Name() string     { return "cv" }
func (p *cvCmd) Synopsis() string { return p.createVaultCmd.Synopsis() }
func (p *cvCmd) Usage() string {
	return p.createVaultCmd.Usage()
}

func (p *cvCmd) SetFlags(f *flag.FlagSet) {
	p.createVaultCmd.SetFlags(f)
}

func (p *cvCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return p.createVaultCmd.Execute(ctx, f)
}
