package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)
/*
tile N - top, left, right, bottom
tile M - t, l, r, b

for N to sit with M, n.top = m.bottom
flip thru' horiz axis => bits reverse for left & right
flip thru' vert axis => bits reversed for top & bottom
*/

type TileEdge struct {
    normal int32
    rotated int32
    matchingEdges int
}

func (te TileEdge) String() string {
    return fmt.Sprintf("%v/%v, %v", te.normal, te.rotated, te.matchingEdges)
}

type Tile struct {
    n int
    t *TileEdge
    b *TileEdge
    l *TileEdge
    r *TileEdge
}

func (t Tile) isCorner() bool {
    return (t.t.matchingEdges == 0 && t.l.matchingEdges == 0) ||
        (t.t.matchingEdges == 0 && t.r.matchingEdges == 0) ||
        (t.b.matchingEdges == 0 && t.l.matchingEdges == 0) ||
        (t.b.matchingEdges == 0 && t.r.matchingEdges == 0)
}

func (t Tile) String() string {
    var corner string
    if t.isCorner() {
        corner = " (CORNER?)"
    }
    return fmt.Sprintf("%v%v: %v, %v, %v, %v", t.n, corner, t.t, t.b, t.l, t.r)
}

func readTile(lines []string) (*Tile , error) {
    n, err := strconv.Atoi(lines[0][5:len(lines[0])-1])
    if err != nil {
        return nil, err
    }
    return &Tile{
        n: n,
        t: convertRow(lines[1]),
        b: convertRow(lines[10]),
        l: convertColumn(lines[1:11], 0),
        r: convertColumn(lines[1:11], 9),
    }, nil
}

func convertRow(line string) *TileEdge {
    var n int32
    var r int32
    var i uint
    for i = 0;  i < 10;  i++ {
        if line[i] == '#' {
            n += 1 << i
            r += 1 << (9 - i)
        }
    }
    return &TileEdge{ normal: n, rotated: r }
}

func convertColumn(lines []string, column int) *TileEdge {
    //log.Printf("converting column %v: %v", column, lines)
    var n int32
    var r int32
    var i uint
    for i = 0;  i < 10;  i++ {
        //log.Printf("checking line %v column %v: %v", i, column, lines[i][column])
        if lines[i][column] == '#' {
            n += 1 << i
            r += 1 << (9 - i)
        }
    }
    return &TileEdge{ normal: n, rotated: r}
}

func countMatchingEdgesForTile(tile1 *Tile, tile2 *Tile) {
    countMatchingEdgesForEdge(tile1.t, tile2)
    countMatchingEdgesForEdge(tile1.b, tile2)
    countMatchingEdgesForEdge(tile1.l, tile2)
    countMatchingEdgesForEdge(tile1.r, tile2)
}

// Any edge may match any other edge using combination of flip / rotate
func countMatchingEdgesForEdge(e *TileEdge, tile *Tile) {
    countMatchingEdgesForEdge2(e, tile.t)
    countMatchingEdgesForEdge2(e, tile.b)
    countMatchingEdgesForEdge2(e, tile.l)
    countMatchingEdgesForEdge2(e, tile.r)
}

func countMatchingEdgesForEdge2(e *TileEdge, f *TileEdge) {
    if e.normal == f.normal || e.normal == f.rotated {
        e.matchingEdges++
        f.matchingEdges++
    }
}

func countMatchingEdges(tileArray []*Tile) {
    for i, t := range tileArray {
        for j := i + 1;  j < len(tileArray);  j++ {
            countMatchingEdgesForTile(t, tileArray[j])
        }
    }
}

func dumpTile(tile *Tile) {
    log.Printf("%v", *tile)
}

func doit(filename string) {
    var input = readFile(filename)
    var tiles map[int]*Tile = make(map[int]*Tile)
    var tileArray []*Tile
    for i := 0;  i < len(input);  i += 12 {
        tile, err := readTile(input[i:i+11])
        if err != nil {
            log.Fatal("Cannot read tile from %v", input[i:i+12])
        }
        tiles[tile.n] = tile
        tileArray = append(tileArray, tile)
    }
    countMatchingEdges(tileArray)

    var total int = 1
    for _, tile := range tileArray {
        if tile.isCorner() {
            dumpTile(tile)
            total *= tile.n
        }
    }
    log.Printf("Total = %v", total)
}

func main() {
    if testMode() {
        doit("test.txt")
    } else {
        doit("input.txt")
    }
}

func testMode() bool {
    if len(os.Args) > 1 {
        if os.Args[1] == "-t" {
            return true
        }
        if os.Args[1] == "-T" {
            runTests()
            return true
        }
    }
    return false
}

func readFile(f string) []string {
    file, err := os.Open(f)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var input []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return input
}

func runTests() {
    assertEquals("convertRow #.........", convertRow("#.........").normal, 1)
    assertEquals("convertRow #........#", convertRow("#........#").normal, 513)
    var lines []string
    row := "#........."
    for i := 0;  i < 10;  i++ {
        lines = append(lines, row)
    }
    assertEquals(fmt.Sprintf("convertColumn %v", lines), convertColumn(lines, 0).normal, 1023)
    assertEquals(fmt.Sprintf("convertColumn %v", lines), convertColumn(lines, 9).normal, 0)
}

func assertEquals(t string, got int32, expected int32) {
    if got != expected {
        log.Fatalf("%s test failed, expected %v got %v", t, expected, got)
    }
}
