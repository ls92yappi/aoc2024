package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	"strings"
	. "github.com/ls92yappi/aoc" // problems with go proxy not being even remotely up-to-date
	"github.com/ls92yappi/aoc/deq"
)

///////////////////////////////////////////////////////////


///////////////////////////////////////////////////////////

const TooBig = 987654321

type Pos struct {
	R int
	C int
	D int // distance
}

// Dijkstra computes shortest path from begin to target
// assumes begin and target are the szWide-based indices within grid
// assumes every move that doesn't hit a wall adds 1 to the distance
func Dijkstra(grid []int, szWide int, begin, target int) int {
	// no moves means distance=0
	if begin == target {
		return 0
	}
	// shortcircuit known impossible source or destination
	if grid[begin] == TooBig || grid[target] == TooBig {
		return -1
	}
	beginR := begin/szWide
	beginC := begin%szWide
	distance := -1 // not found indicator, as final cost >= 0
	mc := 1 // movement cost
	numRows := len(grid)/szWide

	// Dijkstra's algorithm below
	var q deq.Deq[Pos]      // distances queue
	v := make([]bool,len(grid)) // visited list
	v[begin] = true
	start := Pos{beginR,beginC,0} // initial cost = 0
	q.PushFront(start)

	// find shortest path from begin to target
	for q.Len() > 0 {
		p := q.PopBack() // current position in queue
		dirs := []Pos{Pos{p.R+1,p.C,p.D+mc}, Pos{p.R,p.C+1,p.D+mc}, Pos{p.R-1,p.C,p.D+mc}, Pos{p.R,p.C-1,p.D+mc}, }
		// down,right,up,left
		for _,loc := range(dirs) {
			idx := loc.R*szWide+loc.C
			//can't operate beyond edges
			if loc.R < 0 || loc.C < 0 || loc.R > numRows-1 || loc.C > szWide-1 {
				continue
			}
			// skip corruption walls
			if grid[idx] == TooBig {
				continue
			}
			// skip already visited spots
			if v[idx] {
				continue
			}
			// check if we are at the destination
			if idx == target {
				return loc.D // distance is the target location's distance
			}
			// mark the spot as visited
			v[idx] = true
			// treat as a queue
			q.PushFront(loc)
		}
	}
	return distance // should only be -1 if it gets here, which means no valid path
}


// in 3450 71x71 grid
// ex 25   7x7 grid
func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, _, _ := ReadWholeFile(fname)
	answer := 0
	begin := -1
	target := -1

	sz := len(lines[0]) // 141 input or 15 ex

	grid := make([]int, sz*sz)

	// construct hollow maze
	for i := range(sz) {
		line := lines[i]

		// find start or endpoints as we go
		begSpot := strings.Index(line, "S")
		endSpot := strings.Index(line, "E")
		if begSpot > -1 { begin = i*sz+begSpot }
		if endSpot > -1 { target= i*sz+endSpot }

		// initial placement of walls
		for j := range(sz) {
			if string(line[j]) == "#" {
				grid[i*sz+j] = TooBig
			}
		}
	} // for i

	baseDistance := Dijkstra(grid, sz, begin, target) 
	//answer = baseDistance
	Threshold := 100 // minimum improvement to count
	// skip outer walls
	for i := 1; i < sz-1; i++ {
		for j := 1; j < sz-1; j++ {
			loc := i*sz+j
			if grid[loc] == TooBig {
				// temporarily remove a single wall at a time
				grid[loc] = 0
				distance := Dijkstra(grid, sz, begin, target)
				if distance > -1 && distance <= baseDistance-Threshold {
					//fmt.Printf("Removing (%d,%d) save %d steps\n", i, j, baseDistance-distance)
					answer++
				}
				// then restore that wall
				grid[loc] = TooBig
			}
		}
	}

	return answer

	// ran in about 2.6seconds
	// Submissions:
	// (1) 1311 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
