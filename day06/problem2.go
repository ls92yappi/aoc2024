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

const Debugging bool = false

func debug(msg string) {
	if !Debugging { return }
	fmt.Print(msg)
}

func debugf(format string, dets ...interface{}) {
	if !Debugging { return }
	fmt.Printf(format, dets...)
}

///////////////////////////////////////////////////////////

const Empt = "."
const Obst = "#"
const North = 0
const East  = 1
const South = 2
const West  = 3
const Wall = 2
const Visited = 1
// prechecked input size was 130x130 grid, example was 10x10
const Size = 130 //10 // 130 // hard-coded per run since using arrays rather than slices
const TooBig = Size*Size*Size
const TimedOut = (Size+4)*(Size+4)

type Guard struct {
	Row int
	Col int
	Face int
}

var Maze [Size][Size]int
var Solv [Size][Size]int
var Orig [Size][Size]int
var G Guard
var SaveG Guard
//guardStartRow = 87 of 130x130// 7 of 10x10

func travelGuard() int {
	final := -1
	// make the guard travel until  he leaves (or times out, lazy loop detection)
	for step := 0;step<TimedOut;step++ {
		if G.Row < 0 || G.Col < 0 || G.Row >= Size || G.Col >= Size {
			fmt.Printf("Jail Break Warning! This should not occur. Should be caught in the switch statement.")
			break
		}
		Solv[G.Row][G.Col] = Visited
		switch G.Face {
		case North:
			if G.Row == 0 { final = step; step = TooBig; break }
			if Maze[G.Row-1][G.Col] == Wall {
				G.Face = East
			} else {
				G.Row--
			}
		case East:
			if G.Col == Size-1 { final = step; step = TooBig; break }
			if Maze[G.Row][G.Col+1] == Wall {
				G.Face = South
			} else {
				G.Col++
			}
		case South:
			if G.Row == Size-1 { final = step; step = TooBig; break }
			if Maze[G.Row+1][G.Col] == Wall {
				G.Face = West
			} else {
				G.Row++
			}
		case West:
			if G.Col == 0 { final = step; step = TooBig; break }
			if Maze[G.Row][G.Col-1] == Wall {
				G.Face = North
			} else {
				G.Col--
			}
		}
		//fmt.Printf("Step %d Guard @ (%d,%d) facing %d\n", step, G.Row,G.Col,G.Face)
		_ = step
	}
	return final
}

func resetMazeFromOrig(n int) {
	for i := range(n) {
		for j := range(n) {
			Maze[i][j] = Orig[i][j]
			Solv[i][j] = Orig[i][j]
		}
	}
}

func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, n, _ := ReadWholeFile(fname)
	answer := 0

	for i := range(n) {
		line := lines[i]
		guardAt := strings.Index(line,"^")
		if (guardAt > -1) {
			G = Guard{i,guardAt,0}
			SaveG = G
			Maze[i][guardAt] = Visited
			Solv[i][guardAt] = Visited
			Orig[i][guardAt] = Visited
		}

		// construct the maze
		mazeRow := strings.Split(line,"")
		for j := range(n) {
			if mazeRow[j] == Obst {
				Maze[i][j] = Wall
				Orig[i][j] = Wall
			}
		}
	} // for i


	// testing -- known no loops pass took under a second to process travelGuard n^2 times
	// actual run should be less than about 3 times that long, so still fast enough
	for i := range(n) {
		for j := range(n) {
			// skip instances where a wall initially started
			if Orig[i][j] == Wall {
				continue
			}
			// skip guard starting spot
			if i == SaveG.Row && j == SaveG.Col {
				continue
			}

			// reset the Guard and the maze/solution
			G = SaveG
			resetMazeFromOrig(n)

			// try to trap the guard in a loop by imposing a new Wall at this spot
			Maze[i][j] = Wall
			Solv[i][j] = Wall
			//fmt.Printf("Trying Wall @ (%d,%d)\n", i, j)

			// Normal guard running
			outStep := travelGuard()
			if outStep != -1 {
				//fmt.Printf("Guard Exit on step %d\n", outStep)
			} else {
				fmt.Printf("Loop at (%d,%d)\n", i, j)
				answer++
			}
		}
	}
	// ex G={6,4,0} Bads={6,3; 7,6; 7,7; 8,1; 8,3; 9,7}
	// input G={86,45,0}

	return answer

	// Submissions:
	// (1) 1767 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
