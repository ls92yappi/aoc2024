package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////

//var L []int
//var R []int

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
	//L = make([]int, 1024)
	//R = make([]int, 1024)
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

		if Safe(levels) {
			numsafe++
		} else {
			for j := range(numlevels) {
				leftSlice := levels[:j]
				rightSlice := levels[j+1:]
				isRemovingJSafe := Safe(leftSlice) && Safe(rightSlice)
				if isRemovingJSafe {
					// single spot end removals are good
					if j < 1  || j > numlevels-2 {
						numsafe++
						debugf("Row %d is fixable by removing entry %d from %v\n", i, j, levels)
						break
					}

					// verify the bridge of the removal first
					diff := math.Abs(levels[j+1] - levels[j-1])
					if diff >= 1 && diff <= 3 {
						// then check that increasing/decreasing is correct
						// and be careful of almost first and almost last elements
						//[40 38 43 44 46 48 49] should be fixable
						incStart := levels[1] > levels[0]
						if j==1 {
							incStart = levels[2] > levels[0]
						}
						incEnd := levels[numlevels-1] > levels[numlevels-2]
						if (j==numlevels-2) {
							incEnd = levels[numlevels-1] > levels[numlevels-3]
						}
						if incStart != incEnd {
							continue
						}

						// finally check gap direction
						//[23 26 23 24 27 28] is not fixable at the 2 spot
						incBridge := levels[j+1] > levels[j-1]
						if incStart != incBridge {
							continue
						}

						numsafe++
						debugf("Row %d is fixable by removing entry %d from %v\n", i, j, levels)
						break
					}
				}
			}
		}
	}

	return numsafe

	// Submissions:
	// (1) 670 = Incorrect
	// (2) 666 = Too Low
	// (3) 671 = Too Low
	// (4) 676 = Incorrect ... 5 minutes after 1:37
	// (5) 730 = Incorrect ... 5 minutes after 2:05
	// (6) 689
	// The key to getting it right was leveraging the Safe() function from Problem 1
	// solution, followed by building the right set of checks
}

// copied over from problem1, but with initial numlevels check added
func Safe(levels []float64) bool {
	safe := true
	numlevels := len(levels)
	if numlevels < 2 {
		return safe
	}
	increasing := levels[1] > levels[0]
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
	return safe
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
