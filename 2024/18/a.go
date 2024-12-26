package day18

import (
    _ "embed"
    "fmt"
    "log"

    "github.com/baldrick/aoc/common/aoc"
    "github.com/urfave/cli"
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day18A",
        Aliases: []string{"day18a"},
        Usage: "Day 18 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day18B",
        Aliases: []string{"day18b"},
        Usage: "Day 18 part B",
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
    for _, line := range puzzle {
        log.Print(line)
    }
    return 0, fmt.Errorf("Not yet implemented")
}

func processB(puzzle []string) (int, error) {
    return 0, fmt.Errorf("Not yet implemented")
}
