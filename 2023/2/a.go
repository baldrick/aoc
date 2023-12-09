package day2

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
    day = 2
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day2A",
        Usage: "Day 2 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day2B",
        Usage: "Day 2 part B",
        Action: partB,
    }

	bag = map[string]int{"red": 12, "green": 13, "blue": 14}
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
	total := 0
	for _, line := range puzzle {
		id, possible, err := possibleGame(line)
		if err != nil {
			return 0, err
		}
		if possible {
			total += id
		}
	}
	return total, nil
}

func processB(puzzle []string) (int, error) {
	total := 0
	for _, line := range puzzle {
		_, m, err := getMaxBalls(line)
		if err != nil {
			return 0, err
		}
		power := 1
		for _, v := range m {
			power *= v
		}
		total += power
	}
	return total, nil
}

func possibleGame(line string) (int, bool, error) {
	id, m, err := getMaxBalls(line)
	if err != nil {
		return 0, false, err
	}
	if len(m) != len(bag) {
		fmt.Printf("m: #%v, bag: #%v\n", len(m), len(bag))
		return id, false, nil
	}
	for c, v := range m {
		count, ok := bag[c]
		if !ok || count < v {
			fmt.Printf("%v not found or too few (%v < %v)\n", c, count, v)
			return id, false, nil
		}
	}
	return id, true, nil
}

func getMaxBalls(line string) (int, map[string]int, error) {
	// Game N: x colour<,x colour>;x color...
	re := regexp.MustCompile(`Game ([0-9]*): (.*)`)
	matches := re.FindStringSubmatch(line)
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, nil, err
	}
	games := matches[2] + ";"

	m := make(map[string]int)
	for ; games != "" ; {
		var game string
		game, games = getNextGame(games)
		if err := addGame(m, game); err != nil {
			return 0, nil, err
		}
	}

	return id, m, nil
}

func getNextGame(games string) (string, string) {
	re := regexp.MustCompile(`([0-9, a-z]*);(.*)`)
	matches := re.FindStringSubmatch(games)
	if len(matches) < 3 {
		return matches[1], ""
	}
	return matches[1], matches[2]
}

func addGame(m map[string]int, game string) error {
	game += ", "
	re := regexp.MustCompile(`([0-9]*) ([a-z]*), (.*)`)
	for ; game != "" ; {
		matches := re.FindStringSubmatch(game)
		if len(matches) == 1 {
			break
		}
		n, err := strconv.Atoi(matches[1])
		if err != nil {
			return err
		}
		c, ok := m[matches[2]]
		if !ok || n > c {
			m[matches[2]] = n
		}
		game = matches[3]
	}
	return nil
}
