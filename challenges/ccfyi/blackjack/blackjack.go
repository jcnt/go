package main

import (
	"fmt"
	"math/rand"
	"time"
)

var used map[string]int
var cards []string
var users []int

func main() {

	cards = append(cards, "2", "3", "4", "5", "6", "7", "8", "9", "j", "q", "k", "a")
	rand.NewSource(time.Now().UnixNano())

	used = map[string]int{
		"2":  0,
		"3":  0,
		"4":  0,
		"5":  0,
		"6":  0,
		"7":  0,
		"8":  0,
		"9":  0,
		"10": 0,
		"j":  0,
		"q":  0,
		"k":  0,
		"a":  0,
	}

	fmt.Println("\nWelcom to Black Jack! \n")
	c := pullCard()
	fmt.Println(cards[c])
	used[cards[c]]++
	fmt.Println(used)

	for i := 0; i < 2; i++ {
		users = append(users, 0)
	}

	// start game
	users[rand.Intn(2)] = 1
}

func pullCard() int {
	return rand.Intn(len(cards))
}
