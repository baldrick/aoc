package day18

import (
    _ "embed"
    "fmt"
    "log"
    "math"
    "regexp"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/baldrick/aoc/2023/grid"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 18
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day18A",
        Usage: "Day 18 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day18B",
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

// func processA(puzzle []string) (int, error) {
//     g := grid.Empty(800, 800)
//     x,y := g.Width()/2, g.Height()/2
//     for _, line := range puzzle {
//         log.Printf("adding path %q to %v,%v", line, x,y)
//         x,y = addPath(g, x,y, line)
//     }
//     g.Dump()
//     g.Fill(g.Width()/2+1, g.Height()/2+1, "X", "#")
//     g.Dump()
//     return g.CountIf("X") + g.CountIf("#"), nil
// }

func processA(puzzle []string) (int, error) {
    m := make(map[int]leftRight)
    x,y := 0,0
    for _, line := range puzzle {
        re := regexp.MustCompile(`([A-Z]) ([0-9]*) \(#([a-z0-9]*)\)`)
        matches := re.FindStringSubmatch(line)
        length := aoc.MustAtoi(matches[2])
        switch matches[1] {
        case "R": // right
            addMinMax(m, x, y, x+length, y)
            x+=length
        case "D": // down
            addMinMax(m, x, y, x, y+length)
            y+=length
        case "L": // left
            addMinMax(m, x, y, x-length, y)
            x-=length
        case "U": // up
            addMinMax(m, x, y, x, y-length)
            y-=length
        }
    }
    return dumpFill(m), nil
}

func addMinMax(m map[int]leftRight, x1,y1, x2,y2 int) {
    for y := aoc.MinInt(y1,y2);  y <= aoc.MaxInt(y1,y2);  y++ {
        left := aoc.MinInt(x1,x2)
        right := aoc.MaxInt(x1,x2)
        lr, ok := m[y]
        if ok {
            left = aoc.MinInt(left, lr.left)
            right = aoc.MaxInt(right, lr.right)
        }
        m[y] = leftRight{left:left, right:right}
    }
}

func dumpFill(m map[int]leftRight) int {
    total := 0
    sy := math.MaxInt
    ey := math.MinInt
    for y, _ := range m {
        sy = aoc.MinInt(sy, y)
        ey = aoc.MaxInt(ey, y)
    }
    log.Printf("dumping %v lines: %v-%v", ey-sy, sy, ey)
    for y:=sy; y<=ey; y++ {
        lr, ok := m[y]
        if !ok {
            panic(fmt.Sprintf("could not find entry for y=%v", y))
        }
        f := aoc.AbsInt(lr.right - lr.left)+1
        //log.Printf("%v: %v to %v = %v", y, lr.left, lr.right, f)
        total += f
    }
    return total
}

type leftRight struct {
    left, right int
}

func processB(puzzle []string) (int, error) {
    m := make(map[int]leftRight)
    x,y := 0,0
    for _, line := range puzzle {
        re := regexp.MustCompile(`([A-Z]) ([0-9]*) \(#([a-z0-9]*)\)`)
        matches := re.FindStringSubmatch(line)
        hex := matches[3]
        length := aoc.MustXtoi(hex[:5])
        switch aoc.MustAtoi(string(hex[5])) {
        case 0: // right
            addMinMax(m, x, y, x+length, y)
            x+=length
        case 1: // down
            addMinMax(m, x, y, x, y+length)
            y+=length
        case 2: // left
            addMinMax(m, x, y, x-length, y)
            x-=length
        case 3: // up
            addMinMax(m, x, y, x, y-length)
            y-=length
        }
    }
    return dumpFill(m), nil
}

type direction struct {
    dx,dy int
}

var (
    directionMap = map[string]direction{
        "L": direction{-1,0},
        "R": direction{1,0},
        "U": direction{0,-1},
        "D": direction{0,1},
    }
    directionNMap = map[int]direction{
        2: direction{-1,0},
        0: direction{1,0},
        3: direction{0,-1},
        1: direction{0,1},
    }
)

func addPath(g *grid.Grid, x,y int, line string) (int, int) {
    re := regexp.MustCompile(`([A-Z]) ([0-9]*) \(#([a-z0-9]*)\)`)
    matches := re.FindStringSubmatch(line)
    d, ok := directionMap[matches[1]]
    if !ok {
        panic(fmt.Sprintf("cannot find direction %q", matches[1]))
    }
    length := aoc.MustAtoi(matches[2])
    for n := 0;  n < length;  n++ {
        log.Printf("setting %v,%v", x+n*d.dx, y+n*d.dy)
        g.Set(x+n*d.dx, y+n*d.dy, "#")
    }
    return x+length*d.dx, y+length*d.dy
}
