package main

// See README.md for problem description
// Usage: problem1 input1.txt

import (
	"fmt"
	//"regexp"
	//"strconv"
	"sort"
	"strings"
	. "github.com/ls92yappi/aoc"
)

///////////////////////////////////////////////////////////

type Gate struct {
	A string // First Input Wire
	B string // Second Input Wire
	C string // Output Wire
	F string // Operation
	Done bool
}

type msb map[string]bool
type mssl map[string][]string
type msf map[string]Gate

var Dict msb
var Gates msf
var Follows mssl
var Permits mssl
var ZRem msb // list of Z-wires not processed yet

///////////////////////////////////////////////////////////

func (m msb)SortedKeys() []string {
	sl := make([]string, 0)
	for k,_ := range(m) {
		sl = append(sl, k)
	}
	sort.Strings(sl)
	return sl
}

func (g Gate)resolveGate(av, bv bool) bool {
	res := false
	switch g.F {
	case "AND":
		res = av && bv
	case "OR":
		res = av || bv
	case "XOR":
		res = av != bv
	default:
		panic(fmt.Sprintf("Invalid Gate: %+v\n", g))
	}
	g.Done = true
	return res
}

func processInput(fname string) int {
	lines, n, _ := ReadWholeFile(fname)
	n--
	answer := 0
	gatesMode := false

	Dict = make(msb, 0)
	Gates = make(msf, 0)
	Follows = make(mssl, 0)
	Permits = make(mssl, 0)
	ZRem = make(msb, 0)

	// gather base vals into Dict, followed by Gates
	for i := range(n) {
		line := lines[i]
		if len(line) < 2 {
			gatesMode = !gatesMode
			continue
		}
		if !gatesMode {
			wire, val, _ := strings.Cut(line, ": ")
			bVal := val=="1"
			Dict[wire] = bVal
		} else {
			sl := strings.Split(line, " ")
			wa := sl[0]
			wf := sl[1]
			wb := sl[2]
			wc := sl[4]
			fn := Gate{wa,wb,wc,wf,false}
			Gates[wc] = fn
			// update Follows[wc]
			Follows[wc] = []string{wa,wb}
			// update Permits[wa,wb]
			if _,ok := Permits[wa]; !ok {
				Permits[wa] = []string{wc}
			} else {
				Permits[wa] = append(Permits[wa], wc)
			}
			if _,ok := Permits[wb]; !ok {
				Permits[wb] = []string{wc}
			} else {
				Permits[wb] = append(Permits[wb], wc)
			}
			// do not create dict[wc] yet, will use its existence later
			// //Dict[wc] = false
			// we are interested in completion of z-wires
			if string(wc[0]) == "z" {
				ZRem[wc] = false
			}
		}
	} // for i

	// given: there are no loops

	// ex   : d10, g36, f36, p33, D45
	// mini : d6, g3, f3, p6, D9
	// input: d90, g222, f222, p266, D312
	fmt.Printf("Dict: %d, Gates: %d, Follows: %d, Permits: %d\n", len(Dict), len(Gates), len(Follows), len(Permits))

	pass := 0
	for len(ZRem) > 0 {
		pass++
		gr := make([]string, 0) // resolved gates list
		// compute resolvable Gates
		for i, g := range(Gates) {
			a := g.A
			b := g.B
			c := g.C
			av, ok1 := Dict[a]
			bv, ok2 := Dict[b]
			// can we resolve this gate currently
			if !ok1 || !ok2 {
				continue
			}
			res := g.resolveGate(av, bv)
			Dict[c] = res
			// add the gate to the resolved list for removal
			gr = append(gr, c)
			if string(c[0]) == "z" {
				delete(ZRem, c)
			}
			_ = i
		}

		if len(gr) == 0 {
			panic(fmt.Sprintf("No Gates were resolved on pass %d!\n", pass))
		}

		// trim resolved Gates
		for _,s := range(gr) {
			delete(Gates, s)
		}

		if len(Gates) < 1 {
			break
		}
	}
	fmt.Printf("Resolved in %d passes\n", pass)

	kl := Dict.SortedKeys()
	zi := 0
	for _,k := range(kl) {
		if string(k[0]) != "z" {
			continue
		}
		answer += If(Dict[k],1,0) << zi
		zi++
	}

	return answer

	// Submissions:
	// (1) 48063513640678 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
