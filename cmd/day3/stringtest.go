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

	fmt.Printf("%T \n", readings)    // prints : []string
	fmt.Printf("%T \n", readings[0]) // prints : string

	for _, reading := range readings {
		fmt.Printf("%T \n", reading) // prints : string
		break
	}

	for _, character := range "100101" {
		fmt.Printf("%T \n", character) // prints : int32
		break
	}

	for _, reading := range readings {
		for _, character := range reading {
			fmt.Printf("%T \n", character) // prints int32
			break
		}
		break
	}

	for _, character := range readings {
		for i := 0; i < (len(character)); i++ {
			fmt.Printf("%T \n", character[i:i+1]) // prints string
			break
		}
		break
	}

	for pos, char := range "日本語" {
		fmt.Printf("character %c starts at byte position %d\n", char, pos) // prints character 日 starts at byte position 0
		fmt.Printf("%T \n", char)                                          // prints : int32
		break
	}
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
