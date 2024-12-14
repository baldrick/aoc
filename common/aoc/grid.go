package grid

import (
	"fmt"
	"log"
	"strings"

	"github.com/baldrick/aoc/common/aoc"
)

type Grid struct {
	values [][]string
	border int
}

func New(input []string) *Grid {
	values := make([][]string, len(input))
	for y, line := range input {
		values[y] = make([]string, len(line))
		for x, char := range line {
			values[y][x] = string(char)
		}
	}
	return &Grid{values: values, border: 0}
}

type GridKey string // so grids can be stored in maps.

func (g *Grid) Key() GridKey {
	s := ""
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			s += g.Get(x, y)
		}
	}
	return GridKey(s)
}

func (g *Grid) CountIf(match string) int {
	count := 0
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if g.Get(x, y) == match {
				count++
			}
		}
	}
	return count
}

func Empty(x, y int) *Grid {
	values := make([][]string, y)
	for n := 0; n < y; n++ {
		values[n] = make([]string, x)
	}
	return &Grid{values: values, border: 0}
}

func (g *Grid) Clone() *Grid {
	values := make([][]string, g.Height())
	for y := 0; y < g.Height(); y++ {
		values[y] = make([]string, g.Width())
		for x := 0; x < g.Width(); x++ {
			values[y][x] = g.Get(x, y)
		}
	}
	return &Grid{values: values, border: g.border}
}

func (g *Grid) Width() int {
	return len(g.values[0])
}

func (g *Grid) Height() int {
	return len(g.values)
}

func (g *Grid) Outside(x, y int) bool {
	return x < g.border || y < g.border || x >= g.Width()-g.border || y >= g.Height()-g.border
}

func (g *Grid) Get(x, y int) string {
	if len(g.values[y][x]) == 0 {
		return "."
	}
	return g.values[y][x]
}

func (g *Grid) Set(x, y int, s string) {
	g.values[y][x] = s
}

func (g *Grid) String() string {
	s := ""
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if g.Get(x, y) == "-1" {
				s += "."
			} else {
				s += g.Get(x, y)
			}
		}
		s += "\n"
	}
	return s
}

func (g *Grid) Dump() {
	g.DumpMsg("")
}

func (g *Grid) DumpMsg(s string) {
	g.DumpMsgLen(s, 1)
}

func (g *Grid) DumpMsgLen(msg string, l int) {
	s := ""
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if g.Get(x, y) == "-1" {
				s += "."
			} else {
				s += g.Get(x, y)
			}
			repeat := l - len(g.Get(x, y))
			if repeat > 0 {
				s += strings.Repeat(" ", repeat)
			}
		}
		s += "\n"
	}
	log.Printf("%v: %v x %v grid:\n%v", msg, g.Width(), g.Height(), s)
}

func (g *Grid) Fill(x, y int, fillWith, edge string) {
	if g.Outside(x, y) {
		return
	}
	if g.Get(x, y) == fillWith {
		return
	}
	if g.Get(x, y) == edge {
		return
	}
	log.Printf("Filling %v,%v to %v from %v", x, y, fillWith, g.Get(x, y))
	g.Set(x, y, fillWith)
	g.Fill(x+1, y, fillWith, edge)
	g.Fill(x-1, y, fillWith, edge)
	g.Fill(x, y-1, fillWith, edge)
	g.Fill(x, y+1, fillWith, edge)
}

func (g *Grid) FillN(x, y int, edge string, step, maxSteps int) {
	if maxSteps <= 0 {
		return
	}
	log.Printf("Filling from %v,%v (step %v, %v to go)", x, y, step, maxSteps)
	if g.Outside(x, y) {
		return
	}
	if g.Get(x, y) != "." && g.Get(x, y) != "S" {
		return
	}
	if g.Get(x, y) == edge {
		return
	}
	if step > 0 {
		g.Set(x, y, fmt.Sprintf("%v", step))
	}
	g.FillN(x+1, y, edge, step+1, maxSteps-1)
	g.FillN(x-1, y, edge, step+1, maxSteps-1)
	g.FillN(x, y-1, edge, step+1, maxSteps-1)
	g.FillN(x, y+1, edge, step+1, maxSteps-1)
}

func (g *Grid) Find(s string) (int, int) {
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if g.Get(x, y) == s {
				return x, y
			}
		}
	}
	return -1, -1
}

func (g *Grid) FindAll(s string) []aoc.PairInt {
	var pi []aoc.PairInt
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if g.Get(x, y) == s {
				pi = append(pi, aoc.NewPairInt(x, y))
			}
		}
	}
	return pi
}

func (g *Grid) Replace(f, r string) {
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if g.Get(x, y) == f {
				g.Set(x, y, r)
			}
		}
	}
}

func (g *Grid) AddBorder(b string) *Grid {
	g.border = 1
	ng := Empty(g.Width()+2, g.Height()+2)
	for x := 0; x < g.Width(); x++ {
		for y := 0; y < g.Width(); y++ {
			ng.Set(x+1, y+1, g.Get(x, y))
		}
	}
	for x := 0; x < ng.Width(); x++ {
		ng.Set(x, 0, b)
		ng.Set(x, ng.Height()-1, b)
	}
	for y := 0; y < ng.Height(); y++ {
		ng.Set(0, y, b)
		ng.Set(ng.Width()-1, y, b)
	}
	return ng
}

func (g *Grid) Increment(x, y int) {
	current := g.Get(x, y)
	n := 1
	if len(current) != 0 && current != "." {
		n = aoc.MustAtoi(current) + 1
	}
	g.Set(x, y, fmt.Sprintf("%v", n))
}
