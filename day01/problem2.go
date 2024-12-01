package main

// See README.md for problem description

// Usage: problem01 input01.txt
// Currently templated off of my aoc2023/day23 code

/* Analysis: ***************************
Left ID, Right ID
Sort left list small-to-large
Sort right list small-to-large
Using a single pass through each of the left and right
 lists, count the number of occurrences of each left
 list member's appearances in the right list.
Sum the product of the left value with num occurrences
************************************* */

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////

var L []int
var R []int

var Size int

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


// msg is interface{} because cannot convert error to string
func die(msg interface{}) {
	log.Println(msg)
	os.Exit(1)
}

///////////////////////////////////////////////////////////


func init() {
	//Routes = make(map[string]int,0)
	L = make([]int, 1024)
	R = make([]int, 1024)
}


///////////////////////////////////////////////////////////

func processInput(fname string) int {
	// open file
	file, err := os.Open(fname)
	if err != nil { die(err) }
	defer file.Close()

	// read the whole file in
	srcbuf, err := ioutil.ReadAll(file)
	if err != nil { die(err) }
	src := string(srcbuf)


	lines := strings.Split(src, "\n")
	numLines := len(lines)
	Size = numLines-1


	// Fill in L and R using ParseLine()
	for i:=0;i<Size;i++ {
		line := lines[i]
		if len(line) < 1 {
			continue
		}
		left, right, found := strings.Cut(line, "   ")
		if !found {
			die("improper input line")
		}
		l, _ := strconv.Atoi(left)
		r, _ := strconv.Atoi(right)

		L[i] = l
		R[i] = r
	}

	sort.Ints(L)
	sort.Ints(R)

	var cur int
	var similarity int
	var sum int
	var rightIndex int = 1024-Size
	for i:=1024-Size;i<1024;i++ {
		// skip empty rows
		if L[i] == 0 {
			continue
		}

		// jump ahead rightIndex to correct value
		for j := rightIndex; j<1024 && R[j] < L[i]; j++ {
			rightIndex = j
		}

		// use previous similarity if a repeated left number
		if L[i] == cur {
			sum += cur*similarity
			debugf("%d * %d = %d\t%d\n", cur, similarity, cur*similarity, sum)
			continue
		}

		// new current value to check
		cur = L[i]

		// compute similarity here
		similarity = 0
		for j := rightIndex; j<1024 && R[j] <= cur; j++ {
			rightIndex = j
			if (R[j] == cur) {
				similarity++
			}
		}

		sum += cur*similarity
		debugf("%d * %d = %d\t%d\n", cur, similarity, cur*similarity, sum)
	}

	// Submissions:
	// (1) 20520794 = Correct

	return sum
}

func main() {
	var argc int = len(os.Args)
	var argv []string = os.Args


	if argc < 2 {
		fmt.Printf("Usage: %s [inputfile]\n", argv[0])
		os.Exit(1)
	}

	inputFileName := argv[1]
	output := processInput(inputFileName)
	fmt.Println(output)
	os.Exit(0)
}
