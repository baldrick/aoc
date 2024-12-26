package day25

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/baldrick/aoc/common/grid"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day25A",
		Aliases: []string{"day25a"},
		Usage:   "Day 25 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day25B",
		Aliases: []string{"day25b"},
		Usage:   "Day 25 part B",
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
	var keys []key
	var locks []lock
	for n := 0; n < len(puzzle); n += 8 {
		if puzzle[n] == "#####" {
			locks = append(locks, getLock(puzzle[n:n+7]))
		} else {
			keys = append(keys, getKey(puzzle[n:n+7]))
		}
	}
	log.Printf("locks: %v", locks)
	log.Printf("keys: %v", keys)
	fit := 0
	for _, l := range locks {
		for _, k := range keys {
			if l.fits(k) {
				log.Printf("lock %v fits key %v", l, k)
				fit++
			}
		}
	}
	return fit, nil
}

func processB(puzzle []string) (int, error) {
	return 0, fmt.Errorf("Not yet implemented")
}

type key []int

type lock []int

func getLock(l []string) lock {
	g := grid.New(l)
	g.DumpMsg("lock")

	var p lock
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			if g.Get(x, y) == "." {
				p = append(p, y-1)
				break
			}
		}
	}
	return p
}

func getKey(k []string) key {
	g := grid.New(k)
	g.DumpMsg("key")

	var p key
	for x := 0; x < g.Width(); x++ {
		for y := g.Height() - 1; y >= 0; y-- {
			if g.Get(x, y) == "." {
				p = append(p, y+1)
				break
			}
		}
	}
	return p
}

func (l lock) fits(k key) bool {
	for n := 0; n < len(l); n++ {
		if k[n] <= l[n] {
			return false
		}
	}
	return true
}
