package day8

import (
    _ "embed"
    "fmt"
    "log"
    "regexp"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 8
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day8A",
        Usage: "Day 8 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day8B",
        Usage: "Day 8 part B",
        Action: partB,
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
	moves := puzzle[0]
	nodes := make(map[string]*node)
	for i:=2; i<len(puzzle); i++ {
		loc, n := decode(puzzle[i])
		nodes[loc] = n
	}
	return makeMoves(moves, nodes)
}

func processB(puzzle []string) (int, error) {
	moves := puzzle[0]
	nodes := make(map[string]*node)
	var ghosts []string
	for i:=2; i<len(puzzle); i++ {
		loc, n, ghost := decodeGhosts(puzzle[i])
		if len(ghost) > 0 {
			ghosts = append(ghosts, ghost)
		}
		nodes[loc] = n
	}
	var pathLengths []int
	for _, ghost := range ghosts {
		pl := findEnd(ghost, moves, nodes)
		pathLengths = append(pathLengths, pl)
	}
	return aoc.LCM(pathLengths[0], pathLengths[1], pathLengths[2:]...), nil
}

type node struct {
	left, right string
}

func (n *node) String() string {
	return fmt.Sprintf("%v,%v", n.left, n.right)
}

func decode(line string) (string, *node) {
	// AAA = (BBB, CCC)
	re := regexp.MustCompile(`(...) = \((...), (...)\)`)
	matches := re.FindStringSubmatch(line)
	return matches[1], &node{left: matches[2], right: matches[3]}
}

func decodeGhosts(line string) (string, *node, string) {
	// AAA = (BBB, CCC)
	re := regexp.MustCompile(`(...) = \((...), (...)\)`)
	matches := re.FindStringSubmatch(line)
	ghost := ""
	if matches[1][2] == 'A' {
		ghost = matches[1]
	}
	return matches[1], &node{left: matches[2], right: matches[3]}, ghost
}

func makeMoves(moves string, nodes map[string]*node) (int, error) {
	step := -1
	stepCount := len(moves)
	left := "LR"[0]
	right := "LR"[1]
	loc := "AAA"
	for ; loc != "ZZZ"; {
		n, ok := nodes[loc]
		if !ok {
			return 0, fmt.Errorf("could not find node %v", loc)
		}
		step++
		switch moves[step % stepCount] {
		case left:
			log.Printf("moving left from %v to %v", loc, n.left)
			loc = n.left
		case right:
			log.Printf("moving right from %v to %v", loc, n.right)
			loc = n.right
		}
	}
	step++
	return step, nil
}

func findEnd(loc, moves string, nodes map[string]*node) int {
	step := -1
	stepCount := len(moves)
	left := "LR"[0]
	right := "LR"[1]
	for ; loc[2] != 'Z'; {
		n, ok := nodes[loc]
		if !ok {
			panic(fmt.Sprintf("could not find node %v", loc))
		}
		step++
		switch moves[step % stepCount] {
		case left:
			log.Printf("moving left from %v to %v", loc, n.left)
			loc = n.left
		case right:
			log.Printf("moving right from %v to %v", loc, n.right)
			loc = n.right
		}
	}
	step++
	return step
}
