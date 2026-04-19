// Advent of Code, 2025. Day 9, part 1.

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
		fmt.Println("unable to open file")
	}
	input := strings.Split(string(in), "\n")
	input = input[0 : len(input)-1]
	fmt.Printf("%q\n", input)

	for i := range input {
		for j := i + 1; j < len(input); j++ {
			fmt.Println(get_area(input[i], input[j]))
		}
	}
}

func get_area(a, b string) int {
	la := strings.Split(a, ",")
	lb := strings.Split(b, ",")
	a1, _ := strconv.Atoi(la[0])
	a2, _ := strconv.Atoi(la[1])
	b1, _ := strconv.Atoi(lb[0])
	b2, _ := strconv.Atoi(lb[1])
	t := (a1 - b1) * (a2 - b2)
	if t < 0 {
		return t * -1
	} else {
		return t
	}
}
