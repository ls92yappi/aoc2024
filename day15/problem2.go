package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	"strings"
	. "github.com/ls92yappi/aoc" // problems with go proxy not being even remotely up-to-date
)

///////////////////////////////////////////////////////////


///////////////////////////////////////////////////////////

// doublewide boxes modification
const box = 'O'
const robot = '@'
const wall = '#'
const empty = '.'
const boxL = '['
const boxR = ']'

const moveU = '^'
const moveD = 'v'
const moveL = '<'
const moveR = '>'

var warehouse [][]rune
var Size = 0

func findRobot(w [][]rune) (int,int) {
	for i := range(Size) {
		c := strings.IndexByte(string(w[i]),robot)
		if c > -1 {
			r := i
			return r,c
		}
	}
	return 0,0
}

func display(w [][]rune) string {
	s := ""
	for i := range(Size) {
		for j := range(Size*2) {
			s += string(w[i][j])
		}
		s += "\n"
	}
	return s
}

// dir = -1 for up and 1 for down
func noVerticalBoxCollision( w [][]rune, rx,ry int, upto int, dir int) bool {
	immedBoxHalf := w[rx+dir][ry]
	// known that immedBoxHalf column has no blockers

	// TODO: code this for possibly incomplete side growth for partial columns


	if immedBoxHalf == boxL {
		// check boxR for conflict
		// and any boxes above
	} else {
		// check boxL for conflict
		// and any boxes above
	}

	// verify that no side boxes cause issues in other columns

	return true
}


// need to modify for double wide
func canMove(w [][]rune, rx,ry int, move byte) (bool,int) {
	ok := false
	upto := -1
	// verify there is a dot between the robot and the nearest wall in that direction
	switch move {
	case moveU:
		for i := rx-1; i > 0; i-- {
			if w[i][ry] == wall {
				break
			}
			if w[i][ry] == empty {
				upto = i
				// now need to check rest of this box plus other pushed boxes
				ok = noVerticalBoxCollision(w, rx,ry, upto, -1)
				break
			}
		}
	case moveD:
		for i := rx+1; i <Size; i++ {
			if w[i][ry] == wall {
				break
			}
			if w[i][ry] == empty {
				upto = i
				// now need to check rest of this box plus other pushed boxes
				ok = noVerticalBoxCollision(w, rx,ry, upto, 1)
				break
			}
		}
	case moveL:
		// left moves are unchanged
		for i := ry-1; i > 0; i-- {
			if w[rx][i] == wall {
				break
			}
			if w[rx][i] == empty {
				upto = i
				ok = true
				break
			}
		}
	case moveR:
		// right moves are unchanged
		for i := ry+1; i <Size; i++ {
			if w[rx][i] == wall {
				break
			}
			if w[rx][i] == empty {
				upto = i
				ok = true
				break
			}
		}
	default:
		fmt.Printf("Should never get here in canMove()\n")
		return false,upto
	}
	return ok,upto
}


// need to modify for double wide
func doMove(w [][]rune, rx,ry int, upto int, move byte) (int,int,[][]rune) {
	// move the boxes and robot upto the found empty
	switch move {
	case moveU:
		// TODO: update this for doublewide
		for i := upto; i < rx; i++ {
			w[i][ry] = w[i+1][ry]
		}
	case moveD:
		// TODO: update this for doublewide
		for i := upto; i > rx; i-- {
			w[i][ry] = w[i-1][ry]
		}
	case moveL:
		// left moves are unchanged
		for i := upto; i < ry; i++ {
			w[rx][i] = w[rx][i+1]
		}
	case moveR:
		// right moves are unchanged
		for i := upto; i > ry; i-- {
			w[rx][i] = w[rx][i-1]
		}
	default:
	}
	// oldl robot spot is now empty
	w[rx][ry] = empty
	//update robot coords
	rx,ry = findRobot(w) // feeling lazy
	return rx, ry, w
}

func performMoves(w [][]rune, moves []byte) [][]rune {
	//fmt.Printf("OkMoves: ")
	rx,ry := findRobot(w)
	for _,m := range(moves) {
		if ok, upto := canMove(w, rx, ry, m); ok {
			//fmt.Printf(string(m))
			rx, ry, w = doMove(w, rx, ry, upto, m)
			//fmt.Printf("%s\n%s\n", string(m), display(w))
		}
	}
	//fmt.Printf("\n")
	return w
}

func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, _, _ := ReadWholeFile(fname)
	answer := 0
	Size = len(lines[0])

	// init the warehouse
	warehouse = make([][]rune, 0)
	for i := range(Size) {
		line := lines[i]

		// new rows are doublewide
		row := make([]rune, 0)
		for j := range(Size) {
			rn := line[j]
			switch rn {
			case box:
				row = append(row, boxL)
				row = append(row, boxR)
			case robot:
				row = append(row, robot)
				row = append(row, empty)
			case wall:
				row = append(row, wall)
				row = append(row, wall)
			case empty:
				row = append(row, empty)
				row = append(row, empty)
			default:
				fmt.Printf("Should never get here in construct warehouse.\n")
			}
		}
		warehouse = append(warehouse, row)
	}

	// init the moves
	moves := []byte(strings.Join(lines[Size+1:],""))

	fmt.Printf("%s\n", display(warehouse))

	after := performMoves(warehouse, moves)
	fmt.Printf("\n%s\n", display(after))

	for i := range(Size) {
		for j := range(Size) {
			if after[i][j] == boxL {
				gps := 100*i + j
				answer += gps
			}
		}
	}
	//fmt.Printf("%+v\n\nMoves: %+v\n", warehouse,moves)
	//fmt.Printf("%+v\n\n", after)

	return answer

	// exampleS = 9021
	// Submissions:
	// (1) 1441031 = Correct
}

func main() {
	//AllRules = make([][]int, 0)
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
