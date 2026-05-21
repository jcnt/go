// Advent of Code, 2024. Day 2, part 1.

package main

import (
	"fmt"
	"os"
	"slices"
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

	var reactor [][]int
	for _, v := range input {
		sti := strings.Split(v, " ")
		t := []int{}
		for _, w := range sti {
			a, _ := strconv.Atoi(w)
			t = append(t, a)
		}
		reactor = append(reactor, t)
	}

	for _, line := range reactor {
		safe := 0
		fail := 0
		for i := 0; i < len(line)-1; i++ {
			if line[i]-line[i+1] > 0 && line[i]-line[i+1] < 4 {
				safe += 1
			} else {
				fail = i
				break
			}
		}
		if safe == len(line)-1 {
			sum++
		} else if safe > 1 {
			fmt.Println(fsecond(line, fail))
		}

		safe = 0
		if fail == 0 {
			for i := 0; i < len(line)-1; i++ {
				if line[i+1]-line[i] > 0 && line[i+1]-line[i] < 4 {
					safe++
				} else {
					fail = i
					break
				}
			}
			if safe == len(line)-1 {
				sum++
			} else {
				fmt.Println(fsecond(line, fail))
			}
		}
	}

	fmt.Println("sum is ", sum)
}

func fsecond(s []int, f int) bool {
	r := false
	for i := 0; i < len(s)-1; i++ {
		t := make([]int, len(s))
		copy(t, s)
		current := slices.Delete(t, f, f+1)
		fmt.Println(current)
	}
	return r
}
