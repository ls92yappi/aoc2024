package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	//"regexp"
	//"strconv"
	//"strings"
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

// assumes begin and target are the sz-based indices within grid
// assumes every move that doesn't hit a wall adds 1 to the distance
func Dijkstra(grid []int, sz int, begin, target int) int {
	beginR := begin/sz
	beginC := begin%sz
	distance := -1 // not found indicator, as final cost >= 0
	mc := 1 // movement cost

	// Dijkstra's algorithm below
	var q deq.Deq[Pos]      // distances queue
	v := make([]bool,sz*sz) // visited list
	v[0] = true
	start := Pos{beginR,beginC,0} // initial cost = 0
	q.PushFront(start)

	// find shortest path from begin to target
	for q.Len() > 0 {
		p := q.PopBack() // current position in queue
		dirs := []Pos{Pos{p.R+1,p.C,p.D+mc}, Pos{p.R,p.C+1,p.D+mc}, Pos{p.R-1,p.C,p.D+mc}, Pos{p.R,p.C-1,p.D+mc}, }
		// down,right,up,left
		for _,loc := range(dirs) {
			idx := loc.R*sz+loc.C
			//can't operate beyond edges
			if loc.R < 0 || loc.C < 0 || loc.R > sz-1 || loc.C > sz-1 {
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
	lines, n, _ := ReadWholeFile(fname)
	answer := -1

	limit := 1024 // 12 ex
	sz := 71 // 7 ex

	grid := make([]int, sz*sz)

	// construct hollow maze
	for i := range(Min2(n,limit)) {
		line := lines[i]
		if len(line) < 2 { continue }
		// drop the corruption wall
		coords, _, _ := IntSlice(line,",")
		x := coords[0]
		y := coords[1]
		grid[x*sz+y] = TooBig
	} // for i

	answer = Dijkstra(grid, sz, 0, sz*sz-1)

	return answer

	// Submissions:
	// (1) 340 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
