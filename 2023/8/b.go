package main

import (
	"fmt"
	"flag"
	"log"
	"regexp"

	"github.com/baldrick/aoc/AdventOfCode/2023/aoc"
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
	var ghosts []string
	for i:=2; i<len(puzzle); i++ {
		loc, n, ghost := decode(puzzle[i])
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
	lcm := LCM(pathLengths[0], pathLengths[1], pathLengths[2:]...)
	logger.Printf("path lengths: %v, lcm: %v", pathLengths, lcm)
	return nil
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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
			logger.Printf("moving left from %v to %v", loc, n.left)
			loc = n.left
		case right:
			logger.Printf("moving right from %v to %v", loc, n.right)
			loc = n.right
		}
	}
	step++
	return step
}

type node struct {
	left, right string
}

func (n *node) String() string {
	return fmt.Sprintf("%v,%v", n.left, n.right)
}

func decode(line string) (string, *node, string) {
	// AAA = (BBB, CCC)
	re := regexp.MustCompile(`(...) = \((...), (...)\)`)
	matches := re.FindStringSubmatch(line)
	ghost := ""
	if matches[1][2] == 'A' {
		ghost = matches[1]
	}
	return matches[1], &node{left: matches[2], right: matches[3]}, ghost
}

func makeMoves(ghosts []string, moves string, nodes map[string]*node) error {
	step := -1
	stepCount := len(moves)
	left := "LR"[0]
	right := "LR"[1]
	for ; !zzz(ghosts); {
		step++
		move := moves[step % stepCount]
		for g, ghost := range ghosts {
			n, ok := nodes[ghost]
			if !ok {
				return fmt.Errorf("could not find node %v", ghost)
			}
			switch move {
			case left:
				if step % 10_000_000 == 0 {
					logger.Printf("step %v ghost %v moving left from %v to %v", step, g, ghost, n.left)
				}
				ghosts[g] = n.left
			case right:
				if step % 10_000_000 == 0 {
					logger.Printf("step %v ghost %v moving right from %v to %v", step, g, ghost, n.right)
				}
				ghosts[g] = n.right
			}
		}
	}
	step++
	logger.Printf("#steps: %v", step)
	return nil
}

func zzz(ghosts []string) bool {
	for _, ghost := range ghosts {
		if ghost[2] != 'Z' {
			return false
		}
	}
	return true
}