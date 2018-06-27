package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lyderic/tools"
	"github.com/urfave/cli"
)

func init() {
}

func main() {
	appfile := filepath.Join(os.Getenv("HOME"), ".bixru.json")
	app := cli.NewApp()
	app.Name = "bixru"
	app.Usage = "Rubik's Cube Timer for the Command Line"
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "`path` to application file",
			Value:       appfile,
			Destination: &appfile,
		},
	}
	app.Before = func(c *cli.Context) (err error) {
		if len(os.Args) == 1 {
			return
		}
		return
	}
	app.After = func(c *cli.Context) (err error) {
		return
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Printf("%s: command not found, try '%s --help'\n", app.Name, command)
	}
	app.Commands = []cli.Command{
		{
			Name:    "timer",
			Usage:   "start timer",
			Aliases: []string{"t"},
			Action: func(c *cli.Context) (err error) {
				return timer()
			},
		},
		{
			Name:    "scramble",
			Usage:   "produce a scramble",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) (err error) {
				return scramble()
			},
		},
		{
			Name:    "performance",
			Usage:   "manage performances",
			Aliases: []string{"perf", "p"},
			Subcommands: cli.Commands{
				cli.Command{
					Name:  "add",
					Usage: "add a performance",
					Action: func(c *cli.Context) (err error) {
						if err = setup(appfile); err != nil {
							  return
						}
						return addPerformance(appfile)
					},
				},
				cli.Command{
					Name:  "show",
					Usage: "show performances",
					Action: func(c *cli.Context) (err error) {
						if err = setup(appfile); err != nil {
							  return
						}
						return showPerformances()
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		tools.PrintColorf(tools.RED, "[ERROR] %v\n", err)
		return
	}
}
