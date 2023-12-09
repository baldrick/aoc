package day9

import (
    "fmt"
    "log"
    "strings"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 9
)

var (
    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day9A",
        Usage: "Day 9 part A",
        Flags: aoc.Flags,
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day9B",
        Usage: "Day 9 part B",
        Flags: aoc.Flags,
        Action: partB,
    }
    logger = log.Default()
)

func partA(ctx *cli.Context) error {
    puzzle, err := aoc.GetPuzzleInput(ctx.String("f"), year, day)
    if err != nil {
        return fmt.Errorf("oops: %v\n", err)
    }
    answer, err := processA(puzzle)
    if err != nil {
        return err
    }
    log.Printf("Answer A: %v", answer)
    return nil
}

func partB(ctx *cli.Context) error {
    puzzle, err := aoc.GetPuzzleInput(ctx.String("f"), year, day)
    if err != nil {
        return fmt.Errorf("oops: %v\n", err)
    }

    answer, err := processB(puzzle)
    if err != nil {
        return err
    }
    log.Printf("Answer B: %v", answer)
    return nil
}

func processA(puzzle []string) (int, error) {
    total := 0
    for _, line := range puzzle {
        next := nextNumberFromString(line)
        logger.Printf("%v -> %v", line, next)
        total += next
    }
    return total, nil
}

func processB(puzzle []string) (int, error) {
    total := 0
    for _, line := range puzzle {
        next := nextNumberFromStringB(line)
        logger.Printf("%v -> %v", line, next)
        total += next
    }
    return total, nil
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
