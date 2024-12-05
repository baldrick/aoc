package day2

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day2A",
		Aliases: []string{"day2a"},
		Usage:   "Day 2 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day2B",
		Aliases: []string{"day2b"},
		Usage:   "Day 2 part B",
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
	safe := 0
	for _, line := range puzzle {
		r := createRecord(strings.Split(line, " "))
		if r.isSafe() {
			safe++
		}
	}
	return safe, nil
}

func processB(puzzle []string) (int, error) {
	safe := 0
	for _, line := range puzzle {
		r := createRecord(strings.Split(line, " "))
		if r.isSafe() {
			safe++
			continue
		}
		for n := 0; n < len(r.numbers); n++ {
			r2 := createRecord(strings.Split(line, " "))
			start := r2.numbers[:n]
			end := r2.numbers[n+1:]
			var n []int
			n = append(append(n, start...), end...)
			r2.numbers = n
			if r2.isSafe() {
				safe++
				break
			}
		}
	}
	return safe, nil
}

type record struct {
	numbers   []int
	direction bool
}

func createRecord(s []string) record {
	return record{numbers: aoc.MustAtoiAll(s)}
}

func (r record) createSubrecord(first, remainderStartIndex int) record {
	var n []int
	n = append(n, first)
	n = append(n, r.numbers[remainderStartIndex:]...)
	log.Printf("subrecord: %v", n)
	return record{numbers: n, direction: first > r.numbers[remainderStartIndex]}
}

func (r record) String() string {
	return fmt.Sprintf("%v", r.numbers)
}

func (r record) isSafe() bool {
	log.Printf("%v - checking", r.numbers)
	r.direction = r.numbers[0] > r.numbers[1]
	for i := 1; i < len(r.numbers); i++ {
		if r.safePair(i, i-1, i) {
			continue
		}
		return false
	}
	log.Printf("%v - safe", r.numbers)
	return true
}

func (r record) safePair(i, first, second int) bool {
	aDiff := aoc.AbsInt(r.numbers[first] - r.numbers[second])
	safe := aDiff >= 1 && aDiff <= 3 && (r.direction == (r.numbers[first] > r.numbers[second]))
	log.Printf("i=%v, %v vs %v (direction=%v): %v", i, r.numbers[first], r.numbers[second], r.direction, safeStr(safe))
	return safe
}

func safeStr(s bool) string {
	ss := "unsafe"
	if s {
		ss = "safe"
	}
	return ss
}
