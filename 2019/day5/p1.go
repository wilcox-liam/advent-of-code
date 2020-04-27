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

	dat, err := ioutil.ReadFile("input.txt")
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
	intcode[intcode[pos+1]] = intInput
	return intcode
}

func step_intcodes(intcode []int) {
	for i := 0; i < len(intcode); i++ {
		var parm1, parm2, parm3 int
		opcode := intcode[i]
		opcodeString := padOpCode(opcode)
		instruction := getInstruction(opcodeString)

		if instruction == 99 {
			break
		} else {
			parm1, parm2, parm3 = getParms(intcode, opcodeString, instruction, i)
		}

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
		case 99:
			break
		default:
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
	if instruction == 1 || instruction == 2 {
		switch parm2Mode {
		case 0:
			parm2 = intcode[intcode[i+2]]
		case 1:
			parm2 = intcode[i+2]
		default:
			os.Exit(1)
		}

		switch parm3Mode {
		case 0:
			parm3 = intcode[i+3]
		default:
			os.Exit(1)
		}
	}
	return parm1, parm2, parm3
}
