package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var input []string
	in, err := os.ReadFile("example")
	if err != nil {
		print("failed")
	}
	input = strings.Split(string(in), "\n")
	input = input[0 : len(input)-1] // remove last empty column

	for i, v := range input {
		for j := range v {
			fmt.Printf("%v, %v:\t\t", i, j)
			for k := i - 1; k <= i+1; k++ {
				for l := j - 1; l <= j+1; l++ {
					if (l >= 0) && (l < 10) && (k >= 0) && (k < 10) {
						if (l == j) && (k == i) {
							fmt.Printf("...\t")
						} else {
							fmt.Printf("%v, %v \t", k, l)
						}
					}
				}
			}
			fmt.Println()
		}
	}
}
