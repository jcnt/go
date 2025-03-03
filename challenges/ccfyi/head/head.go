package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input []string

func main() {

	alen := len(os.Args)
	var lines, bytes int
	var err error

	if alen > 1 {

		if os.Args[1] != "-n" && os.Args[1] != "-c" {
			// head filename or head -10 (stdin)
			s := string(string(os.Args[1]))

			if string(s[0]) == "-" {
				// head -10 stdin
				n := string(s[1:])
				lines, err = strconv.Atoi(n)
				if err != nil {
					fmt.Println("head: illegal line count --", n)
					return
				}
				readstdin()
				argn(lines)

			} else {
				// head filename
				lines = 10
				readfile(1)
				argn(lines)
			}

		} else if os.Args[1] == "-n" || os.Args[1] == "-c" {
			// head -n/-c

			if alen < 3 {
				fmt.Println("head: option requires an argument --", os.Args[1])
				fmt.Println("usage: head [-n lines | -c bytes] [file ...]")
				return

			} else if alen == 3 {

				if os.Args[1] == "-n" {
					lines, err = strconv.Atoi(os.Args[2])
					if err != nil {
						fmt.Println("head: illegal line count --", os.Args[2])
						return
					}
					readstdin()
					argn(lines)

				} else if os.Args[1] == "-c" {
					bytes, err = strconv.Atoi(os.Args[2])

					if err != nil {
						fmt.Println("head: illegal byte count --", os.Args[2])
						return
					}
					readstdin()
					argc(bytes)
				}

			} else if alen == 4 {

				if os.Args[1] == "-n" {
					lines, err = strconv.Atoi(os.Args[2])

					if err != nil {
						fmt.Println("head: illegal line count --", os.Args[2])
						return
					}
					readfile(3)
					argn(lines)

				} else if os.Args[1] == "-c" {
					bytes, err = strconv.Atoi(os.Args[2])
					if err != nil {
						fmt.Println("head: illegal byte count --", os.Args[2])
						return
					}
					readfile(3)
					argc(bytes)
				}
			}

		} else {
			// head -10
			fmt.Println(os.Args[1])
		}

	} else {
		// head (stdin)
		lines = 10
		readstdin()
		argn(lines)
	}

}

func readfile(a int) {
	// read file to input
	in, err := os.ReadFile(os.Args[a])
	if err != nil {
		fmt.Println("head:", os.Args[a], "No such file or directory")
		return
	}
	input = strings.Split(string(in), "\n")
}

func readstdin() {
	// read stdin to input
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		input = append(input, in.Text())
	}
}

func argn(l int) {
	if l > len(input) {
		l = len(input)
	}
	for i := 0; i < l; i++ {
		fmt.Println(input[i])
	}
}

func argc(b int) {
	i := 0
	for {
		if len(input[i]) < b {
			fmt.Println(input[i])
			b -= len(input[i])
			i++
		} else {
			for j := 0; j < b; j++ {
				fmt.Print(string(input[i][j]))
			}
			return
		}
	}
}
