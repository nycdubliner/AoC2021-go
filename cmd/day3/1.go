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
	// Width of input changes with file.
	// s := "/home/tdeburca/git/AoC2021-go/cmd/day3/testfile.txt"
	// var mostCommonBitArray [5]int
	s := "/home/tdeburca/git/AoC2021-go/cmd/day3/input.txt"
	var mostCommonBitArray [12]int

	// assign filehandler to f.
	f, err := os.Open(s)
	check(err)
	// create scanner for file
	scanner := bufio.NewScanner(f)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Type of var
	//fmt.Printf("%T", v)

	for _, v := range input {
		for i, q := range v {
			bit, _ := strconv.Atoi(string(q))
			mostCommonBitArray[i] = mostCommonBitArray[i] + bit
		}
	}
	// most common bit is the bit in more than half the records.
	target := (len(input) / 2)

	gamma := ""
	epsilon := ""
	for i := 0; i < len(mostCommonBitArray); i++ {
		if mostCommonBitArray[i] > target {
			gamma = gamma + string("1")
			epsilon = epsilon + string("0")
		} else {
			gamma = gamma + string("0")
			epsilon = epsilon + string("1")
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(g * e)
}
