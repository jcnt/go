package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input []string
var output []string

func main() {

	alen := len(os.Args)
	var sub string

	if alen == 1 {
		// error
	} else if alen == 2 {
		// grep string stdin
		readstdin()
		sub = os.Args[1]
		noarg(sub)
		printout()
	} else if alen == 3 {
		if os.Args[1] == "-i" {
			// grep -i string stdin
			readstdin()
			sub = os.Args[2]
			argi(sub)
			printout()
		} else if os.Args[1] == "-v" {
			// grep -v string stdin
			readstdin()
			sub = os.Args[2]
			argv(sub)
			printout()
		} else {
			// grep string file
			readfile(2)
			sub = os.Args[1]
			noarg(sub)
			printout()
		}
	} else if alen == 4 {
		// grep -i string file
		if os.Args[1] == "i" {
			readfile(3)
			sub = os.Args[2]
			fmt.Println(sub) // <--- DOESNT WORK HERE!!!!
			argi(sub)
			printout()
		} else if os.Args[1] == "v" {

		} else if os.Args[1] == "c" {

		} else {
			//something error
		}
	}
}

func readstdin() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		input = append(input, in.Text())
	}
}

func readfile(f int) {
	in, err := os.ReadFile(os.Args[f])
	if err != nil {
		fmt.Printf("Error comes here\n")
	}
	input = strings.Split(string(in), "\n")
}

func printout() {
	for _, v := range output {
		fmt.Println(v)
	}
}

func noarg(s string) {
	for _, v := range input {
		if strings.Contains(v, s) {
			output = append(output, v)
		}
	}
}

func argi(s string) {
	s = strings.ToLower(s)
	for _, v := range input {
		v = strings.ToLower(v)
		if strings.Contains(v, s) {
			output = append(output, v)
		}
	}
}

func argv(s string) {
	for _, v := range input {
		if !strings.Contains(v, s) {
			output = append(output, v)
		}
	}
}

func argc() {

}
