package main

import (
	"fmt"
	"os"

	"github.com/lyderic/tools"
	"github.com/urfave/cli"
)

func main() {
	appfile := "bix.json"
	app := cli.NewApp()
	app.Name = "bix"
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
				if err = setup(appfile); err != nil {
					return
				}
				return timer(appfile)
			},
		},
		{
			Name:  "scramble",
			Usage: "produce a scramble",
			//Aliases: []string{"s"},
			Action: func(c *cli.Context) (err error) {
				return scramble()
			},
		},
		{
			Name:    "show",
			Usage:   "show recorded times",
			Aliases: []string{"s"},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "show all",
				},
			},
			Action: func(c *cli.Context) (err error) {
				if err = setup(appfile); err != nil {
					return
				}
				return showPerformances(10)
			},
		},
		{
			Name:    "manual",
			Usage:   "manually add a performance",
			Aliases: []string{"m"},
			Action: func(c *cli.Context) (err error) {
				if err = setup(appfile); err != nil {
					return
				}
				return inputPerformance(appfile)
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		tools.PrintColorf(tools.RED, "[ERROR] %v\n", err)
		return
	}
}
