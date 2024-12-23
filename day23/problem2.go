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
type Trip []string // [3] makes array, better as a slice
type TripSet map[string]bool
type DegMap map[int][]string

///////////////////////////////////////////////////////////

// only have to search N[pc] for those pc's with dm[i>=l]
// bah, assumption above is useless since deg(pc1)=deg(pc2) always
// ex is all deg=4, input is all deg=13
func FindUniquePassword(N Net, PCs []string, l int) (string, bool) {
	bestSoFar := make([]string, 0)
	for i,Start := range(PCs) {
		// start a clique containing this PC
		clique := []string{Start}
		// check starting PC's neighbors
		for j,neighbor := range(N[Start]) {
			ok := true
			// verify this neighbor is part of the clique constructued so far
			for _,t := range(clique) {
				if !slices.Contains(N[neighbor], t) {
					ok = false
				}
			}
			// add it to the click if it passed verification
			if ok {
				clique = append(clique, neighbor)
			}
			// check if this is the best clique so far
			if len(bestSoFar) == 0 || len(clique) > len(bestSoFar) {
				bestSoFar = clique
				fmt.Printf("Trial %d %q NeighborIndex=%d BestLength=%d %+v\n", i, Start, j, len(bestSoFar), bestSoFar)
				// check if we have reached the target length
				if len(bestSoFar) == l {
					sort.Strings(bestSoFar)
					str := strings.Join(bestSoFar,",")
					return str, true
				}
			}
		}
		_ = i
	}
	return "", false
}

// was hoping to get some nice filtering based on degree(PC, but
// sadly all PCs are of uniform degree)
func AnalyzeDegreeCounts(N Net) (int,DegMap) {
	degCounts := make(map[int]int, 0)
	degMap := make(DegMap, 0)
	for pc,vs := range(N) {
		deg := len(vs)
		fmt.Printf("deg(%s) = %d:  %+v\n", pc, len(vs), vs)
		if _,ok := degCounts[deg]; ok {
			degCounts[deg]++
			degMap[deg] = append(degMap[deg], pc)
		} else {
			degCounts[deg] = 1
			degMap[deg] = []string{pc}
		}
	}

	degs := make([]int, 0)
	for k,_ := range(degCounts) {
		degs = append(degs, k)
	}
	sort.Ints(degs)
	//sort.Sort(sort.Reverse(sort.Ints(degs)))
	//degs = sort.Reverse(degs)

	fmt.Printf("Analysis: degCounts=%+v\n", degCounts)
	for i := len(degs)-1; i >= 0; i-- {
		d := degs[i]
		// must be at least d items of degree d to have a K_d subgraph
		if degCounts[d] < d {
			continue
		}
		return d,degMap
	}
	return 0,degMap
}

func processInput(fname string) string {
	lines, n, _ := ReadWholeFile(fname)
	answer := ""

	N := make(Net, 0)
	T := make(TripSet, 0)
	_ = T

	// construct networks N
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

	PCs := make([]string, 0)
	for k,_ := range(N) {
		PCs = append(PCs, k)
	}
	//PCs := maps.Keys(N)
	sort.Strings(PCs)
	//fmt.Printf("AllKeys = %+v\n", PCs)

	/*
	// construct all triples
	for _,k := range (PCs) {
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
	// 11011 Trips from P1 ignoring "t"-prefix requirement

	// convert into a sorted slice of strings
	AllTrips := make([]string, 0)
	for k,_ := range(T) {
		AllTrips = append(AllTrips, k)
	}
	sort.Strings(AllTrips)
	*/

	//fmt.Printf("NumPCs = %d\n", len(PCs))
	// get an upper bound first
	//maxPossLengthPW, dm := AnalyzeDegreeCounts(N)
	// ran AnalyzeDegreeCounts(). For the example, each of 16 PCs appeared precisely 4 times
	// for the input each of 520 PCs appeared precisely 13 times

	maxPossLengthPW := len(N[PCs[0]]) // they are known to all be the same
	//fmt.Printf("Max PW Length possible = %d\n", maxPossLengthPW)
	for l := maxPossLengthPW; l > 3; l-- {
		//fmt.Printf("Looking for a password of length %d ... \n", l)
		// find the unique password of length l
		sUnique, done := FindUniquePassword(N, PCs, l)
		if done {
			answer = sUnique
			return answer
		}
	}

	//fmt.Printf("AllTrips Count = %d\nPCs count = %d\n", len(AllTrips), len(PCs))

	return answer

	// Submissions:
	// (1) cw,dy,ef,iw,ji,jv,ka,ob,qv,ry,ua,wt,xz = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
