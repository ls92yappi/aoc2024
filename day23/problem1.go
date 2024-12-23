package main

// See README.md for problem description
// Usage: problem1 input1.txt

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	. "github.com/ls92yappi/aoc"
)

///////////////////////////////////////////////////////////

type Net map[string][]string
type Trip []string // [3] makes array, better if a slice
type TripSet map[string]bool

///////////////////////////////////////////////////////////


func processInput(fname string) int {
	lines, n, _ := ReadWholeFile(fname)
	answer := 0

	N := make(Net, 0)
	T := make(TripSet, 0)

	for i := range(n) {
		line := lines[i]
		if len(line) < 3 { continue }

		// find network connection
		pc1,pc2,_ := strings.Cut(line,"-")
		// update pc1's network
		if _, ok := N[pc1]; !ok {
			N[pc1] = []string{pc2}
		} else {
			N[pc1] = append(N[pc1], pc2)
		}

		// update pc1's network
		if _, ok := N[pc2]; !ok {
			N[pc2] = []string{pc1}
		} else {
			N[pc2] = append(N[pc2], pc1)
		}
	} // for i

	tPCs := make([]string, 0)
	for k,_ := range(N) {
		if string(k[0]) != "t" {
			continue
		}
		tPCs = append(tPCs, k)
	}
	//PCs := maps.Keys(N)
	sort.Strings(tPCs)
	fmt.Printf("AllKeys = %+v\n", tPCs)

	// construct all triples that contain a t-starting PC
	for _,k := range (tPCs) {
		for _,s := range(N[k]) {
			for _,t := range(N[s]) {
				if slices.Contains(N[t], k) {
					// put them in alpha-order for consistency
					tr := Trip{k,s,t}
					sort.Strings(tr)
					tk := strings.Join(tr,",")
					// treat as a set to auto-dedup
					T[tk] = true
				}
			}
		}
	}

	fmt.Printf("AllTrips Count = %d\n", len(T))

	answer = len(T)
	return answer

	// Submissions:
	// (1) 893 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
