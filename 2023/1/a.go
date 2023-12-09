package day1

import (
	"fmt"
	"flag"
	"regexp"
	"strconv"

	"github.com/baldrick/aoc/2023/aoc"
	"github.com/urfave/urfave_cli"
)

const (
	year = 2023
	day = 1
)

var (
//	inputFile = flag.String("f", "test", "Puzzle file (partial name) to use")
	cmd := &cli.Command{
		Name:  "1A",
		Usage: "Day 1 part A",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "f",
				Default: "test",
			}
		},
		Action: partA,
	}
)

func partA(ctx context.Context, cmd *cli.Command) error {
//	flag.Parse()
	puzzle, err := aoc.GetPuzzleInput(cmd.GetStringFlag(f), year, day)
	if err != nil {
		fmt.Printf("oops: %v\n", err)
		return
	}
	if err := process(puzzle); err != nil {
		fmt.Printf("oops: %v", err)
	}
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