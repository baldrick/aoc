package ggrid

import (
	"log"
	"strings"
)

type GridValueFactory[V GridValue] interface {
	New(s string) V
}

type GridValue interface {
	String() string
	Clone() GridValue
}

type Grid[V GridValue] struct {
	values [][]V
}

type GridKey string // so grids can be stored in maps.

func (g *Grid[V]) Key() GridKey {
	s := ""
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			s += g.Get(x,y).String()
		}
	}
	return GridKey(s)
}

func New[V GridValue](gf GridValueFactory[V], input []string) *Grid[V] {
	values := make([][]V, len(input))
	for y, line := range input {
		values[y] = make([]V, len(line))
		for x, char := range line {
			values[y][x] = gf.New(string(char))
		}
	}
	return &Grid[V]{values}
}

func (g *Grid[V]) Clone() *Grid[V] {
	values := make([][]V, g.Height())
	for y := 0;  y < g.Height();  y++ {
		values[y] = make([]V, g.Width())
		for x := 0;  x < g.Width();  x++ {
			values[y][x] = g.Get(x, y).Clone().(V)
		}
	}
	return &Grid[V]{values}
}

func (g *Grid[V]) Width() int {
	return len(g.values[0])
}

func (g *Grid[V]) Height() int {
	return len(g.values)
}

func (g *Grid[V]) Outside(x,y int) bool {
	return x<0 || y<0 || x>=g.Width() || y>=g.Height()
}

func (g *Grid[V]) Get(x,y int) V {
	return g.values[y][x]
}

func (g *Grid[V]) Set(x,y int, v V) {
	g.values[y][x] = v
}

func (g *Grid[V]) String() string {
	var s []string
	for _, line := range g.values {
		sline := ""
		for _, v := range line {
			sline += v.String() + ","
		}
		s = append(s, sline)
	}
	return strings.Join(s, "\n")
}

func (g *Grid[V]) Dump() {
	log.Printf("%v", g.String())
}
