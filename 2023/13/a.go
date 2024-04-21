package day13

import (
    _ "embed"
//    "fmt"
    "log"
    "math"
//    "strings"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 13
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day13A",
        Usage: "Day 13 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day13B",
        Usage: "Day 13 part B",
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

func processA(p []string) (int, error) {
    var subPuzzle []string
    p = append(p, "")
    total := 0
    for _, line := range p {
        if len(line) > 0 {
            subPuzzle = append(subPuzzle, line)
        } else {
            rows := findReflections(subPuzzle)
            log.Printf("Found horizontal reflection point at %v", rows)
            cols := findReflections(transpose(subPuzzle))
            log.Printf("Found vertical reflection point at %v", cols)
            total += cols + (100*rows)
            subPuzzle = nil
        }
    }
    return total, nil
}

func processB(p []string) (int, error) {
    var subPuzzle []string
    p = append(p, "")
    total := 0
    for _, line := range p {
        if len(line) > 0 {
            subPuzzle = append(subPuzzle, line)
        } else {
            rows := findReflections(subPuzzle)
            originalRows := rows
            max := len(subPuzzle) * len(subPuzzle[0])
            n := 0
            for ; n<max-1; n++ {
                flip(subPuzzle, n)
                //log.Printf("n=%v, max=%v", n, max)
                r := findReflectionsIgnore(subPuzzle, originalRows)
                if r != 0 && r != rows {
                    x,y := coords(subPuzzle,n)
                    log.Printf("Found unsmudged horizontal reflection point at %v (x,y=%v,%v) (different to %v)", r, x,y, rows)
                    rows = r
                    flip(subPuzzle, n)
                    break
                }
                flip(subPuzzle, n)
            }
            if n >= max-1 {
                flip(subPuzzle, max-1)
            }
            if rows == originalRows {
                log.Printf("Didn't find new horizontal unsmudged reflection point")
            }
            //log.Printf("Subpuzzle returned to original state:\n%v", strings.Join(subPuzzle, "\n"))
            total += (100*rows)
            subPuzzle = nil
        }
    }
    return total, nil
}

func findReflections(p []string) int {
    //log.Printf("Looking for reflections in:\n%v", strings.Join(p, "\n"))
    for y := 0;  y < len(p)-1;  y++ {
        if p[y] == p[y+1] {
            // Possible reflection point.
            if isReflection(p, y) {
                return y+1
            }
        }
    }
    return 0
}

func findReflectionsIgnore(p []string, ignore int) int {
    //log.Printf("Looking for reflections in:\n%v", strings.Join(p, "\n"))
    for y := 0;  y < len(p)-1;  y++ {
        if y == ignore {
            continue
        }
        if p[y] == p[y+1] {
            // Possible reflection point.
            if isReflection(p, y) {
                return y+1
            }
        }
    }
    return 0
}

func isReflection(p []string, y int) bool {
    i := 1
    for ;; {
        if i>y || y+i+1 >= len(p) {
            return true
        }
        if p[y-i] != p[y+i+1] {
            //log.Printf("no reflection: #%v,%v: %v != %v", y-i, y+i+1, p[y-i], p[y+i+1])
            return false
        }
        i++
    }
}

func transpose(p []string) []string {
    var pt []string
    for x := 0; x < len(p[0]); x++ {
        s := ""
        for y := 0; y < len(p); y++ {
            s += string(p[y][x])
        }
        pt = append(pt, s)
    }
    return pt
}

func flip(p []string, n int) {
    x, y := coords(p, n)
    replacement := "#"
    if p[y][x] == '#' {
        replacement = "."
    }
    p[y] = p[y][:x] + replacement + p[y][x+1:]
}

func coords(p []string, n int) (int, int) {
    x := math.Mod(float64(n), float64(len(p[0])))
    y := (float64(n) - x)/float64(len(p[0]))
    //log.Printf("coords for n=%v len(p[0])=%v, y=%v, x=%v", n, len(p[0]), y, x)
    return int(x), int(y)
}
