package day1

import (
	_ "embed"
	"log"
	"sort"
	"strings"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day1A",
		Aliases: []string{"day1a"},
		Usage:   "Day 1 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day1B",
		Aliases: []string{"day1b"},
		Usage:   "Day 1 part B",
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
	var left, right []int
	for _, line := range puzzle {
		numbers := strings.Split(line, "   ")
		log.Printf("%v -> %v", line, numbers)
		left = append(left, aoc.MustAtoi(numbers[0]))
		right = append(right, aoc.MustAtoi(numbers[1]))
	}
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for n := 0; n < len(left); n++ {
		sum += aoc.AbsInt(left[n] - right[n])
	}
	return sum, nil
}

func processB(puzzle []string) (int, error) {
	var left, right []int
	for _, line := range puzzle {
		numbers := strings.Split(line, "   ")
		left = append(left, aoc.MustAtoi(numbers[0]))
		right = append(right, aoc.MustAtoi(numbers[1]))
	}
	counts := make(map[int]int)
	for _, r := range right {
		c, ok := counts[r]
		if !ok {
			c = 0
		}
		counts[r] = c + 1
	}
	log.Printf("counts: %v", counts)
	log.Printf("left (#%v items): %v", len(left), left)
	sum := 0
	for n := 0; n < len(left); n++ {
		count, ok := counts[left[n]]
		if !ok {
			count = 0
		}
		sum += (left[n] * count)
		log.Printf("#%v, %v * %v = %v, sum = %v", n, left[n], count, left[n]*count, sum)
	}
	return sum, nil
}
