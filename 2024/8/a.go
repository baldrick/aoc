package day8

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
		Name:    "day8A",
		Aliases: []string{"day8a"},
		Usage:   "Day 8 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day8B",
		Aliases: []string{"day8b"},
		Usage:   "Day 8 part B",
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
	g.DumpMsg("puzzle")
	antennaLocations := make(map[string][]aoc.PairInt)
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			freq := g.Get(x, y)
			if freq == "." {
				continue
			}
			locations, _ := antennaLocations[freq]
			locations = append(locations, aoc.NewPairInt(x, y))
			antennaLocations[freq] = locations
		}
	}
	antinodesByFrequency := findAntinodesByFrequency(g, antennaLocations, addAntinode)
	antinodes := aoc.NewStringSet()
	antinodesGraph := grid.New(puzzle)
	for f, locations := range antinodesByFrequency {
		faGraph := grid.New(puzzle)
		for _, location := range locations {
			antinodes.Add(location.String())
			faGraph.Set(location.X(), location.Y(), "#")
			antinodesGraph.Set(location.X(), location.Y(), "#")
		}
		faGraph.DumpMsg(fmt.Sprintf("antinodes for %v", f))
	}
	antinodesGraph.DumpMsg("all antinodes")
	return antinodes.Len(), nil
}

func processB(puzzle []string) (int, error) {
	g := grid.New(puzzle)
	g.DumpMsg("puzzle")
	antennaLocations := make(map[string][]aoc.PairInt)
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			freq := g.Get(x, y)
			if freq == "." {
				continue
			}
			locations, _ := antennaLocations[freq]
			locations = append(locations, aoc.NewPairInt(x, y))
			antennaLocations[freq] = locations
		}
	}
	antinodesByFrequency := findAntinodesByFrequency(g, antennaLocations, addMultiAntinodes)
	antinodes := aoc.NewStringSet()
	antinodesGraph := grid.New(puzzle)
	for f, locations := range antinodesByFrequency {
		faGraph := grid.New(puzzle)
		for _, location := range locations {
			antinodes.Add(location.String())
			faGraph.Set(location.X(), location.Y(), "#")
			antinodesGraph.Set(location.X(), location.Y(), "#")
		}
		faGraph.DumpMsg(fmt.Sprintf("antinodes for %v", f))
	}
	antinodesGraph.DumpMsg("all antinodes")
	return antinodes.Len(), nil
}

func findAntinodesByFrequency(g *grid.Grid, antennaLocations map[string][]aoc.PairInt, addAntinodes func(g *grid.Grid, antinodesByFrequency map[string][]aoc.PairInt, frequency string, start, end aoc.PairInt)) map[string][]aoc.PairInt {
	antinodesByFrequency := make(map[string][]aoc.PairInt)
	for f, locations := range antennaLocations {
		log.Printf("Finding antinodes for %v", f)
		for n := 0; n < len(locations)-1; n++ {
			for n2 := n + 1; n2 < len(locations); n2++ {
				addAntinodes(g, antinodesByFrequency, f, locations[n], locations[n2])
			}
		}
		log.Printf("antinodes for %v = %v", f, antinodesByFrequency[f])
	}
	return antinodesByFrequency
}

func addAntinode(g *grid.Grid, antinodesByFrequency map[string][]aoc.PairInt, frequency string, start, end aoc.PairInt) {
	locations, _ := antinodesByFrequency[frequency]
	dx := start.X() - end.X()
	dy := start.Y() - end.Y()
	a1 := aoc.NewPairInt(start.X()+dx, start.Y()+dy)
	a2 := aoc.NewPairInt(start.X()-dx, start.Y()-dy)
	a3 := aoc.NewPairInt(end.X()-dx, end.Y()-dy)
	a4 := aoc.NewPairInt(end.X()-dx, end.Y()-dy)
	if !g.Outside(a1.X(), a1.Y()) && !start.Equals(a1) && !end.Equals(a1) {
		locations = append(locations, a1)
	}
	if !g.Outside(a2.X(), a2.Y()) && !start.Equals(a2) && !end.Equals(a2) {
		locations = append(locations, a2)
	}
	if !g.Outside(a3.X(), a3.Y()) && !start.Equals(a3) && !end.Equals(a3) {
		locations = append(locations, a3)
	}
	if !g.Outside(a4.X(), a4.Y()) && !start.Equals(a4) && !end.Equals(a4) {
		locations = append(locations, a4)
	}
	antinodesByFrequency[frequency] = locations
}

func addMultiAntinodes(g *grid.Grid, antinodesByFrequency map[string][]aoc.PairInt, frequency string, start, end aoc.PairInt) {
	locations, _ := antinodesByFrequency[frequency]
	dx := start.X() - end.X()
	dy := start.Y() - end.Y()
	locations = antinodesByFrequency[frequency]
	nx := start.X()
	ny := start.Y()
	for {
		nx += dx
		ny += dy
		if g.Outside(nx, ny) {
			break
		}
		nl := aoc.NewPairInt(nx, ny)
		locations = append(locations, nl)
	}
	nx = start.X()
	ny = start.Y()
	for {
		nx -= dx
		ny -= dy
		if g.Outside(nx, ny) {
			break
		}
		nl := aoc.NewPairInt(nx, ny)
		locations = append(locations, nl)
	}
	locations = append(locations, start)
	locations = append(locations, end)
	antinodesByFrequency[frequency] = locations
}
