package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var intcodes []int
	//var intcodes_raw []int

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


	//copy(intcodes_raw, intcodes)

	set_initial_state(intcodes)
	step_intcodes(intcodes)

	fmt.Println(intcodes[0])
}

func set_initial_state(intcode []int) {
	intcode[1] = 12
	intcode[2] = 2
}

func step_intcodes(intcode []int) {
	fmt.Println(intcode)
	intcode_ops := len(intcode) / 3
	op_count := 0
	fmt.Println("Number of Operations expected:", intcode_ops)

	for i := 0; i < len(intcode); i++ {	
		opcode := intcode[i]
		op_count++
		if opcode == 1 {
			intcode[intcode[i+3]] = intcode[intcode[i+1]] + intcode[intcode[i+2]]
			i += 3
		} else if opcode == 2 {
			intcode[intcode[i+3]] = intcode[intcode[i+1]] * intcode[intcode[i+2]]
			i += 3
		} else if opcode == 99 {
			break
		} else {
			os.Exit(1)
		}
		//fmt.Println(intcode)
	}
	fmt.Println("Operations stepped:", op_count)
	return
}
