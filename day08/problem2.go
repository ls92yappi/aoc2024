package main

// See README.md for problem description

import (
	"fmt"
	//"strconv"
	//"strings"
	. "github.com/ls92yappi/aoc"
)

///////////////////////////////////////////////////////////

type Pos struct {
	R int
	C int
}

var A map[byte][]Pos // antennas as a hash map by letter // uint8
var a [][]int    // antinodes grid

const Ignore = '.'

// ex 12x12 grid
// in 50x50 grid
func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, n, _ := ReadWholeFile(fname)
	answer := 0

	// allocate memory for the data structures
	A = make(map[byte][]Pos, 0)
	a = make([][]int, n)
	for i := range(n) {
		a[i] = make([]int, n)
	}

	// create Antennas map first
	for i := range(n) {
		line := lines[i]
		if len(line) < 1 {
			continue
		}
		for j := range(n) {
			var ch byte = line[j]
			if ch == Ignore {
				continue
			}
			p := Pos{i,j}
			A[ch] = append(A[ch], p)
		}
	} // for i

	// populate antinodes grid
	for k,list := range(A) {
		numAnts := len(list)
		// antinodes are formed pairwise, can skip singletons
		if numAnts < 2 {
			continue
		}
		// form all pairs
		for x := range(numAnts-1) {
			p := list[x]
			for y := range(numAnts-x-1) {
				z := x+y+1 // pair index for y
				q := list[z]
				// create antinode candidate locations
				for d := range(n-1) {
					// variable distance antinodes
					c1 := Pos{p.R+d*(p.R-q.R), p.C+d*(p.C-q.C)}
					c2 := Pos{q.R+d*(q.R-p.R), q.C+d*(q.C-p.C)}
					// verify that are within the grid
					offgrid1 := c1.R < 0 || c1.C < 0 || c1.R >= n || c1.C >= n
					offgrid2 := c2.R < 0 || c2.C < 0 || c2.R >= n || c2.C >= n
					//fmt.Printf("c1(%s,%d,%d) @ (%d,%d):(%d,%d) = Pos{%d,%d} %v\n", string(k),x,y, p.R,p.C, q.R,q.C, c1.R, c1.C, !offgrid1)
					//fmt.Printf("c2(%s,%d,%d) @ (%d,%d):(%d,%d) = Pos{%d,%d} %v\n", string(k),x,y, p.R,p.C, q.R,q.C, c2.R, c2.C, !offgrid2)
					//fmt.Println()
					// mark antinodes that are ok
					if !offgrid1 {
						a[c1.R][c1.C]++
					}
					if !offgrid2 {
						a[c2.R][c2.C]++
					}
				}
			}
		}
		_ = k
	}

	// Count up ongrid antinode locations
	for i := range(n) {
		for j := range(n) {
			if a[i][j] > 0 {
				answer++
			}
		}
	}
	return answer

	// Submissions:
	// (1) 809 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
