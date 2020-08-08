package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var AppVersion = "Undefined"

func main() {
	app := cli.NewApp()
	app.Name = "assets-database"
	app.Version = AppVersion
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Mike Savage King",
			Email: "mike@savageking.io",
		},
	}
	app.Description = "Assets Database for Indie Game Developers"
	app.Copyright = "Copyright 2020 Mike Savage King"

	app.Commands = []*cli.Command{
		{
			Name:  "daemon",
			Usage: "Run assets-database service",
			Flags: []cli.Flag{},
			Action: func(c *cli.Context) error {
				log.SetLevel(log.TraceLevel)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
