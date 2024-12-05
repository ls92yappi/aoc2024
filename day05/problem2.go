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

func applyRule(orig []int, i int, j int, rn int) []int {
	fixed := make([]int, 0)
	//fmt.Printf("Apply '%v', %d, %d rule number %d = (%d|%d)...\n", orig, i, j, rn, AllRules[rn][0], AllRules[rn][1])

	fixed = append(fixed, orig[:i]...)
	fixed = append(fixed, orig[j])
	fixed = append(fixed, orig[i])
	if (i+1 <= j-1) {
		fixed = append(fixed,orig[i+1:j]...) // j not j-1 was the bug I squashed
	}
	if (j+1 < len(orig)) {
		fixed = append(fixed,orig[j+1:]...)
	}
	//fmt.Printf("\t\t %v \n", fixed)
	_ = rn
	return fixed
}

func fixSequence(orig []int) []int {
	//fmt.Printf("\n-----\n\nFixing: %v\n", orig)
	numRules := len(AllRules)
	nu := len(orig)
	// treat this like an insertion sort where applyRule() is an in-order-preserving swap
	fixed := make([]int, 0)

	for i := range(nu-1)  {
		cur := orig[i]
		// preserve order of before cur
		for j := i+1; j < nu; j++ {
			right := orig[j]
			for rn := range(numRules) {
				befor := AllRules[rn][0]
				after := AllRules[rn][1]
				if (cur==after && right==befor) {
					// clear solution
					fixed = fixed[:0]
					fixed = applyRule(orig, i, j, rn)
					//fmt.Printf("Pass %d %d: %v\n", i, j, fixed)

					// overwrite the orig with the first i spots fixed copy
					copy(orig, fixed)
					// rerun the current j pass with updated elements
					cur = orig[i]
					break
				}
			}
		}
	}
	fmt.Printf("Fixed : %v\n", fixed)

	return fixed
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
			if !verifyRules(update) {
				fixed := fixSequence(update)
				// verified that every list of updates contains an odd number of entries
				// find middle page number
				middle := fixed[nu/2] // ex: 7/2 = 3 = correct 0-based index
				// add to answer
				answer += middle
			}
		}
	} // for i

	return answer

	// Submissions:
	// (1) 5346 = Correct
}

func main() {
	AllRules = make([][]int, 0)
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
