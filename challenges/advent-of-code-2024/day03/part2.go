// Advent of Code, 2024. Day 3, part 2.

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	mulsl := []string{}
	finsl := []string{}
	sum := 0

	in, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.SplitAfter(string(in), ")")
	input = input[0 : len(input)-1]

	is_on := 1
	for _, v := range input {
		fmt.Println("LINE---> ", v)
		if strings.Contains(v, "don't") {
			is_on = 1
			fmt.Println("IS_OFF ----> ")
		} else if strings.Contains(v, "do") {
			is_on = 0
			fmt.Println("IS_ON ----> ")
		}
		if strings.Contains(v, "mul") && is_on == 1 {
			mul := strings.Split(v, "mul")
			if len(mul) > 1 {
				mulsl = append(mulsl, mul[len(mul)-1])
			}
		}
	}

	for _, v := range mulsl {
		matched, err := regexp.MatchString(`^\(\d+,\d+\)$`, v)
		if err == nil && matched == true {
			finsl = append(finsl, v)
		}
	}
	fmt.Println(finsl)

	for _, v := range finsl {
		t := strings.Trim(v, "()")
		tlist := strings.Split(t, ",")
		a, _ := strconv.Atoi(tlist[0])
		b, _ := strconv.Atoi(tlist[1])
		sum += a * b
	}
	fmt.Println(sum)
}
