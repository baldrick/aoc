package day1

import (
    _ "embed"
    "fmt"
    "log"
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
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day1A",
        Usage: "Day 1 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day1B",
        Usage: "Day 1 part B",
        Action: partB,
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
	r := regexp.MustCompile(`[a-z]`)
	total := 0
	for _, in := range puzzle {
		numbers := r.ReplaceAllString(in, "")
		n, err := strconv.Atoi(string(numbers[0]))
		if err != nil {
			return 0, err
		}
		if len(numbers) > 1 {
			n = 10*n
			u, err := strconv.Atoi(string(numbers[len(numbers)-1]))
			if err != nil {
				return 0, err
			}
			n += u
		} else {
			n = 10*n + n
		}
		total += n
		fmt.Printf("%v -> %v => added %v (total %v)\n", in, numbers, n, total)
	}
	return total, nil
}

func processB(puzzle []string) (int, error) {
	r := regexp.MustCompile(`[a-z]`)
	total := 0
	for _, in := range puzzle {
		line := convertWords(in)
		numbers := r.ReplaceAllString(line, "")
		n, err := strconv.Atoi(string(numbers[0]))
		if err != nil {
			return 0, err
		}
		if len(numbers) > 1 {
			n = 10*n
			u, err := strconv.Atoi(string(numbers[len(numbers)-1]))
			if err != nil {
				return 0, err
			}
			n += u
		} else {
			n = 10*n + n
		}
		total += n
		fmt.Printf("%v -> %v -> %v => added %v (total %v)\n", in, line, numbers, n, total)
	}
	return total, nil
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
