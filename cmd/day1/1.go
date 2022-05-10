package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// String for filename
	// s := "/home/tdeburca/git/AoC2022-go/cmd/day1/testfile.txt"
	s := "/home/tdeburca/git/AoC2022-go/cmd/day1/input.txt"
	// assign filehandler to f.
	f, err := os.Open(s)
	check(err)
	// create scanner for file
	scanner := bufio.NewScanner(f)
	// Get first number
	scanner.Scan()
	// convert returned text to integer.
	current, _ := strconv.Atoi(scanner.Text())
	// create counter
	counter := 0
	// Get all subsequent numbers
	for scanner.Scan() {
		next, _ := strconv.Atoi(scanner.Text())
		if next > current {
			counter++
		}
		current = next
	}
	fmt.Println(counter)
}
