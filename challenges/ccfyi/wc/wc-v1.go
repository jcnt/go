package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var input []string

func main() {

	alen := len(os.Args)

	if alen == 1 {
		// wc ...
		readstdin()
		noarg(0)
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
			argw(0)
		} else if os.Args[1] == "-L" {
			readstdin()
			argL(0)
		} else {
			// wc file
			readfile(1)
			noarg(1)
		}
	} else if alen == 3 {
		// wc -l file
		if os.Args[1] == "-l" {
			readfile(2)
			argl(2)
		} else if os.Args[1] == "-c" {
			readfile(2)
			argc(2)
		} else if os.Args[1] == "-m" {
			readfile(2)
			argm(2)
		} else if os.Args[1] == "-w" {
			readfile(2)
			argw(2)
		} else if os.Args[1] == "-L" {
			readfile(2)
			argL(2)
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
		fmt.Printf("%s: %s: open: No such file or directory\n", os.Args[0], os.Args[a])
		return
	}
	input = strings.Split(string(in), "\n")
}

func noarg(a int) {
	var l, w, c int
	l = len(input)
	for _, v := range input {
		w += len(strings.Split(v, " "))
		c += len(v)
		c += 1
	}
	if a == 0 {
		fmt.Println("      ", l, "    ", w, "   ", c)
	} else {
		fmt.Println("      ", l, "    ", w, "   ", c, os.Args[a])
	}
}

func argl(a int) {
	if a == 0 {
		fmt.Println("     ", len(input))
	} else {
		fmt.Println("     ", len(input), os.Args[a])
	}
}

func argm(a int) {
	var m int
	for _, v := range input {
		m += utf8.RuneCountInString(v)
		m += 1
	}
	if a == 0 {
		fmt.Println("    ", m)
	} else {
		fmt.Println("    ", m, os.Args[a])
	}
}

func argc(a int) {
	var c int
	for _, v := range input {
		c += len(v)
		c += 1
	}
	if a == 0 {
		fmt.Println("    ", c)
	} else {
		fmt.Println("    ", c, os.Args[a])
	}
}

func argw(a int) {
	var w int
	for _, v := range input {
		w += len(strings.Split(v, " "))
	}
	if a == 0 {
		fmt.Println("    ", w)
	} else {
		fmt.Println("    ", w, os.Args[a])
	}
}

func argL(a int) {
	var c int
	for _, v := range input {
		if len(v) > c {
			c = len(v)
		}
	}
	if a == 0 {
		fmt.Println("    ", c)
	} else {
		fmt.Println("    ", c, os.Args[a])
	}
}
