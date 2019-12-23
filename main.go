package main

import "fmt"

const box = "  "

const (
	g4 = "\x1b[48;5;22m" + box + "\x1b[0m"
	g3 = "\x1b[48;5;28m" + box + "\x1b[0m"
	g2 = "\x1b[48;5;40m" + box + "\x1b[0m"
	g1 = "\x1b[48;5;120m" + box + "\x1b[0m"
	g0 = "\x1b[48;5;7m" + box + "\x1b[0m"
)

func main() {
	fmt.Println(g0 + g0 + g0 + g0 + g0 + g0 + g4)
	fmt.Println(g4 + g4 + g3 + g2 + g1 + g1 + g2)
	fmt.Println(g1 + g3 + g2 + g2 + g2 + g1 + g2)
	fmt.Println(g0 + g2 + g4 + g2 + g1 + g0 + g2)
	fmt.Println(g0 + g4 + g0 + g0 + g0 + g0 + g3)
	fmt.Println(g2 + g1 + g3 + g2 + g1 + g2 + g3)
}
