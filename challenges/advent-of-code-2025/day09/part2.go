// Advent of Code, 2025. Day 9, part 2.

package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rgdb := [][]int{}
	var maxx, maxy int
	type rect struct {
		a int
		b int
		s int
	}

	in, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("cannot open file")
	}
	input := strings.Split(string(in), "\n")
	input = input[0 : len(input)-1]
	//fmt.Printf("%q\n", input)

	reddb := make([][]int, 0, len(input))
	fmt.Println("appending reddb")
	for _, v := range input {
		intred := []int{}
		red := strings.Split(v, ",")
		for j := range red {
			r, _ := strconv.Atoi(red[j])
			intred = append(intred, r)
			if j == 0 && r > maxx {
				maxx = r
			} else if j == 1 && r > maxy {
				maxy = r
			}
		}
		reddb = append(reddb, intred)
	}
	//fmt.Printf("%d\n\n", reddb)

	fmt.Println("maxy", maxy, "maxx", maxx)
	fmt.Println("appending tiledb")
	tiledb := make([][]string, maxy+2, maxy+2)
	tline := strings.Split(strings.Repeat(".", maxx+2), "")
	for i := 0; i < maxy+2; i++ {
		tl2 := slices.Clone(tline)
		tiledb[i] = tl2
		fmt.Println(i)
	}

	reddb = append(reddb, reddb[0])

	fmt.Println("doing rgdb")
	for i := 1; i < len(reddb); i++ {
		c0 := reddb[i-1][0]
		c1 := reddb[i-1][1]
		n0 := reddb[i][0]
		n1 := reddb[i][1]
		if c0 == n0 {
			if c1 < n1 {
				for j := c1 + 1; j < n1; j++ {
					rgdb = append(rgdb, []int{c0, j})
				}
			} else {
				for j := n1 + 1; j < c1; j++ {
					rgdb = append(rgdb, []int{c0, j})
				}
			}
		} else {
			if c0 < n0 {
				for j := c0 + 1; j < n0; j++ {
					rgdb = append(rgdb, []int{j, c1})
				}
			} else {
				for j := n0 + 1; j < c0; j++ {
					rgdb = append(rgdb, []int{j, c1})
				}
			}
		}
	}

	fmt.Println("doing the # fillup")
	for _, t := range reddb {
		tiledb[t[1]][t[0]] = "#"
	}
	for _, t := range rgdb {
		tiledb[t[1]][t[0]] = "#"
	}

	fmt.Println("doing the x fillup")
	for i := range tiledb {
		for j := range tiledb[i] {
			if tiledb[i][j] == "." {
				tiledb[i][j] = "x"
			} else {
				break
			}
		}
		for j := len(tiledb[i]) - 1; j >= 0; j-- {
			if tiledb[i][j] == "." {
				tiledb[i][j] = "x"
			} else {
				break
			}
		}
	}

	for i := range tiledb[0] {
		for j := range tiledb {
			if tiledb[j][i] == "." {
				tiledb[j][i] = "x"
			} else {
				break
			}
		}
		for j := len(tiledb) - 1; j >= 0; j-- {
			if tiledb[j][i] == "." {
				tiledb[j][i] = "x"
			} else {
				break
			}
		}
	}

	rsum := 0
	for i := range input {
		rsum += i + 1
	}

	rectdb := make([]rect, 0, rsum)
	fmt.Println("running get_area")
	for i := range reddb {
		for j := i + 1; j < len(reddb); j++ {
			var c rect
			c.s = get_area(reddb[i], reddb[j])
			c.a = i
			c.b = j
			rectdb = append(rectdb, c)
		}
	}
	sort.Slice(rectdb, func(i, j int) bool { return rectdb[i].s > rectdb[j].s })
	fmt.Println(rectdb)

	fmt.Println("testing rectangles")
	for _, r := range rectdb {
		var x0, x1, y0, y1 int
		a0 := reddb[r.a][0]
		a1 := reddb[r.a][1]
		b0 := reddb[r.b][0]
		b1 := reddb[r.b][1]
		//fmt.Println(r)
		//fmt.Println(reddb[r.a], reddb[r.b])
		//fmt.Println(a0, a1, b0, b1)
		if a0 <= b0 {
			x0 = a0
			y0 = b0
		} else {
			x0 = b0
			y0 = a0
		}
		if a1 <= b1 {
			x1 = a1
			y1 = b1
		} else {
			x1 = b1
			y1 = a1
		}
		out := false
		for i := x0; i <= y0; i++ {
			for j := x1; j <= y1; j++ {
				//fmt.Println(tiledb[j][i])
				if tiledb[j][i] == "x" {
					out = true
				}
			}
		}
		if out == false {
			fmt.Println("winner is", r.s)
			break
		}
		fmt.Println("current is", r.s)
	}
}

func get_area(a, b []int) int {
	t := (a[0] - b[0] + 1) * (a[1] - b[1] + 1)
	if t < 0 {
		return t * -1
	} else {
		return t
	}
}
