package main

import (
	"strings"

	"github.com/chennqqi/goutils/jsonconfig"
	openapi "gogs.fastapi.org/gitadmin/cas/go"
)

func loadConf(name string) (*openapi.Configuration, error) {
	if name == "" {
		name = DEFAULT_CONFIG_FILE
	}
	conf := openapi.NewConfiguration()
	err := jsonconfig.Load(conf, name)
	if err != nil {
		jsonconfig.Save(*conf, name)
	}
	return conf, err
}

func saveConf(name string, conf *openapi.Configuration) error {
	if name == "" {
		name = DEFAULT_CONFIG_FILE
	}
	return jsonconfig.Save(*conf, name)
}

func parseVaultName(vault string) string {
	if strings.HasPrefix("cas://", vault) {
		return vault[len("cas://"):]
	}
	return vault
}
