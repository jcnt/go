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
			tl := slices.Delete(line, fail+1, fail+2)
			second = append(second, tl)
			fmt.Println("-----> _1_ line is ", line, "tl is ", tl)
			sec = 1
		}
		fmt.Println(line, "pass", psd, "safe", safe, "fail", fail)
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
				if sec == 0 {
					tl := slices.Delete(line, fail+1, fail+2)
					second = append(second, tl)
					fmt.Println("-----> _2_ line is ", line, "tl is ", tl)
				}
			}
		}
		fmt.Println(line, "pass", psd, "safe", safe, "fail", fail)
	}

	for _, line := range second {
		psd := 0
		safe := 0
		for i := 0; i < len(line)-1; i++ {
			if line[i]-line[i+1] > 0 && line[i]-line[i+1] < 4 {
				safe += 1
			}
		}
		if safe == len(line)-1 {
			sum += 1
			psd = 1
			fmt.Println("SAFE", line)
		}
		safe = 0
		if psd == 0 {
			for i := 0; i < len(line)-1; i++ {
				if line[i+1]-line[i] > 0 && line[i+1]-line[i] < 4 {
					safe += 1
				}
			}
			if safe == len(line)-1 {
				sum += 1
				psd += 1
				fmt.Println("SAFE", line)
			}
		}
		if psd == 0 {
			fmt.Println("NOT SAFE", line)
		}
	}
	fmt.Println("sum is ", sum)
}
