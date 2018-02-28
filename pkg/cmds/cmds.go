package cmds

import (
	"fmt"
	"github.com/donkeysharp/gocho/pkg/config"
	"github.com/donkeysharp/gocho/pkg/info"
	"github.com/donkeysharp/gocho/pkg/node"
	"github.com/urfave/cli"
)

func ConfigureAction(c *cli.Context) error {
	err := config.ConfigureWizard()
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}

func StartAction(c *cli.Context) error {
	fmt.Println("Starting Gocho Node...")
	conf := &config.Config{}
	conf.Debug = c.Bool("debug")
	conf.LocalPort = c.String("local-port")
	conf.WebPort = c.String("share-port")
	conf.ShareDirectory = c.String("dir")
	conf.NodeId = c.String("id")

	if conf.NodeId == "" || conf.ShareDirectory == "" {
		fmt.Println("Both --dir and --id should be set.")
		fmt.Println("Checking config file.")
		var err error
		conf, err = config.LoadConfig()
		if err != nil {
			return cli.NewExitError(err, 1)
		}
	}

	fmt.Println("Configuration loaded")
	fmt.Println("---")
	fmt.Println(conf)
	fmt.Println("---")

	node.Serve(conf)

	return nil
}

func New() *cli.App {
	app := cli.NewApp()
	app.Name = info.APP_NAME
	app.Usage = "Auto-discovery local area network file sharing"
	app.Version = info.VERSION
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Sergio Guillen Mantilla",
			Email: "serguimant@gmail.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start Gocho node",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "debug",
					Usage: "Start gocho in debug mode",
				},
				cli.StringFlag{
					Name:   "id",
					Usage:  "Node ID that will be shared to other peers",
					EnvVar: "GOCHO_ID",
				},
				cli.StringFlag{
					Name:   "dir",
					Usage:  "Directory to share",
					EnvVar: "GOCHO_DIR",
				},
				cli.StringFlag{
					Name:   "share-port",
					Usage:  "Port that will be exposed for file sharing",
					EnvVar: "GOCHO_SHARE_PORT",
					Value:  "5555",
				},
				cli.StringFlag{
					Name:   "local-port",
					Usage:  "Port for local dashboard",
					EnvVar: "GOCHO_LOCAL_PORT",
					Value:  "1337",
				},
			},
			Action: StartAction,
		},
		{
			Name:   "configure",
			Usage:  "Create a configuration file for Gocho node",
			Action: ConfigureAction,
		},
	}

	return app
}
