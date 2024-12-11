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

const None = -1

// use a memo-ized inspired approach, given the order independence of blink()
var Blink map[int][]int

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

func blink(v int) []int {
	// rule 1: 0 -> 1
	// blink(0) = {1, None} precomputed
	if numDigits(v)%2 == 0 {
		// rule 2: evennumdigits -> leftdigits, rightdigits
		nd := numDigits(v)
		mid := tenToTheN(nd/2)
		ld := v/mid
		rd := v-(ld*mid)
		return []int{ld, rd}
	}
	// rule 3: old*2024
	return []int{v*2024, None}
}

func processInput(fname string) int {
	lines, _, _ := ReadWholeFile(fname)
	answer := 0
	/*
	for i := range(n) {
		line := lines[i]
		if len(line) < 2 {
			continue
		}
	}
	*/
	beforeSlice, _, _ := IntSlice(lines[0], " ")

	// convert from slice to occurrence count map, since order independent
	before := make(map[int]int, 0) // number of occurrences v of key k
	for _,v := range(beforeSlice) {
		before[v]++
	}

	fmt.Printf("BEGIN: %v\n", before)

	// memoized version of blink()
	Blink = make(map[int][]int, 0)
	Blink[0] = []int{1, None} // precomputing blink(0)

	// number of blink passes to run
	for i := range(75) {
		after := make(map[int]int, 0) // number of occurrences v of key k
		for k,v := range(before) {
			_, ok := Blink[k]
			if !ok {
				Blink[k] = blink(k)
			}
			a := Blink[k][0]
			b := Blink[k][1]

			after[a] += v
			if b != None {
				after[b] += v
			}
		}
		before = after
		//fmt.Printf("BLINK %d: %d\n", i+1, len(before))
		_ = i
	}

	// compute the answer
	for _,v := range(before) {
		answer += v
	}

	return answer

	// Submissions:
	// (1) 250783680217283 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
