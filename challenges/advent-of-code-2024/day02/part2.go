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
		ffail := 0
		fail := 0
		for i := 0; i < len(line)-1; i++ {
			if line[i]-line[i+1] > 0 && line[i]-line[i+1] < 4 {
				safe += 1
			} else {
				if ffail == 0 {
					ffail = i
				}
				fail += 1
				break
			}
		}
		if safe == len(line)-1 {
			sum += 1
		} else if safe > 0 && fail > 0 {
			tl := slices.Delete(line, ffail+1, ffail+2)
			second = append(second, tl)
		}
		safe = 0
		ffail = 0
		fail = 0
		for i := 0; i < len(line)-1; i++ {
			if line[i+1]-line[i] > 0 && line[i+1]-line[i] < 4 {
				safe += 1
			} else {
				if ffail == 0 {
					ffail = i
				}
				fail += 1
				break
			}
		}
		if safe == len(line)-1 {
			sum += 1
		} else if safe > 0 && fail > 0 {
			tl := slices.Delete(line, ffail+1, ffail+2)
			second = append(second, tl)
		}
	}

	for _, line := range second {
		safe := 0
		for i := 0; i < len(line)-1; i++ {
			if line[i]-line[i+1] > 0 && line[i]-line[i+1] < 4 {
				safe += 1
			} else {
				break
			}
		}
		if safe == len(line)-1 {
			sum += 1
		}
		safe = 0
		for i := 0; i < len(line)-1; i++ {
			if line[i+1]-line[i] > 0 && line[i+1]-line[i] < 4 {
				safe += 1
			} else {
				break
			}
		}
		if safe == len(line)-1 {
			sum += 1
		}
	}
	fmt.Println("sum is ", sum)
}
