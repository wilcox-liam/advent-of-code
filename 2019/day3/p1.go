package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var lines [][]Coordinate
	var origin = Coordinate{0, 0}
	var smallest_distance = 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		instructions := strings.Split(input, ",")
		line := make([]Coordinate, 0)
		for _, instruction := range instructions {
			target := instruction_to_coord(origin, instruction)
			line = append(line, calculate_points_between_coords(origin, target)...)
			line = append(line, target)
			origin = target
		}
		lines = append(lines, line)
		origin.X = 0
		origin.Y = 0
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(lines) < 2 {
		log.Fatal("Only 1 Line!")
	}

	fmt.Println("Line 1:", len(lines[0]))
	fmt.Println("Line 2:", len(lines[1]))

	//Compares every point in every line to every point in every other line.
	for k := 0; k < len(lines[0]); k++ {
		for l := 0; l < len(lines[1]); l++ {
			if lines[0][k] == lines[1][l] {
				distance := calculate_manhattan_distance(lines[0][k])
				fmt.Println(lines[0][k])
				if distance < smallest_distance || smallest_distance == 0 {
					smallest_distance = distance

				}

			}
		}
	}

	//fmt.Println(lines)
	fmt.Println(smallest_distance)
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculate_manhattan_distance(point Coordinate) int {
	return Abs(point.X) + Abs(point.Y)
}

//Converts a contextual direction (Up, Down, Left, Right), steps it and returns the coordinate at the end
func instruction_to_coord(origin Coordinate, instruction string) Coordinate {
	direction := instruction[0:1]
	steps, err := strconv.Atoi(string([]byte(instruction[1:])))
	check(err)

	var target = Coordinate{origin.X, origin.Y}

	if direction == "U" {
		target.Y += steps
	} else if direction == "R" {
		target.X += steps
	} else if direction == "D" {
		target.Y -= steps
	} else if direction == "L" {
		target.X -= steps
	} else {
		log.Fatal(err)
	}

	return target
}

//Given 2 coordinates, calculates every point between them and returns them as a slice.
//Coordinates must have the same x axis or y axis
func calculate_points_between_coords(origin Coordinate, target Coordinate) []Coordinate {
	if origin.X == target.X {
		if origin.Y > target.Y {
			return calculate_points_y_axis(target.Y, origin.Y, origin.X)
		} else {
			return calculate_points_y_axis(origin.Y, target.Y, origin.X)
		}
	} else if origin.Y == target.Y {
		if origin.X > target.X {
			return calculate_points_x_axis(target.X, origin.X, origin.Y)
		} else {
			return calculate_points_x_axis(origin.X, target.X, origin.Y)
		}
	} else {
		log.Fatal()
	}
	return nil
}

//Calculates every point between 2 points along the x axis and returns them as a slice.
func calculate_points_x_axis(origin int, target int, y int) []Coordinate {
	if origin == target {
		log.Fatal()
	}

	var points []Coordinate
	for origin += 1; origin < target; origin++ {
		point := Coordinate{origin, y}
		points = append(points, point)
	}
	return points
}

//Calculates every point between 2 points along the y axis and returns them as a slice.
func calculate_points_y_axis(origin int, target int, x int) []Coordinate {
	if origin == target {
		log.Fatal()
	}

	var points []Coordinate
	for origin += 1; origin < target; origin++ {
		point := Coordinate{x, origin}
		points = append(points, point)
	}
	return points
}
