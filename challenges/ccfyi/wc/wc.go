package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input []string

func main() {

	alen := len(os.Args)

	if alen == 1 {
		// wc ...
	} else if alen == 2 {
		// wc -l ...
		if os.Args[1] == "-l" {
			readstdin()
			argl(0)
		} else if os.Args[1] == "-c" {
			readstdin()
			argc(0)
		} else if os.Args[1] == "-m" {
			readstdin()
			argm(0)
		} else if os.Args[1] == "-w" {
			readstdin()

		} else if os.Args[1] == "-L" {
			readstdin()

		} else {
			// wc file
		}
	} else if alen == 3 {
		// wc -l file
		if os.Args[1] == "-l" {
			readfile(2)
			argl(2)
		} else if os.Args[1] == "-c" {

		} else if os.Args[1] == "-m" {

		} else if os.Args[1] == "-w" {

		} else if os.Args[1] == "-L" {

		} else {
			// wc typo
		}
	} else {

	}
}

func readstdin() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		input = append(input, in.Text())
	}
}

func readfile(a int) {
	in, err := os.ReadFile(os.Args[a])
	if err != nil {
		fmt.Println("some error")
		return
	}
	input = strings.Split(string(in), "\n")
}

func noarg() {

}

func argl(a int) {
	if a == 0 {
		fmt.Println("     ", len(input))
	} else {
		fmt.Println("     ", len(input), os.Args[a])
	}
}

func argc(a int) {
}

func argm(a int) {

}

func argw(a int) {

}

func argL(a int) {
	var c int
	for _, v := range input {
		c += len(v)
	}
	fmt.Println(c)
}
