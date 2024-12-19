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


///////////////////////////////////////////////////////////

const white = "w"
const blue = "u"
const black = "b"
const red = "r"
const green = "g"

var Patterns []string

func possible(s string) string {
	ok := ""

	for _, vs := range(Patterns) {
		// return true on match
		if s == vs {
			return fmt.Sprintf("[%s]", vs)
		}

		// eliminate compositions that do not start out correct
		if !strings.HasPrefix(s, vs) {
			continue
		}

		// attempt substrings
		sub := strings.TrimPrefix(s, vs)
		soln := possible(sub)
		if len(soln) > 0 {
			return fmt.Sprintf("[%s]%s", vs, soln)
		}
	}

	return ok
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

		soln := possible(line)
		if len(soln) > 0 {
			//fmt.Printf("%d\t%s\n", i, soln)
			answer++
		}
	} // for i

	return answer

	// Submissions:
	// (1) 363 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
