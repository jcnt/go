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
var newcard string

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

	for i := 0; i < 2; i++ {
		for j := 0; j < players; j++ {
			c := pullCard()
			playerdb[j][c]++
		}
	}

	// remove this later
	fmt.Println("new.....")
	fmt.Println("used", used)
	fmt.Println("players", playerdb)
	// remove this

	for {
		fmt.Println()
		fmt.Println("Current player is: ", cplayer)
		fmt.Println("Your current cards are: ")
		var s int
		for k, v := range playerdb[cplayer-1] {
			if v > 0 {
				fmt.Println(k, v)
				s += v
			}
		}
		fmt.Printf("You currently have %v cards in your hand\n", s)
		csum := calcCards(playerdb[cplayer-1])
		fmt.Println("Sum of your cards are ", csum)

		fmt.Print("Do you want a card? Y/N ")
		if _, err := fmt.Scan(&newcard); err != nil {
			fmt.Println("read failed")
			return
		}
		fmt.Println(newcard)

		if newcard == "Y" {
			c := pullCard()
			playerdb[cplayer-1][c]++
		}
		fmt.Println(playerdb[cplayer-1])

		s = 0

		for k, v := range playerdb[cplayer-1] {
			if v > 0 {
				fmt.Println(k, v)
				s += v
			}
		}
		fmt.Printf("You currently have %v cards in your hand\n", s)
		csum = calcCards(playerdb[cplayer-1])
		fmt.Println("Sum of your cards are ", csum)

		nextPlayer()

	}

}

func pullCard() string {
	r := rand.Intn(len(cards))
	if used[cards[r]] < 4 {
		used[cards[r]]++
	} else {
		pullCard()
	}
	return cards[r]
}

func nextPlayer() {
	if cplayer == players {
		cplayer = 1
	} else {
		cplayer++
	}
}

func calcCards(c map[string]int) int {
	s := 0
	for k, v := range c {
		if k == "2" {
			s = s + 2*v
		} else if k == "3" {
			s = s + 3*v
		} else if k == "4" {
			s = s + 4*v
		} else if k == "5" {
			s = s + 5*v
		} else if k == "6" {
			s = s + 6*v
		} else if k == "7" {
			s = s + 7*v
		} else if k == "8" {
			s = s + 8*v
		} else if k == "9" {
			s = s + 9*v
		} else if k == "a" {
			if s == 10 {
				s += 11
			} else {
				s = s + 1*v
			}
		} else {
			s = s + 10*v
		}
		//handle the ace case
		if s == 11 && c["a"] > 0 {
			s = 21
		}
	}
	return s
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

// todo
// - handle double aces
