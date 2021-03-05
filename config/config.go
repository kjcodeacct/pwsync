package config

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Service    string `yaml:"service"`
	Executable string `yaml:"exec"`
}

func Open(filepath string) (*Config, error) {

	buff, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	newConfig := &Config{}
	err = yaml.Unmarshal(buff, &newConfig)
	if err != nil {
		return nil, err
	}

	return newConfig, nil
}
