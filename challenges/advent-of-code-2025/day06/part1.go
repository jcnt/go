// Advent of Code, 2025. Day 6, part 1.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var source [][]string

	in, err := os.ReadFile("example")
	if err != nil {
		fmt.Println("unable to open file")
	}
	input := strings.Split(string(in), "\n")
	input = input[0 : len(input)-1] // remove last empty column

	for _, v := range input {
		a := strings.Split(v, " ")
		l := []string{}
		for _, w := range a {
			if w != " " && w != "" {
				l = append(l, strings.TrimSpace(w))
			}
		}
		source = append(source, l)
	}

	for _, v := range source {
		fmt.Println(v)
	}
}
