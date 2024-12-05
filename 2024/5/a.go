package day5

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
		Name:    "day5A",
		Aliases: []string{"day5a"},
		Usage:   "Day 5 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day5B",
		Aliases: []string{"day5b"},
		Usage:   "Day 5 part B",
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
	/*
	   47|53
	   97|13
	   97|61

	   53 must come after 47
	   so pagesBefore[53] = 47
	*/
	pagesBefore := make(map[string]aoc.StringSet)
	var n int
	var line string
	for n, line = range puzzle {
		mapping := strings.Split(line, "|")
		if len(mapping) == 1 {
			log.Printf("breaking at line %v", n)
			break
		}
		before := mapping[0]
		after := mapping[1]
		beforeList, ok := pagesBefore[after]
		if !ok {
			beforeList = *aoc.NewStringSet()
			pagesBefore[after] = beforeList
		}
		beforeList.Add(before)
	}
	log.Print(pagesBefore)
	sum := 0
	for ; n < len(puzzle); n++ {
		log.Printf("page #%v: %v", n, puzzle[n])
		pages := strings.Split(puzzle[n], ",")
		valid := true
		for p := 0; p < len(pages)-1; p++ {
			if checkPagesBefore(pagesBefore, pages[p], aoc.MakeStringSet(pages[p+1:])) {
				//log.Printf("Invalid: %v", pages)
				valid = false
				break
			}
		}
		if valid {
			middle := pages[(len(pages)-1)/2]
			log.Printf("Valid, middle number=%v", middle)
			sum += aoc.MustAtoi(middle)
		}
	}

	return sum, nil
}

// checkPagesBefore returns true if any of the beforePages should be after the after page.
func checkPagesBefore(pagesBefore map[string]aoc.StringSet, after string, beforePages *aoc.StringSet) bool {
	pb, ok := pagesBefore[after]
	if !ok {
		return false
	}
	i := pb.Intersect(beforePages)
	//log.Printf("intersection of %v and %v = %v", beforePages, pb, i)
	return i.Len() > 0
}

type pair struct {
	before, after string
}

func processB(puzzle []string) (int, error) {
	pagesBefore := make(map[string]aoc.StringSet)
	ordering := make(map[pair]struct{})
	var n int
	var line string
	for n, line = range puzzle {
		mapping := strings.Split(line, "|")
		if len(mapping) == 1 {
			log.Printf("breaking at line %v", n)
			break
		}
		before := mapping[0]
		after := mapping[1]
		newPair := pair{before, after}
		ordering[newPair] = struct{}{}
		beforeList, ok := pagesBefore[after]
		if !ok {
			beforeList = *aoc.NewStringSet()
			pagesBefore[after] = beforeList
		}
		beforeList.Add(before)
	}
	log.Print(ordering)
	sum := 0
	for ; n < len(puzzle); n++ {
		log.Printf("page #%v: %v", n, puzzle[n])
		pages := strings.Split(puzzle[n], ",")
		for p := 0; p < len(pages)-1; p++ {
			if checkPagesBefore(pagesBefore, pages[p], aoc.MakeStringSet(pages[p+1:])) {
				sort.Slice(pages, func(i, j int) bool {
					if pages[i] == pages[j] {
						return false
					}
					ij := pair{before: pages[i], after: pages[j]}
					_, ok := ordering[ij]
					return ok
				})
				sum += aoc.MustAtoi(pages[len(pages)/2])
			}
		}
	}
	return sum, nil
}
