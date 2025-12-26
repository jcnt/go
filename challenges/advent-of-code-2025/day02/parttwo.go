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
	in, err := os.ReadFile("input")
	if err != nil {
		print("failed")
	}
	input = strings.Split(string(in), ",")

	var count int
	counter := map[string]int{}
	for _, v := range input {
		a := strings.Split(v, "-")
		start, _ := strconv.Atoi(a[0])
		end, _ := strconv.Atoi(a[1])
		for i := start; i <= end; i++ {
			//			fmt.Println(i)
			s := strconv.Itoa(i)
			fmt.Println(s)
			l := strings.Split(s, "")
			fmt.Println(l)
			for j := 2; j <= 10; j++ {
				var fin [][]string
				if len(l)%j == 0 {
					eq := true
					//		fmt.Println("yes", i)
					nsl := len(l) / j
					for k := 0; k < len(l); k += nsl {
						fin = append(fin, l[k:k+nsl])
					}
					fmt.Println(fin)
					for m := 1; m < len(fin); m++ {
						eq = eq && slices.Equal(fin[0], fin[m])
					}
					if eq == true {
						fmt.Println(eq)
						counter[s]++
					}
				}
			}
		}
	}
	for k, _ := range counter {
		fn, _ := strconv.Atoi(k)
		count += fn
	}
	fmt.Println(count)
}
