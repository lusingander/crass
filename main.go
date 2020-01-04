package main

import (
	"log"
	"os"

	"github.com/lusingander/crass/grass"
	"github.com/urfave/cli/v2"
)

func action(c *cli.Context) error {
	if c.NArg() == 0 {
		return nil // TODO: error
	}
	grasses, err := grass.Mow(c.Args().Get(0))
	if err != nil {
		return err
	}
	printGrasses(grasses)
	return nil
}

func main() {
	app := &cli.App{
		Name:   "crass",
		Usage:  "growing grass on CUI",
		Action: action,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
