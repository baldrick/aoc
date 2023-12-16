package day16

import (
    _ "embed"
    "log"
    "sync"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/baldrick/aoc/2023/16/lbgrid"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 16
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day16A",
        Usage: "Day 16 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day16B",
        Usage: "Day 16 part B",
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
    g := lbgrid.New(puzzle)
    g.Dump()
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {defer wg.Done(); lightbeam(&wg, g, 0,0, 1,0)}()
    wg.Wait()
    return g.Energized(), nil
}

func lightbeam(wg *sync.WaitGroup, g *lbgrid.Grid, x,y, dx,dy int) {
    log.Printf("lightbeam cell %v,%v moving %v,%v", x,y, dx,dy)
    for ;!g.Outside(x,y); {
        d := lbgrid.Direction{dx,dy}
        if _, ok := g.Get(x,y).DirectionsSeen[d.Hash()]; ok {
            // A light beam is already travelling from this cell in the direction we need.
            log.Printf("beam already at %v,%v travelling %v", x,y, d)
            return
        }
        g.Get(x,y).DirectionsSeen[d.Hash()] = d
        log.Printf("processing %v,%v (%v)", x,y, g.Get(x,y).Cell)
        switch g.Get(x,y).Cell {
        case ".":
            // Do nothing
        case "-":
            if dy!=0 {
                log.Printf("Split left/right")
                // Moving vertically, split the beam!
                wg.Add(1)
                go func() {defer wg.Done(); lightbeam(wg, g, x-1,y, -1,0)}()
                wg.Add(1)
                go func() {defer wg.Done(); lightbeam(wg, g, x+1,y, 1,0)}()
                return
            }
        case "|":
            if dx != 0 {
                log.Printf("Split up/down")
                // Moving horizontally, split the beam!
                wg.Add(1)
                go func() {defer wg.Done(); lightbeam(wg, g, x,y-1, 0,-1)}()
                wg.Add(1)
                go func() {defer wg.Done(); lightbeam(wg, g, x,y+1, 0,1)}()
                return
            }
        case "/":
            // moving down => reflect left so 0,1 -> -1,0
            // moving right => reflect up so 1,0 -> 0,-1
            log.Printf("reflect /")
            dx,dy = -dy,-dx
        case `\`:
            // moving down => reflect right so 0,1 -> 1,0
            // moving left => reflect up so -1,0 -> 0,-1
            // mpve up => reflect left so 0,-1 -> -1,0
            log.Printf(`reflect \`)
            dx,dy = dy,dx
        }
        x+=dx
        y+=dy
    }
}

func processB(puzzle []string) (int, error) {
    g := lbgrid.New(puzzle)
    m := 0
    var wg sync.WaitGroup
    for x := 0; x < g.Width(); x++ {
        wg.Add(1)
        go func() {defer wg.Done(); lightbeam(&wg, g, x,0, 0,1)}()
        wg.Wait()
        m = aoc.MaxInt(m,g.Energized())
        g = lbgrid.New(puzzle)

        wg.Add(1)
        go func() {defer wg.Done(); lightbeam(&wg, g, g.Width()-x-1,0, 0,-1)}()
        wg.Wait()
        m = aoc.MaxInt(m,g.Energized())
        g = lbgrid.New(puzzle)
    }
    for y := 0; y < g.Height(); y++ {
        wg.Add(1)
        go func() {defer wg.Done(); lightbeam(&wg, g, 0,y, 1,0)}()
        wg.Wait()
        m = aoc.MaxInt(m,g.Energized())
        g = lbgrid.New(puzzle)

        wg.Add(1)
        go func() {defer wg.Done(); lightbeam(&wg, g, 0,g.Height()-y-1, -1,0)}()
        wg.Wait()
        m = aoc.MaxInt(m, g.Energized())
        g = lbgrid.New(puzzle)
    }
    return m, nil
}
