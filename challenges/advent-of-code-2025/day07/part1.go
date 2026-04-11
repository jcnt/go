// Advent of Code, 2025. Day 7, part 1.

package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	tree := [][]string{}
	count := 0

	in, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("unable to open the file")
	}
	input := strings.Split(string(in), "\n")
	input = input[0 : len(input)-1]

	for _, v := range input {
		tree = append(tree, strings.Split(string(v), ""))
	}

	status := tree[0]

	s := slices.Index(status, "S")
	status[s] = "|"

	for _, v := range tree[1:] {
		for j, w := range v {
			if w == "^" && status[j] == "|" {
				status[j-1] = "|"
				status[j+1] = "|"
				status[j] = "."
				count++
			}
			//fmt.Println(status)
		}
	}
	fmt.Println("final answer is", count)
}
