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

	for noun := 0; noun <= 99; noun++ {	
		for verb := 0; verb <= 99 ; verb++ {			
			intcodes_trial := make([]int, len(intcodes))
			copy(intcodes_trial, intcodes)
			set_initial_state(intcodes_trial, noun, verb)
			step_intcodes(intcodes_trial)

			if intcodes_trial[0] == 19690720 {
				fmt.Println("Suceess:", 100 * noun + verb)
				os.Exit(0)
			}
		}
	}

	fmt.Println("You suck")
}

func set_initial_state(intcode []int, noun int, verb int) {
	intcode[1] = noun
	intcode[2] = verb
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
