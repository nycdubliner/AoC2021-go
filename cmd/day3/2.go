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
	s := "/home/tdeburca/git/AoC2021-go/cmd/day3/testfile.txt"
	// s := "/home/tdeburca/git/AoC2021-go/cmd/day3/input.txt"
	readings := getReadings(s)
	fmt.Println(getOxegenGeneratorRating(readings))

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

func getOxegenGeneratorRating(readings []string) int64 {
	// len(string) != number of runes in string. Converting it to a list of runes fixes this.
	// not required here, but going for 'correct'.
	bit_width := len([]rune(readings[0]))
	for digitIndex := 0; digitIndex <= bit_width; digitIndex++ {
		commonDigit := getCommonDigit(digitIndex, readings)
		readings = getMatchingRecords(digitIndex, commonDigit, readings)
		if len(readings) == 1 {
			break
		}
	}
	//should never happen
	out, _ := strconv.ParseInt(readings[0], 2, 64)
	return out
}

func getMatchingRecords(position int, value int, l []string) []string {
	var output []string
	for _, record := range l {
		// record[position] returns a byte, converted to string, converted to int.
		t, _ := strconv.Atoi(string(record[position]))
		if t == value {
			output = append(output, record)
		}
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
