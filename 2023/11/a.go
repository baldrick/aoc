package day11

import (
    _ "embed"
    "fmt"
    "log"
    "strings"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 11
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day11A",
        Usage: "Day 11 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day11B",
        Usage: "Day 11 part B",
        Action: partB,
    }
)

func partA(ctx *cli.Context) error {
    answer, err := processA(aoc.PreparePuzzle(puzzle), 1)
    if err != nil {
        return err
    }
    log.Printf("Answer A: %v", answer)
    return nil
}

func partB(ctx *cli.Context) error {
    answer, err := processB(aoc.PreparePuzzle(puzzle), 1_000_000)
    if err != nil {
        return err
    }
    log.Printf("Answer B: %v", answer)
    return nil
}

type galaxy struct {
    g, x, y, dx, dy int
}

func (g *galaxy) String() string {
    return fmt.Sprintf("#%v %v+%v, %v+%v", g.g, g.x, g.dx, g.y, g.dy)
}

func (g galaxy) distance(end galaxy) int {
    d := aoc.AbsInt((g.x+g.dx) - (end.x+end.dx)) + aoc.AbsInt((g.y+g.dy) - (end.y+end.dy))
    //log.Printf("#%v %v,%v to #%v %v,%v = %v", g.g, g.x, g.y, end.g, end.x, end.y, d)
    return d
}

func processA(puzzle []string, _ int) (int, error) {
    grid, err := decode(puzzle)
    //dump("grid", grid)
    if err != nil {
        return 0, err
    }
    universe := expandColumns(grid)
    //dump("Universe", universe)
    galaxies := findGalaxies(universe)
    //log.Printf("galaxies: %v", galaxies)
    dist := 0
    for start, startPos := range galaxies {
        for end, endPos := range galaxies {
            if end <= start {
                continue
            }
            dist += startPos.distance(endPos)
        }
    }
    return dist, nil
}

func decode(puzzle []string) ([][]int, error) {
    galaxy := 1
    var grid [][]int
    var lineWithNoGalaxy []int
    for x:=0; x<len(puzzle[0]); x++ {
        lineWithNoGalaxy = append(lineWithNoGalaxy, 0)
    }
    for y:=0; y<len(puzzle); y++ {
        if !strings.Contains(puzzle[y], "#") {
            grid = append(grid, lineWithNoGalaxy)
            grid = append(grid, lineWithNoGalaxy)
            continue
        }
        var line []int
        for x:=0; x<len(puzzle[y]); x++ {
            switch puzzle[y][x] {
            case '.':
                line = append(line, 0)
            case '#':
                line = append(line, galaxy)
                galaxy++
            default:
                return nil, fmt.Errorf("failed to decode %q at %v,%v", puzzle[y][x], x, y)
            }
        }
        grid = append(grid, line)
    }
    return grid, nil
}

func expandColumns(grid [][]int) [][]int {
    extraColumns := aoc.NewIntSet()
    for x:=0; x<len(grid[0]); x++ {
        foundGalaxy := false
        for y:=0; y<len(grid) && !foundGalaxy; y++ {
            if grid[y][x] != 0 {
                foundGalaxy = true
            }
        }
        if !foundGalaxy {
            extraColumns.Add(x)
        }
    }
    var eGrid [][]int
    for y:=0; y<len(grid); y++ {
        line := grid[y]
        var expandedLine []int
        for x:=0; x<len(grid[y]); x++ {
            if extraColumns.Contains(x) {
                expandedLine = append(expandedLine, 0)
            }
            expandedLine = append(expandedLine, line[x])
        }
        eGrid = append(eGrid, expandedLine)
    }
    return eGrid
}

func findGalaxies(u [][]int) map[int]galaxy {
    galaxies := make(map[int]galaxy)
    for y:=0; y<len(u); y++ {
        for x:=0; x<len(u[y]); x++ {
            if u[y][x] != 0 {
                galaxies[u[y][x]] = galaxy{x:x, y:y, g:u[y][x]}
            }
        }
    }
    return galaxies
}

func dump(title string, grid [][]int) {
    log.Printf("** %v **", title)
    for y:=0; y<len(grid); y++ {
        log.Printf("%v", grid[y])
    }
}

// Heh, should've done it this way in the first place...
func processB(puzzle []string, expansionRate int) (int, error) {
    galaxies := getGalaxies(puzzle)
    expand(galaxies, expansionRate, puzzle)
    log.Printf("galaxies: %v", galaxies)
    dist := 0
    for start, startPos := range galaxies {
        for end, endPos := range galaxies {
            if end <= start {
                continue
            }
            d := startPos.distance(*endPos)
            log.Printf("%v->%v: %v", startPos, endPos, d)
            dist += d
        }
    }
    return dist, nil
}

func getGalaxies(puzzle []string) map[int]*galaxy {
    galaxies := make(map[int]*galaxy)
    g := 1
    for y:=0; y<len(puzzle); y++ {
        for x:=0; x<len(puzzle[y]); x++ {
            if puzzle[y][x] == '#' {
                galaxies[g] = &galaxy{g:g, x:x, y:y}
                g++
            }
        }
    }
    return galaxies
}

// expansionRate is how much the universe *replaces* existing empty areas
// with, not how much it *adds* (unlike the first part of the puzzle)...
func expand(galaxies map[int]*galaxy, expansionRate int, puzzle []string) {
    extraRows := aoc.NewIntSet()
    for y:=0; y<len(puzzle); y++ {
        if !strings.Contains(puzzle[y], "#") {
            extraRows.Add(y)
        }
    }
    extraRows.MapOver(func(y int) {
        for _, galaxy := range galaxies {
            if galaxy.y > y {
                galaxy.dy += expansionRate-1
            }
        }
    })

    extraCols := aoc.NewIntSet()
    for x:=0; x<len(puzzle[0]); x++ {
        found := false
        for y:=0; y<len(puzzle); y++ {
            if puzzle[y][x] == '#' {
                found = true
                break
            }
        }
        if !found {
            extraCols.Add(x)
        }
    }
    extraCols.MapOver(func(x int) {
        for _, galaxy := range galaxies {
            if galaxy.x > x {
                galaxy.dx += expansionRate-1
            }
        }
    })
}
