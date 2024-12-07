package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	"strconv"
	"strings"
	. "github.com/ls92yappi/aoc" // problems with go proxy not being even remotely up-to-date
)

///////////////////////////////////////////////////////////

func doable(goal int, vals []int, sz int) int {
	r := 0
	// try additions and mults to achieve goal, if so return goal

	// binary representation of 2 ^ (sz-1) power
	pows := 2 << (sz-1)

	//fmt.Printf("%d -> %d\n", sz, pows)
	for ops := range(pows) {
		trial := vals[0]
		for i := range(sz-1) {
			val := vals[i+1]
			op := 2 << i
			if ops&op == op {
				// mult attempt
				trial *= val
			} else {
				// add attempt
				trial += val
			}
		}
		if trial == goal {
			return goal
		}
	}

	return r
}

func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, n, _ := ReadWholeFile(fname)
	answer := 0

	for i := range(n) {
		line := lines[i]
		if len(line) < 2 {
			continue
		}
		//
		stv, eqn, _ := strings.Cut(line, ": ")
		tv, _ := strconv.Atoi(stv)
		vals, num, _ := IntSlice(eqn, " ")
		dv := doable(tv, vals, num)
		answer += dv
	} // for i

	return answer

	// Submissions:
	// (1) 66343330034722 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
