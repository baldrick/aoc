package day10

import (
	_ "embed"
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
		Name:    "day10A",
		Aliases: []string{"day10a"},
		Usage:   "Day 10 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day10B",
		Aliases: []string{"day10b"},
		Usage:   "Day 10 part B",
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
	g := grid.New(puzzle)
	g.Replace(".", "-1")
	g.Dump()
	trailheads := g.FindAll("0")
	log.Printf("trailheads: %v", trailheads)
	sum := 0
	for _, t := range trailheads {
		apexes := make(map[aoc.PairInt]int)
		findApexes(g, t, apexes)
		sum += len(apexes)
	}
	return sum, nil
}

func processB(puzzle []string) (int, error) {
	g := grid.New(puzzle)
	g.Replace(".", "-1")
	g.Dump()
	trailheads := g.FindAll("0")
	log.Printf("trailheads: %v", trailheads)
	sum := 0
	for _, t := range trailheads {
		apexes := make(map[aoc.PairInt]int)
		findRoutesToApexes(g, t, apexes)
		for apex, routes := range apexes {
			log.Printf("%v routes to %v found from %v", routes, apex, t)
			sum += routes
		}
	}
	return sum, nil
}

func findApexes(g *grid.Grid, start aoc.PairInt, apexes map[aoc.PairInt]int) {
	next := findNext(g, start, aoc.MustAtoi(g.Get(start.X(), start.Y()))+1)
	for {
		if len(next) == 0 {
			return
		}
		nextNext := make(map[aoc.PairInt]int)
		for p := range next {
			if g.Get(p.X(), p.Y()) == "9" {
				apexes[p] = 0
			} else {
				nn := findNext(g, p, aoc.MustAtoi(g.Get(p.X(), p.Y()))+1)
				for np := range nn {
					rating, ok := nextNext[np]
					if !ok {
						nextNext[np] = 0
						rating = 0
					}
					nextNext[np] = rating + 1
				}
			}
		}
		next = nextNext
	}
}

func findNext(g *grid.Grid, start aoc.PairInt, searchHeight int) map[aoc.PairInt]int {
	next := make(map[aoc.PairInt]int)
	directions := make(map[aoc.PairInt]struct{})
	directions[aoc.NewPairInt(1, 0)] = struct{}{}
	directions[aoc.NewPairInt(-1, 0)] = struct{}{}
	directions[aoc.NewPairInt(0, 1)] = struct{}{}
	directions[aoc.NewPairInt(0, -1)] = struct{}{}
	for d := range directions {
		nx, ny := start.X()+d.X(), start.Y()+d.Y()
		if g.Outside(nx, ny) {
			continue
		}
		if aoc.MustAtoi(g.Get(nx, ny)) == searchHeight {
			next[aoc.NewPairInt(nx, ny)] = 0
		}
	}
	log.Printf("next from %v = %v", start, next)
	return next
}

func findRoutesToApexes(g *grid.Grid, start aoc.PairInt, routes map[aoc.PairInt]int) {
	next := findAllNext(g, start, aoc.MustAtoi(g.Get(start.X(), start.Y()))+1)
	for {
		if len(next) == 0 {
			return
		}
		var nextNext []aoc.PairInt

		for _, p := range next {
			if g.Get(p.X(), p.Y()) == "9" {
				rating, ok := routes[p]
				if !ok {
					rating = 0
					routes[p] = rating
				}
				routes[p] = rating + 1
			} else {
				nn := findAllNext(g, p, aoc.MustAtoi(g.Get(p.X(), p.Y()))+1)
				nextNext = append(nextNext, nn...)
			}
		}
		next = nextNext
	}
}

func findAllNext(g *grid.Grid, start aoc.PairInt, searchHeight int) []aoc.PairInt {
	var next []aoc.PairInt
	directions := make(map[aoc.PairInt]struct{})
	directions[aoc.NewPairInt(1, 0)] = struct{}{}
	directions[aoc.NewPairInt(-1, 0)] = struct{}{}
	directions[aoc.NewPairInt(0, 1)] = struct{}{}
	directions[aoc.NewPairInt(0, -1)] = struct{}{}
	for d := range directions {
		nx, ny := start.X()+d.X(), start.Y()+d.Y()
		if g.Outside(nx, ny) {
			continue
		}
		if aoc.MustAtoi(g.Get(nx, ny)) == searchHeight {
			next = append(next, aoc.NewPairInt(nx, ny))
		}
	}
	log.Printf("next from %v = %v", start, next)
	return next
}
