package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// String for filename
	// s := "/home/tdeburca/git/AoC2021-go/cmd/day2/testfile.txt"
	s := "/home/tdeburca/git/AoC2021-go/cmd/day2/input.txt"

	// assign filehandler to f.
	f, err := os.Open(s)
	check(err)
	// create scanner for file
	scanner := bufio.NewScanner(f)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	depth := 0
	distance := 0
	for _, direction := range input {
		q := strings.Split(direction, " ")
		vector, _ := strconv.Atoi(q[1])
		if q[0] == "forward" {
			distance = distance + vector
		} else if q[0] == "down" {
			depth = depth + vector
		} else if q[0] == "up" {
			depth = depth - vector
		}

	}
	fmt.Print(depth*distance, "\n")

}
