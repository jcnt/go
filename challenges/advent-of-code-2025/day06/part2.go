// Advent of Code, 2025. Day 6, part 2.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	source := [][]string{}
	final := 0

	in, err := os.ReadFile("example")
	if err != nil {
		fmt.Println("unable to open file")
	}
	input := strings.Split(string(in), "\n")
	input = input[0 : len(input)-1] // remove last empty column

	for _, v := range input {
		a := strings.Split(v, "")
		source = append(source, a)
	}

	for _, v := range source {
		fmt.Printf("%+q\n", v)
	}

	o := ""
	l := []string{}
	//f := [][]string{}
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
			fmt.Printf("%+q\n", l)
			fmt.Println(l[0:len(l)-1], l[len(l)-1:])
			//do_math(l[0:len(l)-1], l[len(l)-1:])
			l = []string{}
		}
		//fmt.Printf("%+q\n", l)
	}

	fmt.Println("final answer is", final)
}

func do_math(l []int, o string) int {
	sum := 0
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
