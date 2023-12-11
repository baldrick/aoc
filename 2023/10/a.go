package day10

import (
    _ "embed"
    "fmt"
    "log"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 10
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day10A",
        Usage: "Day 10 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day10B",
        Usage: "Day 10 part B",
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

type pipeType int

const (
    ground pipeType = iota
    start
    tube
)

type location struct {
    x, y int
}

type where struct {
    x, y int
    h heading
}

type heading int

const (
    // north means we're moving north, i.e. from south to north
    nowhere heading = iota
    north
    south
    east
    west
    noExit
)

func (h heading) String() string {
    switch h {
    case north: return "N"
    case south: return "S"
    case east: return "E"
    case west: return "W"
    case noExit: return "x"
    case nowhere: return "-"
    }
    return "?"
}

type pipe struct {
    pt pipeType
    headings map[heading]heading
}

func newPipe(a, b heading) pipe {
    p := pipe{}
    p.headings = make(map[heading]heading)
    p.headings[a] = b
    p.headings[b] = a
    p.pt = tube
    return p
}

func (p pipe) String() string {
    if p.pt == start { return "S" }
    if p.pt == ground { return "."}
    return fmt.Sprintf("%v", p.headings)
}

func (p pipe) from(h heading) heading {
    exitHeading, ok := p.headings[h]
    if !ok {
        log.Printf("No exit from %v", h)
        return noExit
    }
    return exitHeading
}

func next(pipes [][]pipe, from where) where {
    ty := from.y
    tx := from.x
    var entryPoint heading
    switch from.h {
    case north:
        ty--
        entryPoint = south
    case south:
        ty++
        entryPoint = north
    case east:
        tx++
        entryPoint = west
    case west:
        tx--
        entryPoint = east
    }
    if ty < 0 || tx < 0 || ty >= len(pipes) || tx >= len(pipes[0]) {
        return where{h: noExit}
    }
    target := pipes[ty][tx]
    return where{x: tx, y: ty, h: target.from(entryPoint)}
}

func processA(puzzle []string) (int, error) {
    sx, sy, pipes := getPipes(puzzle)
    if sx == -1 || sy == -1 {
        return 0, fmt.Errorf("Failed to find start: %v,%v", sx, sy)
    }
    log.Printf("start %v,%v; pipes: %v", sx, sy, pipes)

    for _, w := range []where{{sx,sy,north},{sx,sy,south},{sx,sy,west},{sx,sy,east}} {
        loc := w
        steps := 0
        log.Printf("starting at %v, steps=%v", loc, steps)
        for ;; {
            n := next(pipes, loc)
            log.Printf("%v -> %v", loc, n)
            steps++
            if loc.h == noExit || loc.h == nowhere {
                log.Printf("No exit %v from %v,%v", loc.h, loc.x, loc.y)
                break
            }
            if n.x == sx && n.y == sy {
                log.Printf("Back to the start(from=%v, to=%v) !  Steps=%v", loc, n, steps)
                return int(steps/2), nil
            }
            loc = n
        }
    }
    return 0, fmt.Errorf("Failed to find loop from %v,%v", sx, sy)
}

func processB(puzzle []string) (int, error) {
    sx, sy, pipes := getPipes(puzzle)
    if sx == -1 || sy == -1 {
        return 0, fmt.Errorf("Failed to find start: %v,%v", sx, sy)
    }
    //log.Printf("start %v,%v; pipes: %v", sx, sy, pipes)

    var startHeading heading
    var path map[location]status
    for _, w := range []where{{sx,sy,north},{sx,sy,south},{sx,sy,west},{sx,sy,east}} {
        loc := w
        steps := 0
        path = make(map[location]status)
        //log.Printf("starting at %v, steps=%v", loc, steps)
        for ;; {
            path[location{x:loc.x, y:loc.y}] = pathPart
            n := next(pipes, loc)
            //log.Printf("%v -> %v", loc, n)
            steps++
            if loc.h == noExit || loc.h == nowhere {
                log.Printf("No exit %v from %v,%v", loc.h, loc.x, loc.y)
                break
            }
            if n.x == sx && n.y == sy {
                log.Printf("Back to the start(from=%v, to=%v) !  Steps=%v", loc, n, steps)
                startHeading = w.h
                break
            }
            loc = n
        }
        if startHeading != nowhere {
            break
        }
    }
    dump(pipes, path)
    count := countEnclosed(pipes, path)
    log.Printf("%v cells enclosed by the path", count)
    dump(pipes, path)
    return count, nil
}

func dump(pipes [][]pipe, highlight map[location]status) {
    for y := 0; y < len(pipes); y++ {
        s := ""
        for x := 0; x < len(pipes[y]); x++ {
            st, ok := highlight[location{x:x, y:y}]
            if ok {
                s += st.String()
            } else {
                s += pipes[y][x].String()
            }
        }
        log.Printf("%v", s)
    }
}

type status int

const (
    unknown status = iota
    pathPart
    inside
    outside
    maybeInside
)

func (s status) String() string {
    switch s {
    case pathPart: return "*"
    case inside: return "I"
    case outside: return "O"
    case maybeInside: return "i"
    }
    return "?"
}

func countEnclosed(pipes [][]pipe, path map[location]status) int {
    // Fill the outside, we'll count what's left as enclosed.
    floodFill(pipes, 0, 0, path)
    counts := make(map[status]int)
    for loc, status := range path {
        log.Printf("%v=%v", loc, status)
        n, ok := counts[status]
        if !ok {
            n = 0
        }
        counts[status] = n+1
    }
    for s, c := range counts {
        log.Printf("%v: %v", s, c)
    }
    log.Printf("%v * %v = %v, path = %v", len(pipes), len(pipes[0]), len(pipes) * len(pipes[0]), len(path))
    count := (len(pipes) * len(pipes[0])) - len(path)
    return count
}

/*try switching to flood fill using a queue*/

func floodFill(pipes [][]pipe, x, y int, path map[location]status) {
    if y<0 || y>=len(pipes) || x<0 || x>=len(pipes[0]) {
        log.Printf("%v,%y is outside 0-%v,0-%v", x, y, len(pipes[0]), len(pipes))
        return
    }
    log.Printf("floodFill from %v,%v (%v) (path %v)", x, y, pipes[y][x], path[location{x,y}])
    s, ok := path[location{x,y}]
    if ok && s == outside {
        log.Printf("outside - %v s:%v", path[location{x,y}], s)
        return
    }
    if s != pathPart {
        path[location{x,y}] = outside
    }
    if !connected(pipes, x+1, y, x+1, y+1) { floodFill(pipes, x+1, y, path) }
    log.Printf("x+1 done, now doing y+1 from %v,%v", x, y)
    //if !connected(pipes, x-1, y, x-1, y+1) { floodFill(pipes, x-1, y, path) }
    if !connected(pipes, x, y+1, x+1, y+1) { floodFill(pipes, x, y+1, path) }
    //if !connected(pipes, x, y-1, x+1, y-1) { floodFill(pipes, x, y-1, path) }
}

func connected(pipes [][]pipe, x, y, x2, y2 int) bool {
    log.Printf("connected? %v; %v,%v - %v,%v", pipes, x, y, x2, y2)
    if outsideGrid(x,y,pipes) || outsideGrid(x2,y2,pipes) {
        log.Printf("%v,%v - %v,%v not connected", x, y, x2, y2)
        return false
    }
    if y==5 && y2==4 {
        log.Printf("checking whether %v,%v and %v,%v are connected", x,y, x2,y2)
    }
    var ex2,ey2 int
    var in, out heading
    if x == x2 {
        // Moving north/south, check for east/west connection.
        ex2,ey2 = x2+1,y2
        in,out = east,west
    } else {
        // Moving east/west, check for north/south connection.
        ex2,ey2 = x2,y2+1
        in,out = north,south
    }
    if outsideGrid(ex2,ey2,pipes) {
        // One cell is outside the grid therefore not connected.
        log.Printf("%v,%v outside grid => not connected", ex2, ey2)
        return false
    }
    if len(pipes[y2][x2].headings) == 0 || len(pipes[ey2][ex2].headings) == 0 {
        // One of them is ground so not connected.
        log.Printf("%v-%v; ground => not connected", pipes[y2][x2], pipes[ey2][ex2])
        return false
    }
    for inA, outA := range pipes[y2][x2].headings {
        for inB, outB := range pipes[ey2][ex2].headings {
            log.Printf("moving %v,%v -> %v,%v, blocked if %v-%v, %v -> %v, in=%v, out=%v, outA=%v, inA=%v, outB=%v, inB=%v", x, y, x2, y2, in, out, pipes[y2][x2], pipes[ey2][ex2], in, out, outA, inA, outB, inB)
            if ((outA == in || inA == in) && (outB == out || inB == out)) ||
                ((outB == in || inB == in) && (outA == out || inA == out)) {
                log.Printf("returning true")
                return true
            }
            return false
        }
    }
    return false
}

func outsideGrid(x,y int, pipes [][]pipe) bool {
    return x<0 || y<0 || y>=len(pipes) || x>=len(pipes[0])
}

func getPipes(puzzle []string) (int,int,[][]pipe) {
    var pipes [][]pipe
    startx := -1
    starty := -1
    for y, line := range puzzle {
        var pipeline []pipe
        for x, p := range line {
            sx,sy,s := section(x, y, p)
            // Puzzle / test pipelines don't start at 0,0.
            if sx!=0 || sy!=0 {
                startx = sx
                starty = sy
            }
            pipeline = append(pipeline, s)
        }
        pipes = append(pipes, pipeline)
    }
    return startx, starty, pipes
}

func section(x,y int, pipeType rune) (int,int,pipe) {
    var section pipe
    var startx,starty int
    switch pipeType {
    case 'S':
        startx = x
        starty = y
        section = pipe{pt:start}
    case '-':
        section = newPipe(east, west)
    case '|':
        section = newPipe(north, south)
    case 'L':
        section = newPipe(north, east)
    case 'J':
        section = newPipe(north, west)
    case '7':
        section = newPipe(south, west)
    case 'F':
        section = newPipe(south, east)
    case '.':
        section = pipe{pt:ground}
    }
    return startx,starty,section
}