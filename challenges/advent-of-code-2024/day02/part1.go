// Advent of Code, 2024. Day 2, part 1.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	in, err := os.ReadFile("example")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(in), "\n")
	input = input[:len(input)-1]

	reactor := make([][]int, 0, len(input))
	for _, v := range input {
		sti := strings.Split(v, " ")
		t := make([]int, 0, len(sti))
		for _, w := range sti {
			a, _ := strconv.Atoi(w)
			t = append(t, a)
		}
		reactor = append(reactor, t)
	}
	for _, v := range reactor {
		fmt.Printf("%d\n", v)
	}
}
