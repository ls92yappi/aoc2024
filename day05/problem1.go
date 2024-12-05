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

var AllRules [][]int

// numRules = 1176, numUpdates = 197
// if these numbers were bigger, I'd do a topological sort
// since they are small enough, simply going to naively check all pairs against all rules
func verifyRules(update []int) bool {
	numRules := len(AllRules)
	nu := len(update)
	for i := range(nu-1)  {
		left := update[i]
		for j := i+1; j < nu; j++ {
			right := update[j]
			for rn := range(numRules) {
				befor := AllRules[rn][0]
				after := AllRules[rn][1]
				if (left==after && right==befor) {
					return false
				}
			}
		}
	}

	return true
}

func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, n, _ := ReadWholeFile(fname)
	numRules := 0
	numUpdates := 0
	answer := 0

	for i := range(n) {
		line := lines[i]
		if len(line) < 3 {
			numRules = i
			numUpdates = n - numRules - 1
			fmt.Printf("numRules = %d, numUpdates = %d\n", numRules, numUpdates)
			continue
		}

		if numRules < 1 {
			rule,_,_ := IntSlice(line,"|")
			AllRules = append(AllRules, rule)
		}

		if numUpdates > 0 {
			update,nu,_ := IntSlice(line,",")
			if verifyRules(update) {
				// verified that every list of updates contains an odd number of entries
				// find middle page number
				middle := update[nu/2] // ex: 7/2 = 3 = correct 0-based index
				// add to answer
				answer += middle
			}
		}
	} // for i

	return answer

	// Submissions:
	// (1) 6260 = Correct
}

func main() {
	AllRules = make([][]int, 0)
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
