package day14

import (
    _ "embed"
    "fmt"
    "log"
    "math"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/baldrick/aoc/2023/grid"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 14
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day14A",
        Usage: "Day 14 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day14B",
        Usage: "Day 14 part B",
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
    g := grid.New(puzzle)
    log.Printf("After A:\n%v", g)
    return calculateLoad(g), nil
}

func calculateLoad(g *grid.Grid) int {
    load := 0
    for x := 0; x < g.Width(); x++ {
        space := 0
        for y := 0; y < g.Height(); y++ {
            switch g.Get(x,y) {
            case "O":
                load += g.Height() - y + space
            case ".":
                space++
            case "#":
                space = 0
            }
        }
    }
    return load
}

// Cycle is unlikely to get to a state where rocks don't move after a N/E/S/W
// cycle.  But they are likely to get into an N-long loop.
// So find the loop then figure out where we'll be in that loop after 1bn
// cycles and calculate load on that grid.
func processB(puzzle []string) (int, error) {
    type gridAndIndex struct {
        g *grid.Grid
        n int
    }
    g := grid.New(puzzle)
    log.Printf("B loaded (%v):\n%v", g.Key(), g)
    gridMap := make(map[grid.GridKey]gridAndIndex)
    cycledGrid := g.Clone()
    log.Printf("cycledGrid load = %v", calculateLoad(cycledGrid))
    // for i:=0; i<100; i++ {
    //     cycle(cycledGrid, 0, -1)
    //     log.Printf("%v: cycle N, load = %v, key=%v", i, calculateLoad(cycledGrid), cycledGrid.Key())
    //     cycle(cycledGrid, -1, 0)
    //     log.Printf("%v: cycle W, load = %v, key=%v", i, calculateLoad(cycledGrid), cycledGrid.Key())
    //     cycle(cycledGrid, 0, 1)
    //     log.Printf("%v: cycle S, load = %v, key=%v", i, calculateLoad(cycledGrid), cycledGrid.Key())
    //     cycle(cycledGrid, 1, 0)
    //     log.Printf("%v: cycle E, load = %v, key=%v", i, calculateLoad(cycledGrid), cycledGrid.Key())
    //     cg, ok := gridMap[cycledGrid.Key()]
    //     if ok {
    //         log.Printf("cycle found %v-%v", cg.n, i)
    //         break
    //     }
    //     gridMap[cycledGrid.Key()] = gridAndIndex{cycledGrid.Clone(), i}
    // }


    for cycles := 0; cycles < 1e5; cycles++ {
        for _, s := range []struct{
            dx int
            dy int
        }{
            {0,-1},
            {-1,0},
            {0, 1},
            {1, 0},
        } {
            cycle(cycledGrid, s.dx, s.dy)
            log.Printf("B cycle %v after %v,%v:\n%v", cycles, s.dx, s.dy, cycledGrid)
        }
        loopStart, ok := gridMap[cycledGrid.Key()]
        if ok {
            // We've found a loop between cycles and loopStart.
            cyclesToGo := 1e9 - cycles
            loopLength := cycles - loopStart.n
            remainder := math.Mod(float64(cyclesToGo), float64(loopLength))
            destGridIndex := loopStart.n + int(remainder)
            log.Printf("Found cycle from %v to %v", loopStart.n, cycles)
            log.Printf("cyclesToGo:%v, loopLength:%v, remainder:%v, destGridIndex:%v", cyclesToGo, loopLength, remainder, destGridIndex)
            for _, v := range gridMap {
                log.Printf("%v load=%v:\n%v", v.n, calculateLoad(v.g), v.g)
            }
            for _, v := range gridMap {
                if v.n == destGridIndex {
                    return calculateLoad(v.g), nil 
                }
            }
        }
        gridMap[cycledGrid.Key()] = gridAndIndex{g:cycledGrid.Clone(), n:cycles+1}
    }
    return 0, fmt.Errorf("failed to find cycle")
}

func cycle(g *grid.Grid, dx, dy int) {
    if dx == 0 {
        for x:=0; x<g.Width(); x++ {
            cycleColumn(g, x, dy)
        }
    }
    if dy == 0 {
        for y:=0; y<g.Height(); y++ {
            cycleRow(g, y, dx)
        }
    }
    //log.Printf("post cycle %v,%v:\n%v", dx, dy, g)
}

func cycleColumn(g *grid.Grid, x, dy int) {
    starty := 0
    endy := g.Height()-1
    if dy > 0 {
        starty = g.Height()-1
        endy = 0
    }

    for y:=starty; y!=endy; y-=dy {
        if g.Get(x,y) == "." {
            // Find first rolling rock against the direction we're iterating.
            if rx,ry:=rockNearby(g, x,y, 0,-dy); rx!=-1 && ry!=-1 {
                //log.Printf("Rock found at %v,%v", rx,ry)
                g.Set(x,y,"O")
                g.Set(rx,ry,".")
            }
        }
    }
}

func cycleRow(g *grid.Grid, y, dx int) {
    startx := 0
    endx := g.Width()-1
    if dx > 0 {
        startx = g.Width()-1
        endx = 0
    }

    for x:=startx; x!=endx; x-=dx {
        if g.Get(x,y) == "." {
            // Find first rolling rock against the direction we're iterating.
            if rx,ry:=rockNearby(g, x,y, -dx,0); rx!=-1 && ry!=-1 {
                //log.Printf("Rock found at %v,%v", rx,ry)
                g.Set(x,y,"O")
                g.Set(rx,ry,".")
            }
        }
    }
}

func rockNearby(g *grid.Grid, x,y, dx,dy int) (int, int) {
    for ;; {
        x+=dx
        y+=dy
        //log.Printf("Checking %v,%v", x, y)
        if g.Outside(x,y) {
            return -1,-1
        }
        switch g.Get(x,y) {
        case "O":
            //log.Printf("Found rolling rock at %v,%v", x, y)
            return x,y
        case "#":
            //log.Printf("Found static rock at %v,%v", x, y)
            return -1,-1
        }
        //log.Printf("%v found at %v,%v", g.Get(x,y), x, y)
    }
}
