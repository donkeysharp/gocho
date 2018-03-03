package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func configureWizard() error {
	reader := bufio.NewReader(os.Stdin)

	config, err := getDefaultConfig()
	if err != nil {
		return err
	}

	fmt.Println("Gocho Configure Wizard")
	fmt.Println("It will reset previous \"Gocho\" configure file")
	var (
		shareDirectory string
		webPort        string
		localPort      string
		nodeId         string
	)

	fmt.Printf("Node Id: (%s) ", config.NodeId)
	fmt.Scanf("%s", &nodeId)
	fmt.Printf("Share Directory: ")
	// In windows it fails using fmt.Scanf
	lineRaw, _, err := reader.ReadLine()
	fmt.Println(string(lineRaw))
	if err != nil || strings.Trim(string(lineRaw), " \t") == "" {
		fmt.Println("Invalid value for \"Share Directory\"")
		os.Exit(1)
	}
	shareDirectory = string(lineRaw)
	fmt.Printf("Share Port: (%s) ", config.WebPort)
	fmt.Scanf("%s", &webPort)
	fmt.Printf("Dashboard Port: (%s) ", config.LocalPort)
	fmt.Scanf("%s", &localPort)

	if nodeId != "" {
		config.NodeId = nodeId
	}
	if shareDirectory != "" {
		config.ShareDirectory = CleanPath(shareDirectory)
	}
	if webPort != "" {
		config.WebPort = webPort
	}
	if localPort != "" {
		config.LocalPort = localPort
	}

	if fileExists(config.ConfigFile) {
		err := os.Remove(config.ConfigFile)
		if err != nil {
			return err
		}
	}

	return writeConfigToFile(config, config.ConfigFile)
}
