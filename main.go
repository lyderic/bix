package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/lyderic/tools"
	"github.com/urfave/cli"
)

var (
	appfile string
)

func init() {
	var err error
	var me *user.User
	if me, err = user.Current(); err != nil {
		log.Fatal(err)
	}
	appfile = filepath.Join(me.HomeDir, ".bixru.json")
}

func main() {
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
			Name:    "parse64",
			Usage:   "parse base64 encoded files of ruwix.com",
			Aliases: []string{"64"},
			Action: func(c *cli.Context) (err error) {
				return  parse64()
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		tools.PrintColorf(tools.RED, "[ERROR] %v\n", err)
		return
	}
}
