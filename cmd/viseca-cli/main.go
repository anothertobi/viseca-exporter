package main

import (
	"log"
	"os"

	"github.com/anothertobi/viseca-exporter/internal/app"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "viseca-cli",
		Usage: "query Viseca one",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "username",
				Usage:    "username",
				Required: true,
				EnvVars:  []string{"VISECA_CLI_USERNAME"},
			},
			&cli.StringFlag{
				Name:     "password",
				Usage:    "password",
				Required: true,
				EnvVars:  []string{"VISECA_CLI_PASSWORD"},
			},
		},
		Commands: []*cli.Command{
			app.NewTransactionsCommand(),
			app.NewCardsCommand(),
			app.NewUserCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
