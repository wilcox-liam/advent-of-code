package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	password_count := 0

	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	input_range := string(dat)

	passwords := strings.Split(input_range, "-")
	password_low, err := strconv.Atoi(passwords[0])
	check(err)
	password_high, err := strconv.Atoi(passwords[1])
	check(err)

	fmt.Println("LOW:", password_low, "HIGH:", password_high)
	for ; password_low <= password_high; password_low++ {
		if is_six_digits(password_low) && contains_number_pair(password_low) && numbers_increase(password_low) {
			password_count++
		}
	}
	fmt.Println(password_count)
}

//Check that the password is 6 characters long
func is_six_digits(password int) bool {
	pass := strings.Split(strconv.Itoa(password), "")

	if len(pass) == 6 {
		return true
	}
	return false
}

//Checks a string of numbers whether it contains a pair of adjacent identical numbers
func contains_number_pair(password int) bool {
	pass := strings.Split(strconv.Itoa(password), "")

	//Start at 1 and compare to previous, end
	for i := 1; i < len(pass); i++ {
		if pass[i] == pass[i-1] {
			return true
		}
	}
	return false
}

//Checks a string of numbers, that each individual number is more than or equal to the previous
func numbers_increase(password int) bool {
	pass := strings.Split(strconv.Itoa(password), "")

	//Start at 1 and compare to previous, end
	for i := 1; i < len(pass); i++ {
		if pass[i] < pass[i-1] {
			return false
		}
	}
	return true
}
