package main

import (
	"fmt"
	"flag"
	"log"
	"regexp"
	"strconv"

	"github.com/baldrick/aoc/AdventOfCode/2023/aoc"
)

const (
	year = 2023
	day = 6
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
	//Time:      7  15   30
	//Distance:  9  40  200
	times := parseLine2(puzzle[0])
	distances := parseLine2(puzzle[1])
	logger.Printf("times: %v", times)
	logger.Printf("distances: %v", distances)
	total := 1
	for n, r := range times {
		beatenRecord := 0
		for h := 1; h < r; h++ {
			dist := (r-h)*h
			if dist > distances[n] {
				beatenRecord++
			}
		}
		if beatenRecord > 0 {
			total *= beatenRecord
		}
	}
	logger.Printf("total=%v", total)
	return nil
}

func parseLine(s string) []int {
	re := regexp.MustCompile(`[A-Za-z: ]*([0-9]*) *([0-9]*) *([0-9]*) *([0-9]*)`)
	matches := re.FindStringSubmatch(s)
	return []int{atoi(matches[1]), atoi(matches[2]), atoi(matches[3]), atoi(matches[4])}
}

func parseLine2(s string) []int {
	re := regexp.MustCompile(`[A-Za-z: ]*([0-9]*) *([0-9]*) *([0-9]*) *([0-9]*)`)
	matches := re.FindStringSubmatch(s)
	return []int{atoi(matches[1] + matches[2] + matches[3] + matches[4])}
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to translate %q to number: %v", s, err))
	}
	return n
}