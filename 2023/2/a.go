package main

import (
	"fmt"
	"flag"

	"github.com/baldrick/aoc/AdventOfCode/2023/aoc"
)

const (
	year = 2023
	day = 2
)

var (
	inputFile = flag.String("f", "test", "Puzzle file (partial name) to use")
)

func main() {
	flag.Parse()
	puzzle, err := aoc.GetPuzzleInput(*inputFile, year, day)
	if err != nil {
		fmt.Printf("oops: %v\n", err)
		return
	}
	if err := process(puzzle); err != nil {
		fmt.Printf("oops: %v\n", err)
	}
}

func process(puzzle []string) error {
	return fmt.Errorf("Not yet implemented")
}
