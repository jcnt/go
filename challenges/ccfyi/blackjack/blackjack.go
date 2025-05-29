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

	fmt.Print("How many players? (max 7): ")
	if _, err := fmt.Scan(&players); err != nil {
		fmt.Println("read failed")
		return
	}
	fmt.Printf("There are %v players in this game. Let's go!\n", players)
	cplayer = rand.Intn(players) + 1
	for i := 0; i < players; i++ {
		playerdb = append(playerdb, map[string]int{"q": 0})
	}

	for i := 0; i < players; i++ {
		c := pullCard()
		used[c]++
		playerdb[i][c]++
	}

	fmt.Println("new.....")
	fmt.Println(used)
	fmt.Println(playerdb)

	fmt.Println(playerdb)
	fmt.Println(len(playerdb))
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

func pullCard() string {
	return rand.Intn(len(cards))
}

func nextPlayer() {
	if cplayer == players {
		cplayer = 1
	} else {
		cplayer++
	}
}

// gameflow
// welcome message
// ask for number of players
// distibute cards (pullCard 2x for every player)
// mark who is starting the game
// loop:
// - print current player
// - ask if more card
// - pull card
// - check if that card is gone full (all 4 is on the table)
// - check if there's a 21
