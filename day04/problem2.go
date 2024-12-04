package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	//"regexp"
	//"strconv"
	//"strings"
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

func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, n, _ := ReadWholeFile(fname)

	xmas_rmm := 0
	xmas_rss := 0
	xmas_dmm := 0
	xmas_dss := 0

	for i := range(n-2) {
		line := lines[i]
		for j := range(n-2) {
			ul := line[j]
			ur := line[j+2]
			cc := lines[i+1][j+1]
			dl := lines[i+2][j]
			dr := lines[i+2][j+2]
			//fmt.Printf("%d %d %q\t",i, j, ch)
			if cc != 'A' {
				continue
			}
			// M.M.A.S.S check
			if (ul=='M' && ur=='M' && dl=='S' && dr=='S') {
				xmas_rmm++
			}
			// S.S.A.M.M check
			if (ul=='S' && ur=='S' && dl=='M' && dr=='M') {
				xmas_rss++
			}
			// M.S.A.M.S check
			if (ul=='M' && ur=='S' && dl=='M' && dr=='S') {
				xmas_dmm++
			}
			// S.M.A.S.M check
			if (ul=='S' && ur=='M' && dl=='S' && dr=='M') {
				xmas_dss++
			}
		}
		fmt.Printf("Row %d: %d right, %d left, %d down, %d up\n", i, xmas_rmm, xmas_rss, xmas_dmm, xmas_dss)
	} // for i

	xmas_total := xmas_rmm + xmas_rss + xmas_dmm + xmas_dss
	return xmas_total

	// Submissions:
	// (1) 1923 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
