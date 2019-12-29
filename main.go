package main

import (
	"log"
	"os"

	"github.com/lusingander/crass/grass"
)

func run(args []string) error {
	if len(args) == 0 {
		return nil // TODO: error
	}
	grasses, err := grass.Mow(args[0])
	if err != nil {
		return err
	}
	printGrasses(grasses)
	return nil
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
