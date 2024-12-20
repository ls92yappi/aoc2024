package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	//"regexp"
	//"strconv"
	"strings"
	. "github.com/ls92yappi/aoc" // problems with go proxy not being even remotely up-to-date
	"github.com/ls92yappi/aoc/deq"
)

///////////////////////////////////////////////////////////


///////////////////////////////////////////////////////////

const TooBig = 987654321
const MaxCheatLen = 20 // 2 for Part 1 20 for Part 2
var Threshold = 100 // minimum improvement to qualify

type Pos struct {
	R int
	C int
	D int // distance
}

// BfsDijkstra computes shortest path from begin to everywhere
// can filter to maxDist, may choose to ignore Walls
// assumes begin and target are the szWide-based indices within grid
// assumes every move that doesn't hit a wall adds 1 to the distance
// assumes Dijkstra(begin,target) = Dijkstra(target,begin)
func BfsDijkstra(grid []int, szWide int, begin int, ignoreWalls bool, maxDist int) []int {
	beginR := begin/szWide
	beginC := begin%szWide
	numRows := len(grid)/szWide
	mc := 1 // movement cost
	if maxDist < 0 {
		maxDist = len(grid)
	}

	// Dijkstra's algorithm below
	var q deq.Deq[Pos]      // distances queue
	v := make([]bool,len(grid)) // visited list
	v[begin] = true

	// start out that every location is a wall
	distance := make([]int, len(grid))
	for i := range(len(distance)) { distance[i] = -1 } // -1 = not found indicator
	// but distance from source to itself is always 0
	distance[begin] = 0

	start := Pos{beginR,beginC,0} // initial cost = 0
	q.PushFront(start)

	// find shortest path from begin to target
	for q.Len() > 0 {
		p := q.PopBack() // current position in queue
		dirs := []Pos{Pos{p.R+1,p.C,p.D+mc}, Pos{p.R,p.C+1,p.D+mc}, Pos{p.R-1,p.C,p.D+mc}, Pos{p.R,p.C-1,p.D+mc}, }
		// down,right,up,left
		// filter to maxDist here
		if p.D+mc > maxDist {
			continue
		}
		for _,loc := range(dirs) {
			idx := loc.R*szWide+loc.C
			//can't operate beyond edges
			if loc.R < 0 || loc.C < 0 || loc.R > numRows-1 || loc.C > szWide-1 {
				continue
			}
			// skip corruption walls
			if !ignoreWalls && grid[idx] == TooBig {
				continue
			}
			// skip already visited spots
			if v[idx] {
				continue
			}

			// only keep the best distance, may not need this condition
			if distance[idx] >= loc.D {
				continue
			}

			// mark the spot as visited
			v[idx] = true
			distance[idx] = loc.D
			// treat as a queue
			q.PushFront(loc)
		}
	}
	return distance
}

// Computes the Manhattan distance (grid-only travel) from begin to target
// assuming sz-based indices for a grid
func Manhattan(sz, begin, target int) int {
	// grab relevant coordinates
	fx, fy :=  begin/sz,  begin%sz
	tx, ty := target/sz, target%sz
	lx, ly := Min2(fx,tx), Min2(fy,ty)
	hx, hy := Max2(fx,tx), Max2(fy,ty)
	dist := hx-lx + hy-ly
	return dist
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

// Count number of Cheats up to MaxCheatLen given min Threshold improvement
func Cheatstra(grid []int, sz int, begin, target int) int {
	numGoodCheats := 0

	// find reference distance
	baseDistance := Dijkstra(grid, sz, begin, target)
	if baseDistance < 1 {
		// no cheats required
		return 0
	}

	// compute distances from start to every point, and from every point to end
	ds := BfsDijkstra(grid, sz, begin, false, -1)
	de := BfsDijkstra(grid, sz, target, false, -1)

	// good starting point candidates
	candidates := make([]int, 0)
	for i := range(len(grid)) {
		if ds[i] != -1 {
			candidates = append(candidates, i)
		}
	}

	bestCheatingScore := baseDistance
	fmt.Printf("%dx%d grid, runnable in %d steps, with %d candidates\n", sz, sz, baseDistance, len(candidates))
	for _,i := range(candidates) {
		// construct a bfs ignoring walls with max distance of 20 from the candidate
		ignoreWalls := BfsDijkstra(grid, sz, i, true, MaxCheatLen)

		if false { // verification sampling
			distStart := ignoreWalls[begin]
			distTarget := ignoreWalls[target]
			fmt.Printf("%d (S:%d, I: %d, T:%d, E:%d) %s\n", i, ds[i], distStart, distTarget, de[i], If(grid[i]==TooBig, "#", "."))
		}

		// now check candidates reachability
		for _,j := range(candidates) {
			// can't cheat from one spot to itself
			if i == j {
				continue
			}
			// grab the clean distance from i to j
			jumpDist := ignoreWalls[j]
			if 1 <= jumpDist && jumpDist <= MaxCheatLen {
				// filter out unfinishable spots, if any
				if de[j] == -1 {
					continue
				}
				
				// compute cheating score, and save the best for trivia
				cheatScore := ds[i] + jumpDist + de[j]
				if cheatScore < bestCheatingScore {
					bestCheatingScore = cheatScore
				}

				// count the cheat if exceeds/meets Threshold
				if baseDistance - cheatScore >= Threshold {
					numGoodCheats++
				}
				
				//manDist := Manhattan(sz, i, j)
				//if manDist != jumpDist {
				//	// this correctly never happens
				//	fmt.Printf("Man=%d, Jump=%d, from i=%d to j=%d\n", manDist, jumpDist, i, j)
				//}
			}
		}
	}
	fmt.Printf("Best cheat with MaxLen %d cheats reduces time from %d down to %d\n", MaxCheatLen, baseDistance, bestCheatingScore)

	return numGoodCheats
}

func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines,n, _ := ReadWholeFile(fname)
	answer := 0
	begin := -1
	target := -1

	// assign Threshold based on example or input
	Threshold = If(n>30, 100, 50)

	sz := len(lines[0]) // 141 input or 15 ex

	grid := make([]int, sz*sz)
	ds := make([]int, sz*sz) // Dijkstra distances from start
	de := make([]int, sz*sz) // Dijkstra distances from end

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

	if false { // verification sampling
		for i := range(len(ds)) {
			if i%200 != 16 { continue } // sample every 10th // 164 end, 16 beg
			fmt.Printf("%d: (S:%d, E:%d) given(%d,%d) %s\n", i, ds[i], de[i], begin, target, If(grid[i]==TooBig, "#", "."))
		}
		return 0
	}

	// initial approach
	if false {
		baseDistance := Dijkstra(grid, sz, begin, target)
		if baseDistance < 1 {
			// no cheats required
			return 0
		}

		// compute distances from start to every point, and from every point to end
		ds = BfsDijkstra(grid, sz, begin, false, sz*sz)
		de = BfsDijkstra(grid, sz, target, false, sz*sz)

		for i := range(sz*sz-1) {
			ix, iy := i/sz, i%sz
			// skip invalid cheating starting spots
			if ds[i] == -1 {
				continue
			}
			for j := i+1; j < sz*sz; j++ {
				jx, jy := j/sz, j%sz
				// skip invalid cheating ending spots
				if de[j] == -1 {
					continue
				}

				// respect MaxCheatLen filter
				manhattan := Abs(jx-ix) + Abs(jy-iy) // distance ignoring walls between i and j
				//if manhattan > MaxCheatLen {
				//	continue
				//}

				cheatDistance := ds[i] + de[j] + manhattan

				if cheatDistance  <= baseDistance-Threshold {
					//distBetween := Dijkstra(grid, sz, i, j)
					//if distBetween == -1 {
						answer++
					//}
				}
				_ = cheatDistance
			}
		}
	}

	// 285 ex @ 50+
	// 961364 input @ 100+
	answer = Cheatstra(grid, sz, begin, target)

	return answer

	// Submissions:
	// (1) 596906 = Too Low, Known < 1032776, > 596906
	// (2) 961363 = Too Low
	// (3) 961364 = Correct
	// my Go solution finds the answer in under a second
	// 961364 in 40 seconds with forked https://github.com/PeterCullenBurbery/advent-of-code-002/blob/main/jupyter-notebook-python/2024/2024_020/2024_020.ipynb
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
