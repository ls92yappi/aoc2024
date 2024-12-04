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

	xmas_rr := 0
	xmas_ll := 0
	xmas_dd := 0
	xmas_uu := 0
	xmas_dr := 0
	xmas_dl := 0
	xmas_ur := 0
	xmas_ul := 0

	for i := range(n) {
		line := lines[i]
		for j := range(n) {
			ch := line[j]
			//fmt.Printf("%d %d %q\t",i, j, ch)
			// forward check
			if (j<n-3 && ch=='X' && line[j+1]=='M' && line[j+2]=='A' && line[j+3]=='S') {
				xmas_rr++
			}
			// backward check
			if (j>=3 && ch=='X' && line[j-1]=='M' && line[j-2]=='A' && line[j-3]=='S') {
				xmas_ll++
			}
			// down check
			if (i<n-3 && ch=='X' && lines[i+1][j]=='M' && lines[i+2][j]=='A' && lines[i+3][j]=='S') {
				xmas_dd++
			}
			// up check
			if (i>=3 && ch=='X' && lines[i-1][j]=='M' && lines[i-2][j]=='A' && lines[i-3][j]=='S') {
				xmas_uu++
			}
			// down right check
			if (j<n-3 && i<n-3 && ch=='X' && lines[i+1][j+1]=='M' && lines[i+2][j+2]=='A' && lines[i+3][j+3]=='S') {
				xmas_dr++
			}
			// down left check
			if (j>=3 && i<n-3 && ch=='X' && lines[i+1][j-1]=='M' && lines[i+2][j-2]=='A' && lines[i+3][j-3]=='S') {
				xmas_dl++
			}
			// up right check
			if (j<n-3 && i>=3 && ch=='X' && lines[i-1][j+1]=='M' && lines[i-2][j+2]=='A' && lines[i-3][j+3]=='S') {
				xmas_ur++
			}
			// up left check
			if (j>=3 && i>=3 && ch=='X' && lines[i-1][j-1]=='M' && lines[i-2][j-2]=='A' && lines[i-3][j-3]=='S') {
				xmas_ur++
			}
		}
		fmt.Printf("Row %d: %d right, %d left, %d down, %d up\n", i, xmas_rr, xmas_ll, xmas_dd, xmas_uu)
	} // for i

	xmas_total :=xmas_rr + xmas_ll + xmas_dd + xmas_uu + xmas_dr + xmas_dl + xmas_ur + xmas_ul
	return xmas_total

	// Submissions:
	// (1) 2504 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
