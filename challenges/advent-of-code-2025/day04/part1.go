// Advent of Code, 2025. Day 4, part 1.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var input []string
	var rolls [][]string

	in, err := os.ReadFile("example")
	if err != nil {
		print("failed to open input")
	}
	input = strings.Split(string(in), "\n")
	input = input[0 : len(input)-1] // remove last empty column

	for _, row := range input {
		rolls = append(rolls, strings.Split(row, ""))
	}

	for row := range rolls {
		//		mlen := len(rolls)
		for item := range rolls[row] {
			fmt.Print(rolls[row][item])
			fmt.Println()
		}
	}

}
