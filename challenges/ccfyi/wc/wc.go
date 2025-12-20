package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input []string
var args []string

func main() {

	if len(os.Args) == 1 {
		readstdin()
		noarg()
	} else {

		_, err := os.Stat(os.Args[len(os.Args)-1])
		if err != nil {
			readstdin()
			args = os.Args[1:]
		} else {
			readfile(os.Args[len(os.Args)-1])
			args = os.Args[1 : len(os.Args)-1]
		}
	}

	fmt.Printf("%#v", input)
	fmt.Println(args)
}

func readstdin() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		input = append(input, in.Text())
	}
}

func readfile(f string) {
	in, err := os.ReadFile(f)
	if err != nil {
		fmt.Printf("%s: %s: open: No such file or directory\n", os.Args[0], f)
		return
	}
	input = strings.Split(string(in), "\n")
}

func noarg() {
	fmt.Println(input)
}
