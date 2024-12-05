package day3

import (
	_ "embed"
	"log"
	"regexp"
	"strings"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day3A",
		Aliases: []string{"day3a"},
		Usage:   "Day 3 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day3B",
		Aliases: []string{"day3b"},
		Usage:   "Day 3 part B",
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
	sum := 0
	for _, line := range puzzle {
		log.Print(line)
		re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		muls := re.FindAllString(line, -1)
		log.Print(muls)
		for _, mul := range muls {
			re = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
			calcs := re.FindAllStringSubmatch(mul, -1)
			for _, calc := range calcs {
				sum += aoc.MustAtoi(calc[1]) * aoc.MustAtoi(calc[2])
			}
			log.Print(calcs)
		}
	}
	return sum, nil
}

func processB(puzzle []string) (int, error) {
	allMuls := ""
	enabled := true
	line := puzzle[0]
	for {
		if enabled {
			// already enabled, find where we start being disabled.
			before, after, found := strings.Cut(line, "don't()")
			log.Printf("\n-----------------------------\nline: %v\n\nbefore (enabled): %v\n\nafter (disabled): %v\n\nallMuls: %v", line, before, after, allMuls)
			allMuls += before
			if !found {
				break
			}
			enabled = false
			line = after
		} else {
			// disabled, find where it's next enabled.
			before, after, found := strings.Cut(line, "do()")
			log.Printf("\n########################\nline: %v\n\nbefore (disabled): %v\n\nafter (enabled): %v", line, before, after)
			if !found {
				break
			}
			line = after
			enabled = true
		}
	}

	log.Printf("%v -> all muls: %v", puzzle[0], allMuls)
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	calcs := re.FindAllStringSubmatch(allMuls, -1)
	sum := 0
	for _, calc := range calcs {
		sum += (aoc.MustAtoi(calc[1]) * aoc.MustAtoi(calc[2]))
	}
	log.Print(calcs)
	// incorrect answers:
	// 85770822
	// 7092082
	return sum, nil
}
