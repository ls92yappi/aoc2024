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

const box = 'O'
const robot = '@'
const wall = '#'
const empty = '.'

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
		for j := range(Size) {
			s += string(w[i][j])
		}
		s += "\n"
	}
	return s
}

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
				ok = true
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
				ok = true
				break
			}
		}
	case moveL:
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


func doMove(w [][]rune, rx,ry int, upto int, move byte) (int,int,[][]rune) {
	// move the boxes and robot upto the found empty
	switch move {
	case moveU:
		for i := upto; i < rx; i++ {
			w[i][ry] = w[i+1][ry]
		}
	case moveD:
		for i := upto; i > rx; i-- {
			w[i][ry] = w[i-1][ry]
		}
	case moveL:
		for i := upto; i < ry; i++ {
			w[rx][i] = w[rx][i+1]
		}
	case moveR:
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
		//row := make([]byte, Size)
		//slChars := strings.Split(line, "")
		row := []rune(line)
		warehouse = append(warehouse, row)
	}

	// init the moves
	moves := []byte(strings.Join(lines[Size+1:],""))

	fmt.Printf("%s\n", display(warehouse))

	after := performMoves(warehouse, moves)
	fmt.Printf("\n%s\n", display(after))

	for i := range(Size) {
		for j := range(Size) {
			if after[i][j] == box {
				gps := 100*i + j
				answer += gps
			}
		}
	}
	//fmt.Printf("%+v\n\nMoves: %+v\n", warehouse,moves)
	//fmt.Printf("%+v\n\n", after)

	return answer

	// smallSum = 2028
	// exampleS = 10092
	// Submissions:
	// (1) 1441031 = Correct
}

func main() {
	//AllRules = make([][]int, 0)
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
