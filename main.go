package main

import "fmt"

const (
	g4 = "\x1b[31m■\x1b[0m"
	g3 = "\x1b[33m■\x1b[0m"
	g2 = "\x1b[32m■\x1b[0m"
	g1 = "\x1b[34m■\x1b[0m"
	g0 = "\x1b[37m■\x1b[0m"
)

func main() {
	fmt.Println("Less", g0, g1, g2, g3, g4, "More")
	fmt.Println()
	fmt.Println(g0, g0, g0, g0, g0, g0, g4)
	fmt.Println(g4, g4, g3, g2, g1, g1, g2)
	fmt.Println(g1, g3, g2, g2, g2, g1, g2)
	fmt.Println(g0, g2, g4, g2, g1, g0, g2)
	fmt.Println(g0, g4, g0, g0, g0, g0, g3)
	fmt.Println(g2, g1, g3, g2, g1, g2, g3)
}
