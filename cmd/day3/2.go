package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	s := "/home/tdeburca/git/AoC2021-go/cmd/day3/testfile.txt"
	// s := "/home/tdeburca/git/AoC2021-go/cmd/day3/input.txt"
	readings := getReadings(s)
	getCommonDigit(0, readings)

	fmt.Println(getMatchingRecords(1, 1, readings))
}

func getReadings(filename string) []string {
	f, err := os.Open(filename)
	check(err)
	// create scanner for file
	scanner := bufio.NewScanner(f)
	var readings []string
	for scanner.Scan() {
		readings = append(readings, scanner.Text())
	}
	return readings
}

//func get_oxegen_generator_rating(readings []string)

func getMatchingRecords(position int, value int, l []string) []string {
	var output []string
	for _, record := range l {
		fmt.Print(record, position, record[position], value)
	}
	return output
}

func getCommonDigit(position int, l []string) int {
	ones := 0
	zeroes := 0
	for _, record := range l {
		if record[position] == '1' {
			ones++
		} else if record[position] == '0' {
			zeroes++
		} else {
			fmt.Print(record, position, "error")
			break
		}
	}
	if ones >= zeroes {
		return 1
	} else {
		return 0
	}
}
