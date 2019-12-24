package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lusingander/crass/grass"
)

const cell = "  "

var grassCells = map[int]string{
	4: "\x1b[48;5;22m" + cell + "\x1b[0m",
	3: "\x1b[48;5;28m" + cell + "\x1b[0m",
	2: "\x1b[48;5;40m" + cell + "\x1b[0m",
	1: "\x1b[48;5;120m" + cell + "\x1b[0m",
	0: "\x1b[48;5;7m" + cell + "\x1b[0m",
}

func printGrasses(grasses []*grass.Grass) {
	week := [7]string{}
	for i, g := range grasses {
		week[i%7] += grassCells[g.Growth()]
	}
	for _, w := range week {
		fmt.Println(w)
	}
}

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
