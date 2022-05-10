package main

import (
	"os"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	s := "/home/tdeburca/git/AoC2022-go/cmd/day1/part1/testfile.txt"
	dat, err := os.ReadFile(s)
	check(err)
	fmt.Print(string(dat))
	fmt.Print("-----\n")

}

