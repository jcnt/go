// first attempt to blackjack. don't expect anything pretty, trying to get it working
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var used map[string]int
var playerdb []map[string]int
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

	fmt.Print("How many players? (max 7): ")
	if _, err := fmt.Scan(&players); err != nil {
		fmt.Println("read failed")
		return
	}
	fmt.Println("players: ", players)

	// start game
	cplayer = rand.Intn(players) + 1
	for i := 0; i < players; i++ {
		playerdb = append(playerdb, map[string]int{"q": 0})
	}
	fmt.Println(playerdb)
	fmt.Println("cplayer: ", cplayer)
	fmt.Println(playerdb[cplayer-1])
	//	playerdb[cplayer]map{"q": 1;}
	nextPlayer()
	fmt.Println("cplayer: ", cplayer)
	nextPlayer()
	fmt.Println("cplayer: ", cplayer)
	nextPlayer()
	fmt.Println("cplayer: ", cplayer)
	nextPlayer()
	fmt.Println("cplayer: ", cplayer)
	playerdb[cplayer-1]["q"]++
	fmt.Println(playerdb)
}

func pullCard() int {
	return rand.Intn(len(cards))
}

func nextPlayer() {
	if cplayer == players {
		cplayer = 1
	} else {
		cplayer++
	}
}
