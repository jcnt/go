// Advent of Code, 2025. Day 9, part 1.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	answer := 0

	in, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("unable to open file")
	}
	input := strings.Split(string(in), "\n")
	input = input[0 : len(input)-1]
	fmt.Printf("%q\n", input)

	for i := range input {
		for j := i + 1; j < len(input); j++ {
			current := get_area(input[i], input[j])
			if current > answer {
				answer = current
			}
		}
	}
	fmt.Println("final answer is: ", answer)
}

func get_area(a, b string) int {
	la := strings.Split(a, ",")
	lb := strings.Split(b, ",")
	a1, _ := strconv.Atoi(la[0])
	a2, _ := strconv.Atoi(la[1])
	b1, _ := strconv.Atoi(lb[0])
	b2, _ := strconv.Atoi(lb[1])
	t := (a1 - b1 + 1) * (a2 - b2 + 1)
	if t < 0 {
		return t * -1
	} else {
		return t
	}
}
