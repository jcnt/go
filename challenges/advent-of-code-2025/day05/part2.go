// Advent of Code, 2025. Day 5, part 2.

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	type fresh_range struct {
		s int
		e int
	}
	var fresh, sfresh []fresh_range
	var answer int

	in, err := os.ReadFile("example")
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
			var fr fresh_range
			fr.s = a
			fr.e = b
			fresh = append(fresh, fr)
		}
	}

	fmt.Println(fresh)
	fmt.Println(sfresh)

	changed := true
	for changed {
		changed = false
		for i := 0; i < len(fresh)-1; i++ {
			cfr := fresh[i]
			nfr := fresh[i+1]
			//fmt.Println("comparing ", cfr, nfr)
			if cfr.s < nfr.s && cfr.e >= nfr.s {
				//fmt.Println("comparing", cfr.s, nfr.s, "and", cfr.e, nfr.e)
				fresh[i+1].s = cfr.s
				if cfr.e > nfr.e {
					fresh[i+1].e = cfr.e
				}
				changed = true
				fresh = slices.Delete(fresh, i, i+1)
				//fmt.Println("current fresh: ", fresh)
			} else if cfr.s >= nfr.s && cfr.s <= nfr.e && cfr.e > nfr.e {
				//fmt.Println("comparing 2", cfr.s, nfr.s, "and", cfr.e, nfr.e)
				fresh[i+1].e = cfr.e
				//fmt.Println("after change", nfr.e, cfr.e)
				changed = true
				fresh = slices.Delete(fresh, i, i+1)
				//fmt.Println("current fresh: ", fresh)
			}
		}
	}

	fmt.Println(fresh)

	for _, r := range fresh {
		count := r.e - r.s + 1
		answer += count
	}

	fmt.Println("final answer is: ", answer)
}

/*
comparing ranges:
- if start < current start and end >= current start
	-> start = c start
	-> if end > c end
		-> end = c end

- if start >= c start and start <= c end and end > c end
	-> end = c end
*/
