package grid

import (
	"log"
	//"strings"
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

func (g *Grid) CountIf(match string) int {
	count := 0
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if g.Get(x,y) == match {
				count++
			}
		}
	}
	return count
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

func Empty(x,y int) *Grid {
	values := make([][]string, y)
	for n := 0;  n < y;  n++ {
		values[n] = make([]string, x)
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
	if len(g.values[y][x]) == 0 {
		return "."
	}
	return g.values[y][x]
}

func (g *Grid) Set(x,y int, s string) {
	g.values[y][x] = s
}

func (g *Grid) String() string {
	s := ""
	for y:=0; y<g.Height(); y++ {
		for x:=0; x<g.Width(); x++ {
			s += g.Get(x,y)
		}
		s += "\n"
	}
	return s
	// var s []string
	// for _, line := range g.values {
	// 	s = append(s, strings.Join(line, ""))
	// }
	// return strings.Join(s, "\n")
}

func (g *Grid) Dump() {
	log.Printf("%v", g.String())
}

func (g *Grid) Fill(x,y int, fillWith, edge string) {
	if g.Outside(x,y) {
		return
	}
	if g.Get(x,y) == fillWith {
		return
	}
	if g.Get(x,y) == edge {
		return
	}
	log.Printf("Filling %v,%v to %v from %v", x,y, fillWith, g.Get(x,y))
	g.Set(x,y,fillWith)
	g.Fill(x+1,y,fillWith,edge)
	g.Fill(x-1,y,fillWith,edge)
	g.Fill(x,y-1,fillWith,edge)
	g.Fill(x,y+1,fillWith,edge)
}