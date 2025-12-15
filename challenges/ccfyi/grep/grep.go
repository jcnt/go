package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input []string

// var output []string

func main() {
	//	fmt.Println(os.Args)

	if len(os.Args) == 1 {
		readstdin()
		noarg()
	}
	_, err := os.Stat(os.Args[len(os.Args)-1])
	if err != nil {
		readstdin()
	} else {
		readfile(os.Args[len(os.Args)-1])
	}

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
		fmt.Printf("%v: %v: No such file or directory\n", os.Args[0], f)
	}
	input = strings.Split(string(in), "\n")
}

func noarg() {

}
