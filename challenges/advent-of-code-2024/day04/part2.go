// Advent of Code, 2024. Day 4, part 1.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var count int
	var xmas [][]string
	type loc struct {
		l int
		i int
	}
	var alist []loc

	in, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	for _, v := range strings.Split(string(in), "\n") {
		xmas = append(xmas, strings.Split(v, ""))
	}
	xmas = xmas[0 : len(xmas)-1]

	for line := 1; line < len(xmas)-1; line++ {
		for i := 1; i < len(xmas[0])-1; i++ {
			if xmas[line][i] == "A" {
				var ai loc
				ai.i = i
				ai.l = line
				alist = append(alist, ai)
			}
		}
	}

	for _, v := range alist {
		mas := ""
		mas += xmas[v.l-1][v.i-1]
		mas += xmas[v.l-1][v.i+1]
		mas += xmas[v.l+1][v.i-1]
		mas += xmas[v.l+1][v.i+1]
		//fmt.Println(mas)
		if mas == "MSMS" || mas == "MMSS" || mas == "SMSM" || mas == "SSMM" {
			//fmt.Println("bingo!!")
			count++
		}
	}

	fmt.Println("final number is ", count)
}
