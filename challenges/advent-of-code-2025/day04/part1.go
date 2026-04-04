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
	var answer int

	in, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("failed to open input")
	}
	input = strings.Split(string(in), "\n")
	input = input[0 : len(input)-1] // remove last empty column

	for _, row := range input {
		rolls = append(rolls, strings.Split(row, ""))
	}

	for row := range rolls {
		mlen := len(rolls)
		for item := range rolls[row] {
			if rolls[row][item] == "@" {
				//				fmt.Println("yes")
				counter := 0
				for r := row - 1; r < row+2; r++ {
					for c := item - 1; c < item+2; c++ {
						if r >= 0 && r < mlen && c >= 0 && c < mlen {
							if r != row || c != item {
								if rolls[r][c] == "@" {
									counter++
								}
							}
						}
					}
				}
				if counter < 4 {
					answer++
				}
			}
		}
	}

	fmt.Println("The final answer is", answer)

}
