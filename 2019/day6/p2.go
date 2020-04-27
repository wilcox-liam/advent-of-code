package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	galaxy := make(map[string]string)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		planetas := strings.Split(line, ")")
		galaxy[planetas[1]] = planetas[0]
	}

	distanceToYou := make(map[string] int)
	o := galaxy["YOU"]
	distance := 0
	ok := false
	for {
		distanceToYou[o] = distance
		o, ok = galaxy[o]
		if ok {
			distance++
		} else {
			break
		}
	}
	
	o = galaxy["SAN"]
	distance = 0
	for {
		pathValue, ok := distanceToYou[o]
		if ok {
			distance += pathValue
			break
		}
		
		o, ok = galaxy[o]
		distance++
	}
	
	fmt.Println("Distance:", distance);
}
