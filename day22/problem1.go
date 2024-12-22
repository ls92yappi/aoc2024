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



// old*64, mix, prune
// /32, mix, prune
// *2048, mix, prune

// mix = bitwise xor old and cur
// prune = %16777216

// 123 -> ex1.txt

// 2000th

const Magic = 16777216
// will probably need to memoize next()

func next(num int) int {
	cur := num*64
	num ^= cur
	num %= Magic
	cur  = num/32
	num ^= cur
	num %= Magic
	cur = num*2048
	num ^= cur
	num %= Magic
	return num
}


func processInput(fname string) int {
	lines, n, _ := ReadWholeFile(fname)
	answer := 0

	for i := range(n) {
		line := lines[i]
		if len(line) < 1 { continue }

		vals,_,_ := IntSlice(line,",")
		v := vals[0]
		//fmt.Printf("%d:%d -> ", i, v)
		for range(2000) {
			v = next(v)
		}
		//fmt.Printf("%s:   %d\n", line, v)
		answer += v
	} // for i

	return answer

	// Submissions:
	// (1) 13429191512 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
