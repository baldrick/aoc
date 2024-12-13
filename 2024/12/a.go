package day12

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/baldrick/aoc/common/grid"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day12A",
		Aliases: []string{"day12a"},
		Usage:   "Day 12 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day12B",
		Aliases: []string{"day12b"},
		Usage:   "Day 12 part B",
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
	g = g.AddBorder("-")
	garden := Garden{g: g}
	garden.g.Dump()
	sum := 0
	for x := 1; x < garden.g.Width()-1; x++ {
		for y := 1; y < garden.g.Height()-1; y++ {
			garden.Reset(g)
			plant := garden.g.Get(x, y)
			garden.Flood(x, y, 0, 0, plant)
			sum += garden.area * garden.perimeter
			log.Printf("%v area %v x perimeter %v = %v", plant, garden.area, garden.perimeter, garden.area*garden.perimeter)
			garden.g.Dump()
		}
	}
	return sum, nil
}

func processB(puzzle []string) (int, error) {
	return 0, fmt.Errorf("Not yet implemented")
}

type Garden struct {
	g                     *grid.Grid
	area, perimeter       int
	visited, plantVisited *aoc.StringSet
}

func (gn *Garden) Reset(g *grid.Grid) {
	gn.g = g.Clone()
	gn.area, gn.perimeter = 0, 0
	gn.visited = aoc.NewStringSet()
	if gn.plantVisited == nil {
		gn.plantVisited = aoc.NewStringSet()
	}
}

func (gn *Garden) Flood(x, y, dx, dy int, plant string) {
	nx, ny := x+dx, y+dy
	loc := gn.loc(x, y, nx, ny)
	// if gn.g.Outside(x, y) {
	// 	log.Printf("%v is outside, visited=%v", loc, gn.visited)
	// } else {
	// 	log.Printf("%v = %v, visited=%v", loc, gn.g.Get(x, y), gn.visited)
	// }

	if gn.visited.Contains(loc) {
		//log.Printf("already visited %v", loc)
		return
	}
	log.Printf("not visited %v, visited: %v", loc, gn.visited)
	gn.visited.Add(loc)
	if gn.g.Outside(nx, ny) || gn.g.Get(nx, ny) != plant {
		// if gn.g.Outside(x, y) {
		// 	log.Printf("++ perimeter at %v (outside)", loc)
		// } else {
		log.Printf("++ perimeter at %v (%v)", loc, gn.g.Get(nx, ny))
		n, err := strconv.Atoi(gn.g.Get(nx, ny))
		if err != nil {
			n = 0
		}
		gn.g.Set(nx, ny, fmt.Sprintf("%v", n+1))
		//}
		gn.perimeter++
		return
	}
	gn.plantVisited.Add(loc)
	gn.area++
	gn.g.Set(nx, ny, ".")
	gn.Flood(nx, ny, 1, 0, plant)
	gn.Flood(nx, ny, -1, 0, plant)
	gn.Flood(nx, ny-1, 0, -1, plant)
	gn.Flood(nx, ny+1, 0, 1, plant)
}

func (gn *Garden) loc(x, y, nx, ny int) string {
	return fmt.Sprintf("%v.%v-%v.%v", x, y, nx, ny)
}
