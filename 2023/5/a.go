package day5

import (
    _ "embed"
    "fmt"
    "log"
    "math"
    "time"

    "github.com/baldrick/aoc/2023/aoc"
	"github.com/baldrick/aoc/2023/rangemap"
	"github.com/dustin/go-humanize"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 5
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day5A",
        Usage: "Day 5 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day5B",
        Usage: "Day 5 part B",
        Action: partB,
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
    for _, line := range puzzle {
        log.Print(line)
    }
    return 0, fmt.Errorf("Not yet implemented")
}

func processB(puzzle []string) (int, error) {
	seeds, err := aoc.StripAndParse("seeds:", puzzle[0])
	if err != nil {
		return 0, err
	}
	log.Printf("seeds: %v", seeds)
	var maps []*rangemap.TheMap
	var i int
	var rm *rangemap.TheMap
	for i=1; i<len(puzzle); {
		i, rm = readRangeMap(puzzle, i)
		maps = append(maps, rm)
		log.Println(rm)
	}
	start := time.Now()
	min := math.MaxInt
	for i=0; i<len(seeds); i+=2 {
		log.Printf("checking %v (%v - %v)", humanize.Comma(int64(seeds[i+1])), seeds[i], seeds[i]+seeds[i+1])
		min = aoc.MinInt(min, findClosestSeedInRange(seeds[i], seeds[i+1], maps))
		elapsed := time.Since(start)
		log.Printf("%v/sec", float64(seeds[i+1])/elapsed.Seconds())
		start = time.Now()
	}
	return min, nil
}

func findClosestSeed(seeds []int, maps []*rangemap.TheMap) int {
	min := math.MaxInt
	for _, seed := range seeds {
		n := seed
		for _, m := range maps {
			x := m.Map(n)
			log.Printf("%v mapped %v to %v", m.Name, n, x)
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
			//log.Printf("%v mapped %v to %v", m.Name, n, x)
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