package main

// See README.md for problem description

import (
	"fmt"
	//"strconv"
	//"strings"
	. "github.com/ls92yappi/aoc"
)

///////////////////////////////////////////////////////////

var G [][]int

type Pos struct {
	R int
	C int
}

var T []Pos // trail heads
var S []Pos // summits
var R []int // routes per trail head

var N int

func ShowAttempt(t, g Pos, i int, h int) string {
	s := fmt.Sprintf("%d_%d:", i, h)
	for range(h) {
		s += "."
	}
	s += fmt.Sprintf(" %v", t)
	if h==9 {
		s += fmt.Sprintf(" =? %v", g)
	}
	s += "\n"
	return s
}

// route from trailhead t with index i at height h to goal g at height 9
func routeFrom(t, g Pos, i int, h int) bool {
	//fmt.Print(ShowAttempt(t,g,i,h))
	if (h == 9) {
		//fmt.Printf("Summit Check: t=%v g=%v h=%d\n", t, g, h)
		// found the goal
		return t.R==g.R && t.C==g.C
	}

	var n, s, e, w *Pos
	if t.R > 0 { n = &Pos{t.R-1,t.C} }
	if t.R < N { s = &Pos{t.R+1,t.C} }
	if t.C > 0 { w = &Pos{t.R,t.C-1} }
	if t.C < N { e = &Pos{t.R,t.C+1} }

	up := h+1
	nok := n!=nil && (G[n.R][n.C] == up)
	sok := s!=nil && (G[s.R][s.C] == up)
	eok := e!=nil && (G[e.R][e.C] == up)
	wok := w!=nil && (G[w.R][w.C] == up)

	ok := ((nok && routeFrom(*n, g, i, up)) ||
	       (sok && routeFrom(*s, g, i, up)) ||
	       (eok && routeFrom(*e, g, i, up)) ||
	       (wok && routeFrom(*w, g, i, up)))

	return ok
}

func climbEmAll() {
	for i := range(len(T)) {
		t := T[i]
		tx := t.R + t.C
		for _,s := range(S) {
			// quick filter reachability by distance
			sx := s.R + s.C
			if Abs(sx-tx) > 9 {
				continue
			}
			//fmt.Printf("BEGIN %d %v -> %v\n", i, t, s)
			// check for a route between t and s
			if routeFrom(t,s,i,0) {
				//fmt.Printf("SUCCESS %d %v -> %v\n", i, t, s)
				R[i]++
			}
		}
	}
	return
}

// ex 8x8 grid
// in 45x45 grid
func processInput(fname string) int {
	lines, n, _ := ReadWholeFile(fname)
	n--
	N = n-1 // final index of row/col

	answer := 0
	G = make([][]int, 0)
	T = make([]Pos, 0)
	S = make([]Pos, 0)
	R = make([]int, 0)

	// construct grid G
	for i := range(n) {
		line := lines[i]
		if len(line) < 2 {
			continue
		}
		row, _, _ := IntSlice(line, "")
		G = append(G, row)
		// find trailheads and summits
		for j := range(n) {
			if G[i][j] == 0 {
				t := Pos{i,j}
				T = append(T, t)
				R = append(R, 0)
			}
			if G[i][j] == 9 {
				s := Pos{i,j}
				S = append(S, s)
			}
		}
	}

	// climb every mountain per trailhead
	climbEmAll()

	// score trailheads
	for i := range(len(T)) {
		answer += R[i]
	}

	return answer

	// Submissions:
	// 35802=221T*162S, 9745 within distance, 1647 total routes, 566 valid
	// (1) 566 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
