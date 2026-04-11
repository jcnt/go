// Advent of Code, 2025. Day 6, part 2.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	source := [][]string{}
	final := 0

	in, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("unable to open file")
	}
	input := strings.Split(string(in), "\n")
	input = input[0 : len(input)-1] // remove last empty column

	for _, v := range input {
		a := strings.Split(v, "")
		source = append(source, a)
	}

	o := ""
	l := []string{}
	for i := len(source[0]) - 1; i >= 0; i-- {
		s := ""
		turn := false
		for _, v := range source {
			if v[i] == "*" || v[i] == "+" {
				o = v[i]
				turn = true
			} else {
				s += v[i]
			}
		}
		if turn != true {
			l = append(l, s)
		} else if turn {
			l = append(l, s)
			l = append(l, o)
			final += do_math(l)
			l = []string{}
		}
		//fmt.Printf("%+q\n", l)
	}

	fmt.Println("final answer is", final)
}

func do_math(l []string) int {
	sum := 0
	o := l[len(l)-1:][0]  // last item
	tl := l[0 : len(l)-1] // all items minus last
	for _, v := range tl {
		tv := strings.TrimSpace(v)
		if tv != "" {
			fv, _ := strconv.Atoi(tv)
			if o == "+" {
				sum += fv
			} else if o == "*" {
				if sum == 0 {
					sum += fv
				} else {
					sum *= fv
				}
			}
		}
	}
	return sum
}
