// Advent of Code, 2024. Day 1, part 1.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in, err := os.ReadFile("example")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(in), "\n")
	input = input[:len(input)-1]
	fmt.Println(input)
	fmt.Println(len(input))
	for _, v := range input {
		fmt.Printf("%q\n", v)
	}
}
