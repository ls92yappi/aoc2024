package main

// See README.md for problem description

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	. "github.com/ls92yappi/aoc"
)

///////////////////////////////////////////////////////////

// Linux ANSI graphics
var CLS = "\x1b[H\x1b[2J"
const Reset = "\x1b[0m"
//const Show = "\x1b[97;42m█\x1b[0m"
const Show = "\x1b[97;42m \x1b[0m"
const Hide = "\x1b[97;40m.\x1b[0m"
//const Hide = "\x1b[97;40m \x1b[0m"

func ReadOneChar() int {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	// restore the echoing state when exiting
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	//fmt.Println("Press 'q' to quit.")
	//fmt.Print("> ")
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		//fmt.Println("I got the byte", b, "("+string(b)+")")
		if b[0] == byte('q')  || b[0] == byte('Q') {
			return 1
		}
		if b[0] == byte('n')  || b[0] == byte('N') {
			return 2
		}
		if b[0] == byte('p')  || b[0] == byte('P') {
			return 3
		}
		if b[0] == byte('t')  || b[0] == byte('T') {
			return 4
		}
	}
	//exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	return 0
}


///////////////////////////////////////////////////////////

type Robot struct {
	I int // index
	R int // loc
	C int
	X int // veloc
	Y int
}

// input grid is 101 Wide x 103 Tall
const GX int = 101
const GY int = 103
//const GX int = 11
//const GY int = 7

var Tags []int

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

func TakeTurns(t int) {
	for i,r := range(Army) {
		ex := ((r.R + t*r.X) % GX + GX) % GX
		ey := ((r.C + t*r.Y) % GY + GY) % GY
		r.R = ex
		r.C = ey
		Army[i] = r
	}
}

// AssumeBinary() calls TakeTurns() until it detects a binary pattern
func AssumeBinary() int {
	for t := range(GX*GY) { // 10403 = 101*103 = gridsize
		ok := true
		for row := range(GY) {
			for spot := range(GX) {
				num := 0
				for _,r := range(Army) {
					if r.R == row && r.C == spot {
						num++
					}
				}
				if num > 1 {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			return t
		}
		TakeTurns(1)
	}
	return -1
}

// find n such that n%d1=r1 && n%d2=r2
// solves r1 base repeats every d1 passes and r2 base repeats every d2 passes
func ModSolver(d1, r1, d2, r2 int) int {
	// verify we have valid modular arithmetic numbers
	if d1 <= 0 || d2 <= 0 || r1 < 0 || r2 < 0 || r1 >= d1 || r2 >= d2 {
		return -1
	}

	// limited trial and error
	guess := r1
	for range(d2) {
		guess += d1 // maintains guess%d1=r1
		if guess%d2 == r2 {
			return guess
		}
	}

	// This should never occur given the top filter above
	return -1
}

func DisplayPic() {
	for row := range(GY) {
		for spot := range(GX) {
			display := Hide
			num := 0
			for _,r := range(Army) {
				if r.R == row && r.C == spot {
					display = Show
					num++
					//break
				}
			}
			if num > 0 {
				//display = fmt.Sprintf("\x1b[97;42m%d\x1b[0m", num)
				//display = fmt.Sprintf("\x1b[97;40m%d\x1b[0m", num)
			}

			fmt.Print(display)
		}
		fmt.Print(Reset)
		fmt.Println()
	}
	fmt.Print(Reset)
}

var Army []Robot

// example is 12 lines
// input is 500 lines
func processInput(fname string) int {
	lines, n, _ := ReadWholeFile(fname)
	n--
	answer := 0

	//quads := make([]int, 5) // quadrant 0 is the center cross, then 4 quadrants
	Army = make([]Robot, 0)
	Tags = make([]int, 0)

	// Read in the army of robots
	for i := range n {
		line := lines[i]
		if len(line) < 2 {
			continue
		}
		robot := ReadRobot(line, i)
		Army = append(Army, robot)
	}

	// took about 5 seconds to find the solution
	// Analytical assuming solution is all 1s and 0s
	if true {
		answer = AssumeBinary()
		DisplayPic()
	}

	//guess := ModSolver(GX, 10, GY, 70) // 10 and 70 from observations using visual approach
	//fmt.Printf("ModSolver guess = %d\n", guess)

	// Interactive visual comparison approach
	if false {
		t := 0
		inc := 1
		for {
			fmt.Printf("%sTime %d\n====================\n\n", CLS, t)
			DisplayPic()
			TakeTurns(inc) // early by 1 ... oops -- should be After DisplayPic()

			// @t=10 -- mostly horizontal @t=70 -- mostly horizontal
			// horizontal'ish recurs every t=101 , vertical'ish recurs every t=103
			// guesstimating from this pattern, we note that
			// 7383 % 101 = 10 and 7383 % 103 = 70

			code := ReadOneChar()
			if code == 1 {
				fmt.Println("Quit")
				break
			}
			if code == 2 {
				fmt.Println("Next")
				time.Sleep(1 * time.Millisecond)
				inc = 1
				t++
			}
			if code == 3 {
				fmt.Println("Prev")
				time.Sleep(1 * time.Millisecond)
				inc = -1
				t--
			}
			if code == 4 {
				fmt.Println("Tagged")
				time.Sleep(1 * time.Millisecond)
				inc = 0
				Tags = append(Tags, t)
			}
			if t > 10403 { // 10403 = 101*103 = gridsize
				fmt.Println("Fail at 10403!")
				break
			}
		}

		fmt.Println("\n--------------\n")
		for k,v := range(Tags) {
			fmt.Printf("%d\t",v)
			_ = k
		}
		fmt.Println()
	} // if visual comparison
	
	return answer

	// Submissions:
	// (1) 69 = Wrong -- requires visual recognition without a visual example to compare against, nor a useful example
	// (2) 7382 = Too Low ... off by 1 DisplayPic() should be before TakeTurns(inc)
	// (3) 7383 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
