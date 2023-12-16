package lbgrid

import (
	// "fmt"
	"log"
	// "sort"
	"strings"
)

type Direction struct {
    DX,DY int
}

func (d Direction) Hash() int {
    /*
    -1,0 = -2
    1,0 = 2
    0,1 = 1
    0,-1 = -1
    */
    return d.DX*2 + d.DY
}

type GridValue struct {
    Cell string
    DirectionsSeen map[int]Direction
}

func NewGridValue(s string) *GridValue {
    return &GridValue{Cell: s, DirectionsSeen: make(map[int]Direction)}
}

func (gv *GridValue) String() string {
	s := gv.Cell
	if len(gv.DirectionsSeen) > 0 {
		s = "#"
	}
	return s
    // var ds []Direction
    // for _,d := range gv.DirectionsSeen {
    //     ds = append(ds, d)
    // }
    // sort.Slice(ds, func(a,b int) bool {
    //     if ds[a].DX == ds[b].DX {
    //         return ds[a].DY < ds[b].DY
    //     }
    //     return ds[a].DX < ds[b].DX
    // })
    // for _, d := range ds {
    //     s += fmt.Sprintf(", (%v,%v)", d.DX, d.DY)
    // }
    // return s
}

type Grid struct {
	values [][]*GridValue
}

type GridKey string // so grids can be stored in maps.

func (g *Grid) Key() string {
	s := ""
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			s += g.Get(x,y).String()
		}
	}
	return s
}

func New(input []string) *Grid {
	values := make([][]*GridValue, len(input))
	for y, line := range input {
		values[y] = make([]*GridValue, len(line))
		for x, char := range line {
			values[y][x] = NewGridValue(string(char))
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

func (g *Grid) Wrap(x,y int) (int,int) {
	if x<0 {
		x = g.Width()-1
	}
	if y<0 {
		y = g.Height()-1
	}
	if x>=g.Width() {
		x = 0
	}
	if y>=g.Height() {
		y = 0
	}
	return x,y
}

func (g *Grid) Get(x,y int) *GridValue {
	return g.values[y][x]
}

func (g *Grid) Set(x,y int, v *GridValue) {
	g.values[y][x] = v
}

func (g *Grid) String() string {
	var s []string
	for _, line := range g.values {
		sline := ""
		for _, v := range line {
			sline += v.String()
		}
		s = append(s, sline)
	}
	return strings.Join(s, "\n")
}

func (g *Grid) Dump() {
	log.Printf("\n%v", g.String())
}

func (g *Grid) Energized() int {
	g.Dump()
	energized := 0
	for _, line := range g.values {
		for _, cell := range line {
			if len(cell.DirectionsSeen) > 0 {
				energized++
			}
		}
	}
	return energized
}
