package main

import (
	"errors"

	"github.com/urfave/cli/v2"
)

const (
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
	id     string
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

func readOptions(c *cli.Context) (*options, error) {
	if c.NArg() == 0 {
		return nil, errors.New("Error: command requires GitHub ID")
	}
	opt := parseFlags(c)
	opt.id = c.Args().Get(0)
	return opt, nil
}
