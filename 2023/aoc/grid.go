package grid

import (
	"log"
	"strings"
)

type Grid struct {
	values [][]string
}

type GridKey string // so grids can be stored in maps.

func (g *Grid) Key() GridKey {
	s := ""
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			s += g.Get(x,y)
		}
	}
	return GridKey(s)
}

func New(input []string) *Grid {
	values := make([][]string, len(input))
	for y, line := range input {
		values[y] = make([]string, len(line))
		for x, char := range line {
			values[y][x] = string(char)
		}
	}
	return &Grid{values}
}

func (g *Grid) Clone() *Grid {
	values := make([][]string, g.Height())
	for y := 0;  y < g.Height();  y++ {
		values[y] = make([]string, g.Width())
		for x := 0;  x < g.Width();  x++ {
			values[y][x] = g.Get(x, y)
		}
	}
	return &Grid{values}
}

func (g *Grid) Width() int {
	return len(g.values[0])
}

func (g *Grid) Height() int {
	return len(g.values)
}

func (g *Grid) Outside(x,y int) bool {
	return x<0 || y<0 || x>=g.Width() || y>=g.Height()
}

func (g *Grid) Get(x,y int) string {
	return g.values[y][x]
}

func (g *Grid) Set(x,y int, s string) {
	g.values[y][x] = s
}

func (g *Grid) String() string {
	var s []string
	for _, line := range g.values {
		s = append(s, strings.Join(line, ""))
	}
	return strings.Join(s, "\n")
}

func (g *Grid) Dump() {
	log.Printf("%v", g.String())
}
