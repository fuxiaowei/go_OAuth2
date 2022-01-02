package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spring2go/gravitee/cmd"
	"github.com/urfave/cli"
)

var (
	cliApp     *cli.App
	configFile string
)

func init() {
	// Initialize a CLI app
	cliApp = cli.NewApp()
	cliApp.Name = "gravitee-oauth2-server"
	cliApp.Usage = "Gravitee OAuth 2.0 Server"
	cliApp.Author = "雨中伞"
	cliApp.Email = "780106978@@qq.com"
	cliApp.Version = "0.5.0"
	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "configFile",
			Value:       "config.yml",
			Destination: &configFile,
		},
	}
}

func main() {
	// Set the CLI app commands
	cliApp.Commands = []cli.Command{
		{
			Name:  "migrate",
			Usage: "run migrations",
			Action: func(c *cli.Context) error {
				err := cmd.Migrate(configFile)
				fmt.Print("migrate==============", err)
				return err
			},
		},
		{
			Name:  "loaddata",
			Usage: "load data from fixture",
			Action: func(c *cli.Context) error {
				err := cmd.LoadData(c.Args(), configFile)
				fmt.Print("loaddata==============", err)
				return err
			},
		},
		{
			Name:  "runserver",
			Usage: "run web server",
			Action: func(c *cli.Context) error {
				err := cmd.RunServer(configFile)
				fmt.Print("runserver==============", err)
				return err
			},
		},
	}

	// Run the CLI app
	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal("=====================================", err)
	}
}
