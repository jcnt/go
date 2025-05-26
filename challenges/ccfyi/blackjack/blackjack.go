package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	// Fix this part, need a simple stdin -> int conversion
	temp := string(char)
	fmt.Println("temp: ", temp)
	players, err = strconv.Atoi(temp)
	fmt.Println("players: ", players)
	// fix this part

	// start game
	cplayer = rand.Intn(players) + 1
	for i := 0; i < cplayer; i++ {
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
