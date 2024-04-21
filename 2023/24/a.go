package day24

import (
    _ "embed"
    "fmt"
    "log"
    "regexp"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/urfave/cli"
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day24A",
        Aliases: []string{"day24a"},
        Usage: "Day 24 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day24B",
        Aliases: []string{"day24b"},
        Usage: "Day 24 part B",
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

type hailstone struct {
    px,py,pz int
    vx,vy,vz int
}

func processA(puzzle []string) (int, error) {
    var hs []hailstone
    for _, line := range puzzle {
        hs = append(hs, decode(line))
    }
    return countIntersections(hs), nil
}

func decode(line string) hailstone {
    re := regexp.MustCompile(`([\-0-9]*),\s*([ \-0-9]*),\s*([\-0-9]*)\s*@\s*([\-0-9]*),\s*([\-0-9]*),\s*([\-0-9]*)`)
    matches := re.FindStringSubmatch(line)
    return hailstone{
        px: aoc.MustAtoi(matches[1]),
        py: aoc.MustAtoi(matches[2]),
        pz: aoc.MustAtoi(matches[3]),
        vx: aoc.MustAtoi(matches[4]),
        vy: aoc.MustAtoi(matches[5]),
        vz: aoc.MustAtoi(matches[6]),
    }
}

func countIntersections(hs []hailstone) int {
    intersectionCount := 0
    for i, h := range hs {
        for j:=i+1; j<len(hs); j++ {
            if intersects(h, hs[j]) {
                intersectionCount++
            }
        }
    }
    return intersectionCount
}

func intersects(h1, h2 hailstone) bool {
    sx1,sy1,_,ex1,ey1,_ := extend(h1)
    sx2,sy2,_,ex2,ey2,_ := extend(h2)
    ir := checkLineIntersection(sx1,sy1,ex1,ey1,sx2,sy2,ex2,ey2)
    log.Printf("%v / %v -> %v", h1, h2, ir)
    return ir.x >= 7 && ir.x <= 27 && ir.y >= 7 && ir.y <= 27
}

func extend(h hailstone) (int,int,int,int,int,int) {
    nanos := 100
    return h.px - h.vx * nanos,
        h.py - h.vy * nanos,
        h.pz - h.vz * nanos,
        h.px + h.vx * nanos,
        h.py + h.vy * nanos,
        h.pz + h.vz * nanos
}

func processB(puzzle []string) (int, error) {
    return 0, fmt.Errorf("Not yet implemented")
}

type intersectionResult struct {
    x, y int
    onLine1, onLine2 bool
}

func (ir intersectionResult) String() string {
    return fmt.Sprintf("%v,%v (%v,%v)", ir.x, ir.y, ir.onLine1, ir.onLine2)
}

// http://jsfiddle.net/justin_c_rounds/Gd2S2/light/
// from https://stackoverflow.com/a/60368757/752411
// via reddit...
func checkLineIntersection(line1StartX, line1StartY, line1EndX, line1EndY, line2StartX, line2StartY, line2EndX, line2EndY int) intersectionResult {
    // If the lines intersect, the result contains the x and y of the
    // intersection (treating the lines as infinite) and booleans for
    // whether line segment 1 or line segment 2 contain the point.
    result := intersectionResult{
        x: 0,
        y: 0,
        onLine1: false,
        onLine2: false,
    }
    denominator := ((line2EndY - line2StartY) * (line1EndX - line1StartX)) - ((line2EndX - line2StartX) * (line1EndY - line1StartY));
    if (denominator == 0) {
        return result
    }
    a := line1StartY - line2StartY
    b := line1StartX - line2StartX
    numerator1 := ((line2EndX - line2StartX) * a) - ((line2EndY - line2StartY) * b)
    numerator2 := ((line1EndX - line1StartX) * a) - ((line1EndY - line1StartY) * b)
    a = numerator1 / denominator
    b = numerator2 / denominator

    // if we cast these lines infinitely in both directions, they intersect here:
    result.x = line1StartX + (a * (line1EndX - line1StartX))
    result.y = line1StartY + (a * (line1EndY - line1StartY))

    // if line1 is a segment and line2 is infinite, they intersect if:
    if (a > 0 && a < 1) {
        result.onLine1 = true
    }
    // if line2 is a segment and line1 is infinite, they intersect if:
    if (b > 0 && b < 1) {
        result.onLine2 = true
    }
    // if line1 and line2 are segments, they intersect if both of the above are true
    return result
}
