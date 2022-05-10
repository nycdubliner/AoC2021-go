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

	// As we need to use the numbers more than once and move back and forth it will be easier to get all records in a slice than to try and handle eveything off a single file read.
	var input []int
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
	}
	var windows []int
	i := 0
	for i < (len(input) - 2) {
		windows = append(windows, (input[i] + input[i+1] + input[i+2]))
		i++
	}

	current := windows[0]
	counter := 0
	j := 1

	for j < (len(windows)) {
		next := windows[j]
		if next > current {
			counter++
		}
		current = windows[j]
		j++
	}
	fmt.Println(counter)
}
