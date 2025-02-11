package main

import (
	"fmt"
	"os"
	"strconv"
)

var line, byte int
var file string
var err error

func main() {
	if len(os.Args) == 1 {
		fmt.Println("for now needs an argument")
		return
	}
	if len(os.Args) == 2 {
		fmt.Println("invalid number of arguments")
		return
	}
	if len(os.Args) > 4 {
		fmt.Println("usage: head [-n lines | -c bytes] [file ...]")
		return
	}
	if os.Args[1] == "-n" {
		line, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("parameter for -n should be the line of numbers")
			return
		}
		file = os.Args[3]
	} else if os.Args[1] == "-c" {
		byte, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("parameter for -c should be the amount of bytes")
			return
		}
		file = os.Args[3]
	} else {
		file = os.Args[1]
		line = 10
	}

	fmt.Println(line)
	fmt.Println(byte)
	fmt.Println(file)
}
