package main

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Port              int      `yaml:"port"`
	Nodes             []string `yaml:"nodes"`
	Interval          int      `yaml:"check_interval"`
	BlockThreshold    int64    `yaml:"block_treshold"`
	ConnectionTimeout int      `yaml:"connection_timeout"`
}

func ParseConfig(configPath string) (Config, error) {
	yamlFile, err := ioutil.ReadFile(configPath)

	if err != nil {
		return Config{}, errors.Errorf("Given path doesn't exist: %v", configPath)
	}

	config := Config{}

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return Config{}, errors.Errorf("Unable to parse yaml: %v", configPath)
	}

	if len(config.Nodes) == 0 {
		return Config{}, errors.Errorf("Nodes are not defined")
	}

	return config, nil
}

func ParseConfigWPanic(configPath string) Config {
	config, err := ParseConfig(configPath)

	if err != nil {
		panic(err)
	}

	return config
}
