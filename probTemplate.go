package main

// See README.md for problem description
// Usage: problem1 input1.txt

import (
	"fmt"
	//"regexp"
	//"strconv"
	//"strings"
	. "github.com/ls92yappi/aoc"
)

///////////////////////////////////////////////////////////


///////////////////////////////////////////////////////////


func processInput(fname string) int {
	lines, n, _ := ReadWholeFile(fname)
	answer := 0

	for i := range(n) {
		line := lines[i]
		//rule,num,err := IntSlice(line,"|")
		_ = line
	} // for i

	return answer

	// Submissions:
	// (1) 5346 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
