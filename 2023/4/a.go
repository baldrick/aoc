package day4

import (
    _ "embed"
    "fmt"
    "log"
    "math"
    "regexp"
    "strconv"
    "strings"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 4
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day4A",
        Usage: "Day 4 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day4B",
        Usage: "Day 4 part B",
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
	total := float64(0)
	for _, line := range puzzle {
		re := regexp.MustCompile(`Card *[0-9]*: ([ 0-9]*)\|([0-9 ]*)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			return 0, fmt.Errorf("len(matches)=%v, want 3", len(matches))
		}
		cards := numberSet(matches[1])
		winning := numberSet(matches[2])
		overlap := cards.Intersect(winning).Len()
		score := float64(0)
		if overlap > 0 {
			score = math.Pow(float64(2), float64(overlap-1))
			total += score
		}
		//log.Printf("%v -> cards %v, winning %v, overlap %v, score %v", line, cards, winning, overlap, score)
    }
    log.Printf("%v puzzle lines processed", len(puzzle))
    return int(total), nil
}

func processB(puzzle []string) (int, error) {
	cardCount := make(map[string]int)
	for n, _ := range puzzle {
		cardCount[fmt.Sprintf("%v", n+1)] = 1
	}
	for _, line := range puzzle {
		re := regexp.MustCompile(`Card *([0-9]*): ([ 0-9]*)\|([0-9 ]*)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			return 0, fmt.Errorf("len(matches)=%v, want 3", len(matches))
		}
		card := matches[1]
		countOfThisCard, ok := cardCount[card]
		if !ok {
			return 0, fmt.Errorf("failed to find card %v", card)
		}
		cards := numberSet(matches[2])
		winning := numberSet(matches[3])
		overlap := cards.Intersect(winning).Len()
		for n:=1; n<=overlap && n<len(puzzle); n++ {
			cardCount[addOne(card,n)] += countOfThisCard
		}
		log.Printf("%v -> cards %v, winning %v, overlap %v", line, cards, winning, overlap)
	}
	total := 0
	for _, v := range cardCount {
		total += v
    }
    return total, nil
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
