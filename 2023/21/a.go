package day21

import (
    _ "embed"
    "fmt"
    "log"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/baldrick/aoc/2023/grid"
    "github.com/urfave/cli"
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day21A",
        Aliases: []string{"day21a"},
        Usage: "Day 21 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day21B",
        Aliases: []string{"day21b"},
        Usage: "Day 21 part B",
        Action: partB,
    }
)

func partA(ctx *cli.Context) error {
    answer, err := processA(aoc.PreparePuzzle(puzzle), 64)
    if err != nil {
        return err
    }
    log.Printf("Answer A: %v", answer)
    return nil
}

func partB(ctx *cli.Context) error {
    answer, err := processB(aoc.PreparePuzzle(puzzle), 64)
    if err != nil {
        return err
    }
    log.Printf("Answer B: %v", answer)
    return nil
}

func processA(puzzle []string, steps int) (int, error) {
    g := grid.New(puzzle)
    x, y := g.Find("S")
    if x<0 || y<0 {
        panic(fmt.Sprintf("Could not find 'S' in grid:\n%v", g.String()))
    }
    g.FillN(x,y,"#",0,steps+1)
    // for n := 0;  n <= steps;  n++ {
    //     g.Replace(fmt.Sprintf("%v",n), ".")
    // }
    g.Dump()
    return g.CountIf(fmt.Sprintf("%v",steps)), nil
}

func processB(puzzle []string, steps int) (int, error) {
    return 0, fmt.Errorf("Not yet implemented")
}
