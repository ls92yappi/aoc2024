package main

// See README.md for problem description

// Usage: problem01 input01.txt
// Currently templated off of my aoc2023/day23 code

/* Analysis: ***************************
Forks @ b(5,3) a(3,11) c(13,5) d(13,13)
Merge @ A(11,21) B(19,13) C(19,19)

Start(α)->b
b->a|c
c->B|d
d->A|B
B->C
C->Final(β)
A->C
a->d|A

Label lengths of each segment
|α,a| = 15 ignore from spot
|a,b| = 
|C,β| = 5

|K,9| is end run for input23

Observe outer walls on all sides, except Entry and Exit.
Observe, Junstions are NOT dense. Never more than 2 in any 5x5 window.
Trivia from analysis:
Trait ex23  input23
Size    23      141
MustR   10       58
MustD   12       60
Path   191     9246
Tree   316    10517
Forks    4       25
Merge    3        9
TJunc    7       34
************************************* */
/* Computes: ************************
For part 2, we may ignore the directionality arrows.
Everything remains the same, except ComputeAllRoutes().
It must now include a "have we visited this node before" to prevent cycles.

With this restriction removed, the longest route for ex23 should be 154.
************************************* */

/* Example output: ******************
154
************************************* */

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

///////////////////////////////////////////////////////////


const (
	Path  uint8 = '.' //' ' //'.' // 9246
	Wall  uint8 = '#' //'█' //'#' // 10517 
	MustR uint8 = '>' //'▸' //'>' // 58
	MustD uint8 = 'v' //'▾' //'v' // 60
	BegCh uint8 = '0' //'α'
	EndCh uint8 = '9' //'β'
	//SteepR uint8 = '>' // none exist of either of these two
	//SteepU uint8 = '^'
)

const ForksNames string = "abcdefghijklmnopqrstuvwxyz"
const MergeNames string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var Start PoI
var Final PoI

var Size int // 141 or 23

var MaxSteps int // 94

const (
	// Point-of-Interest Types
	StartPoI int = 1+iota // 1
	ForksBoth             // 2
	ForksTop              // 3
	ForksLeft             // 4
	MergeDown             // 5
	MergeRight            // 6
	FinalPoI              // 7
)

type PoI struct {
	// Known attributes at identification time
	N string // Name/label
	R int // Row
	C int // Column
	T int // PoI type (one of StartPoI .. FinalPoI)
	// Stats computed by FindPath()
	RN string // Right Name
	DN string // Down Name
	RL int    // Right Length
	DL int    // Down Length
	// New Stats computed by FindPath()
	LN string // Right Name
	UN string // Down Name
	LL int    // Right Length
	UL int    // Down Length
}
type PoIMap map[string]PoI
type CoordsMap map[int]string // map[1000*PoI.R+PoI.C]PoI.N


type Row []uint8
type Maze []Row

var P  Maze      // Puzzle Maze
var PM PoIMap    // PoI map
var CM CoordsMap // Coordinates of PoI map entries

var Routes map[string]int

//type Direction int
const (
	North int = 1
	East  int = 2
	South int = 4
	West  int = 8
)


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
	PM = make(PoIMap,0)
	CM = make(CoordsMap,0)
	Routes = make(map[string]int,0)
}


///////////////////////////////////////////////////////////

func Max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func (p *PoI) UpdateFieldsDown(dest string, length int) {
	p.DN = dest
	p.DL = length
	// Update the global variable here
	PM[p.N] = *p
}

func (p *PoI) UpdateFieldsRight(dest string, length int) {
	p.RN = dest
	p.RL = length
	// Update the global variable here
	PM[p.N] = *p
}


func (p *PoI) UpdateFieldsUp(dest string, length int) {
	p.UN = dest
	p.UL = length
	// Update the global variable here
	PM[p.N] = *p
}

func (p *PoI) UpdateFieldsLeft(dest string, length int) {
	p.LN = dest
	p.LL = length
	// Update the global variable here
	PM[p.N] = *p
}

func (p PoIMap) String() string {
	sb := strings.Builder{}

	keys := make([]string, 0, len(p))

	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := p[k]
	//for k, v := range p {
		sb.WriteString(fmt.Sprintf("%s: ",k))
		if v.RL > 0 { sb.WriteString(fmt.Sprintf("R:%s%d ",v.RN,v.RL)) }
		if v.DL > 0 { sb.WriteString(fmt.Sprintf("D:%s%d ",v.DN,v.DL)) }
		if v.LL > 0 { sb.WriteString(fmt.Sprintf("L:%s%d ",v.LN,v.LL)) }
		if v.UL > 0 { sb.WriteString(fmt.Sprintf("U:%s%d ",v.UN,v.UL)) }
		sb.WriteString("\n")
		//sb.WriteString(fmt.Sprintf("%s: R:%s%d D:%s%d L:%s%d U:%s%d\n",k,v.RN,v.RL,v.DN,v.DL,v.LN,v.LL,v.UN,v.UL))
	}
	return sb.String()	
}




// Starting moving in direction dir, provide row and column offsets, and new direction
func TurnLeft(dir int) (int, int, int) {
	var rOff, cOff int
	newDir := dir
	switch dir {
		case North: cOff = -1; newDir = West
		case South: cOff =  1; newDir = East
		case West:  rOff =  1; newDir = South
		case East:  rOff = -1; newDir = North
		default:
			die("Impossible direction to TurnLeft() from\n")
			return 0,0,dir
	}
	return rOff,cOff,newDir
}

// Starting moving in direction dir, provide row and column offsets, and new direction
func TurnRight(dir int) (int, int, int) {
	var rOff, cOff int
	newDir := dir
	switch dir {
		case North: cOff =  1; newDir = East
		case South: cOff = -1; newDir = West
		case West:  rOff = -1; newDir = North
		case East:  rOff =  1; newDir = South
		default:
			die("Impossible direction to TurnLeft() from\n")
			return 0,0,dir
	}
	return rOff,cOff,newDir
}

// Starting moving in direction dir, provide row and column offsets, and new direction
func Straight(dir int) (int, int, int) {
	var rOff, cOff int
	switch dir {
		case North: rOff = -1
		case South: rOff =  1
		case West:  cOff = -1
		case East:  cOff =  1
		default:
			die("Impossible direction to TurnLeft() from\n")
			return 0,0,dir
	}
	return rOff,cOff,dir
}


func FindPointsOfInterest() int {
	// Known Start and final positions
	Start = PoI{N:"0",R:0,C:1,T:StartPoI}
	Final = PoI{N:"9",R:Size-1,C:Size-2,T:FinalPoI}

	// update both the PoI Map and ...
	PM["0"] = Start
	PM["9"] = Final
	// ... also the coordinates map
	CM[1] = "0"
	CM[1000*(Size-1)+Size-2] = "9"

	var numPoI int = 2
	var numForks int = 0
	var numMerge int = 0

	for r := 1; r < Size-1; r++ {
		for c := 1; c < Size-1; c++ {
			here  := P[r][c]
			above := P[r-1][c]
			below := P[r+1][c]
			leftt := P[r][c-1]
			right := P[r][c+1]

			//debug(fmt.Sprintf("\tPoI check ...  at (%d,%d)\n",r,c))

			var poiType int
			var label string
			switch {
				case here!=Path:
					// Not a Point of Interest
					continue

				case above==Path || below==Path || leftt==Path || right==Path:
					// Not a Point of Interest
					continue

				// #v#
				// >.>
				// #v#
				case above==MustD && below==MustD && leftt==MustR && right==MustR:
					poiType = ForksBoth
					label = string(ForksNames[numForks])
					numForks++
					//debug(fmt.Sprintf("ForksBoth '%s' at (%d,%d)\n",label,r,c))

				// #v#
				// #.>
				// #v#
				case above==MustD && below==MustD && leftt==Wall && right==MustR:
					poiType = ForksTop
					label = string(ForksNames[numForks])
					numForks++
					//debug(fmt.Sprintf("ForksTop '%s' at (%d,%d)\n",label,r,c))

				// ###
				// >.>
				// #v#
				case above==Wall && below==MustD && leftt==MustR && right==MustR:
					poiType = ForksLeft
					label = string(ForksNames[numForks])
					numForks++
					//debug(fmt.Sprintf("ForksLeft '%s' at (%d,%d)\n",label,r,c))

				// #v#
				// >.>
				// ###
				case above==MustD && below==Wall && leftt==MustR && right==MustR:
					poiType = MergeRight
					label = string(MergeNames[numMerge])
					numMerge++
					//debug(fmt.Sprintf("MergeRight '%s' at (%d,%d)\n",label,r,c))

				// #v#
				// >.#
				// #v#	
				case above==MustD && below==MustD && leftt==MustR && right==Wall:
					poiType = MergeDown
					label = string(MergeNames[numMerge])
					numMerge++
					//debug(fmt.Sprintf("MergeDown '%s' at (%d,%d)\n",label,r,c))

				default:
					// Not a Point of Interest
					die(fmt.Sprintf("IMPOSSIBLE PoI at (%d,%d)\n",r,c))
					continue
			}
			// Add the found, categorized and labeled Point of Interest
			var poi PoI = PoI{N:label,R:r,C:c,T:poiType}
			PM[label] = poi
			numPoI++
			// Add it to the Coordinates Map also, for easy finding later
			CM[1000*r+c] = label
		}
	}

	return numPoI
}


// Returns the label of the PoI at the end of this path along with
// returns the length of the path just followed
func FollowPath(label string, dir int) (string, int, int) {
	var numSteps int
	var p PoI = PM[label]
	var poiDir int
	r := p.R
	c := p.C

	// Proceed straight first
	rOffset,cOffset,prev := Straight(dir)
	r += rOffset
	c += cOffset

	// Update step count
	numSteps++
	coords := 1000*r+c
	dest, ok := CM[coords]
	for !ok {
		// move in appropriate direction
		// Try straight first ...
		rOffset,cOffset,dir = Straight(prev)
		coords = 1000*(r+rOffset)+(c+cOffset)
		dest, ok = CM[coords]
		if ok {
			// Found a Point of Interest
			poiDir = prev
			numSteps++
			break
		} 
		if P[r+rOffset][c+cOffset] != Wall {
			// Continue Straight
			numSteps++
			r += rOffset
			c += cOffset
			continue
		}

		// Hit a wall, try left ...
		rOffset,cOffset,dir = TurnLeft(prev)
		coords = 1000*(r+rOffset)+(c+cOffset)
		dest, ok = CM[coords]
		if ok {
			// Found a Point of Interest
			poiDir = prev
			numSteps++
			break
		} 
		if P[r+rOffset][c+cOffset] != Wall {
			// Continue Left
			numSteps++
			r += rOffset
			c += cOffset
			prev = dir
			continue
		}


		// Hit a wall, turn right ...
		rOffset,cOffset,dir = TurnRight(prev)
		coords = 1000*(r+rOffset)+(c+cOffset)
		dest, ok = CM[coords]
		if ok {
			// Found a Point of Interest
			poiDir = prev
			numSteps++
			break
		} 
		if P[r+rOffset][c+cOffset] != Wall {
			// Continue Left
			numSteps++
			r += rOffset
			c += cOffset
			prev = dir
			continue
		}
		die(fmt.Sprintf("FollowPath(%s) at hopeless location (%d,%d) from %d.\n",label,r,c,prev))
	}
	// at the PoI dest
	return dest, numSteps, poiDir
}


func (p *PoI) UpdateFieldsAsNeeded(dest string, length int, dir int) {
	switch dir {
		case South:
			p.UpdateFieldsUp(dest,length)
		case East:
			p.UpdateFieldsLeft(dest,length)
		// The below two cases should never occur
		//case North:
		//	p.UpdateFieldsDown(dest,length)
		//case East:
		//	p.UpdateFieldsRight(dest,length)
		default:
			die(fmt.Sprintf("UpdateFieldsAsNeeded() %+v -> %s ~ %d Going %d\n",p, dest, length, dir))
	}
}

func FindPaths(label string) int {
	var numC int // numConnectionsFound
	var p PoI = PM[label]
	var d PoI
	//var dest string
	//var length int

	// Short circuit nodes that have already been visited
	if len(p.RN)>0 || len(p.DN)>0 {
		return 0
	}

	var dest string
	var length int
	var dir int

	// DONE: Add UN/UL and LN/LL as needed
	// DONE: Update FollowPath() to also return the final direction of the path

	// Perform a search of the PoI Map, filling in RN,RL,DN,DL as needed, and for children
	switch p.T {
		case StartPoI:
			// the original way happens to work for StartPoI
			dest, length, dir = FollowPath(label, South)
			p.UpdateFieldsDown(dest, length)
			d = PM[p.DN]
			d.UpdateFieldsUp(label, length)
			return 1 + FindPaths(PM[label].DN)

		case ForksBoth:
			dest, length, dir = FollowPath(label, South)
			p.UpdateFieldsDown(dest, length)
			d = PM[p.DN]
			d.UpdateFieldsAsNeeded(label, length, dir)

			dest, length, dir = FollowPath(label, East)
			p.UpdateFieldsRight(dest, length)
			d = PM[p.RN]
			d.UpdateFieldsAsNeeded(label, length, dir)

			numC += 1 + FindPaths(PM[label].DN)
			numC += 1 + FindPaths(PM[label].RN)

		case ForksTop:
			dest, length, dir = FollowPath(label, South)
			p.UpdateFieldsDown(dest, length)
			d = PM[p.DN]
			d.UpdateFieldsAsNeeded(label, length, dir)

			dest, length, dir = FollowPath(label, East)
			p.UpdateFieldsRight(dest, length)
			d = PM[p.RN]
			d.UpdateFieldsAsNeeded(label, length, dir)

			numC += 1 + FindPaths(PM[label].DN)
			numC += 1 + FindPaths(PM[label].RN)

		case ForksLeft:
			dest, length, dir = FollowPath(label, South)
			p.UpdateFieldsDown(dest, length)
			d = PM[p.DN]
			d.UpdateFieldsAsNeeded(label, length, dir)

			dest, length, dir = FollowPath(label, East)
			p.UpdateFieldsRight(dest, length)
			d = PM[p.RN]
			d.UpdateFieldsAsNeeded(label, length, dir)

			numC += 1 + FindPaths(PM[label].DN)
			numC += 1 + FindPaths(PM[label].RN)

		case MergeDown:
			dest, length, dir = FollowPath(label, South)
			p.UpdateFieldsDown(dest, length)
			d = PM[p.DN]
			//debugf("%s %d -> %s %d [%d]\n",label, East, dest, dir, length)
			d.UpdateFieldsAsNeeded(label, length, dir)
			return 1 + FindPaths(PM[label].DN)

		case MergeRight:
			dest, length, dir = FollowPath(label, East)
			p.UpdateFieldsRight(dest, length)
			d = PM[p.RN]
			//debugf("%s %d -> %s %d [%d]\n",label, East, dest, dir, length)
			d.UpdateFieldsAsNeeded(label, length, dir)
			return 1 + FindPaths(PM[label].RN)

		case FinalPoI:
			return 0

		default:
			die("Impossible to FindPaths() for Label '"+label+"'")
			return 0
	}

	_ = d
	_ = dir
	return numC
}

// Returns the longest route found and populates Routes[] map as it goes
func ComputeAllRoutes(from string, curName string, curVal int) int {
	var maxRoute int
	var rlThis int
	var p PoI = PM[from]

	// if we've reached the Final goal
	if p.T == FinalPoI {
		// Update the Routes list
		Routes[curName] = curVal
		// return this route length
		debug(fmt.Sprintf("Route '%s' = %d\n", curName, curVal))
		return curVal
	}

	// Cancel out this route if we have hit a loop
	stops := strings.Split(curName, "~")
	if len(stops)>2 {
		justAdded := stops[len(stops)-1]
		for i := 0; i < len(stops)-1; i++ {
			if stops[i] == justAdded {
				// cycle detected. Abandon this route
				return 0
			}
		}
	}
	

	// attempt to follow Down route
	if p.DL > 0 {
		rlThis = ComputeAllRoutes(p.DN, curName+"~"+p.DN, curVal+p.DL)
		maxRoute = rlThis
	}

	// attempt to follow Right route
	if p.RL > 0 {
		rlThis = ComputeAllRoutes(p.RN, curName+"~"+p.RN, curVal+p.RL)
		maxRoute = Max(maxRoute, rlThis)
	}


	// DONE: Add UN/UL and LN/LL Routes in here
	// attempt to follow Up route
	if p.UL > 0 {
		rlThis = ComputeAllRoutes(p.UN, curName+"~"+p.UN, curVal+p.UL)
		maxRoute = Max(maxRoute, rlThis)
	}

	// attempt to follow Left route
	if p.LL > 0 {
		rlThis = ComputeAllRoutes(p.LN, curName+"~"+p.LN, curVal+p.LL)
		maxRoute = Max(maxRoute, rlThis)
	}

	return maxRoute
}


// Side Effect: Builds P Maze
func ParseLine(linenum int, line string) {
	var row Row = make(Row,Size)
	for i := 0; i < Size; i++ {
		row[i] = line[i]
	}
	P = append(P, row)
	return
}

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


	// Fill in P using ParseLine()
	for i:=0;i<Size;i++ {
		line := lines[i]
		if len(line) < 1 {
			continue
		}
		ParseLine(i,line)
	}

	// PrettifyMap() // Optional visualization feature
	numPoI := FindPointsOfInterest()
	numCnxn := FindPaths("0")
	debug(fmt.Sprintf("Found %d connections for %d Points of Interest\n", numCnxn, len(PM)))


	debug("-----\n")
	debug(PM.String())
	debug("-----\n")

	// Goal is to find the longest possible route without revisiting any location.
	// Returns length of the longest Route it found, and builds the Routes[] map
	score := ComputeAllRoutes("0", "0", 0)

	//FindPaths() αa=15,ab=22,bd=24,dA=18,AC=10,Cβ=5 Total=94 (so exclude start spot)
	//ComputeAllRoutesAndLengths() = 94, 90, 86, 82, 82, 74
	//ReportLongestLength() = 94 α,a▸,b▾,d▸A,C,β

	// We could debug the Routes[] map to see all the routes and their respective lengths

	// Submissions:
	// (1) xxxx = Correct

	_ = numCnxn
	_ = numPoI
	return score
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
