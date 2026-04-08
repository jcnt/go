// Advent of Code, 2025. Day 6, part 1.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var source [][]string
	var final int

	in, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("unable to open file")
	}
	input := strings.Split(string(in), "\n")
	input = input[0 : len(input)-1] // remove last empty column

	for _, v := range input {
		a := strings.Split(v, " ")
		l := []string{}
		for _, w := range a {
			if w != " " && w != "" {
				l = append(l, strings.TrimSpace(w))
			}
		}
		source = append(source, l)
	}

	for i := range source[0] {
		var mlist []int
		var s int
		op := source[len(source)-1][i]
		for _, v := range source[0 : len(source)-1] {
			a, _ := strconv.Atoi(v[i])
			mlist = append(mlist, a)
		}
		s = do_math(mlist, op)
		final += s
	}

	fmt.Println("final answer is", final)
}

func do_math(l []int, o string) int {
	var sum int
	for _, v := range l {
		if o == "+" {
			sum += v
		} else if o == "*" {
			if sum == 0 {
				sum += v
			} else {
				sum *= v
			}
		}
	}
	return sum
}
