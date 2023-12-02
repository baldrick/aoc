package main

import (
	"fmt"
	"flag"
	"log"

	"github.com/baldrick/aoc/AdventOfCode/2023/aoc"
)

const (
	year = {{YEAR}}
	day = {{DAY}}
)

var (
	inputFile = flag.String("f", "test", "Puzzle file (partial name) to use")
	logger = log.Default()
)

func main() {
	flag.Parse()
	puzzle, err := aoc.GetPuzzleInput(*inputFile, year, day)
	if err != nil {
		logger.Fatalf("oops: %v\n", err)
		return
	}
	if err := process(puzzle); err != nil {
		logger.Fatalf("oops: %v\n", err)
	}
}

func process(puzzle []string) error {
	for _, line := range puzzle {
		logger.Print(line)
	}
	return logger.Fatal("Not yet implemented")
}
