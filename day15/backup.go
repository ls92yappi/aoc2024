package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	. "github.com/ls92yappi/aoc" // problems with go proxy not being even remotely up-to-date
)

///////////////////////////////////////////////////////////


///////////////////////////////////////////////////////////

const Reset = "\x1b[0m"
const Show = "\x1b[97;42m@\x1b[0m"


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

type boxRow struct {
	level int
	locs []int
}
// a boxSet is the sorted slice of slices of box parts to move
// in vertical direction dir prior to moving the robot
// it is ignored for horizontal moves
type boxSet []boxRow

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
			if w[i][j] == robot {
				s += Show + Reset
			} else {
				s += string(w[i][j])
			}
		}
		s += "\n"
	}
	return s
}

// dir = -1 for up and 1 for down
func noVerticalBoxCollision( w [][]rune, rx,ry int, upto int, dir int) (bool, boxSet) {
	bs := make(boxSet, 0)
	cur := boxRow{rx+dir, make([]int, 0)}

	// check immediate spot above robot
	// if empty, simple move
	immedAbove := w[rx+dir][ry]
	if immedAbove == empty {
		return true, bs
	}

	// verify that no side boxes cause issues in other columns and return boxSet
	// of boxes that will ultimately be moved in order of moving them

	// must be either boxL or boxR
	// construct cur = initial boxRow directly from robot
	//cur.depth = rx+dir
	cur.locs = append(cur.locs, ry)

	if immedAbove == boxL {
		cur.locs = append(cur.locs, ry+1)
	} else {
		// immedAbove must be boxR
		cur.locs = append(cur.locs, ry-1)
	}

	// for loop terminates when either entire row safe to move up OR
	// when any wall is found which blocks the move
	// guaranteed to terminate since entire warehouse is surrounded by walls
	for {
		numCurEmpty := 0 // number of spots above cur.locs that are empty 
		next := boxRow{cur.level+dir, make([]int, 0)}

		// sort and remove duplicates
		sort.Ints(cur.locs)
		cur.locs = slices.Compact(cur.locs)

		// verify cur boxRow to see if it resolves, constructing the
		// next = contingent boxRow while checking
		for _,pos := range(cur.locs) {
			aboveSpot := w[cur.level][pos]
			if aboveSpot == wall {
				// resolves as BAD when any wall found, it blocks the entire move
				return false, bs
			}
			if aboveSpot == empty {
				numCurEmpty++
				continue
			}
			if aboveSpot == boxL {
				next.locs = append(next.locs, pos)
				next.locs = append(next.locs, pos+1)
			} else {
				// aboveSpot must be boxR
				next.locs = append(next.locs, pos)
				next.locs = append(next.locs, pos-1)
			}
		}

		// PULLING VERSION - this pulls an extra row
		if true {
			newBs := make(boxSet, 0)
			newBs = append(newBs, cur) // next instead of cur???
			newBs = append(newBs, bs...)
			bs = newBs // yes, we prepend to the boxSet
		}

		// THIS VERSION DOES NOT PULL on test2.txt
		// but does not end up with the right answer, and 700 moves without
		// a worked out example to compare against is waaay too many
		// in order to be able to tell what is wrong
		if false {
			if cur.level != rx+dir {
				// prepend cur boxRow to the boxSet for correct processing order
				newBs := make(boxSet, 0)
				newBs = append(newBs, cur) // next instead of cur???
				newBs = append(newBs, bs...)
				bs = newBs // yes, we prepend to the boxSet
			}
		}

		if numCurEmpty == len(cur.locs) {
			// resolves as GOOD when ALL cur elements covered by empty
			return true, bs
		}

		// continue constructing boxSet by checking next boxRow otherwise
		cur.locs = cur.locs[:]
		cur = next
	}

	fmt.Printf("Should never get here in noVerticalBoxCollision()!\n")
	return true, bs
}


// need to modify for double wide
func canMove(w [][]rune, rx,ry int, move byte) (bool,int,boxSet) {
	ok := false
	upto := -1
	bs := make(boxSet, 0)

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
				ok, bs = noVerticalBoxCollision(w, rx,ry, upto, -1)
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
				ok, bs = noVerticalBoxCollision(w, rx,ry, upto, 1)
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
		return false,upto, bs
	}
	return ok,upto,bs
}


// need to modify for double wide
func doMove(w [][]rune, rx,ry int, upto int, bs boxSet, move byte) (int,int,[][]rune) {
	// move the boxes and robot upto the found empty
	switch move {
	case moveU:
		dir := -1
		//fmt.Printf("UP: boxSet=%+v robot@(%d,%d)\n", bs, rx, ry)
		// move the boxSet first, already sorted and simplified
		for _, br := range(bs) {
			depthTo := br.level
			depthFrom := br.level - dir
			//if len(br.locs) > 2 {
			//	fmt.Printf("UP: boxSet=%+v robot@(%d,%d)\n", bs, rx, ry)
			//	fmt.Printf("\tUP: T=%d F=%d width=%d\n", depthTo, depthFrom, len(br.locs))
			//	fmt.Printf("\nBEFORE:\n%s\n", display(w))
			//}
			for _,spot := range(br.locs) {
				//fmt.Printf("\tw[%d][%d] = w[%d][%d]\n", spot, depthTo, spot, depthFrom)
				w[depthTo][spot] = w[depthFrom][spot]
				// and backfill dots
				if (spot != ry) {
					w[depthFrom][spot] = empty
				}
			}
			//if len(br.locs) > 2 {
			//	fmt.Printf("\nAFTER:\n%s\n", display(w))
			//}
		}
		// then move the robot
		//fmt.Printf("UP: robot To (%d,%d) From (%d,%d)\n", rx+dir,ry,rx,ry)
		w[rx+dir][ry] = w[rx][ry]
	case moveD:
		dir := 1
		//fmt.Printf("DN: boxSet=%+v robot@(%d,%d)\n", bs, rx, ry)
		// move the boxSet first, already sorted and simplified
		for _, br := range(bs) {
			depthTo := br.level
			depthFrom := br.level - dir
			//fmt.Printf("\tDN: T=%d F=%d width=%d\n", depthTo, depthFrom, len(br.locs))
			for _,spot := range(br.locs) {
				w[depthTo][spot] = w[depthFrom][spot]
				// and backfill dots
				if (spot != ry) {
					w[depthFrom][spot] = empty
				}
			}
		}
		// then move the robot
		//fmt.Printf("DN: robot To (%d,%d) From (%d,%d)\n", rx+dir,ry,rx,ry)
		w[rx+dir][ry] = w[rx][ry]
	case moveL:
		// left move all boxes then robot
		for i := upto; i < ry; i++ {
			w[rx][i] = w[rx][i+1]
		}
	case moveR:
		// right move all boxes then robot
		for i := upto; i > ry; i-- {
			w[rx][i] = w[rx][i-1]
		}
	default:
	}
	// old robot spot is now empty
	w[rx][ry] = empty
	//update robot coords
	rx,ry = findRobot(w) // feeling lazy
	return rx, ry, w
}

func performMoves(w [][]rune, moves []byte) [][]rune {
	//fmt.Printf("OkMoves: ")
	rx,ry := findRobot(w)
	for mnum,m := range(moves) {
		if ok, upto, bs := canMove(w, rx, ry, m); ok {
			//fmt.Printf(string(m))
			rx, ry, w = doMove(w, rx, ry, upto, bs, m)
			if breakDetection(w, mnum) {
				fmt.Printf("%s\n%s\n", string(m), display(w))
				return w
			}
			//fmt.Printf("%s\n%s\n", string(m), display(w))
		}
		_ = mnum
	}
	//fmt.Printf("\n")
	return w
}

func breakDetection(w [][]rune, mv int) bool {
	bad := false
	for i := range(Size) {
		for j := range(Size*2) {
			if w[i][j] == boxL {
				if w[i][j+1] != boxR {
					fmt.Printf("Break found on move# %d! Left missing Right.\n", mv)
					return true
				}
			} else if w[i][j] == boxR {
				if w[i][j-1] != boxL {
					fmt.Printf("Break found on move# %d! Right missing Left.\n", mv)
					return true
				}
			}
		}
	}
	return bad
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
	// (1) 671557 = Too Low -- this was with the "DOES NOT PULL" version of noVerticalBoxCollision()
	// (x) 666422 -- this was with the PULLING version that does break
}

func main() {
	//AllRules = make([][]int, 0)
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
