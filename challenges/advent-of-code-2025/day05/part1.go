// Advent of Code, 2025. Day 5, part 1.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	type frange struct {
		s int
		e int
	}
	var fresh []frange
	var ing []int
	var answer int

	in, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("failed to open file")
	}

	input := strings.Split(string(in), "\n")
	//	fmt.Println("input", input, "input")

	for _, v := range input {
		if strings.Contains(v, "-") {
			r := strings.Split(v, "-")
			a, _ := strconv.Atoi(r[0])
			b, _ := strconv.Atoi(r[1])
			var fr frange
			fr.s = a
			fr.e = b
			fresh = append(fresh, fr)
			//			fmt.Println("appended range")
		} else if v != "" {
			iv, _ := strconv.Atoi(v)
			ing = append(ing, iv)
		}
	}
	//	fmt.Println(fresh, ing)

	for _, ingr := range ing {
		for _, frn := range fresh {
			//			fmt.Println(ingr, frn)
			if ingr >= frn.s {
				if ingr <= frn.e {
					answer++
					break
				}
			}
		}
	}

	fmt.Println("final answer is: ", answer)
}
