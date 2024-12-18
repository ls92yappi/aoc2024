package main

// See README.md for problem description

// Usage: problem1 input1.txt
// Currently templated off of my aoc2023/day23 code

import (
	"fmt"
	"strconv"
	"strings"
	. "github.com/ls92yappi/aoc" // problems with go proxy not being even remotely up-to-date
)

///////////////////////////////////////////////////////////


///////////////////////////////////////////////////////////

//ex
//Register A: 729
//Register B: 0
//Register C: 0
//
//Program: 0,1,5,4,3,0

//input
//Register A: 52884621
//Register B: 0
//Register C: 0
//
//Program: 2,4,1,3,7,5,4,7,0,3,1,5,5,5,3,0

// input translated
//
// bst A
// bxl 3
// cdv B
// bxc
// adv 3
// bxl 5
// out B
// jnz 0

// hard-coded based on input2.txt
// for {
//	//fmt.Printf("%#o", RegA)
// 	RegB = RegA % 8
// 	RegB ^= 3
// 	RegC = RegA / 1<<B
// 	RegB ^= RegC
// 	RegA /= 8
// 	RegB ^= 5
// 	Output = append(Output,RegB%8)
// 	if RegA==0 { break }
// }

// approach -- brute force a, 3 digits at a time
// a = ast *  1<<9 + 0o789       or whatever octal value you have thus far
// a = ast * 1<<18 + 0o456789    or whatever octal value you have thus far
// a = ast * 1<<27 + 0o123456789 or whatever octal value you have thus far

// literal operand = number
// combo operandType 0..3=number, 4=regA, 5=regB, 6=regC, 7=invalid

const (
	// auto IP increment=2, except jnz actual jump
	adv = 0 // division = regA / 2^combo -- integer division, -> regA
	bxl = 1 // bitwise XOR of regB and literal -> regB
	bst = 2 // mod8 = combo % 8 -> regB
	jnz = 3 // jump non zero = regA==0 nop, else IP=literal
	bxc = 4 // bitwise XOR regB and regC -> regB (ignores operand)
	out = 5 // mod8 = combo % 8, outputs that value comma separated
	bdv = 6 // division = regA / 2^combo -- int div, -> regB
	cdv = 7 // division = regA / 2^combo -- int div, -> regC
)

var IP int // instruction pointer
var RegA int
var RegB int
var RegC int
var Program []int
var Length int
var Output []int
var NumInst int

// Octal computer simulation
// opcode,operand pairs (8~input, 3~ex)

func Combo(operand int) int {
	val := 0
	switch {
	case operand <= 3:
		val = operand
	case operand==4:
		val = RegA
	case operand==5:
		val = RegB
	case operand==6:
		val = RegC
	default:
		fmt.Printf("Invalid Combo detected!\n")
	}
	return val
}

func IntsJoin(il []int, delim string) string {
	if len(il) == 0 {
		return ""
	}
	s := fmt.Sprintf("%d",il[0])
	for i := 1; i < len(il); i++ {
		s += fmt.Sprintf("%s%d", delim, il[i])
	}

	return s
}

func DisplayState(opcode, operand int) {
	sOp := ""
	switch opcode {
	case adv: sOp = "ADV"
	case bdv: sOp = "BDV"
	case cdv: sOp = "CDV"
	case bxl: sOp = "BXL"
	case bxc: sOp = "BXC"
	case bst: sOp = "BST"
	case out: sOp = "OUT"
	case jnz: sOp = "JNZ"
	}
	fmt.Printf("t=%d IP=%d Opcode=%s Operand=%d A={%d} B={%d} C={%d}\n", NumInst, IP, sOp, operand, RegA, RegB, RegC)
}

func processOpcode() bool {
	jumped := false
	val := 0
	use := 0

	opcode := Program[IP]
	operand := Program[IP+1]

	//DisplayState(opcode, operand)

	switch opcode {
	case adv:
		val = Combo(operand)
		if val > 62 {
			fmt.Printf("Power %d of 2 too big!\n", val)
		}
		use = 1 << val
		val = RegA/use
		RegA = val
		//fmt.Printf("ADV opd=%d   val=%d   use=%d   A={%d}\n", Combo(operand), val, use, RegA)
	case bdv:
		val = Combo(operand)
		if val > 62 {
			fmt.Printf("Power %d of 2 too big!\n", val)
		}
		use = 1 << val
		val = RegA/use
		RegB = val
	case cdv:
		val = Combo(operand)
		if val > 62 {
			fmt.Printf("Power %d of 2 too big!\n", val)
		}
		use = 1 << val
		val = RegA/use
		RegC = val
	case bxl:
		val = RegB ^ operand
		RegB = val
	case bxc:
		val = RegB ^ RegC
		RegB = val
	case bst:
		val = Combo(operand)
		use = val % 8
		RegB = use
	case out:
		val = Combo(operand)
		use = val % 8
		Output = append(Output, use)
		//fmt.Printf("%s\n", IntsJoin(Output,","))
	case jnz:
		if RegA != 0 {
			use = IP
			IP = operand
			if use != IP {
				jumped = true
			} else {
				fmt.Printf("Detected an infinite loop maybe with jnz\n")
			}
		}
	}

	return jumped
}

func runProgram(a,b,c int) string {
	// reset the computer state
	RegA = a
	RegB = b
	RegC = c
	IP = 0
	NumInst = 0
	Output = make([]int,0)

	// Length-1 instead of Length since all instructions read 2 values
	for IP < Length-1 {
		NumInst++
		jumped := processOpcode()
		IP += If(!jumped,2,0)


		// what we should do is display RegA, octal(RegA), numdigits correct for
		// each new digit we get, then use that value as an offset with a suitable
		// power of 2 increment to brute force the next set of digits

		// short-circuit on incorrect Output
		if len(Output) > 0 && Output[0] != 2 {
			return "wrong '2'"
		}
		if len(Output) > 1 && Output[1] != 4 {
			return "wrong '4'"
		}
		if len(Output) > 2 && Output[2] != 1 {
			return "wrong '1'"
		}
		if len(Output) > 3 && Output[3] != 3 {
			return "wrong '3'"
		}
		if len(Output) > 4 && Output[4] != 7 {
			return "wrong '7'"
		}
		if len(Output) > 5 && Output[5] != 5 {
			return "wrong '5'"
		}
		if len(Output) > 6 && Output[6] != 4 {
			return "wrong '4'"
		}
		if len(Output) > 7 && Output[7] != 7 {
			return "wrong '7'"
		}
		if len(Output) > 8 && Output[8] != 0 {
			return "wrong '0'"
		}
		if len(Output) > 9 && Output[9] != 3 {
			return "wrong '3'"
		}
		//1,5,5,5,3,0
		if NumInst > 10000 {
			fmt.Printf("\t *** Timed out after 10000 steps! ***\n")
			break
		}
	}

	s := IntsJoin(Output,",")
	return s
}

func processInput(fname string) int {
	// n is both the number of lines and the square grid size
	lines, _, _ := ReadWholeFile(fname)
	answer := 0

	// Read in the computer definition
	_,sRegA,_ := strings.Cut(lines[0], ": ")
	_,sRegB,_ := strings.Cut(lines[1], ": ")
	_,sRegC,_ := strings.Cut(lines[2], ": ")
	_,sProg,_ := strings.Cut(lines[4], ": ")

	RegA, _ = strconv.Atoi(sRegA)
	RegB, _ = strconv.Atoi(sRegB)
	RegC, _ = strconv.Atoi(sRegC)

	//Program := make([]int,0)
	Program,Length,_ = IntSlice(sProg,",")
	Output = make([]int,0)


	// guesstimating that bruteforcing it would take ~ 4000 hours ~= 5.5 months
	// MaxInt   9223372036854775807
	// #11          190615597431823
	jumpStart        := 55128850000 // brute forced to here
	for a := 1; a < 987654321098765; a++ {
		// jump start to a known bad value
		if a < jumpStart {
			a = jumpStart
		}
		if a%10000 == 0 {
			fmt.Printf("%d ",a/10000)
		}
		if a%1000000 == 0 {
			fmt.Printf("\n")
		}
		result := runProgram(a,0,0)
		if result == sProg {
			answer = a
			break
		}
		// test Part 2 against two2.txt instead of ex2.txt
		//if a==729 || (a>117438 && a<117444) {
		//	fmt.Printf("\nPass %d\nResult =%s\nProgram=%s\n", a, result, sProg)
		//}
	}
	fmt.Println("\n")

	return answer

	// Submissions:
	// (1) 0 = Correct
}

func main() {
	//AllRules = make([][]int, 0)
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
