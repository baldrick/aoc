package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetPuzzleInput(inputFile string, year, day int) ([]string, error) {
	if !strings.Contains(inputFile, ".txt") {
		inputFile += ".txt"
	}
	if !strings.Contains(inputFile, fmt.Sprintf("AdventOfCode/%v/%v/", year, day)) {
		inputFile = fmt.Sprintf("AdventOfCode/%v/%v/%v", year, day, inputFile)
	}
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var puzzle []string
	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, line)
	}
	return puzzle, nil
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}