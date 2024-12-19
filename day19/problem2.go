package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/ls92yappi/aoc" // problems with go proxy not being even remotely up-to-date
)

///////////////////////////////////////////////////////////

var Patterns []string

// memoized results from count()
var Count map[string]int

// Memoized calls took care of the performance issue that produced correct results for example
func count(s string) int {
	if len(s) < 1 {
		return 1
	}
	numWays := 0

	for _, vs := range(Patterns) {
		// eliminate compositions that do not start out correct
		if !strings.HasPrefix(s, vs) {
			continue
		}

		// attempt substrings, using a memoized version when able
		sub := strings.TrimPrefix(s, vs)
		_, ok := Count[sub]
		if !ok {
			Count[sub] = count(sub)
		}
		numWays += Count[sub]
	}

	return numWays
}

func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, n, _ := ReadWholeFile(fname)
	answer := 0

	Patterns = strings.Split(lines[0], ", ")
	sort.Strings(Patterns)

	for i := 2; i < n; i++ {
		line := lines[i]
		if len(line) < 1 {
			continue
		}

		// memoized version of count()
		Count = make(map[string]int, 0)
		numWays := count(line)
		if numWays > 0 {
			fmt.Printf("%d\t%d\n", i, numWays)
			answer+=numWays
		}
	} // for i

	return answer

	// Submissions:
	// (1) 642535800868438 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
