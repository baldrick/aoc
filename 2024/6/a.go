package day6

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
		Name:    "day6A",
		Aliases: []string{"day6a"},
		Usage:   "Day 6 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day6B",
		Aliases: []string{"day6b"},
		Usage:   "Day 6 part B",
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
	x, y := g.Find("^")
	visited, _ := walk(g, x, y)
	return visited.Len(), nil
}

func processB(puzzle []string) (int, error) {
	g := grid.New(puzzle)
	sx, sy := g.Find("^")
	loopCount := 0
	// If the guard never visits a square we could put an obstacle,
	// they'll never come across the obstacle so it'd make no difference.
	// So don't even bother in those cases.  Find the path with no extra obstructions.
	// That also handily filters out cells where there's already an obstruction.
	//unobstructed, _ := walk(g, sx, sy)
	for x := 0; x < g.Width(); x++ {
		log.Printf("x=%v", x)
		for y := 0; y < g.Height(); y++ {
			if x == sx && y == sy {
				continue
			}
			// if !unobstructed.Contains(xyKey(x, y)) {
			// 	continue
			// }
			if g.Get(x, y) == "#" {
				continue
			}
			g.Set(x, y, "O")
			_, loop := walk(g, sx, sy)
			if loop {
				loopCount++
				log.Printf("x,y=%v,%v (%v loops found)", x, y, loopCount)
				// ng := g.Clone()
				// for lx := 0; lx < ng.Width(); lx++ {
				// 	for ly := 0; ly < ng.Height(); ly++ {
				// 		if lx == sx && ly == sy {
				// 			continue
				// 		}
				// 		upDown := visited.Contains(xydKey(lx, ly, 0, 1)) || visited.Contains(xydKey(lx, ly, 0, -1))
				// 		leftRight := visited.Contains(xydKey(lx, ly, 1, 0)) || visited.Contains(xydKey(lx, ly, -1, 0))
				// 		if upDown && leftRight {
				// 			ng.Set(lx, ly, "+")
				// 		} else if upDown {
				// 			ng.Set(lx, ly, "|")
				// 		} else if leftRight {
				// 			ng.Set(lx, ly, "-")
				// 		}
				// 	}
				// }
				// ng.Dump()
			}
			g.Set(x, y, ".")
		}
	}
	// 1484 is too low (so is 1485, thought I might be off by one)
	// 1553 is too low
	return loopCount, nil
}

func walk(g *grid.Grid, x, y int) (*aoc.StringSet, bool) {
	dx, dy := 0, -1
	visited := aoc.NewStringSet()
	visitedWithDirection := aoc.NewStringSet()
	for {
		visited.Add(xyKey(x, y))
		vwd := xydKey(x, y, dx, dy)
		if visitedWithDirection.Contains(vwd) {
			return visitedWithDirection, true
		}
		visitedWithDirection.Add(vwd)
		nx, ny := x+dx, y+dy
		if g.Outside(nx, ny) {
			break
		}
		if g.Get(nx, ny) == "#" || g.Get(nx, ny) == "O" {
			if dx == 0 {
				// We've been moving up/down.
				if dy < 0 {
					// We've been moving up so now move right.
					dx, dy = 1, 0
				} else {
					// We've been moving down so now move left.
					dx, dy = -1, 0
				}
			} else {
				// We've been moving left/right.
				if dx < 0 {
					// We've been moving left so now move up.
					dx, dy = 0, -1
				} else {
					// We've been moving right so now move down.
					dx, dy = 0, 1
				}
			}
			vwd = xydKey(x, y, dx, dy)
			if visitedWithDirection.Contains(vwd) {
				return visitedWithDirection, true
			}
			visitedWithDirection.Add(vwd)
		}
		x, y = x+dx, y+dy
	}
	// for x := 0; x < g.Width(); x++ {
	// 	for y := 0; y < g.Height(); y++ {
	// 		if visited.Contains(xyKey(x, y)) {
	// 			g.Set(x, y, "X")
	// 		}
	// 	}
	// }
	// g.Dump()
	return visited, false
}

func xyKey(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func xydKey(x, y, dx, dy int) string {
	return fmt.Sprintf("%v,%v,%v,%v", x, y, dx, dy)
}
