package day6

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
    day = 6
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day6A",
        Usage: "Day 6 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day6B",
        Usage: "Day 6 part B",
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
	//Time:      7  15   30
    //Distance:  9  40  200
    log.Printf("puzzle: %v", puzzle)
	times := parseLine(puzzle[0])
	distances := parseLine(puzzle[1])
    return race(times, distances)
}

func processB(puzzle []string) (int, error) {
	//Time:      71530
	//Distance:  940200
	times := parseLine2(puzzle[0])
    distances := parseLine2(puzzle[1])
    return race(times, distances)
}

func race(times, distances []int) (int, error) {
	log.Printf("times: %v", times)
	log.Printf("distances: %v", distances)
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
	return total, nil
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
    if len(s) == 0 {
        return 0
    }
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to translate %q to number: %v", s, err))
	}
	return n
}
