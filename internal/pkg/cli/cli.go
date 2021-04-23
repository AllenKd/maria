package cli

import (
	"github.com/urfave/cli/v2"
	"maria/internal/pkg/cleaner"
	"maria/internal/pkg/configs"
)

const (
	version = "0.1.0"
)

func InitCli() (app *cli.App) {

	app = &cli.App{
		Name:    "Maria",
		Usage:   "A CLI for routine job",
		Version: version,
	}

	app.Commands = initCommand()
	return
}

func initCommand() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "clean",
			Aliases: []string{"c"},
			Usage:   "Clean folder",
			Flags:   cleanCommandFlag(),
			Action: func(c *cli.Context) error {
				cleaner.New().Clean()
				return nil
			},
		},
	}
}

func cleanCommandFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "root-folder",
			Aliases:     []string{"r"},
			Value:       configs.New().Cleaner.RootFolder,
			Usage:       "root folder path",
			Destination: &configs.New().Cleaner.RootFolder,
		},
	}
}
