package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	//"regexp"
	//"strconv"
	"slices"
	"strings"
	. "github.com/ls92yappi/aoc" // problems with go proxy not being even remotely up-to-date
)

///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////

// places
const Empt = "."
const Obst = "#"
const Start = "S"
const Goal = "E"
// facings
const None = -1
const North = 0
const East  = 1
const South = 2
const West  = 3
// states
const Visited = 1
const Wall = 2
const Dead = 3
const Fork = 4
const Cross = 5
const Begin = 6
const Final = 7
const Corner = 8
const Psgway = 9 // Passageway

var Size = 0 //130 //10 // 130 // hard-coded per run since using arrays rather than slices
var TooBig = Size*Size*Size
var TimedOut = (Size+4)*(Size+4)

var FromsCount = 0

type Guard struct {
	Row int
	Col int
	Face int
}

// Points of interest
type PoI struct {
	R int // row
	C int // col
	S int // state (Begin,Final,Dead,Fork,Cross)
	T []int // touches index list of other PoIs it can directly reach
}

type Best struct {
	Route []int // primary payload
	Score int   // compared criteria
	Done bool   // already computed
	            // ancillary data
}

var SPos PoI
var GPos PoI

type scoreFunc func([]PoI, []int) int

//var Grid [][]string
//var AML []PoI // Abstract Maze List

// construct the maze into a slice of Points of Interest, gives start and goal indices, and validity check
// corners is whether corners are counted as PoI or not; typically true
func abstractMaze(g [][]string, start string, goal string, corners bool) ([]PoI, int, int, bool) {
	a := make([]PoI, 0)
	emptyAML := make([]PoI,0)
	sIdx := -1
	gIdx := -1
	ok := true

	state := None

	sz := len(g[0]) // assumes a square grid
	if sz < 3 { return emptyAML,-1,-1,false }

	// Construct Abstract Maze List in two passes
	// First passes identifies all of the Points of Interest by count of adjacent walls
	for i := range(sz) {
		for j := range(sz) {
			ch := g[i][j]
			emptyT := make([]int,0)
			if (i==0 || i==sz-1 || j==0 || j==sz-1) {
				// if on an outer wall
				switch ch {
				case Start:
					if sIdx != -1 {
						fmt.Printf("Already have a Starting position: %d\n",sIdx)
						return emptyAML,-1,-1,false
					}
					sIdx = len(a)
					poi := PoI{i,j,Begin,emptyT}
					a = append(a, poi)
					continue
				case Goal:
					if gIdx != -1 {
						fmt.Printf("Already have an Ending position: %d\n",gIdx)
						return emptyAML,-1,-1,false
					}
					gIdx = len(a)
					poi := PoI{i,j,Final,emptyT}
					a = append(a, poi)
					continue
				case Obst:
					continue
				case Empt:
					fmt.Printf("Outer edges must be Wall or Begin or Final: (%d,%d)\n", i,j)
					return emptyAML,-1,-1,false
				default:
					fmt.Printf("Outer edges must be Wall or Begin or Final: (%d,%d)\n", i,j)
					return emptyAML,-1,-1,false
				}
			}

			switch ch {
			case Start:
				if sIdx != -1 {
					fmt.Printf("Already have a Starting position: %d\n",sIdx)
					return emptyAML,-1,-1,false
				}
				sIdx = len(a)
				poi := PoI{i,j,Begin,emptyT}
				a = append(a, poi)
				continue
			case Goal:
				if gIdx != -1 {
					fmt.Printf("Already have an Ending position: %d\n",gIdx)
					return emptyAML,-1,-1,false
				}
				gIdx = len(a)
				poi := PoI{i,j,Final,emptyT}
				a = append(a, poi)
				continue
			case Obst:
				continue
			case Empt:
				nWall := If(g[i-1][j]==Obst,1,0)
				sWall := If(g[i+1][j]==Obst,1,0)
				wWall := If(g[i][j-1]==Obst,1,0)
				eWall := If(g[i][j+1]==Obst,1,0)
				wallsTouching := nWall + sWall + wWall + eWall
				switch wallsTouching {
				case 0:
					state = Cross
				case 1:
					state = Fork
				case 3:
					state = Dead
				case 2:
					if corners && (eWall != wWall) {
						// we DO care about corners, just not passageways
						state = Corner
					} else {
						state = Psgway
						continue
					}

				default: continue // normal passageway touches 2
				}
				poi := PoI{i,j,state,emptyT}
				a = append(a, poi)
				continue
			default:
				fmt.Printf("What is this strange character?: (%d,%d)=%q\n", i,j, ch)
				return emptyAML,-1,-1,false
			}
		}
	} // First Pass

	// Second pass goes through the list and updates its touches list
	for fIdx := range(len(a)-1) {
		for tIdx := fIdx+1; tIdx<len(a); tIdx++ {
			f := a[fIdx]
			t := a[tIdx]
			// check all pairs a[fIdx], a[tIdx]
			if f.R != t.R && f.C != t.C {
				// not on same row on or column, no relationship possible
				continue
			}
			// verify there are no Walls in between them
			if f.R == t.R {
				// same Row
				low := Min2(f.C, t.C)
				high := Max2(f.C, t.C)
				clear := true
				for col := low+1; col < high; col++ {
					if g[f.R][col] == Obst {
						clear = false
						break
					}
				}
				//fmt.Printf("Same Row %d,%d %t %+v %+v\n", fIdx, tIdx, clear, f, t)
				if clear {
					a[fIdx].T = append(a[fIdx].T, tIdx)
					a[tIdx].T = append(a[tIdx].T, fIdx)
				}
				continue
			}
			// same Col
			low := Min2(f.R, t.R)
			high := Max2(f.R, t.R)
			clear := true
			for row := low+1; row < high; row++ {
				if g[row][f.C] == Obst {
					clear = false
					break
				}
			}
			if clear {
				a[fIdx].T = append(a[fIdx].T, tIdx)
				a[tIdx].T = append(a[tIdx].T, fIdx)
			}
			continue
		}
	} // Second Pass

	// Third pass removes touches when there is a point in between
	for i := range(a)  {
		f := a[i]
		// create a filter list of items to remove from a[i].T since modifying
		// slices in place that you are iterating through is horribly error prone
		filterList := make([]int,0)
		// identify items to remove from a[i].T
		for j := len(f.T)-1; j > 0; j-- {
			t := a[f.T[j]]
			for k := j-1; k >=0; k-- {
				v := a[f.T[k]]
				if f.R==t.R {
					// row check
					if f.R==v.R {
						if (t.C<f.C && f.C<v.C) || (v.C<f.C && f.C<t.C) {
							// in between is fine
							continue
						}
						// more distant of t,v should be removed
						td := Abs(t.C-f.C)
						vd := Abs(v.C-f.C)
						if vd > td {
							filterList = append(filterList, f.T[k]) // k/v index
						} else {
							filterList = append(filterList, f.T[j]) // j/t index
						}
					}
				} else {
					// col check
					if f.C==v.C {
						if (t.R<f.R && f.R<v.R) || (v.R<f.R && f.R<t.R) {
							// in between is fine
							continue
						}
						// more distant of t,v should be removed
						td := Abs(t.R-f.R)
						vd := Abs(v.R-f.R)
						if vd > td {
							filterList = append(filterList, f.T[k]) // k/v index
						} else {
							filterList = append(filterList, f.T[j]) // j/t index
						}
					}
				}
			}
		}
		// create empty new list that will replace a[i].T
		newT := make([]int,0)
		// fill  it with the items that are not in the filterList
		for _,val := range(a[i].T) {
			if !slices.Contains(filterList, val) {
				newT = append(newT, val)
			}
		}
		// finally replace a[i].T touches list with the cleaned up version
		a[i].T = newT
	} // Third Pass

	ok = sIdx >= 0 && gIdx >= 0 && len(a)>=2

	return a, sIdx, gIdx, ok
}

// find the best route b from index f of existing bl
// len(b) > 0 at first call, so initialize with start pos of maze pre-built
func constructBest(a []PoI, bl []Best, f int, isRoot bool, sf scoreFunc) []Best {
	blAfter := make([]Best, len(bl))
	copy(blAfter,bl)

	// loop prevention
	if !isRoot {
		if bl[f].Done {
			return bl
		}
	}

	FromsCount++
	//fmt.Printf("From %d: %d\n", f, FromsCount)
	//if FromsCount > 200 { Die("Too many!") }
	blAfter[f].Done = true

	for i,t := range(a[f].T) {
		// compute this route to the target
		r := make([]int, len(blAfter[f].Route))
		copy(r,blAfter[f].Route)
		r = append(r,t)
		b := Best{r,sf(a,r),false}

		if blAfter[t].Score == 0 {
			blAfter[t] = b
		} else if b.Score > 0 && b.Score < blAfter[t].Score {
			blAfter[t] = b
		}

		// if we have a new candidate, compute from here
		if (blAfter[t].Score != bl[t].Score) {
			//fmt.Printf("Compute best from %d to %d\n", f, t)
			blAfter = constructBest(a, blAfter, t, false, sf)
		}
		_= i
	}

	return blAfter
}


// a route is a list of PoI indices from s to g, tracking visited status v from given base route 
// returns a list of routes without loops
func findAllRoutes(a []PoI, v []int, base []int, s int, g int) ([][]int, int) {
	rl := make([][]int, 0)
	numRoutes := 0

	if v[s] > 0 {
		// loop detected, bad
		return rl,0
	}

	// current route
	r := make([]int, len(base))
	copy(r,base)
	r = append(r,s)

	// mark this PoI as visited
	v[s] = v[s] + 1

	//fmt.Printf("Trying %v -> %d...\n", r, g)
	if s == g {
		// we have reached the end, and it is good
		fmt.Printf("SUCCESSFUL ROUTE FOUND!  %v\n", RouteString(a,r))
		rl = append(rl,r)
		return rl,1
	}

	if a[s].S == Dead {
		// dead ends are bad
		return rl,0
	}

	for i,idx := range(a[s].T) {
		indent := len(r)
		//RouteString(a,r)
		//fmt.Printf("%s(%d,%d)--(%d,%d): %s ...\n", strings.Repeat("  ",indent-1), a[s].R, a[s].C, a[idx].R, a[idx].C, RouteString(a,a[s].T))
		nrl, found := findAllRoutes(a, v, r, idx, g)
		if found > 0 {
			rl = append(rl, nrl...)
			numRoutes += found
		}
		_ = indent
		_ = i
	}

	return rl,numRoutes
}

func scoreTheRoutes(a []PoI, routes [][]int, sf scoreFunc) []int {
	scores := make([]int, 0)
	for _,route := range(routes) {
		score := sf(a, route)
		scores = append(scores, score)
	}
	return scores
}

func IsRouteSane(a []PoI, r []int) bool {
	// all sane routes must contain both the Begin locations
	sanity := len(r) >= 1 && a[r[0]].S == Begin

	// only true for completed routes, not intermediate routes
	// all sane routes must contain both the Begin and Final locations
	//sanity := len(r) >= 2 && a[r[0]].S == Begin && a[r[len(r)-1]].S == Final
	return sanity
}

func RouteString(a []PoI, r []int) string {
	s := ""
	if len(r) < 1 { return "" }

	s = fmt.Sprintf("(%d,%d)", a[r[0]].R, a[r[0]].C)
	for i := 1; i < len(r); i++ {
		p := a[r[i]]
		s += fmt.Sprintf("-(%d,%d)", p.R, p.C)
	}
	return s
}

func StateType(i int) string {
	s := ""
	switch(i) {
	case Wall: s = "Wall"
	case Dead: s = "Dead"
	case Fork: s = "Fork"
	case Cross: s = "Cross"
	case Begin: s = "BEGIN"
	case Final: s = "FINAL"
	case Corner: s = "Corner"
	case Psgway: s = "Psgway"
	default: s = "None"
	}
	return s
}

func PoIString(a []PoI, i int) string {
	p := a[i]
	s := fmt.Sprintf("PoI %d @(%d,%d) %s touches %s\n", i, p.R, p.C, StateType(p.S), RouteString(a,p.T))
	return s
}

// objective here is fewest turns, then fewest steps
func ScoreRouteFunc(a []PoI, r []int) int {
	score := 0

	// Ensure route goes from Begin to Final
	if !IsRouteSane(a, r) {
		fmt.Printf("Found insane route: %+v\n", RouteString(a,r))
		return -1
	}

	// Problem starts the reindeer facing East
	facing := East
	f := a[r[0]] // from spot

	for i := 1; i < len(r); i++ {
		t := a[r[i]] // top spot

		// num steps between adds 1 per
		steps := Abs(t.R-f.R) + Abs(t.C-f.C)
		score += steps

		newFacing := Direction(f, t)
		if newFacing == None {
			fmt.Printf("Direction facing problem: %+v\n", RouteString(a,r))
			score += 0
		}
		if Abs(newFacing-facing)==2 {
			//fmt.Printf("Can't make a U-turn: %+v\n", RouteString(a,r))
			score += 0
		}
		if newFacing != facing {
			// change in facing adds 1000
			score += 1000
		}

		f = t
		facing = newFacing
	}

	return score
}

func Direction(f PoI, t PoI) int {
	d := North
	// N=0,E=1,S=2,W=3
	// assume a single cardinal direction
	switch {
	case f.R==t.R && f.C == t.C:
		fmt.Printf("Stuck in one place: from=%+v to=%+v\n", f, t)
		d = None
	case f.R==t.R && f.C < t.C:
		d = East
	case f.R==t.R && f.C > t.C:
		d = West
	case f.C==t.C && f.R < t.R:
		d = North
	case f.C==t.C && f.R > t.R:
		d = South
	default:
		fmt.Printf("More than one direction: from=%+v to=%+v\n", f, t)
		d = None
	}
	return d
}

// 17x17 grid maze // goal 11048 (11 turns, 48 steps)
// 15x15 grid ex // goal 7036 (7 turns, 36 steps)
// 141x141 grid input
// Start facing East
func processInput(fname string) int {
	lines, _, _ := ReadWholeFile(fname)
	answer := 0

	Size = len(lines[0])
	TooBig = Size*Size*Size
	TimedOut = (Size+4)*(Size+4)


	// Construct the Maze Grid
	Grid := make([][]string, 0)
	for i := range(Size) {
		slRow := strings.Split(lines[i],"")
		Grid = append(Grid, slRow)
	} // for i


	// Abstract the Maze
	AML, sIdx, gIdx, ok := abstractMaze(Grid, Start, Goal, true) 
	if !ok {
		fmt.Printf("Abstract Maze construction failed!\n")
		return 0
	}

	// Summarize AML for debugging
	//fmt.Printf("\nPoints of Interest\n==================\n")
	count := 0
	for i,p := range(AML) {
		count += len(p.T)
		//fmt.Print(PoIString(AML,i))
		_ = i
	}
	//fmt.Printf("==================\n")
	fmt.Printf("%dx%d Grid has %d Points of Interest in it; BeginIndex=%d, FinalIndex=%d, Connectivity=%d\n", Size,Size, len(AML), sIdx, gIdx, count/2)


	// Optionally filter out Dead ends here
	// skipping since it is running fast enough

	// no longer going this approach to the problem
	if false {
		// Find all routes list without loops
		//Visited := make([]int, len(AML)) // pass empty Visited list for AML
		//emptyRoute := make([]int, 0)
		//Routes, numRoutes := findAllRoutes(AML, Visited, emptyRoute, sIdx, gIdx)
		//fmt.Printf("%d routes found, length=%d\n", numRoutes, len(Routes))
		//_ = numRoutes

		// Score the routes list
		//Scores := scoreTheRoutes(AML, Routes, ScoreRouteFunc)

		//answer = Min(Scores)
	}


	// Construct the set of best routes, scoring them as you go
	routeToStart := []int{sIdx}
	bestStart := Best{routeToStart,0,false}
	initBestList := make([]Best, len(AML))
	initBestList[sIdx] = bestStart
	bl := constructBest(AML, initBestList, sIdx, true, ScoreRouteFunc)

	fmt.Printf("FromsCount: %d\n", FromsCount)
	answer = bl[gIdx].Score // answer is the score of the best route to gIdx

	return answer

	// Submissions:
	// (1) 95448 = Too High (a solution, not the best solution)
	// (2) 92440 = Too High still
	// (3) 92436 = Too High still
	// (4) 92432
}

func main() {
	//AllRules = make([][]int, 0)
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
