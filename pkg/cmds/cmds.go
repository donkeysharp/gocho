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
	conf, err := config.LoadConfig()
	if err != nil {
		return cli.NewExitError(err, 1)
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
			Name:   "start",
			Usage:  "Start Gocho node",
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
