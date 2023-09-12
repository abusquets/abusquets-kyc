package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/abusquets/ab-kyc/internal/app"
	core_commands "github.com/abusquets/ab-kyc/internal/core/adapters/api/cli"
	"github.com/abusquets/ab-kyc/internal/db"
)

func main() {

	cliApp := &cli.App{
		Before: func(cCtx *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "test",
				Aliases: []string{"c"},
				Usage:   "test command",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Cli is working fine")
					return nil
				},
			},
			{
				Name:    "runserver",
				Aliases: []string{"r"},
				Usage:   "Run the server",
				Action: func(cCtx *cli.Context) error {
					config := cCtx.App.Metadata["config"].(*app.Config)
					dbManager := cCtx.App.Metadata["dbManager"].(db.DBManager)
					app.Start(config, dbManager)
					return nil
				},
			},
		},
	}
	cliApp.Commands = append(cliApp.Commands, core_commands.Commands()...)
	if cliApp.Metadata == nil {
		cliApp.Metadata = make(map[string]interface{})
	}

	config, err := app.LoadConfigFromENV()
	if err != nil {
		log.Fatal("cannot load config from ENV", err)
		os.Exit(1)
	}
	dbManager, err := db.NewDBManager(config.DBDsn)
	if err != nil {
		log.Fatal("I can't start the database", err)
		os.Exit(1)
	}
	cliApp.Metadata["config"] = config
	cliApp.Metadata["dbManager"] = dbManager

	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
