// Advent of Code, 2024. Day 1, part 2.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0

	in, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(in), "\n")
	input = input[:len(input)-1]

	left := make([]int, 0, len(input))
	right := make([]int, 0, len(input))

	for _, v := range input {
		t := strings.Split(v, "   ")
		tl, _ := strconv.Atoi(t[0])
		tr, _ := strconv.Atoi(t[1])
		left = append(left, tl)
		right = append(right, tr)
	}

	for _, v := range left {
		count := 0
		for _, w := range right {
			if v == w {
				count++
			}
		}
		sum += (count * v)
	}
	fmt.Println("final answer is ", sum)
}
