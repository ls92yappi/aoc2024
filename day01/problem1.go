package main

// See README.md for problem description

// Usage: problem01 input01.txt
// Currently templated off of my aoc2023/day23 code

/* Analysis: ***************************
Left ID, Right ID
Sort left list small-to-large
Sort right list small-to-large
abs the difference between left and right pairwise
sum the differences
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

	var diff int
	var sum int
	for i:=1024-Size;i<1024;i++ {
		diff = L[i] - R[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
		debugf("%d - %d = %d\t%d\n", L[i], R[i], diff, sum)
	}

	debug("-----\n")


	// Submissions:
	// (1) 1765812 = Correct

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
