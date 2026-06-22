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

	for _, v := range xlist {
		if len(xmas[0])-v.i >= 4 {
			fmt.Println(v)
			if xmas[v.l][v.i+1] == "M" && xmas[v.l][v.i+2] == "A" && xmas[v.l][v.i+3] == "S" {
				fmt.Println("XMAS")
			}
		}
	}

	fmt.Println("----")

	for _, v := range xlist {
		if len(xmas)-v.l >= 4 {
			fmt.Println(v)
			if xmas[v.l+1][v.i] == "M" && xmas[v.l+2][v.i] == "A" && xmas[v.l+3][v.i] == "S" {
				fmt.Println("XMAS")
			}
		}
	}
	fmt.Println("----")
}
