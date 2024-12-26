package day20

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
        Name:  "day20A",
        Aliases: []string{"day20a"},
        Usage: "Day 20 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day20B",
        Aliases: []string{"day20b"},
        Usage: "Day 20 part B",
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
