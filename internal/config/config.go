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
	ForceShutdownTimeout    int    `yaml:"forceShutdownTimeout"`
}

func Init() (*Config, error) {

	rawYaml, err := os.Open("config.yml")
	if err != nil {
		return nil, err
	}

	configData := Config{}
	err = yaml.NewDecoder(rawYaml).Decode(&configData)
	if err != nil {
		return nil, err
	}
	return &configData, nil
}
