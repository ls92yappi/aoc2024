package main

// See README.md for problem description

import (
	"fmt"
	//"strconv"
	//"strings"
	. "github.com/ls92yappi/aoc"
)

///////////////////////////////////////////////////////////

const MaxInt int = 9223372036854775807
const Ten18 int  = 1000000000000000000

func numDigits(n int) int {
	if (n < 0) {
		fmt.Printf("Overflow on %d!!!\n", n)
		return 0
	}
	cmp := Ten18
	dig := 0
	for cmp = Ten18; cmp>=1; cmp/=10 {
		if n >= cmp {
			dig++
		}
	}
	return dig
}

func tenToTheN(n int) int {
	base := 1
	for range(n) {
		base *= 10
	}
	return base
}

func blink(before []int) []int {
	after := make([]int, 0)
	//fmt.Printf("%d %v\n", len(before), before)
	for i,v := range(before) {
		//fmt.Printf("%d %d\n", i, v)
		if v == 0 {
			// rule 1: 0 -> 1
			after = append(after, 1)
		} else if numDigits(v)%2 == 0 {
			// rule 2: evennumdigits -> leftdigits, rightdigits
			nd := numDigits(v)
			mid := tenToTheN(nd/2)
			ld := v/mid
			rd := v-(ld*mid)
			after = append(after, ld)
			after = append(after, rd)
		} else {
			// rule 3: old*2024
			after = append(after, v*2024)
		}
		_ = i
	}
	return after
}

func processInput(fname string) int {
	lines, n, _ := ReadWholeFile(fname)
	n--
	answer := 0
	/*
	for i := range(n) {
		line := lines[i]
		if len(line) < 2 {
			continue
		}
	}
	*/
	before, _, _ := IntSlice(lines[0], " ")
	//fmt.Printf("BEGIN: %v\n", before)

	for b := range(25) {
	//for b := range(25) {
		after := blink(before)
		//fmt.Printf("BLINK %d: %v\n", b, after)
		//copy(before, after)
		before = after
		//fmt.Printf("AFTER %d: %v\n", b, after)
		_ = b
	}

	answer = len(before)

	return answer

	// Submissions:
	// (1) 211306 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
