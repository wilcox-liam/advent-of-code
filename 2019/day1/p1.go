package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var total_fuel int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err == nil {
			total_fuel += calculate_fuel_requirement(mass)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total_fuel)
}

//Function to calculate the fuel requirement, given an objects mass
func calculate_fuel_requirement(mass int) int {
	var fuel int
	fuel = (mass / 3) - 2
	return fuel
}
