package main

import (
	"flag"
	"log"
	"strings"

	"github.com/baldrick/aoc/2023/aoc"
)

const (
	year = 2023
	day = 9
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

/* 0 3 6 9 12 15
3...

1 3 6 10 15 21
2,3,4,5...

10 13 16 21 30 45
3,3,5,9,15
0,2,4,6
2,2,2
*/

func process(puzzle []string) error {
	total := 0
	for _, line := range puzzle {
		next := nextNumberFromStringB(line)
		logger.Printf("%v -> %v", line, next)
		total += next
	}
	logger.Printf("total=%v", total)
	return nil
}

func nextNumberFromString(line string) int {
	return nextNumber(aoc.MustAtoiAll(strings.Split(line, " ")))
}

func nextNumberFromStringB(line string) int {
	return nextNumber(reverse(aoc.MustAtoiAll(strings.Split(line, " "))))
}

func reverse(numbers []int) []int {
	r := make([]int, len(numbers))
	for i:=0; i<len(numbers); i++ {
		r[len(numbers)-1-i] = numbers[i]
	}
	return r
}

func nextNumber(numbers []int) int {
	diffs, allZero := getDiffs(numbers)
	if allZero {
		return numbers[len(numbers)-1]
	}
	n := numbers[len(numbers)-1] + nextNumber(diffs)
	logger.Printf("diffs: %v -> %v", numbers, n)
	return n
}

// Had to resort to reddit for part A of this to find the one (!) part
// of the puzzle where I was wrong.  Root cause was the first implementation
// of "allZero" being dumb by virtue of trying to be "clever" (lazy).  D'oh!
func getDiffs(numbers []int) ([]int, bool) {
	var diffs []int
	allZero := true
	for n:=1; n<len(numbers); n++ {
		diff := numbers[n] - numbers[n-1]
		diffs = append(diffs, diff)
		allZero = allZero && diff == 0
	}
	return diffs, allZero
}
