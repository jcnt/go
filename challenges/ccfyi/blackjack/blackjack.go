package main

import (
	"fmt"
	"math/rand"
	"time"
)

var used map[string]int
var cards []string
var players, cplayer int

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

	players = 2

	// start game
	cplayer = rand.Intn(players) + 1
	fmt.Println(players)
	fmt.Println(cplayer)
	nextPlayer()
	fmt.Println(cplayer)
}

func pullCard() int {
	return rand.Intn(len(cards))
}

func nextPlayer() {
}
