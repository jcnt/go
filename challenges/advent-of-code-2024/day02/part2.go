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

	var reactor, second [][]int
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
		psd := 0
		sec := 0
		for i := 0; i < len(line)-1; i++ {
			if line[i]-line[i+1] > 0 && line[i]-line[i+1] < 4 {
				safe += 1
			} else {
				fail = i
				break
			}
		}
		if safe == len(line)-1 {
			sum += 1
			psd = 1
		} else if safe > 1 {
			tl := slices.Delete(line, fail, fail+1)
			if fsecond(tl) == true {
				psd = 1
				sum += 1
			}
			if psd == 0 {
				tl = slices.Delete(line, fail+1, fail+2)
				if fsecond(tl) == true {
					psd = 1
					sum += 1
				}
			}
		}
		safe = 0
		fail = 0
		if psd == 0 && sec == 0 {
			for i := 0; i < len(line)-1; i++ {
				if line[i+1]-line[i] > 0 && line[i+1]-line[i] < 4 {
					safe += 1
				} else {
					fail = i
					break
				}
			}
			if safe == len(line)-1 {
				sum += 1
			} else {
				tl := slices.Delete(line, fail, fail+1)
				second = append(second, tl)
			}
		}
	}

	fmt.Println("sum is ", sum)
}

func fsecond(s []int) bool {
	safe := 0
	psd := 0
	r := false
	for i := 0; i < len(s)-1; i++ {
		diff := s[i] - s[i+1]
		if diff > 0 && diff < 4 {
			safe += 1
			psd = 1
		}
	}
	if safe == len(s)-1 {
		r = true
	}
	safe = 0
	if psd == 0 {
		for i := 0; i < len(s)-1; i++ {
			diff := s[i+1] - s[i]
			if diff > 0 && diff < 4 {
				safe += 1
			}
		}
		if safe == len(s)-1 {
			r = true
		}
	}
	return r
}
