package day4

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
		Name:    "day4A",
		Aliases: []string{"day4a"},
		Usage:   "Day 4 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day4B",
		Aliases: []string{"day4b"},
		Usage:   "Day 4 part B",
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
	cg := grid.Empty(g.Width(), g.Height())
	total := 0
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			total += countFrom(cg, g, x, y)
		}
	}
	DumpInColour(cg, g)
	return total, nil
}

func countFrom(cg, g *grid.Grid, x, y int) int {
	if g.Get(x, y) != "X" {
		// log.Printf("[%v,%v]=%v != X", x, y, g.Get(x, y))
		return 0
	}
	s := readInGrid(cg, g, "MAS", x, y, 1, 0)
	s += readInGrid(cg, g, "MAS", x, y, -1, 0)
	s += readInGrid(cg, g, "MAS", x, y, 0, 1)
	s += readInGrid(cg, g, "MAS", x, y, 0, -1)
	s += readInGrid(cg, g, "MAS", x, y, 1, 1)
	s += readInGrid(cg, g, "MAS", x, y, 1, -1)
	s += readInGrid(cg, g, "MAS", x, y, -1, 1)
	s += readInGrid(cg, g, "MAS", x, y, -1, -1)
	return s
}

func readInGrid(cg, g *grid.Grid, s string, x, y, dx, dy int) int {
	if g.Outside(x+len(s)*dx, y+len(s)*dy) {
		return 0
	}
	for n := 1; n <= len(s); n++ {
		if g.Get(x+n*dx, y+n*dy)[0] != s[n-1] {
			// log.Printf("[%v,%v]=%v != X", x+n*dx, y+n*dy, g.Get(x+n*dx, y+n*dy))
			return 0
		}
	}
	for n := 0; n <= len(s); n++ {
		cg.Set(x+n*dx, y+n*dy, aoc.Red)
	}
	return 1
}

func DumpInColour(cg, g *grid.Grid) {
	s := ""
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if cg.Get(x, y) == "." {
				s += aoc.Yellow
			} else {
				s += cg.Get(x, y)
			}
			s += g.Get(x, y)
		}
		s += "\n"
	}
	s += aoc.Reset
	log.Printf("\n%v", s)
}

func processB(puzzle []string) (int, error) {
	g := grid.New(puzzle)
	cg := grid.Empty(g.Width(), g.Height())
	total := 0
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			total += isX_MAS(cg, g, x, y)
		}
	}
	DumpInColour(cg, g)
	return total, nil
}

func isX_MAS(cg, g *grid.Grid, x, y int) int {
	if g.Get(x, y) != "A" {
		// log.Printf("[%v,%v]=%v != X", x, y, g.Get(x, y))
		return 0
	}
	if g.Outside(x+1, y+1) || g.Outside(x+1, y-1) || g.Outside(x-1, y+1) || g.Outside(x-1, y-1) {
		return 0
	}
	tl := g.Get(x-1, y-1)
	tr := g.Get(x+1, y-1)
	bl := g.Get(x-1, y+1)
	br := g.Get(x+1, y+1)
	fwd := tl + br
	bwd := bl + tr
	if (fwd == "MS" || fwd == "SM") && (bwd == "MS" || bwd == "SM") {
		cg.Set(x, y, aoc.Red)
		cg.Set(x-1, y-1, aoc.Red)
		cg.Set(x-1, y+1, aoc.Red)
		cg.Set(x+1, y-1, aoc.Red)
		cg.Set(x+1, y+1, aoc.Red)
		return 1
	}
	return 0
}
