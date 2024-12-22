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

const Magic = 16777216 // 2^24
const NumSeqs = 19*19*19*19 // 19^4
const NumSecrets = 2000 // 2000

type Prices [NumSecrets+1]byte
type PDiffs [NumSecrets]int8
type Seq [4]int8 // known sequence length is 4 diffs

// verified that all 1601 values in input1.txt are unique
// wc input1.txt
// #then
// sort -n input1.txt | uniq | wc

func scoreFirstSeqInDiff(seq Seq, pd PDiffs, pr Prices) int {
	score := 0
	for i := range(NumSecrets-4) {
		if pd[i+0] != seq[0] { continue }
		if pd[i+1] != seq[1] { continue }
		if pd[i+2] != seq[2] { continue }
		if pd[i+3] != seq[3] { continue }
		score = int(pr[i+4])
	}
	return score
}


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

// there exist 19*19*19*19 different possible sequences to score against the monkeys = 130,321
// 4 bytes per 130,321 = ~529KB
func processInput(fname string) int {
	lines, n, _ := ReadWholeFile(fname)
	n--
	answer := 0

	var MP = make([]Prices, n) // monkey prices
	var MD = make([]PDiffs, n) // monkey diffs

	// create maps of all monkeys Prices and PDiffs
	for i := range(n) {
		line := lines[i]
		if len(line) < 1 { continue }

		vals,_,_ := IntSlice(line,",")
		v := vals[0]
		var P Prices
		var D PDiffs
		P[0] = byte(v)
		for j := range(NumSecrets) {
			v = next(v)
			price := byte(v%10)
			P[j+1] = price
			D[j] = int8(P[j+1]-P[j])
		}
		MP[i] = P
		MD[i] = D
	} // for i

	// create all sequences of 4 consecutive diffs as the search space
	var AllSeqs [NumSeqs]Seq
	var i,j,k,l int8
	for i = -9; i <=9; i++ {
		for j = -9; j <=9; j++ {
			for k = -9; k <=9; k++ {
				for l = -9; l <=9; l++ {
					idx := int(i+9)*19*19*19 + int(j+9)*19*19 + int(k+9)*19 + int(l+9)
					AllSeqs[idx] = Seq{i,j,k,l}
				}
			}
		}
	}

	// score each of those sequences against the troop of monkeys
	bestScore := 0
	var ScoreSeqs [NumSeqs]int
	// 19^4 = 130_321

	jumpStart := 0 // 65142
	for i, seq := range(AllSeqs) {
		if i < jumpStart {
			continue
		}
		if i%3000==0 { fmt.Printf("%d %d\n", i, bestScore)}
		for j := range(n) {
			score := scoreFirstSeqInDiff(seq, MD[j], MP[j])
			//fmt.Printf("Score: %d Monkey: %s\n", score, lines[j])
			ScoreSeqs[i] += score
		}
		// keep best score
		if ScoreSeqs[i] > bestScore {
			bestScore = ScoreSeqs[i]
			fmt.Printf("New Best %d at i=%d seq=%+v\n", bestScore, i, seq)
			//if bestScore == 1587 {
			//	fmt.Printf("Sequence Number %d %+v\n", i, seq)
			//	return bestScore
			//}
		}
	}

	// answer is the best score of any sequence
	answer = bestScore

	// all MP and MD info fits within 6.5MB
	// AllSeqs within 0.6MB
	// ScoreSeqs within 4.8MB
	// total RAM guess < 12MB

	return answer

	// Submissions:
	// (1) 1587 = Too High // ran in 9min02sec [0 0 -1 1] which was sequence number 65142
	// was off by 1. The problem states there are 2000 price changes (2001 secret numbers)
	// because the provided secret number counts. Just need to prepend that value to the
	// MPs/MDs are re-run ... see you in about 10 minutes ;)
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
