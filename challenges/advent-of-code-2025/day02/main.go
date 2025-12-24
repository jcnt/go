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
	for _, v := range input {
		a := strings.Split(v, "-")
		start, _ := strconv.Atoi(a[0])
		end, _ := strconv.Atoi(a[1])
		for i := start; i <= end; i++ {
			//			fmt.Println(i)
			s := strconv.Itoa(i)
			if len(s)%2 == 0 {
				l := strings.Split(s, "")
				//				fmt.Println(l)
				frs := l[0 : len(l)/2]
				sec := l[len(l)/2:]
				if slices.Equal(frs, sec) {
					elfid, _ := strconv.Atoi(s)
					fmt.Println(elfid)
					count += elfid
					//					fmt.Println("EQUAL-----------------------> ")
				}
			}
		}
	}
	fmt.Println(count)

}
