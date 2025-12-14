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

	if alen == 1 {
		// usage
		fmt.Println("usage: grep [-ivc]")
	} else if alen == 2 {
		// grep string stdin
		readstdin()
		noarg(os.Args[1])
		printout()
	} else if alen == 3 {
		if os.Args[1] == "-i" {
			// grep -i string stdin
			readstdin()
			argi(os.Args[2])
			printout()
		} else if os.Args[1] == "-v" {
			// grep -v string stdin
			readstdin()
			argv(os.Args[2])
			printout()
		} else if os.Args[1] == "-c" {
			// grep -c string stdin
			readstdin()
			noarg(os.Args[2])
			fmt.Println(len(output))
		} else {
			// grep string file
			readfile(2)
			noarg(os.Args[1])
			printout()
		}
	} else if alen == 4 {
		// grep -i string file
		if os.Args[1] == "-i" {
			readfile(3)
			argi(os.Args[2])
			printout()
		} else if os.Args[1] == "-v" {
			readfile(3)
			argv(os.Args[2])
			printout()
		} else if os.Args[1] == "-c" {
			readfile(3)
			noarg(os.Args[2])
			fmt.Println(len(output))
		} else {
			//usage error
			fmt.Printf("%v: invalid option -- %v\n", os.Args[0], os.Args[1])
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
		fmt.Printf("%v: %v: No such file or directory\n", os.Args[0], os.Args[f])
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
