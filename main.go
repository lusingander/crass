package main

import (
	"errors"
	"log"
	"os"

	"github.com/lusingander/crass/grass"
	"github.com/urfave/cli/v2"
)

const (
	appName  = "crass"
	appUsage = "growing grass on CUI"

	flagLegend = "legend"
)

var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:     flagLegend,
		Aliases:  []string{"l"},
		Usage:    "show legend",
		Required: false,
		Value:    false,
	},
}

type options struct {
	legend bool
}

func (o *options) showLegend() bool {
	return o.legend
}

func parseFlags(c *cli.Context) *options {
	return &options{
		legend: c.Bool(flagLegend),
	}
}

func action(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("Error: command requires GitHub ID")
	}
	id := c.Args().Get(0)
	opt := parseFlags(c)

	grasses, err := grass.Mow(id)
	if err != nil {
		return err
	}
	printGrasses(grasses, opt)
	return nil
}

func main() {
	app := &cli.App{
		Name:                  appName,
		Usage:                 appUsage,
		Action:                action,
		Flags:                 flags,
		CustomAppHelpTemplate: appHelpTemplate,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
