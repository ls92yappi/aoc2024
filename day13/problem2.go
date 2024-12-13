package main

// See README.md for problem description

import (
	"fmt"
	"regexp"
	//"sort"
	"strconv"
	. "github.com/ls92yappi/aoc"
)

///////////////////////////////////////////////////////////

type Button struct {
	L string
	R int
	C int
}

type Prize struct {
	A Button
	B Button
	X int
	Y int
}

var P []Prize
const Offset = 10000000000000

// If returns vtrue if cond is true, vfalse otherwise.
//
// Useful to avoid an if statement when initializing variables, for example:
//
//	min := If(i > 0, i, 0)
// If -- the best Ternary operator ever!!! from https://github.com/icza/gox/blob/main/gox/gox.go
func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

// This approach presumes there exists only one such solution of
// (xm*p.A.R+ym*p.B.R == X+Offset) && (xm*p.A.C+ym*p.B.C == Y+Offset)
func solveWithOffset(p Prize) (xm, ym int, ok bool) {
	w := p.X + Offset
	z := p.Y + Offset
	a := p.A.R
	b := p.B.R
	c := p.A.C
	d := p.B.C

	// solve 2 equations with 2 unknowns:
	// a*xm + b*ym = w AND c*xm + d*ym = z
	//
	// which simplifies in the x down to:
	// xm = (w-b*ym)/a AND xm = (z-d*ym)/c
	//
	// solving the simplified xm equations for ym yields:
	// ym = (c*w-a*z)/(b*c-a*d)
	//
	// So it is just a matter of plug'n'chug'ing the solutions, and 
	// verifying that all relevant xm,ym values are indeed integers.
	// Verifying integers just means ensuring each numer/denom pair
	// in the 3 simplified equations mods to 0.

	denom := b*c - a*d
	numer := c*w - a*z

	if numer%denom == 0 {
		// there exists a possible solution for ym
		ym = numer/denom

		xtopA := w - b*ym
		xtopC := z - d*ym
		if xtopA%a == 0 && xtopC%c == 0 {
			// there exists a solution for xm
			xm = xtopA/a
			ok = true
			return
		}
	}

	return
}

// example is 4 eqns
// input is 320 eqns
func processInput(fname string) int {
	lines, n, _ := ReadWholeFile(fname)
	n++
	N := n/4
	answer := 0

	reBtn := regexp.MustCompile(`Button (.): X\+(\d+), Y\+(\d+)`)
	rePrz := regexp.MustCompile(`Prize: X\=(\d+), Y\=(\d+)`)

	//P = make([]Prize, N)
	// Construct the Prizes list P
	//for i := range 1 {
	for i := range N {
		lineA := lines[i*4]
		lineB := lines[i*4+1]
		lineP := lines[i*4+2]

		// construct the equation
		sBtnA := reBtn.FindSubmatch([]byte(lineA))
		l := string(sBtnA[1])
		r, _ := strconv.Atoi(string(sBtnA[2]))
		c, _ := strconv.Atoi(string(sBtnA[3]))
		btnA := Button{l,r,c}
		//fmt.Printf("ButtonA: %+v\n", btnA)

		sBtnB := reBtn.FindSubmatch([]byte(lineB))
		l = string(sBtnB[1])
		r, _ = strconv.Atoi(string(sBtnB[2]))
		c, _ = strconv.Atoi(string(sBtnB[3]))
		btnB := Button{l,r,c}
		//fmt.Printf("ButtonB: %+v\n", btnB)

		sPrz := rePrz.FindSubmatch([]byte(lineP))
		x, _ := strconv.Atoi(string(sPrz[1]))
		y, _ := strconv.Atoi(string(sPrz[2]))

		p := Prize{btnA, btnB, x, y}

		//P[i] = p

		// compute solution and whether it is winnable
		xm,  ym,  ok := solveWithOffset(p)

		// 1000 seconds per eqn*320 equations -> 320,000s to run = 5333m = 88h = 3.7days
		// original method with correctly capped xm would work, but would take 3.7 days to solve
		// score winnable prizes
		if ok {
			fmt.Printf("Solution found for Claw %d\n", i+1)
			//fmt.Printf("xm:%d, ym:%d, ok:%v, eqn:%+v\n", xm, ym, ok, p)
			answer += 3*xm + 1*ym
		}
	}

	_ = N
	return answer

	// Submissions:
	// (1) 103570327981381 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
