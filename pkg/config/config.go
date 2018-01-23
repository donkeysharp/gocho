package config

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	NodeId         string `yaml:"NodeId"`
	WebPort        string `yaml:"WebPort"`
	LocalPort      string `yaml:"LocalPort"`
	ShareDirectory string `yaml:"ShareDirectory"`
	ConfigFile     string `yaml:"-"`
}

func (c *Config) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		return ""
	}
	return string(data)
}

func ConfigureWizard() error {
	return configureWizard()
}

func LoadConfig() (*Config, error) {
	configFile, err := getConfigFileName()
	if err != nil {
		return nil, err
	}

	if !fileExists(configFile) {
		return nil, fmt.Errorf("Error: Config file does not exist\nUse:\n\t$ gocho configure")
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	config := &Config{}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	config.ShareDirectory = CleanPath(config.ShareDirectory)
	return config, nil
}
