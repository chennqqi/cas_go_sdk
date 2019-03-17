package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/chennqqi/goutils/jsonconfig"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
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
	if strings.HasPrefix("cas://", vault) {
		return vault[len("cas://"):]
	}
	return vault
}
