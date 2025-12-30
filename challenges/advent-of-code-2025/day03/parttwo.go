package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	var input []string
	var subcount []int
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
		frst := 0
		fmt.Println(intlist)
		for i := 0; i < 12; i++ {
			//			fmt.Println("frst is", frst)
			sec := len(intlist) - 11 + i
			//			fmt.Println("sec is ", sec)
			cmax := slices.Max(intlist[frst:sec])
			cmaxi := slices.Index(intlist[frst:sec], cmax)
			//			fmt.Println("cmax is", cmax)
			//			fmt.Println()
			subcount = append(subcount, cmax)

			frst = frst + cmaxi + 1

		}
		fmt.Println(subcount)
		subcounti := 0
		for i := 11; i >= 0; i-- {
			j := 11 - i
			subcounti += (subcount[i] * int(math.Pow(10, float64(j))))
		}
		subcount = []int{}
		fmt.Println(subcounti)
		counter += subcounti
	}
	fmt.Println(counter)

}
