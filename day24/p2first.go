package main

// See README.md for problem description
// Usage: problem1 input1.txt

import (
	"fmt"
	//"regexp"
	//"strconv"
	"maps"
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
type mss map[string]string

var Dict msb
var Gates msf
var Follows mssl
var Permits mssl
var ZRem msb // list of Z-wires not processed yet
var Derivs mss // derivations of bits
var Labels mss // schematic labeling of the gates
var Anomalies []string

///////////////////////////////////////////////////////////
//
// Things we know:
//
//   Easy wins:
//     for n<45, a XOR b = z_n. Any z_n gate that is not
//     a XOR is an anomaly
//
//   Looking at the derivations, and remembering digital
//   electronics back in the day, we know z00=x00^y00 is good
//   We also know that x00&y00 is the carry bit (rnv for us)
//   This can be labeled as C00. This is a half-adder.
//
//   Looking at the components of z01, we see a full-adder.
//   We Label some components:
//     Cnn is a Carry bit
//     Ann is a And(x,y) bit x&y
//     Dnn is a Diff(x,y) bit x^y
//     Fnn is a Forced bit D&C_prev
//
//   x01 -- 
//         XOR -- D01 (rbr) --
//   y01 --                  |
//                            XOR -- z01
//   C00 ---------------------
//
//   x01 -- 
//         AND -- A01 (hqp) --
//   y01 --                  |
//                            OR -- C01
//   C00 --                  |
//         AND -- F01 (nbf) --
//   D01 --
//
//
//   Look for this structure for all bits from z01..z44
//
//   z45 should just be C44, which is D44|A44
//
///////////////////////////////////////////////////////////

func ShowDerivations(kl []string, prefix string) {
	for _,k := range(kl) {
		if string(k[0]) != prefix {
			continue
		}
		fmt.Printf("%s = %s\n", k, Derivs[k])
	}
	fmt.Println()
}

func ComputeValue(kl []string, prefix string) int {
	i := 0
	res := 0
	for _,k := range(kl) {
		if string(k[0]) != prefix {
			continue
		}
		res += If(Dict[k],1,0) << i
		i++
	}
	return res
}

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

func processInput(fname string) string {
	lines, n, _ := ReadWholeFile(fname)
	n--
	answer := ""
	gatesMode := false

	Dict = make(msb, 0)
	Gates = make(msf, 0)
	Saved := make(msf, len(Gates))
	Follows = make(mssl, 0)
	Permits = make(mssl, 0)
	ZRem = make(msb, 0)
	Derivs = make(mss, 0)
	Labels = make(mss, 0)
	Anomalies = make([]string, 0)

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
			Derivs[wire] = wire
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

	// make a backup copy of the gates
	maps.Copy(Saved, Gates)
	//copy(Gates, Saved) // slices copy(src,dst) maps.Copy(dst,src)

	// Let's try derivation tracking first
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
			// fully simplified to xNN, yNN primitives:
			//Derivs[c] = fmt.Sprintf("[(%s) %s (%s)]", Derivs[a], g.F, Derivs[b])
			// latest only derivations
			Derivs[c] = fmt.Sprintf("[(%s) %s (%s)]", a, g.F, b)
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
	// fmt.Printf("Resolved in %d passes\n", pass)

	// some research
	if false {
		kl := Dict.SortedKeys()
		xVal := ComputeValue(kl, "x")
		yVal := ComputeValue(kl, "y")
		zVal := ComputeValue(kl, "z")
		//answer = zVal // from Part 1
		sumXY := xVal+yVal
		fmt.Printf("X: %d Y: %d Z: %d Sum: %d\n", xVal, yVal, zVal, sumXY)
		diff := zVal^sumXY
		fmt.Printf("      : 4    43        32        21        1          \n")
		fmt.Printf("      : 5432109876543210987654312098765432109876543210\n")
		fmt.Printf("      : ----------------------------------------------\n")
		fmt.Printf("X-bits: %46b\n", xVal)
		fmt.Printf("Y-bits: %46b\n", yVal)
		fmt.Printf("S-bits: %46b\n", sumXY)
		fmt.Printf("Z-bits: %46b\n", zVal)
		fmt.Printf("D-bits: %46b\n", diff)
		fmt.Printf("      : ----------------------------------------------\n")
		// inspecting X and Y's bits, the low 0..13 bits are carry-safe, except bits 1 and 4
		//ShowDerivations(kl, "x")
		//ShowDerivations(kl, "y")
		//ShowDerivations(kl, "z")
		// Note: all gates in part 1 WERE resolved, none were not needed
		//fmt.Printf("Trivia: %d gates were not resolved (may not have needed to be)\n", len(Gates))
		//fmt.Printf("AllKeys were: %+v\n", kl)
		//fmt.Printf("\n\n\n%+v\n", Follows)
		//fmt.Printf("\n\n\n%+v\n", Permits)
	}

	// restore our backup of the gates
	maps.Copy(Gates, Saved)

	// Find incorrectly structured Z-gates -- Easy wins
	kl := Dict.SortedKeys()
	for _,k := range(kl) {
		// only look at z-Gates
		if string(k[0]) != "z" {
			continue
		}
		if Gates[k].F != "XOR" && k != "z45" {
			Anomalies = append(Anomalies, k) 
		}
	}

	// Find corresponding pair


	Labels["C00"] = "rnv" // observed during analysis

	// Label all A and D gates
	for _,k := range(kl) {
		// ignore all terminal-Gates
		kPrefix := string(k[0])
		if kPrefix == "z"  || kPrefix == "x" || kPrefix == "y" {
			continue
		}
		g := Gates[k]
		// note that all gates that contain an "x#" input also has a "y#" input with same #
		//fmt.Printf("%s: %+v\n", k, g)
		prefix := string(g.A[0])
		// only interested in x# op y# gates on this pass
		if prefix != "x" && prefix != "y" {
			continue
		}
		suffix := string(g.A[1:])
		if g.F == "AND" {
			aLab := "A"+suffix
			// verified no duplicate A's detected
			//if len(Labels[aLab]) > 0 {
			//	fmt.Printf("Duplicate %s: %s or %s\n", aLab, Labels[aLab], k)
			//	Labels[aLab] = k
			//} else {
				Labels[aLab] = k
			//}
			continue
		}
		if g.F == "XOR" {
			dLab := "D"+suffix
			// verified no duplicate D's detected
			//if len(Labels[dLab]) > 0 {
			//	fmt.Printf("Duplicate %s: %s or %s\n", dLab, Labels[dLab], k)
			//	Labels[dLab] = k
			//} else {
				Labels[dLab] = k
			//}
			continue
		}
		_ = suffix
		// invalid "OR" found
		Anomalies = append(Anomalies, k)
	}

	// F and C gates pass
	prevSuffix := "00"
	OutsideFCsearch:
	for i := range(45) {
		// known half-adder in spot 00
		if i==0 { continue }
		suffix := fmt.Sprintf("%02d", i)
		searchC, cok := Labels["C"+prevSuffix]
		searchD, dok := Labels["D"+suffix]
		if !cok || !dok {
			fmt.Printf("Missing either 'C' or 'D' value when trying construct 'F' for %s\n",suffix)
			continue
		}
		for _,k := range(kl) {
			fLab := "F"+suffix
			kPrefix := string(k[0])
			if kPrefix == "z"  || kPrefix == "x" || kPrefix == "y" {
				continue
			}
			g := Gates[k]
			//  && g.F=="AND"
			if (g.A==searchC && g.B==searchD) || (g.A==searchD && g.B==searchC) {
				// Found the forced F gate
				if len(Labels[fLab]) > 0 {
					fmt.Printf("Multiple F's found: %s %s or %s for (%s,%s)\n", fLab, Labels[fLab], k, searchC, searchD)
					Anomalies = append(Anomalies, k)
					break
					//break OutsideFCsearch
				} else {
					Labels[fLab] = k
				}
				// could make faster with a break here after verifying no duplicates exist
			}
		}

		searchF, fok := Labels["F"+suffix]
		searchA, aok := Labels["A"+suffix]
		if !fok || !aok {
			fmt.Printf("Missing either 'F' or 'A' value when trying construct 'C' for %s\n",suffix)
			continue
		}
		for _,k := range(kl) {
			cLab := "C"+suffix
			kPrefix := string(k[0])
			if kPrefix == "z"  || kPrefix == "x" || kPrefix == "y" {
				continue
			}
			g := Gates[k]
			if (g.A==searchF && g.B==searchA) || (g.A==searchA && g.B==searchF) {
				// Found the forced F gate
				if len(Labels[cLab]) > 0 {
					fmt.Printf("Multiple C's found: %s %s or %s for (%s,%s)\n", cLab, Labels[cLab], k, searchF, searchA)
					break OutsideFCsearch
				}
				Labels[cLab] = k
				// could make faster with a break here after verifying no duplicates exist
			}
		}
		// ...
		prevSuffix = suffix
	}


	sort.Strings(Anomalies)
	str := strings.Join(Anomalies, ",")
	answer = str
	return answer

	// Submissions:
	// (1) 48063513640678 = Correct
}

func main() {
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
