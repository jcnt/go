package main

import (
	"fmt"
	"math/rand"
	"time"
)

var used map[string]int
var cards []string
var players []int

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
		players = append(players, 0)
	}

	// start game
	players[rand.Intn(2)] = 1
	fmt.Println(players)
	nextPlayer()
	fmt.Println(players)
}

func pullCard() int {
	return rand.Intn(len(cards))
}

func nextPlayer() {

	for i := range players {
		if i == len(players)-1 {
			players[i] = 0
			players[0] = 1
		} else if players[i] == 1 {
			players[i] = 0
			players[i+1] = 1
		}
	}
}
