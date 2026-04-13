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

	status := []int{}
	for i := range len(tree[0]) {
		status = append(status, 0)
		i++
	}
	s := slices.Index(tree[0], "S")
	status[s] = 1

	for _, v := range tree[1:] {
		for j, w := range v {
			if w == "^" {
				status[j-1] += status[j]
				status[j+1] += status[j]
				status[j] = 0
			}
			//fmt.Println(status)
		}
	}

	for _, v := range status {
		count += v
	}
	fmt.Println("final answer is", count)
}
