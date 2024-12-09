package main

// See README.md for problem description

import (
	"fmt"
	//"strconv"
	//"strings"
	. "github.com/ls92yappi/aoc"
)

///////////////////////////////////////////////////////////

const Empty = '.'
const Skip = -1

func intsCompress(Actual []int) []int {
	nf := 0            // next free spot
	end := len(Actual)-1 // end index being pulled from
	block := Actual[end] // block being pulled from the end

	//fmt.Printf("%d\t%d\t%d\n", nf, block, end)

	for {
		// move nf to next available spot
		for ;nf<=end && Actual[nf]!=Skip;nf++ {
			// nf moved by for loop
		}

		// recognize when done
		if end <= nf {
			Actual = Actual[:nf]
			break
		}

		// move end back to a character to move
		for ;Actual[end]==Skip && end>=nf;end-- {
			// end moved by for loop
		}

		block = Actual[end]

		// move the block to the right spot, trimming off the old block from the end
		before := Actual[:nf]
		between := Actual[nf+1:end]

		//fmt.Print(show(before))
		//fmt.Printf("_%d_", block)
		//fmt.Println(show(between))

		var updated []int = make([]int, 0)
		updated = append(updated, before...)
		updated = append(updated, block)
		updated = append(updated, between...)
		Actual = updated

		// and advance the end spot back
		end--
	}
	return Actual
}

func show(a []int) string {
	sb := ""
	for i := range(a) {
		if (a[i] == Skip) {
			sb += string(Empty)
		} else {
			sb += fmt.Sprintf("%d", a[i])
		}
	}
	return sb
}

// ex 20 chars
// in 20000 chars
func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, _, _ := ReadWholeFile(fname)
	answer := 0

	// allocate memory for the data structures
	dmap := lines[0]
	dml := len(dmap)
	D := make([]int, 0) // treat everything as a slice of ints, not chars

	//cur := 0
	//dig := 0 // digit
	// expand disk contents from dmap into D
	for i := range(dml) {
		ch := dmap[i]
		n := ch - '0'
		id := i/2
		for range(n) {
			if (i%2 == 0) {
				D = append(D, id)
			} else {
				D = append(D, Skip)
			}
		}
	}

	// D is the expansion
	fmt.Println(show(D))

	// compress disk
	Actual := intsCompress(D)

	fmt.Println(show(Actual))
	// expanded input = 241263 chars long
	// compressed = 196201 chars long
	// dots 45062

	// checksum the compressed disk
	for i := range(len(Actual)) {
		check := i*Actual[i]
		// ignore skips, since they are negative
		if check > 0 {
			answer += check
		}
	}

	return answer

	// Submissions:
	// (1) 230090724 = Too Low
	// (2) 92092395920 = Too Low - Note: compress should be one block at a time, not one char at a time
	// (3) 6461289671426 -- Fixed by using []int instead of []byte
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
