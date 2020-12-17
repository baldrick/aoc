package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func dump(cube *Cube) {
    var s string
    for z := range cube {
        s = fmt.Sprintf("%sz=%d\n", s, z)
        for y := range cube[z] {
            for x := range cube[z][y] {
                if cube[z][y][x] {
                    s = fmt.Sprintf("%s#", s)
                } else {
                    s = fmt.Sprintf("%s.", s)
                }
            }
            s = fmt.Sprintf("%s\n", s)
        }
    }
    log.Println(s)
}

func set(cube *Cube, x int, y int, z int, active bool) {
    cube[z][y][x] = active
}

func addToCube(cube *Cube, line string, x int, y int, z int) {
    for i := range line {
        switch (line[i:i+1]) {
        case "#":
            set(cube,x,y,z,true)
        case ".":
            set(cube,x,y,z,false)
        default:
            log.Fatalf("Invalid character %s in %s", line[i:i+1], line)
        }
        x++
    }
}

type Cube [25][25][25]bool

func countActive(cube *Cube) int {
    var active int
    for z := range cube {
        for y := range cube[z] {
            for x := range cube[z][y] {
                if cube[z][y][x] {
                    active++
                }
            }
        }
    }
    return active
}

func countNeighbours(cube *Cube, x int, y int, z int) int {
    var seen [3][3][3]bool
    var n int
    for dx := -1;  dx <= 1;  dx++ {
        for dy := -1;  dy <= 1;  dy++ {
            for dz := -1;  dz <= 1;  dz++ {
                if outsideCube(cube, x, y, z, dx, dy, dz) || (dx == 0 && dy == 0 && dz == 0) {
                    continue
                }
                if !seen[dz+1][dy+1][dx+1] {
                    seen[dz+1][dy+1][dx+1] = true
                    if cube[z+dz][y+dy][x+dx] {
                        n++
                    }
                }
            }
        }
    }
    return n
}

func outsideCube(cube *Cube, x int, y int, z int, dx int, dy int, dz int) bool {
    var size int = len(cube)
    return x+dx < 0 || y+dy < 0 || z+dz < 0 || x+dx >= size || y+dy >= size || z+dz >= size
}

func cycle(cube *Cube) *Cube {
    var newCube Cube
    var s string
    for z := range cube {
        s = fmt.Sprintf("%sz=%d\n", s, z)
        for y := range cube[z] {
            for x := range cube[z][y] {
                n := countNeighbours(cube, x, y, z)
                newCube[z][y][x] = cube[z][y][x]
                if cube[z][y][x] && n != 2 && n != 3 {
                    newCube[z][y][x] = false
                }
                if !cube[z][y][x] && n == 3 {
                    newCube[z][y][x] = true
                }
            }
        }
    }
    return &newCube
}

func doit(filename string) {
    var theCube Cube
    var cube *Cube = &theCube
    var input = readFile(filename)
    var x = 11
    var y = 11
    var z = 11
    for _, line := range input {
        log.Printf("Adding %s", line)
        addToCube(cube, line, x, y, z)
        y++
    }
    dump(cube)
    for i := 0;  i < 6;  i++ {
        log.Printf("Cyle %d", i+1)
        cube = cycle(cube)
    }
    dump(cube)
    log.Printf("Cube contains %d active cells", countActive(cube))
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