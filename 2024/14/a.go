package day14

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/baldrick/aoc/common/grid"
	"github.com/baldrick/aoc/common/terminal"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	startSecond int
	delay       time.Duration

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day14A",
		Aliases: []string{"day14a"},
		Usage:   "Day 14 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day14B",
		Aliases: []string{"day14b"},
		Usage:   "Day 14 part B",
		Action:  partB,
		Flags: []cli.Flag{
			cli.IntFlag{Name: "start", Destination: &startSecond},
			cli.DurationFlag{Name: "delay", Destination: &delay},
		},
	}
)

func partA(ctx *cli.Context) error {
	answer, err := processA(aoc.PreparePuzzle(puzzle), aoc.NewPairInt(101, 103))
	if err != nil {
		return err
	}
	log.Printf("Answer A: %v", answer)
	return nil
}

func partB(ctx *cli.Context) error {
	answer, err := processB(aoc.PreparePuzzle(puzzle), aoc.NewPairInt(101, 103))
	if err != nil {
		return err
	}
	log.Printf("Answer B: %v", answer)
	return nil
}

func processA(puzzle []string, size aoc.PairInt) (int, error) {
	var robots []robot
	for _, line := range puzzle {
		if len(line) == 0 {
			continue
		}
		robots = append(robots, newRobot(line, size))
	}
	var movedRobots []robot
	for _, robot := range robots {
		movedRobots = append(movedRobots, robot.move(100, size))
	}
	//dumpRobots(robots, size)
	//dumpRobots(movedRobots, size)
	robots = movedRobots
	xMid := ((size.X() - 1) / 2) - 1
	yMid := ((size.Y() - 1) / 2) - 1
	c1 := countQuadrant(robots, 0, 0, xMid, yMid, size)                   // top left
	c2 := countQuadrant(robots, xMid+2, 0, size.X(), yMid, size)          // top right
	c3 := countQuadrant(robots, 0, yMid+2, xMid, size.Y(), size)          // bottom left
	c4 := countQuadrant(robots, xMid+2, yMid+2, size.X(), size.Y(), size) // bottom right
	return c1 * c2 * c3 * c4, nil
}

func processB(puzzle []string, size aoc.PairInt) (int, error) {
	var robots []robot
	for _, line := range puzzle {
		if len(line) == 0 {
			continue
		}
		robots = append(robots, newRobot(line, size))
	}
	if startSecond > 0 {
		robots = moveAll(robots, startSecond, size)
	}
	moves := startSecond
	for {
		simpleDumpRobots(robots, size)
		log.Printf("moves: %v", moves)
		time.Sleep(delay)
		robots = moveAll(robots, 1, size)
		moves++
	}
	// This puzzle required (at least for me) visually looking
	// at the grid after each move hence the introduction of
	// start and delay flags to control the start point and delay
	// between frames.  Zoom out on a suitably sized terminal and
	// watch.  Spot the tree, iterate until the exact frame is found.
	return moves, fmt.Errorf("Failed to find solution")
}

func moveAll(robots []robot, n int, size aoc.PairInt) []robot {
	var movedRobots []robot
	for _, robot := range robots {
		movedRobots = append(movedRobots, robot.move(n, size))
	}
	return movedRobots
}

type robot struct {
	position, velocity aoc.PairInt
}

func newRobot(line string, size aoc.PairInt) robot {
	// p=0,4 v=3,-3
	re := regexp.MustCompile(`p=(-?[0-9]+),(-?[0-9]+) v=(-?[0-9]+),(-?[0-9]+)`)
	r := re.FindStringSubmatch(line)
	//log.Printf("%v matches: %v", line, r)
	vx, vy := aoc.MustAtoi(r[3]), aoc.MustAtoi(r[4])
	if vx < 0 {
		vx += size.X()
	}
	if vy < 0 {
		vy += size.Y()
	}
	return robot{
		position: aoc.NewPairInt(aoc.MustAtoi(r[1]), aoc.MustAtoi(r[2])),
		velocity: aoc.NewPairInt(vx, vy),
	}
}

func (r *robot) move(n int, size aoc.PairInt) robot {
	nx := aoc.ModInt(r.position.X()+n*r.velocity.X(), size.X())
	ny := aoc.ModInt(r.position.Y()+n*r.velocity.Y(), size.Y())
	//log.Printf("moving from %v to %v,%v", r.position, nx, ny)
	return robot{position: aoc.NewPairInt(nx, ny), velocity: r.velocity}
}

func countQuadrant(robots []robot, sx, sy, ex, ey int, size aoc.PairInt) int {
	//log.Printf("checking %v,%v - %v,%v: %v", sx, sy, ex, ey, robots)
	count := 0
	//dumpRobots(robots, size)
	for _, r := range robots {
		rx := r.position.X()
		ry := r.position.Y()
		if rx >= sx && rx <= ex && ry >= sy && ry <= ey {
			count++
		}
	}
	return count
}

func dumpRobots(robots []robot, size aoc.PairInt) {
	g := grid.Empty(size.X(), size.Y())
	for _, r := range robots {
		rx := r.position.X()
		ry := r.position.Y()
		g.Increment(rx, ry)
	}
	g.Dump()
}

func simpleDumpRobots(robots []robot, size aoc.PairInt) {
	g := grid.Empty(size.X(), size.Y())
	for _, r := range robots {
		rx := r.position.X()
		ry := r.position.Y()
		g.Set(rx, ry, "#")
	}
	log.Print(terminal.Home)
	g.Dump()
}
