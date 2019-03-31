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
	"fmt"
	"os"
	"runtime"
	"strings"

	openapi "github.com/chennqqi/cas_go_sdk/go"
	"github.com/chennqqi/goutils/jsonconfig"
)

func homeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

func fixPath(name string) string {
	if name == "" {
		name = DEFAULT_CONFIG_FILE
	}
	if strings.HasPrefix(name, "~/") {
		name = strings.Replace(name, "~/", homeDir()+"/", 1)
	}
	return name
}

func loadConf(name string) (*openapi.Configuration, error) {
	name = fixPath(name)
	conf := openapi.NewConfiguration()
	err := jsonconfig.Load(conf, name)
	if err != nil {
		fmt.Println("load error", err)
		jsonconfig.Save(*conf, name)
	}
	fmt.Println("LOAD CONF:", conf)
	if conf.AppId == "" {
		conf.AppId = "-"
	}
	return conf, err
}

func saveConf(name string, conf *openapi.Configuration) error {
	name = fixPath(name)
	return jsonconfig.Save(*conf, name)
}

func parseVaultName(vault string) string {
	if strings.HasPrefix(vault, "cas://") {
		return vault[len("cas://"):]
	}
	return vault
}
