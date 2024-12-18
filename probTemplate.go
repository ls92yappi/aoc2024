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


///////////////////////////////////////////////////////////


func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, n, _ := ReadWholeFile(fname)
	answer := 0

	for i := range(n) {
		line := lines[i]
		//rule,num,err := IntSlice(line,"|")
		//AllRules = append(AllRules, rule)
		_ = line
	} // for i

	return answer

	// Submissions:
	// (1) 5346 = Correct
}

func main() {
	//AllRules = make([][]int, 0)
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
