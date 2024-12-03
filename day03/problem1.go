package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
	lines, numLines, _ := ReadWholeFile(fname)

	re := regexp.MustCompile(`mul\(\d\d?\d?,\d\d?\d?\)`)
	//rn := regexp.MustCompile(`(\d*)`)
	multsum := 0
	for i:=0;i<numLines;i++ {
		line := lines[i]
		if len(line) < 1 {
			continue
		}
		seqs := re.FindAllString(line, -1)
		//levels, numlevels, _ := IntSlice(line, " ")
		numseqs := len(seqs)
		if (numseqs > 0) {
			for j := range(numseqs) {
				mult := seqs[j]
				left,right,_ := strings.Cut(mult,",")
				_,left2,_ := strings.Cut(left,"(")
				right2,_,_ := strings.Cut(right,")")

				l, _ := strconv.Atoi(left2)
				r, _ := strconv.Atoi(right2)
				product := l*r
				multsum += product
				fmt.Printf("%d %d: %q  '%d x %d' = %d  ---> %d\n", i, j, mult, l, r, product, multsum)
			}
		}
		//fmt.Printf("%d: %d = %+v\n", i, len(seqs), seqs)


	}

	return multsum

	// Submissions:
	// (1) 184122457 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
