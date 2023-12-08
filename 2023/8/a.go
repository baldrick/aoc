package main

import (
	"fmt"
	"flag"
	"log"
	"regexp"

	"github.com/baldrick/aoc/2023/aoc"
)

const (
	year = 2023
	day = 8
)

var (
	inputFile = flag.String("f", "test", "Puzzle file (partial name) to use")
	logger = log.Default()
)

func main() {
	flag.Parse()
	puzzle, err := aoc.GetPuzzleInput(*inputFile, year, day)
	if err != nil {
		logger.Fatalf("oops: %v\n", err)
		return
	}
	if err := process(puzzle); err != nil {
		logger.Fatalf("oops: %v\n", err)
	}
}

func process(puzzle []string) error {
	moves := puzzle[0]
	nodes := make(map[string]*node)
	for i:=2; i<len(puzzle); i++ {
		loc, n := decode(puzzle[i])
		nodes[loc] = n
	}
	return makeMoves(moves, nodes)
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

func makeMoves(moves string, nodes map[string]*node) error {
	step := -1
	stepCount := len(moves)
	left := "LR"[0]
	right := "LR"[1]
	loc := "AAA"
	for ; loc != "ZZZ"; {
		n, ok := nodes[loc]
		if !ok {
			return fmt.Errorf("could not find node %v", loc)
		}
		step++
		switch moves[step % stepCount] {
		case left:
			logger.Printf("moving left from %v to %v", loc, n.left)
			loc = n.left
		case right:
			logger.Printf("moving right from %v to %v", loc, n.right)
			loc = n.right
		}
	}
	step++
	logger.Printf("#steps: %v", step)
	return nil
}
