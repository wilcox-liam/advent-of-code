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

	set_initial_state(intcodes)
	step_intcodes(intcodes)

	fmt.Println(intcodes[0])
}

func set_initial_state(intcode []int) {
	intcode[1] = 12
	intcode[2] = 2
}

func step_intcodes(intcode []int) {
	for i := 0; i < len(intcode); i++ {	
		opcode := intcode[i]
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
	}
	return
}
