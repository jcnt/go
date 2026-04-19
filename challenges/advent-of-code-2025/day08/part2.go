// Advent of Code, 2025. Day 8, part 2.

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
	slen := []int{}
	changed := true
	lastrun := []int{}

	in, err := os.ReadFile("input")
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

	i := 1
	for changed {
		changed = false
		//fmt.Println("current run is: ", i, dist_db[i])
		if slices.Contains(boxlist, dist_db[i].a) && slices.Contains(boxlist, dist_db[i].b) {
			// if we have both we need to merge them
			var a, b int
			for j := range conn {
				if slices.Contains(conn[j], dist_db[i].a) {
					a = j
				}
				if slices.Contains(conn[j], dist_db[i].b) {
					b = j
				}
			}
			if a > b {
				conn[b] = append(conn[b], conn[a]...)
				conn = slices.Delete(conn, a, a+1)
				//fmt.Println(conn)
				changed = true
			} else if b > a {
				conn[a] = append(conn[a], conn[b]...)
				conn = slices.Delete(conn, b, b+1)
				changed = true
				//fmt.Println(conn)
			} else if b == a && len(conn) > 1 {
				changed = true
			}
		} else if slices.Contains(boxlist, dist_db[i].a) || slices.Contains(boxlist, dist_db[i].b) {
			// if we have one we need to add the other
			for j := range conn {
				//fmt.Println("boxlist is: ", boxlist)
				if slices.Contains(conn[j], dist_db[i].a) && slices.Contains(conn[j], dist_db[i].b) == false {
					conn[j] = append(conn[j], dist_db[i].b)
					boxlist = append(boxlist, dist_db[i].b)
					if len(boxlist) == len(input) {
						//fmt.Println("WE ARE DONE")
						lastrun = append(lastrun, dist_db[i].a, dist_db[i].b)
					}
					changed = true
					//fmt.Println("boxlist is: ", boxlist)
				} else if slices.Contains(conn[j], dist_db[i].b) && slices.Contains(conn[j], dist_db[i].a) == false {
					conn[j] = append(conn[j], dist_db[i].a)
					boxlist = append(boxlist, dist_db[i].a)
					if len(boxlist) == len(input) {
						//fmt.Println("WE ARE DONE")
						lastrun = append(lastrun, dist_db[i].a, dist_db[i].b)
					}
					changed = true
					//fmt.Println("boxlist is: ", boxlist)
				}
			}
		} else {
			// if we have none, we need to add both into a new junction
			conn = append(conn, []int{dist_db[i].a, dist_db[i].b})
			boxlist = append(boxlist, dist_db[i].a, dist_db[i].b)
			if len(boxlist) == len(input) {
				//fmt.Println("WE ARE DONE")
				lastrun = append(lastrun, dist_db[i].a, dist_db[i].b)
			}
			changed = true
			//fmt.Println("boxlist is: ", boxlist)
		}
		//fmt.Println("conn is: ", conn)
		i++
		if len(boxlist) < len(input) {
			changed = true
		}
	}

	for _, v := range conn {
		slen = append(slen, len(v))
	}
	sort.Slice(slen, func(i, j int) bool { return slen[i] > slen[j] })

	fmt.Println(conn)
	fmt.Println(len(input))
	fmt.Println(i)
	fmt.Println(lastrun)
	fmt.Println(answer(input[lastrun[0]], input[lastrun[1]]))
}

func answer(s1, s2 string) int {
	la := strings.Split(s1, ",")
	lb := strings.Split(s2, ",")
	a, _ := strconv.Atoi(la[0])
	b, _ := strconv.Atoi(lb[0])
	return a * b
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
