// Advent of Code, 2025. Day 7, part 1.

package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {

	type dist struct {
		a int
		b int
		d float64
	}

	dist_db := []dist{}
	conn := [][]int{}
	boxlist := []int{}

	in, err := os.ReadFile("example")
	if err != nil {
		fmt.Println("unable to open file")
	}
	input := strings.Split(string(in), "\n")
	input = input[0 : len(input)-1]
	//fmt.Printf("%q\n", input)

	for i := range input {
		for j := i + 1; j < len(input); j++ {
			var tdb dist
			tdb.a = i
			tdb.b = j
			tdb.d = get_distance(input[i], input[j])
			dist_db = append(dist_db, tdb)
		}
	}
	sort.Slice(dist_db, func(i, j int) bool { return dist_db[i].d < dist_db[j].d })
	conn = append(conn, []int{dist_db[0].a, dist_db[0].b})
	boxlist = append(boxlist, dist_db[0].a, dist_db[0].b)

	for i := 1; i < 10; i++ {
		fmt.Println("current run is: ", dist_db[i])
		if slices.Contains(boxlist, dist_db[i].a) || slices.Contains(boxlist, dist_db[i].b) {
			for j := range conn {
				fmt.Println("conn j is: ", conn[j])
				fmt.Println("boxlist is: ", boxlist)
				if slices.Contains(conn[j], dist_db[i].a) && slices.Contains(conn[j], dist_db[i].b) == false {
					conn[j] = append(conn[j], dist_db[i].b)
					boxlist = append(boxlist, dist_db[i].b)
					fmt.Println("boxlist is: ", boxlist)
				} else if slices.Contains(conn[j], dist_db[i].b) && slices.Contains(conn[j], dist_db[i].a) == false {
					conn[j] = append(conn[j], dist_db[i].a)
					boxlist = append(boxlist, dist_db[i].a)
					fmt.Println("boxlist is: ", boxlist)
				}
			}
		} else {
			conn = append(conn, []int{dist_db[i].a, dist_db[i].b})
			boxlist = append(boxlist, dist_db[i].a, dist_db[i].b)
		}
		fmt.Println("conn is: ", conn)
	}

	fmt.Println(conn)
}

func get_distance(a string, b string) float64 {
	la := strings.Split(a, ",")
	lb := strings.Split(b, ",")
	var ila, ilb []float64
	for _, v := range la {
		t, _ := strconv.Atoi(v)
		ila = append(ila, float64(t))
	}
	for _, v := range lb {
		t, _ := strconv.Atoi(v)
		ilb = append(ilb, float64(t))
	}
	sq := math.Sqrt(math.Pow(ila[0]-ilb[0], 2) + math.Pow(ila[1]-ilb[1], 2) + math.Pow(ila[2]-ilb[2], 2))
	return sq
}
