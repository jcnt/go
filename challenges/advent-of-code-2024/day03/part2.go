// Advent of Code, 2024. Day 3, part 1.

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

	for _, v := range input {
		if strings.Contains(v, "mul") {
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
		matched, err = regexp.MatchString(`do\(\)`, v)
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
