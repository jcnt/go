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
	fmt.Println(dist_db)

	for i := 0; i < 10; i++ {
		for j := range conn {
			ctemp := []int{}
			fmt.Println("pre if")
			if slices.Contains(conn[j], dist_db[i].a) {
				conn[j] = append(conn[j], dist_db[i].b)
			} else if slices.Contains(conn[j], dist_db[i].a) {
				conn[j] = append(conn[j], dist_db[i].a)
			} else {
				ctemp = append(ctemp, dist_db[i].a)
				ctemp = append(ctemp, dist_db[i].b)
				conn = append(conn, ctemp)
				fmt.Println("were ctemp", ctemp, conn)
			}
		}
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
