package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Port                    string `yaml:"port"`
	DbUri                   string `yaml:"dbUri"`
	OpenexchangeratesApiKey string `yaml:"openexchangeratesApiKey"`
	OpenexchangeratesUrl    string `yaml:"openexchangeratesUrl"`
}

var ConfigData ConfigStruct

func Init() error {
	rawYaml, err := os.ReadFile("config.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(rawYaml, &ConfigData)
	if err != nil {
		return err
	}

	return nil
}
