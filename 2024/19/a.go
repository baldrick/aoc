package day19

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day19A",
		Aliases: []string{"day19a"},
		Usage:   "Day 19 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day19B",
		Aliases: []string{"day19b"},
		Usage:   "Day 19 part B",
		Action:  partB,
	}
)

func partA(ctx *cli.Context) error {
	answer, err := processA(aoc.PreparePuzzle(puzzle))
	if err != nil {
		return err
	}
	log.Printf("Answer A: %v", answer)
	return nil
}

func partB(ctx *cli.Context) error {
	answer, err := processB(aoc.PreparePuzzle(puzzle))
	if err != nil {
		return err
	}
	log.Printf("Answer B: %v", answer)
	return nil
}

func processA(puzzle []string) (int, error) {
	towels := aoc.MakeStringSet(strings.Split(puzzle[0], ", "))
	log.Printf("towels: %v", towels)
	patterns := puzzle[2:]
	log.Printf("patterns: %v", patterns)
	sum := 0
	for _, p := range patterns {
		if canCreate(p, towels) {
			sum++
		}
	}
	return sum, nil
}

func processB(puzzle []string) (int, error) {
	return 0, fmt.Errorf("Not yet implemented")
}

func canCreate(p string, towels *aoc.StringSet) bool {
	// pattern might be abc, towels a, b, bc
	return false
}
