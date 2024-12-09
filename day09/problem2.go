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

type F struct {
	id int
	sz int
}

// improvement: could remove empties of size 0 entirely
func combineEmpties(orig []F) []F {
	list := orig
	numCombined := 0
	//fmt.Printf("BEFORE: %s\n", showF(orig))
	for spot := len(orig)-1; spot>0; spot-- {
		b := list[spot]
		a := list[spot-1]
		if a.id != Skip || b.id != Skip { continue }

		numCombined++
		combo := F{Skip, a.sz+b.sz}
		before := list[:spot-1]
		after := list[spot+1:]

		var updated []F = make([]F, 0)
		updated = append(updated, before...)
		updated = append(updated, combo)
		updated = append(updated, after...)

		list = updated
	}
	//fmt.Printf("AFTER : %s\n", showF(list))
	//if numCombined > 0 {
	//	fmt.Printf("Combined %d\n", numCombined)
	//}
	return list
}

func fCompress(Actual []F) []F {
	nf := 0               // next free spot
	end := len(Actual)-1  // end index being pulled from
	bid := Actual[end].id // block being pulled from the end
	var block F
	var empty F

	// approach problem as block to test (bt)
	for bv := bid; bv>0; bv-- {
		//fmt.Printf("\tbv = %d\n", bv)
		//fmt.Printf("%s\n", showF(Actual))

		// find F matching block value bv
		for end = 0; end < len(Actual); end++ {
			block = Actual[end]
			if block.id == bv {
				break
			}
		}
		bs := block.sz

		// look for an empty slot larger enough to hold it
		for nf = 0; nf < end; nf++ {
			empty = Actual[nf]
			if empty.id != Skip { continue }
			if empty.sz >= bs { break }
		}

		// verify a move is possible
		if empty.sz < bs {
			continue
		}

		// verify the gap isn't after the current block
		if nf >= end {
			//fmt.Printf("Gap too late for block %d\n", bv)
			continue
		} 

		// shrink the empty slot
		empty.sz -= bs
		fillEmpty := F{Skip,block.sz}

		// move the block to the right spot, trimming off the old block from the end
		before := Actual[:nf]
		between := Actual[nf+1:end]
		after := Actual[end+1:]

		var updated []F = make([]F, 0)
		updated = append(updated, before...)
		updated = append(updated, block)
		if empty.sz > 0 {
			// keep partial empty if big enough
			updated = append(updated, empty)
		}
		updated = append(updated, between...)
		updated = append(updated, fillEmpty)
		updated = append(updated, after...)

		Actual = combineEmpties(updated)
		_ = fillEmpty
	}
	return Actual
}

func showF(a []F) string {
	sb := ""
	for i := range(a) {
		id := a[i].id
		sb += fmt.Sprintf("[")
		for range(a[i].sz) {
			if id == Skip {
				sb += string(Empty)
			} else {
				sb += fmt.Sprintf("%d", id)
			}
		}
		sb += fmt.Sprintf("]")
	}
	return sb
}

// ex 20 chars
// in 20000 chars
func processInput(fname string) int {
	lines, _, _ := ReadWholeFile(fname)
	answer := 0

	dmap := lines[0]
	dml := len(dmap)
	D := make([]F, 0) // treat everything as a slice of ints, not chars

	// expand disk contents from dmap into D
	for i := range(dml) {
		ch := dmap[i]
		n := int(ch - '0')
		id := i/2
		f := F{Skip, n}
		if (i%2 == 0) {
			f = F{id, n}
		}
		D = append(D, f)
	}

	// D is the expansion
	fmt.Println(showF(D))

	// compress disk
	fmt.Println("__COMPRESS__")
	Actual := fCompress(D)
	fmt.Println(showF(Actual))

	// checksum the compressed disk
	cur := 0
	for i := range(len(Actual)) {
		f := Actual[i]
		for range(f.sz) {
			check := cur*f.id
			if check > 0 {
				answer += check
			}
			cur++
		}
	}

	return answer

	// Submissions:
	// (1) 9645271028911 = Too High (forgot to skip move if gap followed block)
	// (2) 6488291456470 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
