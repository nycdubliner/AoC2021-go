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
		if q[0] == "forward" {
			vector, _ := strconv.Atoi(q[1])
			distance = distance + vector
		} else if q[0] == "down" {
			vector, _ := strconv.Atoi(q[1])
			depth = depth + vector
		} else if q[0] == "up" {
			vector, _ := strconv.Atoi(q[1])
			depth = depth - vector
		}

	}
	fmt.Print(depth*distance, "\n")

}
