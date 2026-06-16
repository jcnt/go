// Advent of Code, 2024. Day 4, part 1.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var xmas [][]string
	type loc struct {
		l int
		i int
	}
	var xlist []loc

	in, err := os.ReadFile("example")
	if err != nil {
		panic(err)
	}

	for _, v := range strings.Split(string(in), "\n") {
		xmas = append(xmas, strings.Split(v, ""))
	}
	xmas = xmas[0 : len(xmas)-1]

	for line := range xmas {
		for i := range xmas[line] {
			if xmas[line][i] == "X" {
				var xi loc
				xi.i = i
				xi.l = line
				xlist = append(xlist, xi)
			}
		}
	}

	fmt.Println(xlist)
}
