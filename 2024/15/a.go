package day15

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/baldrick/aoc/common/grid"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day15A",
		Aliases: []string{"day15a"},
		Usage:   "Day 15 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day15B",
		Aliases: []string{"day15b"},
		Usage:   "Day 15 part B",
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
	g := readGrid(puzzle)
	g.Dump()
	moves := readMoves(puzzle)
	rx, ry := g.Find("@")
	robot := aoc.NewPairInt(rx, ry)
	applyMoves(g, moves, robot, moveRobot1)
	objects := g.FindAll("O")
	log.Printf("objects found at %v", objects)
	sum := 0
	for _, o := range objects {
		sum += o.X() + o.Y()*100
	}

	return sum, nil
}

func processB(puzzle []string) (int, error) {
	g := readGrid(puzzle)
	g.Dump()
	g = doubleGrid(g)
	g.Dump()
	moves := readMoves(puzzle)
	rx, ry := g.Find("@")
	robot := aoc.NewPairInt(rx, ry)
	applyMoves(g, moves, robot, moveRobot2)
	objects := g.FindAll("[")
	log.Printf("objects found at %v", objects)
	sum := 0
	for _, o := range objects {
		sum += o.X() + o.Y()*100
	}

	return sum, nil
}

func readGrid(puzzle []string) *grid.Grid {
	n := 0
	for ; len(puzzle[n]) > 0; n++ {
	}
	g := grid.Empty(len(puzzle[0]), n)
	for n = 0; len(puzzle[n]) > 0; n++ {
		for x := 0; x < len(puzzle[n]); x++ {
			g.Set(x, n, string(puzzle[n][x]))
		}
	}
	return g
}

func doubleGrid(g *grid.Grid) *grid.Grid {
	ng := grid.Empty(g.Width()*2, g.Height())
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Height(); y++ {
			c := g.Get(x, y)
			if c == "O" {
				ng.Set(x*2, y, "[")
				ng.Set(x*2+1, y, "]")
			} else if c == "@" {
				ng.Set(x*2, y, "@")
				ng.Set(x*2+1, y, ".")
			} else {
				ng.Set(x*2, y, c)
				ng.Set(x*2+1, y, c)
			}
		}
	}
	return ng
}

func readMoves(puzzle []string) string {
	n := 0
	for ; len(puzzle[n]) > 0; n++ {
	}
	return strings.Join(puzzle[n:], "")
}

func applyMoves(g *grid.Grid, moves string, robot aoc.PairInt, moveRobot func(g *grid.Grid, robot aoc.PairInt, dx, dy int) aoc.PairInt) {
	for _, move := range moves {
		robot = applyMove(g, string(move), robot, moveRobot)
		g.DumpMsg(fmt.Sprintf("move %v", string(move)))
	}
}

func applyMove(g *grid.Grid, move string, robot aoc.PairInt, moveRobot func(g *grid.Grid, robot aoc.PairInt, dx, dy int) aoc.PairInt) aoc.PairInt {
	//log.Printf("moving %v", move)
	switch move {
	case "^":
		return moveRobot(g, robot, 0, -1)
	case "v":
		return moveRobot(g, robot, 0, 1)
	case "<":
		return moveRobot(g, robot, -1, 0)
	case ">":
		return moveRobot(g, robot, 1, 0)
	default:
		log.Panicf("unhandled move %q", move)
	}
	return robot
}

func moveRobot1(g *grid.Grid, robot aoc.PairInt, dx, dy int) aoc.PairInt {
	nx, ny := robot.X()+dx, robot.Y()+dy
	switch g.Get(nx, ny) {
	case ".":
		g.Set(robot.X(), robot.Y(), "")
		robot = aoc.NewPairInt(nx, ny)
		g.Set(nx, ny, "@")
	case "#":
		return robot
	case "O":
		// Find first empty space in the direction we're moving.
		sx, sy := findSpace(g, robot.X(), robot.Y(), dx, dy)
		if sx >= 0 && sy >= 0 {
			//log.Printf("space found at %v,%v from %v,%v in direction %v,%v", sx, sy, robot.X(), robot.Y(), dx, dy)
			g.Set(sx, sy, "O")
			g.Set(robot.X(), robot.Y(), "")
			robot = aoc.NewPairInt(nx, ny)
			g.Set(nx, ny, "@")
		}
		return robot
	default:
		log.Panicf("unhandled maze item %q", g.Get(nx, ny))
	}
	return robot
}

func moveRobot2(g *grid.Grid, robot aoc.PairInt, dx, dy int) aoc.PairInt {
	nx, ny := robot.X()+dx, robot.Y()+dy
	switch g.Get(nx, ny) {
	case ".":
		g.Set(robot.X(), robot.Y(), "")
		robot = aoc.NewPairInt(nx, ny)
		g.Set(nx, ny, "@")
	case "#":
		return robot
	case "[", "]":
		// Find first empty space in the direction we're moving.
		sx, sy := findSpace(g, robot.X(), robot.Y(), dx, dy)
		if sx >= 0 && sy >= 0 {
			log.Printf("space found at %v,%v from %v,%v in direction %v,%v", sx, sy, robot.X(), robot.Y(), dx, dy)
			if dx != 0 {
				// We're moving in the x direction, flip [] to ][...
				if dx < 0 {
					// We're moving left so replace space with [.
					g.Set(sx, sy, "[")
				} else {
					// We're moving right so replace space with ].
					g.Set(sx, sy, "]")
				}
				flip(g, robot.X()+dx, sx-dx, robot.Y(), dx)
			} else {
				// We're moving in the y direction.  Handle overlapping boxes.
				// Needs to be recursive to move a pile of boxes?
			}
			g.Set(robot.X(), robot.Y(), "")
			robot = aoc.NewPairInt(nx, ny)
			g.Set(nx, ny, "@")
		}
		return robot
	default:
		log.Panicf("unhandled maze item %q", g.Get(nx, ny))
	}
	return robot
}

func flip(g *grid.Grid, startX, endX, y, dx int) {
	for x := startX; x != endX; x += dx {
		if g.Get(x, y) == "[" {
			log.Printf("[ becomes ] at %v,%v", x, y)
			g.Set(x, y, "]")
		} else {
			log.Printf("] becomes [ at %v,%v", x, y)
			g.Set(x, y, "[")
		}
	}
}

func findSpace(g *grid.Grid, rx, ry, dx, dy int) (int, int) {
	for x, y := rx+dx, ry+dy; !g.Outside(x, y); x, y = x+dx, y+dy {
		switch g.Get(x, y) {
		case "#":
			return -1, -1
		case ".":
			return x, y
		case "O", "[", "]":
			// do nothing
		default:
			log.Panicf("unhandled maze item %q", g.Get(x, y))
		}
	}
	return -1, -1
}
