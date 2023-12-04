package main

import (
	"fmt"
	"flag"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/baldrick/aoc/AdventOfCode/2023/aoc"
)

const (
	year = 2023
	day = 4
)

var (
	inputFile = flag.String("f", "test", "Puzzle file (partial name) to use")
	logger = log.Default()
)

func main() {
	flag.Parse()
	puzzle, err := aoc.GetPuzzleInput(*inputFile, year, day)
	if err != nil {
		logger.Fatalf("oops: %v\n", err)
		return
	}
	if err := processB(puzzle); err != nil {
		logger.Fatalf("oops: %v\n", err)
	}
}

func process(puzzle []string) error {
	total := float64(0)
	for _, line := range puzzle {
		re := regexp.MustCompile(`Card *[0-9]*: ([ 0-9]*)\|([0-9 ]*)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			return fmt.Errorf("len(matches)=%v, want 3", len(matches))
		}
		cards := numberSet(matches[1])
		winning := numberSet(matches[2])
		overlap := cards.Intersect(winning).Len()
		score := float64(0)
		if overlap > 0 {
			score = math.Pow(float64(2), float64(overlap-1))
			total += score
		}
		//logger.Printf("%v -> cards %v, winning %v, overlap %v, score %v", line, cards, winning, overlap, score)
	}
	logger.Printf("total = %v", total)
	return nil
}

func processB(puzzle []string) error {
	cardCount := make(map[string]int)
	for n, _ := range puzzle {
		cardCount[fmt.Sprintf("%v", n+1)] = 1
	}
	for _, line := range puzzle {
		re := regexp.MustCompile(`Card *([0-9]*): ([ 0-9]*)\|([0-9 ]*)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			return fmt.Errorf("len(matches)=%v, want 3", len(matches))
		}
		card := matches[1]
		countOfThisCard, ok := cardCount[card]
		if !ok {
			return fmt.Errorf("failed to find card %v", card)
		}
		cards := numberSet(matches[2])
		winning := numberSet(matches[3])
		overlap := cards.Intersect(winning).Len()
		for n:=1; n<=overlap && n<len(puzzle); n++ {
			cardCount[addOne(card,n)] += countOfThisCard
		}
		logger.Printf("%v -> cards %v, winning %v, overlap %v", line, cards, winning, overlap)
	}
	total := 0
	for _, v := range cardCount {
		total += v
	}
	logger.Printf("total = %v", total)
	return nil
}

func numberSet(s string) *aoc.StringSet {
	ss := aoc.NewStringSet()
	for _, n := range strings.Split(s, " ") {
		ss.Add(strings.TrimSpace(n))
	}
	return ss
}

func addOne(a string, b int) string {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %q to int: %v", a, err))
	}
	return fmt.Sprintf("%v", i+b)
}