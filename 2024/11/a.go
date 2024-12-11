package day11

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day11A",
		Aliases: []string{"day11a"},
		Usage:   "Day 11 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day11B",
		Aliases: []string{"day11b"},
		Usage:   "Day 11 part B",
		Action:  partB,
	}
)

func partA(ctx *cli.Context) error {
	answer, err := processA(aoc.PreparePuzzle(puzzle), 25)
	if err != nil {
		return err
	}
	log.Printf("Answer A: %v", answer)
	return nil
}

func partB(ctx *cli.Context) error {
	answer, err := processB(aoc.PreparePuzzle(puzzle), 75)
	if err != nil {
		return err
	}
	log.Printf("Answer B: %v", answer)
	return nil
}

func processA(puzzle []string, blinks int) (int, error) {
	stones := strings.Split(puzzle[0], " ")
	stoneMap := make(map[string]int)
	for _, stone := range stones {
		stoneMap[stone] = 1
	}
	for b := 0; b < blinks; b++ {
		stoneMap = blink(stoneMap)
	}
	sum := 0
	for _, v := range stoneMap {
		sum += v
	}
	return sum, nil
}

func processB(puzzle []string, blinks int) (int, error) {
	return processA(puzzle, blinks)
}

func blink(stoneMap map[string]int) map[string]int {
	newStoneMap := make(map[string]int)
	/*
	   0 -> 1
	   even #digits -> split, left half / right half, ignore leading zeroes in answer
	   else new stone, value = current * 2024
	*/
	for stone, count := range stoneMap {
		switch {
		case stone == "0":
			//log.Printf("0 becomes 1 for %v", count)
			addStone(newStoneMap, "1", count)
		case aoc.ModInt(len(stone), 2) == 0:
			half := len(stone) / 2
			left := fmt.Sprintf("%v", aoc.MustAtoi(stone[:half]))
			right := fmt.Sprintf("%v", aoc.MustAtoi(stone[half:]))
			//log.Printf("split %v into %v and %v for %v", stone, left, right, count)
			addStone(newStoneMap, left, count)
			addStone(newStoneMap, right, count)
		default:
			n := fmt.Sprintf("%v", aoc.MustAtoi(stone)*2024)
			//log.Printf("%v * 2024 = %v for %v", stone, n, count)
			addStone(newStoneMap, n, count)
		}
	}
	//log.Printf("%v blink -> %v", stoneMap, newStoneMap)
	return newStoneMap
}

func addStone(sm map[string]int, stone string, count int) {
	n, ok := sm[stone]
	if !ok {
		sm[stone] = count
	} else {
		sm[stone] = count + n
	}
}
