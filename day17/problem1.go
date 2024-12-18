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

	DisplayState(opcode, operand)

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
		fmt.Printf("%s\n", IntsJoin(Output,","))
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

func processInput(fname string) string {
	// n is both the number of lines and the square grid size
	lines, _, _ := ReadWholeFile(fname)
	answer := ""

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

	// Length-1 instead of Length since all instructions read 2 values
	for IP < Length-1 {
		NumInst++
		jumped := processOpcode()
		IP += If(!jumped,2,0)

		if NumInst > 80 {
			fmt.Printf("\t *** Timed out after 80 steps! ***\n")
			break
		}
	}

	answer = IntsJoin(Output,",")
	return answer

	// Submissions:
	// (1) "1,3,5,1,7,2,5,1,6" = Correct
}

func main() {
	//AllRules = make([][]int, 0)
	inputFileName := InputFileName()
	output := processInput(inputFileName)
	fmt.Println(output)
}
