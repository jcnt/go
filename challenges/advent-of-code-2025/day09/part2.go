// Advent of Code, 2025. Day 9, part 1.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reddb := [][]int{}
	tiledb := [][]string{}
	var maxx, maxy int
	in, err := os.ReadFile("example")
	if err != nil {
		fmt.Println("cannot open file")
	}
	input := strings.Split(string(in), "\n")
	input = input[0 : len(input)-1]
	//fmt.Printf("%q\n", input)

	for _, v := range input {
		intred := []int{}
		red := strings.Split(v, ",")
		for j := range red {
			r, _ := strconv.Atoi(red[j])
			intred = append(intred, r)
			if j == 0 && r > maxx {
				maxx = r
			} else if j == 1 && r > maxy {
				maxy = r
			}
		}
		reddb = append(reddb, intred)
	}
	fmt.Printf("%d\n\n", reddb)

	for i := 0; i <= maxy+1; i++ {
		tline := []string{}
		for j := 0; j <= maxx+1; j++ {
			tline = append(tline, ".")
		}
		tiledb = append(tiledb, tline)
	}

	for _, t := range reddb {
		tiledb[t[1]][t[0]] = "#"
	}

	for i := 1; i < len(reddb); i++ {
		c0 := reddb[i-1][0]
		c1 := reddb[i-1][1]
		n0 := reddb[i][0]
		n1 := reddb[i][1]
		if c0 == n1 {
			if c1 < n1 {
				fmt.Println(c1, n1)
			} else {
				fmt.Println(n1, c1)
			}
		} else {
			if c0 < n0 {
				fmt.Println(c0, n0)
			} else {
				fmt.Println(n0, c0)
			}
		}
	}

	for _, v := range tiledb {
		fmt.Printf("%q\n", v)
	}
}
