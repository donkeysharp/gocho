package config

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	homedir "github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func CleanPath(str string) string {
	return strings.TrimRight(str, string(os.PathSeparator))
}

func writeConfigToFile(c *Config, fileName string) error {
	data := []byte(c.String())
	return ioutil.WriteFile(fileName, data, 0644)
}

func getConfigFileName() (string, error) {
	userHome, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	configFile := fmt.Sprintf("%s%c%s", userHome, os.PathSeparator, ".gocho.conf")
	return configFile, nil
}

func getDefaultConfig() (*Config, error) {
	configFile, err := getConfigFileName()
	if err != nil {
		return nil, err
	}

	defaultWebPort := "5555"
	defaultLocalPort := "1337"
	defaultNodeId := randomdata.SillyName()
	currentUser, err := user.Current()
	if err == nil {
		defaultNodeId = currentUser.Username
	}

	config := &Config{
		ShareDirectory: "",
		WebPort:        defaultWebPort,
		LocalPort:      defaultLocalPort,
		NodeId:         defaultNodeId,
		ConfigFile:     configFile,
	}
	return config, nil
}
