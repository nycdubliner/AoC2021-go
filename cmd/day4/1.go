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
	// s := "/home/tdeburca/git/AoC2021-go/cmd/day4/testfile.txt"
	s := "./testfile.txt"
	// s := "/home/tdeburca/git/AoC2021-go/cmd/day4/input.txt"
	readings := getReadings(s)
	fmt.Println(readings)

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
