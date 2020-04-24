package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var intcodes []int
	//intcodes = append(intcodes, 0)

	dat, err := ioutil.ReadFile("test.txt")
	check(err)
	opcodes := string(dat)

	tokens := strings.Split(opcodes, ",")

	for _, token := range tokens {
		opcode, err := strconv.Atoi(token)
		if err == nil {
			intcodes = append(intcodes, opcode)
		}
	}

	step_intcodes(intcodes)
	fmt.Println("THE END")
}

func intcodeAdd(intcode []int, parm1, parm2, parm3 int) []int {
	intcode[parm3] = parm1 + parm2
	return intcode
}

func intcodeMultiply(intcode []int, parm1, parm2, parm3 int) []int {
	intcode[parm3] = parm1 * parm2
	return intcode
}

func intcodeWrite(parm1 int) {
	fmt.Println("OUTPUT:", parm1)
}

func intcodeRead(intcode []int, pos int) []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("INPUT? ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	intInput, _ := strconv.Atoi(input)
	intcode[intcode[pos + 1]] = intInput
	return intcode
}

func intcodeJumpIfTrue(pos, parm1, parm2 int) int {
	if parm1 != 0 {
		return parm2 - 1
	}
	return pos + 2
}

func intcodeJumpIfFalse(pos, parm1, parm2 int) int {
	if parm1 == 0 {
		return parm2 - 1	
	}
	return pos + 2
}

func intcodeLessThan(intcode []int, parm1, parm2, parm3 int) []int {
	if parm1 < parm2 {
		intcode[parm3] = 1		
	} else {
		intcode[parm3] = 0	
	}
	return intcode
}

func intcodeEquals(intcode []int, parm1, parm2, parm3 int) []int {
	if parm1 == parm2 {
		intcode[parm3] = 1
	} else {
		intcode[parm3] = 0		
	}
	return intcode
}

func step_intcodes(intcode []int) {
	fmt.Println(len(intcode))
	for i := 0; i < len(intcode); i++ {
		var parm1, parm2, parm3 int
		opcode := intcode[i]
		opcodeString := padOpCode(opcode)
		instruction := getInstruction(opcodeString)

		if instruction == 99 {
			break;
		} else {
			parm1, parm2, parm3 = getParms(intcode, opcodeString, instruction, i)
		}

		fmt.Println(instruction, parm1, parm2, parm3)
		fmt.Println(intcode)

		switch instruction {
		case 1:
			intcode = intcodeAdd(intcode, parm1, parm2, parm3)
			i += 3
		case 2:
			intcode = intcodeMultiply(intcode, parm1, parm2, parm3)
			i += 3
		case 3:
			intcode = intcodeRead(intcode, i)
			i += 1
		case 4:
			intcodeWrite(parm1)
			i += 1
		case 5:
			i = intcodeJumpIfTrue(i, parm1, parm2)
		case 6:
			i = intcodeJumpIfFalse(i, parm1, parm2)	
		case 7:
			intcode = intcodeLessThan(intcode, parm1, parm2, parm3)
			i+= 3
		case 8:
			intcode = intcodeEquals(intcode, parm1, parm2, parm3)	
			i += 3
		default:
			fmt.Println("Invalid intcode")
			os.Exit(1)
		}
	}
	return
}

func padOpCode(opcode int) string {
	opcodePadded := fmt.Sprintf("%05d", opcode)
	return opcodePadded
}

func getInstruction(opcode string) int {
	instruction, _ := strconv.Atoi(opcode[len(opcode)-2:])
	return instruction
}

func getParmModes(opcode string) (int, int, int) {
	parm1, _ := strconv.Atoi(opcode[2:3])
	parm2, _ := strconv.Atoi(opcode[1:2])
	parm3, _ := strconv.Atoi(opcode[:1])
	return parm1, parm2, parm3
}

func getParms(intcode []int, opcode string, instruction int, i int) (int, int, int) {
	parm1Mode, parm2Mode, parm3Mode := getParmModes(opcode)
	var parm1, parm2, parm3 int
	switch parm1Mode {
	case 0:
		parm1 = intcode[intcode[i+1]]
	case 1:
		parm1 = intcode[i+1]
	default:

	}
	if instruction == 1 || instruction == 2  || instruction >= 5 {
		switch parm2Mode {
		case 0:
			parm2 = intcode[intcode[i+2]]
		case 1:
			parm2 = intcode[i+2]
		default:
			os.Exit(1)
		}
	}
	if instruction == 1 || instruction == 2  || instruction >= 7  {
		switch parm3Mode {
		case 0:
			parm3 = intcode[i+3] 
		default:
			os.Exit(1)
		}
	}
	return parm1, parm2, parm3
}
