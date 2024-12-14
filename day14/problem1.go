package main

// See README.md for problem description

import (
	"fmt"
	"strconv"
	"strings"
	. "github.com/ls92yappi/aoc"
)

///////////////////////////////////////////////////////////

type Robot struct {
	I int // index
	R int // loc
	C int
	X int // veloc
	Y int
}

// examp grid is 11 Wide x 7 Tall
// input grid is 101 Wide x 103 Tall
//var GX int = 11
//var GY int = 7
var GX int = 101
var GY int = 103

var quads []int

func ReadRobot(s string, idx int) Robot {
	// p=0,4 v=3,-3
	pos, vel, _ := strings.Cut(s, " ")
	_, p, _ := strings.Cut(pos, "=")
	_, v, _ := strings.Cut(vel, "=")
	rs, cs, _ := strings.Cut(p, ",")
	xs, ys, _ := strings.Cut(v, ",")
	r, _ := strconv.Atoi(rs)
	c, _ := strconv.Atoi(cs)
	x, _ := strconv.Atoi(xs)
	y, _ := strconv.Atoi(ys)
	robot := Robot{idx,r,c,x,y}
	return robot
}

func QuadrantAfterTime(r Robot, t int) int {
	ex := ((r.R + t*r.X) % GX + GX) % GX
	ey := ((r.C + t*r.Y) % GY + GY) % GY
	MX := GX/2
	MY := GY/2
	q := 0
	switch {
	case ex == MX || ey == MY:
		q = 0
	case ex > MX && ey < MY:
		q = 1
	case ex < MX && ey < MY:
		q = 2
	case ex < MX && ey > MY:
		q = 3
	case ex > MX && ey > MY:
		q = 4
	default:
		fmt.Printf("Should never happen! {%+v}@%d \n", r, t)
		return 0
	}
	//fmt.Printf("Q%d {%+v}@%d ex:%d ey:%d\n", q, r, t, ex, ey)
	fmt.Printf("Q%d ex:%d ey:%d\n", q, ex, ey)
	return q
}

// example is 12 lines
// input is 500 lines
func processInput(fname string) int {
	lines, n, _ := ReadWholeFile(fname)
	n--
	answer := 0

	quads := make([]int, 5) // quadrant 0 is the center cross, then 4 quadrants

	// where after 100 seconds
	for i := range n {
		line := lines[i]
		if len(line) < 2 {
			continue
		}
		robot := ReadRobot(line, i)

		quad := QuadrantAfterTime(robot, 100)
		quads[quad]++
	}

	fmt.Printf("X:%d\t1:%d\t2:%d\t3:%d\t4:%d\n", quads[0], quads[1], quads[2], quads[3], quads[4])
	answer = quads[1] * quads[2] * quads[3] * quads[4]

	return answer

	// Submissions:
	// (1) 229069152 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
