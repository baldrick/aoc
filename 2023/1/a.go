package day1

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/baldrick/aoc/2023/aoc"
	"github.com/urfave/cli"
)

const (
	year = 2023
	day = 1
)

var (
	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:  "day1A",
		Usage: "Day 1 part A",
		Flags: aoc.Flags,
		Action: partA,
	}
	B = &cli.Command{
		Name:  "day1B",
		Usage: "Day 1 part B",
		Flags: aoc.Flags,
		Action: partB,
	}
)

func partA(ctx *cli.Context) error {
	return fmt.Errorf("needs code adjustment to unwind part B :)")
}

func partB(ctx *cli.Context) error {
	puzzle, err := aoc.GetPuzzleInput(ctx.String("f"), year, day)
	if err != nil {
		return fmt.Errorf("oops: %v\n", err)
	}
	if err := process(puzzle); err != nil {
		return fmt.Errorf("oops: %v", err)
	}
	return nil
}

func process(puzzle []string) error {
	r := regexp.MustCompile(`[a-z]`)
	total := 0
	for _, in := range puzzle {
		line := convertWords(in)
		numbers := r.ReplaceAllString(line, "")
		n, err := strconv.Atoi(string(numbers[0]))
		if err != nil {
			return err
		}
		if len(numbers) > 1 {
			n = 10*n
			u, err := strconv.Atoi(string(numbers[len(numbers)-1]))
			if err != nil {
				return err
			}
			n += u
		} else {
			n = 10*n + n
		}
		total += n
		fmt.Printf("%v -> %v -> %v => added %v (total %v)\n", in, line, numbers, n, total)
	}
	fmt.Printf("total: %v\n", total)
	return nil
}

func convertWords(in string) string {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	out := in
	for i, _ := range out {
		for n, w := range words {
			if len(out)-i < len(w) {
				continue
			}
			if out[i:i+len(w)] == w {
				out = fmt.Sprintf("%v%v%v", out[:i], n+1, out[i+1:])
			}
		}
	}
	return out
}