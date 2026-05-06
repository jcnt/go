// Advent of Code, 2024. Day 1, part 1.

package main

import (
	"fmt"
	"os"
	"sort"
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

	left := make([]int, 0, len(input))
	right := make([]int, 0, len(input))

	for _, v := range input {
		t := strings.Split(v, "   ")
		tl, _ := strconv.Atoi(t[0])
		tr, _ := strconv.Atoi(t[1])
		left = append(left, tl)
		right = append(right, tr)
	}

	sort.Slice(left, func(i, j int) bool { return left[i] < left[j] })
	sort.Slice(right, func(i, j int) bool { return right[i] < right[j] })

	for i := range len(left) {
		sum += get_diff(left[i], right[i])
	}
	fmt.Println("sum is ", sum)

}

func get_diff(a, b int) int {
	s := a - b
	if s < 0 {
		s *= -1
	}
	return s
}
