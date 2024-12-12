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
	B byte
	R int
	C int
}

type Region struct {
	B    byte
	R    int // root Row
	C    int // root Col
	A    int // area
	Locs []*Pos
}

var G [][]*Pos
var R []Region
var N int

var V []bool // number of visits to this location

// If returns vtrue if cond is true, vfalse otherwise.
//
// Useful to avoid an if statement when initializing variables, for example:
//
//	min := If(i > 0, i, 0)
func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

// largest area = 335
func area(r Region) int {
	return r.A
}

func perim(r Region) int {
	p := 0
	for n, l := range r.Locs {
		i := l.R
		j := l.C

		// check edges first
		p += If(i==0, 1, 0) // N
		p += If(j==0, 1, 0) // W
		p += If(i==N-1, 1, 0) // S
		p += If(j==N-1, 1, 0) // E

		// then check other interior side if local edge
		p += If(i>0 && !within(r, i-1, j), 1, 0) // S
		p += If(j>0 && !within(r, i, j-1), 1, 0) // W
		p += If(i<N-1 && !within(r, i+1, j), 1, 0) // N
		p += If(j<N-1 && !within(r, i, j+1), 1, 0) // E
		_ = n
	}
	return p
}

func within(r Region, i, j int) bool {
	for _, l := range r.Locs {
		if i == l.R && j == l.C {
			return true
		}
	}
	return false
}

func grow(r *Region, i, j int) int {
	// don't re-visit any grid spot
	//if r.B == 'I' {
	//	fmt.Printf("Region(%d,%d) @ (%d,%d) size %d\n", r.R, r.C, i, j, r.A)
	//}
	if V[i*N+j] {
		return 0
	}

	V[i*N+j] = true
	// verify not already part of a Region
	if G[i][j] == nil {
		return 0
	}

	// add this spot to the Region if it matches
	b := G[i][j].B
	if r.B == b {
		pos := &Pos{b, i, j}
		r.Locs = append(r.Locs, pos)
		G[i][j] = nil
		V[i*N+j] = true
		r.A++

		//try to grow in all 4 directions
		if i > 0 { grow(r, i-1, j) } // North
		if j > 0 { grow(r, i, j-1) } // West
		if i < N-1 { grow(r, i+1, j) } // South
		if j < N-1 { grow(r, i, j+1) } // East
	}

	return r.A
}


// mini is 4x4 grid
// example is 10x10 grid
// input is 140x140 grid
func processInput(fname string) int {
	lines, n, _ := ReadWholeFile(fname)
	n--
	N = n
	answer := 0

	G = make([][]*Pos, n)
	R = make([]Region, 0)
	// Construct the Grid G
	for i := range n {
		line := lines[i]
		if len(line) < 2 {
			continue
		}
		row := make([]*Pos, n)
		for j := range n {
			pos := &Pos{line[j], i, j}
			row[j] = pos
		}
		G[i] = row
	}

	// Construct Regions
	for i := range n {
		for j := range n {
			if G[i][j] != nil {
				V = make([]bool, N*N) // make and clear out V for visit checks
				b := G[i][j].B
				reg := Region{b, i, j, 0, make([]*Pos, 0)}
				// grow the region, marking spots in G as nil in the process
				grow(&reg, i, j)
				// and add it to the list of Regions
				R = append(R, reg)
			}
		}
	}

	// Compute prices (area*perim)
	for i, r := range R {
		a := area(r)
		p := perim(r)
		price := a * p
		answer += price
		// Given:
		// A(mini)*P(mini) = 4*10 + 4*8 + 4*10 + 1*4 + 3*8 = 24 = 140
		// A(embed)*P(embed) = 21*36=756 + 4*(1*4=4) = 772
		// A(ex)*P(ex) = R12*18 I4*8 C14*28 F10*18 V13*20 J11*20 C1*4 E13*18 I14*22 M5*12 S3*8 = 1930
		fmt.Printf("Region %d(%s) = %d x %d = %d\n", i, string(r.B), a, p, price)
		_ = i
	}

	return answer

	// Submissions:
	// (1) 1488414 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
