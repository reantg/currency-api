package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port                    string `yaml:"port"`
	DbUri                   string `yaml:"dbUri"`
	OpenexchangeratesApiKey string `yaml:"openexchangeratesApiKey"`
	OpenexchangeratesUrl    string `yaml:"openexchangeratesUrl"`
}

var ConfigData Config

func Init() error {
	rawYaml, err := os.Open("config.yml")
	if err != nil {
		return err
	}

	return yaml.NewDecoder(rawYaml).Decode(ConfigData)
}
