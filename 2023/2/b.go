package main

import (
	"fmt"
	"flag"
	"regexp"
	"strconv"

	"github.com/baldrick/aoc/AdventOfCode/2023/aoc"
)

const (
	year = 2023
	day = 2
)

var (
	inputFile = flag.String("f", "test", "Puzzle file (partial name) to use")
	bag = map[string]int{"red": 12, "green": 13, "blue": 14}
)

func main() {
	flag.Parse()
	puzzle, err := aoc.GetPuzzleInput(*inputFile, year, day)
	if err != nil {
		fmt.Printf("oops: %v\n", err)
		return
	}
	if err := process(puzzle); err != nil {
		fmt.Printf("oops: %v\n", err)
	}
}

func process(puzzle []string) error {
	total := 0
	for _, line := range puzzle {
		_, m, err := getMaxBalls(line)
		if err != nil {
			return err
		}
		power := 1
		for _, v := range m {
			power *= v
		}
		total += power
	}
	fmt.Printf("total: %v\n", total)
	return nil
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