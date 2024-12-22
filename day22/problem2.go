package main

// See README.md for problem description
// Usage: solve 2 input
// or   : go run problem2 input2.txt

import (
	"fmt"
	"strconv"
	. "github.com/ls92yappi/aoc"
)

///////////////////////////////////////////////////////////

const Magic = 16777216 // 2^24
const NumSeqs = 19*19*19*19 // 19^4	= 130_321
const NumSecrets = 2000

type Seq [4]int // known sequence length is 4 diffs

///////////////////////////////////////////////////////////

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
	n--
	answer := 0

	var ScoreSeqs = make(map[Seq]int, 0)

	// new approach: do as little as absolutely possible this time through
	// do not brute force things in stages, but instead aim for fewest passes
	// thru the data, keeping only what we absolutely must compute with
	for i := range(n) {
		line := lines[i]
		v, _ := strconv.Atoi(line)

		// create just a list of prices
		p := []int{v%10} // initialize prices with provided secret
		for range(NumSecrets) {
			v = next(v)
			p = append(p, v%10)
		}

		score := make(map[Seq]int, 0) // score for a given sequence for this monkey
		for j := range(NumSecrets+1-4) { // +1 b/c initial secret, -4 b/c Seq length
			// create price diff sequence
			seq := Seq{p[j+1]-p[j], p[j+2]-p[j+1], p[j+3]-p[j+2], p[j+4]-p[j+3]}
			// have we already seen this sequence for this monkey
			_, found := score[seq]
			// only keep the first appearance
			if found { continue }
			// grab the score for this monkey's sequence
			score[seq] = p[j+4]
			// update the shared monkeys sequence sums
			_, ok := ScoreSeqs[seq]
			if !ok { ScoreSeqs[seq] = 0 }
			ScoreSeqs[seq] += score[seq]
		}
	} // for i

	// now look for the biggest sum in the map of sequences
	bestScore := 0
	for key,val := range(ScoreSeqs) {
		if val > bestScore {
			bestScore = val
			fmt.Printf("New best: %d, Seq: %+v\n", val, key)
		}
	}

	// answer is the best score of any sequence
	answer = bestScore

	return answer

	// Submissions:
	// (1) 1587 = Too High // ran in 9min02sec [0 0 -1 1] which was sequence number 65142
	// was off by 1. The problem states there are 2000 price changes (2001 secret numbers)
	// because the provided secret number counts. Just need to prepend that value to the
	// MPs/MDs are re-run ... see you in about 10 minutes ;)
	// Baaah! Didn't get any lower.
	// (2) 1582 = Correct // ran in 1.4seconds [-1 1 -3 3]
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
