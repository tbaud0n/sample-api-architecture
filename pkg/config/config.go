package config

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// Config of the api application
type Config struct {
	// Environment : DEVELOPMENT, STAGING, PRODUCTION
	Environment string `yaml:"environment"`

	LogFilepath string `yaml:"logFilepath"`

	DatabaseDSN string `yaml:"databaseDSN"`
}

// NewConfig generate a new config from the yaml config file
func NewConfig(configFilename string) (*Config, error) {
	config := &Config{}

	filename, err := filepath.Abs(configFilename)
	if err != nil {
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
