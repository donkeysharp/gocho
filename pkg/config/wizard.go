package config

import (
	"fmt"
	"os"
)

func configureWizard() error {
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
	fmt.Printf("Share Directory: (%s) ", config.ShareDirectory)
	fmt.Scanf("%s", &shareDirectory)
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
