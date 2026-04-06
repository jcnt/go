// Advent of Code, 2025. Day 6, part 1.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//var source []int

	in, err := os.ReadFile("example")
	if err != nil {
		fmt.Println("unable to open file")
	}
	input := strings.Split(string(in), "\n")

	for _, v := range input {
		a := strings.Split(v, " ")
		l := []int{}
		for _, w := range a {
			s, err := strconv.Atoi(w)
			if err == nil {
				l = append(l, s)
			}
		}
		for _, x := range l {
			fmt.Println(x)
		}
	}
}
