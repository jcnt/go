// Advent of Code, 2025. Day 7, part 1.

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	in, err := os.ReadFile("example")
	if err != nil {
		fmt.Println("unable to open file")
	}
	input := strings.Split(string(in), "\n")
	input = input[0 : len(input)-1]
	fmt.Printf("%q\n", input)

	for i := range input {
		for j := i + 1; j < len(input); j++ {
			get_distance(input[i], input[j])
		}
	}
}

func get_distance(a string, b string) float64 {
	la := strings.Split(a, ",")
	lb := strings.Split(b, ",")
	var ila, ilb []float64
	for _, v := range la {
		t, _ := strconv.Atoi(v)
		ila = append(ila, float64(t))
	}
	for _, v := range lb {
		t, _ := strconv.Atoi(v)
		ilb = append(ilb, float64(t))
	}
	sq := math.Sqrt(math.Pow(ila[0]-ilb[0], 2) + math.Pow(ila[1]-ilb[1], 2) + math.Pow(ila[2]-ilb[2], 2))
	fmt.Println(sq)
	return 0
}
