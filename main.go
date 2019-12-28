package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/lusingander/crass/grass"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

const cell = "  "

var grassCells = map[int]string{
	4: "\x1b[48;5;22m" + cell + "\x1b[0m",
	3: "\x1b[48;5;28m" + cell + "\x1b[0m",
	2: "\x1b[48;5;40m" + cell + "\x1b[0m",
	1: "\x1b[48;5;120m" + cell + "\x1b[0m",
	0: "\x1b[48;5;7m" + cell + "\x1b[0m",
}

func calcReduceWeeks() int {
	const maxWidth = 4 + (54 * 2) // leftSide + (weeks * 2)
	w, _ := terminal.Width()
	shortage := maxWidth - w
	if shortage <= 0 {
		return 0
	}
	return int(math.Ceil(float64(shortage) / 2.0))
}

func printGrasses(grasses []*grass.Grass) {
	r := calcReduceWeeks()
	header := "    "
	for i, g := range grasses {
		if i%7 > 0 {
			continue
		}
		if (i / 7) < r {
			continue
		}
		if d := g.GetDay(); 1 <= d && d <= 7 {
			header += fmt.Sprintf("%2d", g.GetMonth())
		} else {
			header += "  "
		}
	}
	fmt.Println(header)
	week := [7]string{}
	for i := 0; i < 7; i++ {
		switch i {
		case 1:
			week[i] += "Mon "
		case 3:
			week[i] += "Wed "
		case 5:
			week[i] += "Fri "
		default:
			week[i] += "    "
		}
	}
	for i, g := range grasses {
		if (i / 7) < r {
			continue
		}
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
