package main

import (
	"log"
	"os"

	"github.com/lusingander/crass/grass"
	"github.com/urfave/cli/v2"
)

const (
	appName  = "crass"
	appUsage = "growing grass on CUI"
)

func action(c *cli.Context) error {
	opt, err := readOptions(c)
	if err != nil {
		return err
	}
	grasses, err := grass.Mow(opt.id)
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
