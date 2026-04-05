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

	in, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("failed to open file")
	}

	input := strings.Split(string(in), "\n")

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

	for len(fresh) > 0 {
		ix := 0
		for i := range fresh {
			if fresh[i].s <= fresh[ix].s {
				ix = i
			}
		}
		sfresh = append(sfresh, fresh[ix])
		fresh = slices.Delete(fresh, ix, ix+1)
	}

	changed := true
	for changed {
		changed = false
		for i := 0; i < len(sfresh)-1; i++ {
			cfr := sfresh[i]
			nfr := sfresh[i+1]
			if cfr.s <= nfr.s && cfr.e >= nfr.s {
				sfresh[i+1].s = cfr.s
				if cfr.e > nfr.e {
					sfresh[i+1].e = cfr.e
				}
				changed = true
				sfresh = slices.Delete(sfresh, i, i+1)
			} else if cfr.s >= nfr.s && cfr.s <= nfr.e && cfr.e > nfr.e {
				sfresh[i+1].e = cfr.e
				changed = true
				sfresh = slices.Delete(sfresh, i, i+1)
			}
		}
	}

	for _, r := range sfresh {
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
