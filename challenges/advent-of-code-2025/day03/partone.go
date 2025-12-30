package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	var input []string
	var counter int
	in, err := os.ReadFile("input")
	if err != nil {
		print("failed")
	}
	input = strings.Split(string(in), "\n")
	input = input[0 : len(input)-1] // remove last empty column

	for _, v := range input {
		var intlist []int
		sl := strings.Split(v, "")
		for _, y := range sl {
			in, _ := strconv.Atoi(y)
			intlist = append(intlist, in)
		}
		fmt.Println(intlist)
		frst := slices.Max(intlist[0 : len(intlist)-1])
		sec := slices.Max(intlist[slices.Index(intlist, frst)+1 : len(intlist)])
		fmt.Println(frst, sec)
		final := 10*frst + sec
		fmt.Println(final)
		counter += final
	}
	fmt.Println(counter)

}
