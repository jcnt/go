package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []string
	in, err := os.ReadFile("input")
	if err != nil {
		print("failed")
	}
	input = strings.Split(string(in), "\n")

	var counter int
	var dial int16
	dial = 50

	for _, v := range input {
		var num int16
		a := strings.SplitN(v, "", 2) // I need the first char L/R then the number
		if len(a) != 0 {
			lr := a[0]
			n, err := strconv.Atoi(a[1])
			if err != nil {
				fmt.Println("failed to convert")
			}
			num = int16(n)
			for num > 100 {
				fmt.Println(lr, num)
				num = num - 100
				counter++
				fmt.Println(lr, num)
				fmt.Printf("Counter is: \t %v\n", counter)
			}
			fmt.Printf("%v, %v ==>  ", lr, num)
			if lr == "L" {
				if dial == 0 {
					counter-- // need to decrement because we'll go to negative and will increment unnecessary
				}
				dial -= num
				if dial < 0 {
					dial += 100
					counter++
				}
			} else if lr == "R" {
				dial += num
				if dial > 100 {
					dial -= 100
					counter++
				}
			}
			if dial == 0 {
				counter++
			} else if dial == 100 {
				dial = 0
				counter++
			}
			fmt.Println(dial)
			fmt.Printf("Counter is: \t %v\n", counter)
		}
	}
}
