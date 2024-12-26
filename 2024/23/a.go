package day23

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day23A",
		Aliases: []string{"day23a"},
		Usage:   "Day 23 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day23B",
		Aliases: []string{"day23b"},
		Usage:   "Day 23 part B",
		Action:  partB,
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

func processA(puzzle []string) (int, error) {
	g := newGraph()

	for _, line := range puzzle {
		if len(line) == 0 {
			continue
		}
		s := strings.Split(line, "-")
		g.connect(s[0], s[1])
	}
	log.Printf("connections: %v", g)
	//t := g.triples()
	return 0, fmt.Errorf("Not yet implemented")
}

func processB(puzzle []string) (int, error) {
	return 0, fmt.Errorf("Not yet implemented")
}

type node struct {
	name        string
	connections map[string]*node
}

func newNode(name string) *node {
	return &node{name: name, connections: make(map[string]*node)}
}

func (n *node) String() string {
	return fmt.Sprintf("%v -> %v", n.name, n.connections)
}

func (n *node) connect(name string) {
	if _, ok := n.connections[name]; ok {
		log.Printf("%v and %v are already connected", n.name, name)
		return
	}
	n.connections[name] = newNode(n.name)
}

func (n *node) find(name string) (*node, bool) {
	node, ok := n.connections[name]
	if ok {
		return node, true
	}
	for _, connectedNode := range n.connections {
		cn, ok := connectedNode.find(name)
		if ok {
			return cn, true
		}
	}
	return nil, false
}

type graph struct {
	roots map[string]*node
}

func newGraph() *graph {
	return &graph{roots: make(map[string]*node)}
}

func (g *graph) String() string {
	s := ""
	for root, nodes := range g.roots {
		s += fmt.Sprintf("\n  %v: %v", root, nodes)
	}
	return s
}

func (g *graph) connect(a, b string) {
	g.connectDirect(a, b)
	g.connectDirect(b, a)
}

func (g *graph) connectDirect(from, to string) {
	n, ok := g.find(from)
	if ok {
		n.connect(to)
	} else {
		nn := newNode(from)
		nn.connect(to)
		g.roots[from] = nn
	}
}

func (g *graph) find(n string) (*node, bool) {
	node, ok := g.roots[n]
	if ok {
		return node, true
	}
	for _, node := range g.roots {
		node, ok = node.find(n)
		if ok {
			return node, true
		}
	}
	return nil, false
}

func (g *graph) triples() {
	// Find root
}

type connections struct {
	sets              []*aoc.StringSet
	directConnections map[string]*aoc.StringSet
}

func newConnections() *connections {
	return &connections{directConnections: make(map[string]*aoc.StringSet)}
}

func (c *connections) String() string {
	s := ""
	for _, conn := range c.sets {
		s += fmt.Sprintf("\n  %v", conn)
	}
	s += "\n direct:"
	for from, to := range c.directConnections {
		s += fmt.Sprintf("\n  %v -> %v", from, to)
	}
	return s
}

// a->b->c->a ... so direction connections of a
// whose direct connections loop back to a
// func (c *connections) directlyConnected3() []*aoc.StringSet {
// 	var directConnectionGroups []*aoc.StringSet
// 	for from, to := range c.directConnections {
// 		// in the example commented, "from" is "a" and "to" is "b"
// 		dc := aoc.NewStringSet()
// 		dc.Add(from)
// 		toa := to.AsArray()
// 		for l1 := 0; l1 < len(toa); l1++ {
// 			l1cs := c.directConnections[toa[l1]]
// 			// in the example, l1cs is now direct connections of "b", i.e. "c"
// 		}
// 		directConnectionGroups = append(directConnectionGroups, dc)
// 	}
// 	return directConnectionGroups
// }

func (c *connections) addDirect(a, b string) {
	c.addDirectOneWay(a, b)
	c.addDirectOneWay(b, a)
}

func (c *connections) addDirectOneWay(a, b string) {
	s, ok := c.directConnections[a]
	if !ok {
		ss := aoc.NewStringSet()
		ss.Add(b)
		c.directConnections[a] = ss
	} else {
		s.Add(b)
	}
}

func (c *connections) add(a, b string) {
	for _, conn := range c.sets {
		if conn.Contains(a) || conn.Contains(b) {
			log.Printf("adding %v,%v to %v", a, b, conn)
			conn.Add(a)
			conn.Add(b)
			return
		}
	}
	s := aoc.NewStringSet()
	s.Add(a)
	s.Add(b)
	log.Printf("creating new set %v", s)
	c.sets = append(c.sets, s)
}

func (c *connections) merge() {
	merged := false
	m := 0
	for {
		for n := 0; n < len(c.sets); n++ {
			if n == m || c.sets[n].Len() == 0 || c.sets[m].Len() == 0 {
				continue
			}
			if c.sets[m].ContainsAny(c.sets[n]) {
				log.Printf("merged #%v into #%v: %v -> %v", n, m, c.sets[n], c.sets[m])
				c.sets[m].AddAll(c.sets[n])
				c.sets[n].Clear()
				merged = true
			}
		}
		m++
		if m >= len(c.sets) {
			if !merged {
				return
			}

			log.Printf("looping, m=%v", m)
			m = 0
			merged = false
		}
	}
}
