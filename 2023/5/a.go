package main

import (
	"flag"
	"log"
	"math"
	"time"

	"github.com/baldrick/aoc/2023/aoc"
	"github.com/baldrick/aoc/2023/rangemap"
	"github.com/dustin/go-humanize"
)

const (
	year = 2023
	day = 5
)

var (
	inputFile = flag.String("f", "test", "Puzzle file (partial name) to use")
	logger = log.Default()
)

func main() {
	flag.Parse()
	puzzle, err := aoc.GetPuzzleInput(*inputFile, year, day)
	if err != nil {
		logger.Fatalf("oops: %v\n", err)
		return
	}
	if err := process(puzzle); err != nil {
		logger.Fatalf("oops: %v\n", err)
	}
}

func process(puzzle []string) error {
	seeds, err := aoc.StripAndParse("seeds:", puzzle[0])
	if err != nil {
		return err
	}
	logger.Printf("seeds: %v", seeds)
	var maps []*rangemap.TheMap
	var i int
	var rm *rangemap.TheMap
	for i=1; i<len(puzzle); {
		i, rm = readRangeMap(puzzle, i)
		maps = append(maps, rm)
		logger.Println(rm)
	}
	start := time.Now()
	min := math.MaxInt
	for i=0; i<len(seeds); i+=2 {
		logger.Printf("checking %v (%v - %v)", humanize.Comma(int64(seeds[i+1])), seeds[i], seeds[i]+seeds[i+1])
		min = aoc.MinInt(min, findClosestSeedInRange(seeds[i], seeds[i+1], maps))
		elapsed := time.Since(start)
		logger.Printf("%v/sec", float64(seeds[i+1])/elapsed.Seconds())
		start = time.Now()
	}
	logger.Printf("min=%v", min)
	return nil
}

func findClosestSeed(seeds []int, maps []*rangemap.TheMap) int {
	min := math.MaxInt
	for _, seed := range seeds {
		n := seed
		for _, m := range maps {
			x := m.Map(n)
			logger.Printf("%v mapped %v to %v", m.Name, n, x)
			n = x
		}
		min = aoc.MinInt(min, n)
	}
	return min
}

func findClosestSeedInRange(start, length int, maps []*rangemap.TheMap) int {
	min := math.MaxInt
	for seed := start; seed < start+length; seed++ {
		n := seed
		for _, m := range maps {
			x := m.Map(n)
			//logger.Printf("%v mapped %v to %v", m.Name, n, x)
			n = x
		}
		min = aoc.MinInt(min, n)
	}
	return min
}

func readRangeMap(puzzle []string, startLine int) (int, *rangemap.TheMap) {
	i := startLine
	for ; puzzle[i] ==""; i++ {}
	rm := rangemap.New(puzzle[i])
	i++
	for ; i < len(puzzle) && puzzle[i] !=""; i++ {
		rm.Add(puzzle[i])
	}
	return i, rm
}