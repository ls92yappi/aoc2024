package main

// See README.md for problem description

import (
	"fmt"
	"strconv"
	"strings"
	. "github.com/ls92yappi/aoc" // problems with go proxy not being even remotely up-to-date
)

///////////////////////////////////////////////////////////

func concat(a, b int) int {
	r := a
	// largest b found was 999
	if b >= 100 {
		// 3 digit number
		r = a*1000 + b
	} else if b >= 10 {
		// 2 digit number
		r = a*100 + b
	} else {
		// 1 digit number
		r = a*10 + b
	}
	return r
}

func doable(goal int, vals []int, sz int) int {
	r := 0
	// try additions and mults to achieve goal, if so return goal

	// binary representation of 2 ^ (sz-1) power
	pows := 2 << (sz-1)

	//fmt.Printf("%d -> %d\n", sz, pows)
	for ops := range(pows) {
		trial := vals[0]
		for i := range(sz-1) {
			val := vals[i+1]
			op := 2 << i
			if ops&op == op {
				// mult attempt
				trial *= val
			} else {
				// add attempt
				trial += val
			}
		}
		if trial == goal {
			return goal
		}
	}

	return r
}

func trinaryDoable(goal int, vals []int, sz int) int {
	// try concats, mults and adds to achieve goal, if so return goal

	// represent trinary as twin binary
	// binary representation of 2 ^ (2*(sz-1)) power
	pows := 2 << (2*(sz-1))

	for ops := range(pows) {
		trial := vals[0]
		for i := range(sz-1) {
			// already too big, short circuit this attempt
			if trial > goal {
				break
			}

			val := vals[i+1]
			opPri := 2 << (2*i)
			opSec := 2 << (2*i + 1)
			opTrinary := 0
			if (ops&opPri == opPri) {
				opTrinary += 2
			}
			if (ops&opSec == opSec) {
				opTrinary += 1
			}
			switch opTrinary {
				// pushed concat attempt to top to make it short-circuit quicker
			case 0:
				// concat attempt
				trial = concat(trial,val)
			case 1:
				// mult attempt
				trial *= val
			case 2:
				// add attempt
				trial += val
			case 3:
				// undefined
				trial = trial
			}
		}
		if trial == goal {
			return goal
		}
	}

	return 0
}

// rtv = recursiveTrinaryDoable
func rtv(goal int, vals []int, sz int, cur int, i int) bool {
	// speed optimization shortcircuit if already too big
	// reduces runtime by about 7%
	// computes in 0.149s to 0.164s with cur>goal filter
	// computes in 0.160s to 0.176s without it
	if cur > goal {
		return false
	}

	if i == sz {
		return cur == goal
	}
	// recursion with shortcircuiting || with least desirable last greatly
	// improves performance since not full bruteforcing the problem
	return rtv(goal, vals, sz, cur+vals[i], i+1) || 
		   rtv(goal, vals, sz, cur*vals[i], i+1) ||
		   rtv(goal, vals, sz, concat(cur,vals[i]), i+1)
}

func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, n, _ := ReadWholeFile(fname)
	answer := 0
	for i := range(n) {
		line := lines[i]
		if len(line) < 2 {
			continue
		}
		//
		stv, eqn, _ := strings.Cut(line, ": ")
		tv, _ := strconv.Atoi(stv)
		
		vals, num, _ := IntSlice(eqn, " ")

		dv := 0
		useRecursion := true
		if !useRecursion {
			dv = doable(tv, vals, num)
			// doable() filter shortens from 28s to 17s to compute trinaryDoable()
			if dv == 0 {
				dv = trinaryDoable(tv, vals, num)
			}
			// answer = 637696738133259 in 16.4s Wrong
		} else {
			if rtv(tv, vals, num, vals[0], 1) {
				dv = tv
			}
			// answer = 637696070419031 in 0.16s Wrong
		}
		answer += dv
	} // for i

	return answer

	// Submissions:
	// (1) 637696738133259 = Too High (trinaryDoable method)
	// (2) 310738852181107 = Too Low  (bad 1-digit concats method)
	// (3) 637696070419031 = Correct  (Recursive method)
	// What bothers me is the non-recursive method looks the same
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
