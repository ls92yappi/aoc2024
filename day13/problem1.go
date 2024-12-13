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

var ex = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

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

// We want to find to smallest multiple of p.A xm that satisfies
// (xm*p.A.R+ym*p.B.R == X) && (xm*p.A.C+ym*p.B.C == Y)
func solve(p Prize) (xm, ym int, ok bool) {
	// Given: xm <= 100 from the problem, we quit after trying xm=100
	for xm = range(101) {
		xdiff := p.X - xm*p.A.R
		if xdiff%p.B.R != 0 {
			continue
		}
		ym = xdiff/p.B.R
		// candidate
		//fmt.Printf("X=%d, %d*%d = %d, xdiff=%d, ym=%d\n", p.X, p.A.R,xm,xm*p.A.R, xdiff, ym)

		if p.Y == xm*p.A.C + ym*p.B.C {
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
		xm,  ym,  ok := solve(p)

		// score winnable prizes
		if ok {
			fmt.Printf("xm:%d, ym:%d, ok:%v, eqn:%+v\n", xm, ym, ok, p)
			answer += 3*xm + 1*ym
		}
	}

	_ = N
	return answer

	// Submissions:
	// (1) 29517 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
