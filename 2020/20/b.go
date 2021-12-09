package main

import (
    "bufio"
    "errors"
    "fmt"
    "log"
    "math"
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
    used bool
}

func (t Tile) reorient(o *Orientation) {
    if o == nil {
        return
    }
    for r := 0;  r < o.antiClockwiseRotations;  r++ {

    }
    if o.vFlipped {

    }
    if o.hFlipped {

    }
}

func (t Tile) zeroMatchedEdgeCount() int {
    var n int
    if t.t.matchingEdges == 0 {
        n++
    }
    if t.b.matchingEdges == 0 {
        n++
    }
    if t.l.matchingEdges == 0 {
        n++
    }
    if t.r.matchingEdges == 0 {
        n++
    }
    return n
}

func (t Tile) isCorner() bool {
    return (t.t.matchingEdges == 0 && t.l.matchingEdges == 0) ||
        (t.t.matchingEdges == 0 && t.r.matchingEdges == 0) ||
        (t.b.matchingEdges == 0 && t.l.matchingEdges == 0) ||
        (t.b.matchingEdges == 0 && t.r.matchingEdges == 0)
}

func (t Tile) isEdge() bool {
    return !t.isCorner() &&
        (t.t.matchingEdges == 0 || t.b.matchingEdges == 0 || t.l.matchingEdges == 0 || t.r.matchingEdges == 0)
}

func (t Tile) String() string {
    var corner string
    if t.isCorner() {
        corner = " (CORNER)"
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
    var tileMap map[int]*Tile = make(map[int]*Tile)
    var tileArray []*Tile
    for i := 0;  i < len(input);  i += 12 {
        tile, err := readTile(input[i:i+11])
        if err != nil {
            log.Fatal("Cannot read tile from %v", input[i:i+12])
        }
        tileMap[tile.n] = tile
        tileArray = append(tileArray, tile)
    }
    countMatchingEdges(tileArray)

    var total int = 1
    var tileCount int
    var corners []*Tile
    for _, tile := range tileArray {
        if tile.isCorner() {
            dumpTile(tile)
            total *= tile.n
            corners = append(corners, tile)
        }
        tileCount++
    }
    var imageSize float64 = math.Sqrt(float64(tileCount))
    log.Printf("Total = %v, tileCount = %v for %vx%v image", total, tileCount, imageSize, imageSize)

    assembleImage(tileArray, tileMap, corners, int(math.Round(imageSize)))
}

type ImageRow struct {
    row map[int]*Tile
}

type Image struct {
    image map[int]*ImageRow
}

func (i Image) set(tm map[int]*Tile, x int, y int, tile *Tile, o *Orientation) {
    r, present := i.image[y]
    if !present {
        r = &ImageRow{row: make(map[int]*Tile)}
        i.image[y] = r
    }
    r.row[x] = tile
    tile.reorient(o)
    tm[tile.n].used = true
}

func (i Image) isSet(x int, y int) bool {
    r, present := i.image[y]
    if !present {
        return false
    }
    _, present = r.row[x]
    return present
}

func (i Image) get(x int, y int) *Tile {
    r, present := i.image[y]
    if !present {
        return nil
    }
    t, present := r.row[x]
    if !present {
        return nil
    }
    return t
}

func findTopLeft(corners []*Tile) (*Tile, error) {
    for _, c := range corners {
        if c.t.matchingEdges == 0 && c.l.matchingEdges == 0 {
            return c, nil
        }
    }
    return nil, errors.New("Cannot find top left corner")
}

type EdgeMatch struct {
    t bool
    l bool
    zeros int
}

type Orientation struct {
    antiClockwiseRotations int
    hFlipped bool // filpping is done after rotation
    vFlipped bool // flipped through its vertical axis
}

// em.t true => match tile's top with prevTile's bottom
// em.l true => match tile's left with prevTile's right
func edgeMatchWithRotation(prevTile *Tile, tile *Tile, em *EdgeMatch) *Orientation {
    if em.zeros != tile.zeroMatchedEdgeCount() {
        return nil
    }

    var bottom int32 = -1
    if em.t {
        bottom = prevTile.b.normal
    }

    var right int32 = -1
    if em.l {
        right = prevTile.r.normal
    }

    if em.l && !em.t {
        if bottom == tile.t.normal {
            return &Orientation{antiClockwiseRotations: 0}
        }
        if bottom == tile.r.normal {
            return &Orientation{antiClockwiseRotations: 1}
        }
        if bottom == tile.b.rotated {
            return &Orientation{antiClockwiseRotations: 2}
        }
        if bottom == tile.l.rotated {
            return &Orientation{antiClockwiseRotations: 3}
        }

        if bottom == tile.b.normal {
            return &Orientation{antiClockwiseRotations: 0, vFlipped: true}
        }
        if bottom == tile.l.rotated {
            return &Orientation{antiClockwiseRotations: 1, vFlipped: true}
        }
        if bottom == tile.t.rotated {
            return &Orientation{antiClockwiseRotations: 2, vFlipped: true}
        }
        if bottom == tile.r.rotated {
            return &Orientation{antiClockwiseRotations: 3, vFlipped: true}
        }
    } else if !em.l && em.t {
        if right == tile.l.normal {
            return &Orientation{antiClockwiseRotations: 0}
        }
        if right == tile.t.rotated {
            return &Orientation{antiClockwiseRotations: 1}
        }
        if right == tile.r.rotated {
            return &Orientation{antiClockwiseRotations: 2}
        }
        if right == tile.b.normal {
            return &Orientation{antiClockwiseRotations: 3}
        }

        if right == tile.r.normal {
            return &Orientation{antiClockwiseRotations: 0, hFlipped: true}
        }
        if right == tile.b.rotated {
            return &Orientation{antiClockwiseRotations: 1, hFlipped: true}
        }
        if right == tile.l.rotated {
            return &Orientation{antiClockwiseRotations: 2, hFlipped: true}
        }
        if right == tile.t.rotated {
            return &Orientation{antiClockwiseRotations: 3, hFlipped: true}
        }
    } else /* em.l && em.t */ {

// TODO: flipping (!) flags incorrect here, possibly elsewhere, need h & v flips!  Is this really the right way?

        // simple case - not rotated or flipped
        if bottom == tile.t.normal && right == tile.l.normal {
            return &Orientation{antiClockwiseRotations: 0}
        }
        if bottom == tile.b.normal && right == tile.l.rotated {
            return &Orientation{antiClockwiseRotations: 0, vFlipped: true}
        }
        // rotate once anti-clockwise and/or flipped
        if bottom == tile.r.normal && right == tile.t.rotated {
            return &Orientation{antiClockwiseRotations: 1}
        }
        if bottom == tile.l.normal && right == tile.b.rotated {
            return &Orientation{antiClockwiseRotations: 1, vFlipped: true}
        }
        // rotated twice and/or flipped
        if bottom == tile.b.rotated && right == tile.r.rotated {
            return &Orientation{antiClockwiseRotations: 2}
        }
        if bottom == tile.t.rotated && right == tile.l.rotated {
            return &Orientation{antiClockwiseRotations: 2, hFlipped: true}
        }
        // rotated thrice anti-clockwise and/or flipped
        if bottom == tile.l.rotated && right == tile.b.normal {
            return &Orientation{antiClockwiseRotations: 3}
        }
        if bottom == tile.r.rotated && right == tile.t.normal {
            return &Orientation{antiClockwiseRotations: 3, vFlipped: true}
        }
    }
    return nil
}

func findCorner(corners []*Tile, prevTile *Tile, em *EdgeMatch) (*Tile, *Orientation, error) {
    for _, corner := range corners {
        if !corner.used {
            var o *Orientation = edgeMatchWithRotation(prevTile, corner, em)
            if o != nil {
                return corner, o, nil
            }
        }
    }
    return nil, nil, errors.New("Cannot find corner")
}

func findTile(ta []*Tile, prevTile *Tile, em *EdgeMatch) (*Tile, *Orientation) {
    for _, t := range ta {
        if !t.used {
            var o *Orientation = edgeMatchWithRotation(prevTile, t, em)
            if o != nil {
                return t, o
            }
        }
    }
    return nil, nil
}

func getPrevTile(image Image, x int, y int) *Tile {
    var gx int = x - 1
    var gy int = y - 1
    if gx < 0 {
        gx = 0
    }
    if gy < 0 {
        gy = 0
    }
    return image.get(gx, gy)
}

func assembleImage(ta []*Tile, tm map[int] *Tile, corners []*Tile, imageSize int) *Image {
    var image Image = Image{make(map[int]*ImageRow)}
    topLeftCorner, err := findTopLeft(corners)
    if err != nil {
        log.Fatal(err)
    }
    image.set(tm, 0, 0, topLeftCorner, nil)

    for y := 0;  y < imageSize;  y++ {
        for x := 0;  x < imageSize;  x++ {
            if !image.isSet(x, y) {
                var prevTile *Tile = getPrevTile(image, x, y)
                var t *Tile
                var o *Orientation
                if x == 0 {
                    if y == imageSize - 1 {
                        t, o, err = findCorner(corners, prevTile, &EdgeMatch{t:true, zeros: 2})
                        if err != nil {
                            log.Fatal(err)
                        }
                    } else {
                        t, o = findTile(ta, prevTile, &EdgeMatch{t:true, zeros: 1})
                    }
                } else if y == 0 {
                    if x == imageSize - 1 {
                        t, o, err = findCorner(corners, prevTile, &EdgeMatch{l:true, zeros: 2})
                        if err != nil {
                            log.Fatal(err)
                        }
                    } else {
                        t, o = findTile(ta, prevTile, &EdgeMatch{l:true, zeros: 1})
                    }
                } else if x == imageSize - 1 {
                    if y == imageSize - 1 {
                        t, o, err = findCorner(corners, prevTile, &EdgeMatch{t:true, l:true, zeros: 2})
                        if err != nil {
                            log.Fatal(err)
                        }
                    } else {
                        t, o = findTile(ta, prevTile, &EdgeMatch{t:true, l:true, zeros: 1})
                    }
                }
                if t == nil { // no special cases apply
                    t, o = findTile(ta, prevTile, &EdgeMatch{t:true, l:true, zeros: 0})
                }
                if t == nil {
                    log.Fatal("Failed to find tile!")
                }
                image.set(tm, x, y, t, o)
            }
        }
    }

    return &image
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
