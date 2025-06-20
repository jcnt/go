// first attempt to blackjack. don't expect anything pretty, trying to get it working
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var used map[string]int
var playerdb []map[string]int
var cards []string
var players, cplayer int
var newcard string

func main() {

	played := map[int]int{}
	scores := map[int]int{}
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

	fmt.Println("\nWelcome to Black Jack! \n")

	fmt.Print("How many players? (max 7): ")
	if _, err := fmt.Scan(&players); err != nil {
		fmt.Println("Invalid response")
		return
	}
	fmt.Printf("There are %v players in this game. Let's go!\n\n", players)
	cplayer = rand.Intn(players) + 1
	for i := 0; i <= players; i++ {
		playerdb = append(playerdb, map[string]int{"q": 0})
	}

	for i := 0; i < 2; i++ {
		for j := 1; j < players+1; j++ {
			c := pullCard()
			playerdb[j][c]++
		}
	}

	//dealer
	var d int
	for d = 0; d < 17; {
		c := pullCard()
		playerdb[0][c]++
		d = calcCards(playerdb[0])
	}
	//	fmt.Println("Dealer's score: ", d)
	scores[0] = d

	on := true
	for on == true {
		fmt.Println()
		fmt.Println("Current player is: ", cplayer)

		p := true
		for p {
			fmt.Println("Your current cards are: ")
			var s int
			for k, v := range playerdb[cplayer] {
				if v > 0 {
					fmt.Println(k, v)
					s += v
				}
			}

			fmt.Printf("You currently have %v cards in your hand\n", s)
			csum := calcCards(playerdb[cplayer])
			fmt.Println("Sum of your cards are ", csum)

			fmt.Print("Do you want a card? Y/N ")
			if _, err := fmt.Scan(&newcard); err != nil {
				fmt.Println("Invalid response")
				//				return
			}

			if strings.ToUpper(newcard) == "Y" {
				c := pullCard()
				playerdb[cplayer][c]++
			} else if strings.ToUpper(newcard) == "N" {
				p = false
				scores[cplayer] = csum
			}

			csum = calcCards(playerdb[cplayer])
			if csum >= 21 {
				p = false
				scores[cplayer] = csum
				fmt.Println("Sum of your cards are ", csum)
				fmt.Println("No more cards for you")
			}
		}

		played[cplayer] = 1
		nextPlayer()

		if played[cplayer] == 1 {
			on = false
		}

	}

	fmt.Println()
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println("Summary of the game")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println()

	winner := 0
	for i, v := range scores {
		if v <= 21 && v > winner {
			winner = i
		}

	}

	fmt.Println("=Dealer=:", scores[0])
	for k, v := range scores {
		if k == 0 {
		} else {
			fmt.Printf("Player %v: %v\n", k, v)
		}
	}
	fmt.Println()
	if winner == 0 {
		fmt.Println("winner is the dealer with sum of", scores[winner])
	} else {
		fmt.Println("winner is player", winner, "with sum of", scores[winner])
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

// todo
// - you have to pull card if...
// - handle double aces
// - check if that card is gone full (all 4 is on the table)
// - accept y Y but nothing else
