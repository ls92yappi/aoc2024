package main

// See README.md for problem description

// Usage: problem1 input1.txt
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
	"math"
	"os"
	//"sort"
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

	numsafe := 0
	// Fill in L and R using ParseLine()
	for i:=0;i<=Size;i++ {
		line := lines[i]
		if len(line) < 1 {
			continue
		}
		levs := strings.Split(line, " ")
		numlevels := len(levs)
		levels := make([]float64, numlevels)
		for j := range(numlevels) {
			levels[j], _ = strconv.ParseFloat(levs[j], 64)
		}
		increasing := levels[1] > levels[0]
		safe := true
		for j := range(numlevels-1) {
			diff := math.Abs(levels[j+1] - levels[j])
			if diff < 1 || diff > 3 {
				safe = false
				break
			}
			inc := levels[j+1] > levels[j]
			if inc != increasing {
				safe = false
				break
			}
		}
		if safe {
			//fmt.Printf("Row %d is ok: %v\n", i, levels)
			numsafe++
		} else {
			fmt.Println(levels)
		}
	}

	return numsafe

	// Submissions:
	// (1) 679 = Incorrect
	// (2) 680 = Incorrect
	// (3) 171 = Incorrect
	// (4) 660 = Correct
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
