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


	keep := true // assume mode only resets at beginning of input, not of each line

	re := regexp.MustCompile(`mul\(\d\d?\d?,\d\d?\d?\)`)
	reOnMode := regexp.MustCompile(`do\(\)`)
	reOffMode := regexp.MustCompile(`don't\(\)`)
	multsum := 0
	for i:=0;i<numLines;i++ {
		line := lines[i]
		if len(line) < 1 {
			continue
		}
		seqs := re.FindAllString(line, -1)
		seqlocs := re.FindAllStringIndex(line, -1)
		onnlocs := reOnMode.FindAllStringIndex(line, -1)
		offlocs := reOffMode.FindAllStringIndex(line, -1)
		numseqs := len(seqs)
		maxseqloc := seqlocs[len(seqlocs)-1][0]
		maxonnloc := onnlocs[len(onnlocs)-1][0]
		maxoffloc := offlocs[len(offlocs)-1][0]
		//minseqloc := seqlocs[0][0]
		//minonnloc := onnlocs[0][0]
		//minoffloc := offlocs[0][0]

		// we know that there will be at least one onnloc and one offloc per row
		// we know the longest line is 4000 characters
		if (numseqs > 0) {
			//fmt.Printf("Line %d Seqs: %v\n", i, seqlocs)
			//fmt.Printf("Line %d Onns: %v\n", i, onnlocs)
			//fmt.Printf("Line %d Offs: %v\n", i, offlocs)
			//fmt.Println()
			
			seqdex := 0
			onndex := 0
			offdex := 0

			// progress through the list one character spot at a time, looking for loc matches
			for spot := 0; spot<=Max([]int{maxseqloc,maxonnloc,maxoffloc}); spot++ {
				if spot <= maxonnloc && spot == onnlocs[onndex][0] {
					onndex++
					keep = true
					continue
				}
				if spot <= maxoffloc && spot == offlocs[offdex][0] {
					offdex++
					keep = false
					continue
				}
				if spot <= maxseqloc && spot == seqlocs[seqdex][0] {
					if !keep {
						seqdex++
						continue
					}
					mult := seqs[seqdex]
					left,right,_ := strings.Cut(mult,",")
					_,left2,_ := strings.Cut(left,"(")
					right2,_,_ := strings.Cut(right,")")

					l, _ := strconv.Atoi(left2)
					r, _ := strconv.Atoi(right2)
					product := l*r
					multsum += product
					//fmt.Printf("%d %d @ %d: %q  '%d x %d' = %d  ---> %d\n", i, seqdex, spot, mult, l, r, product, multsum)

					seqdex++
				}
			}
		} // if numseqs
		//fmt.Printf("Line %d: MinSeq@ %d, MinOnn@ %d, MinOff@ %d\n", i, minseqloc, minonnloc, minoffloc)
		//fmt.Printf("Line %d: MaxSeq@ %d, MaxOnn@ %d, MaxOff@ %d\n", i, maxseqloc, maxonnloc, maxoffloc)
	} // for i

	return multsum

	// Submissions:
	// (1) 107862689 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
