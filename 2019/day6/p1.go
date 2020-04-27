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

	count := 0
	for orbit := range galaxy {
		for {
			parent, ok := galaxy[orbit]
			if ok {
				orbit = parent
				count++
			} else {
				break
			}
		}
	}

	fmt.Println("Orbits: ", count)
}
