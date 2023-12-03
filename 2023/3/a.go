package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

	"github.com/baldrick/aoc/AdventOfCode/2023/aoc"
)

const (
	year = 2023
	day = 3
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
/*
x1-x2, y n
x, y symbol

for each number, is there a symbol in box x1-1,y-1 to x2+1,y+1
*/

type Point struct {
	X int
	Y int
}

type Range struct {
	Start, End Point
	N int
}

func process(puzzle []string) error {
	ranges := make(map[int][]Range)
	symbols := make(map[int][]Range)
	var err error
	for y, line := range puzzle {
		ranges[y], symbols[y], err = extractNumbers(line, y)
		if err != nil {
			return err
		}
	}
	dump("Numbers", ranges)
	dump("Symbols", symbols)
	total := 0
	xMax := len(puzzle[0])
	for i:=0; i<len(ranges); i++ {
		r, ok := ranges[i]
		if !ok {
			logger.Fatalf("Range %v not found", i)
		}
		logger.Printf("#%v: %v -> %v", i, puzzle[i], r)
		total += symbolNearby(symbols, r, xMax)
	}
	logger.Printf("total = %v", total)
	return nil
}

func dump(m string, r map[int][]Range) {
	logger.Print(m)
	for i:=0; i<=len(r)-1; i++ {
		v, ok := r[i]
		if !ok {
			logger.Fatalf("%v not found!", i)
		}
		logger.Printf("%v: %v", i, v)
	}
}

func extractNumbers(line string, y int) ([]Range, []Range, error) {
	validNumbers := "0123456789"
	var ranges []Range
	var symbols []Range
	start := -1
	for x, c := range line {
		if strings.Contains(validNumbers, string(c)) {
			if start < 0 {
				// We've found the start of a number.
				start = x
			}
			continue
		}
		if start >= 0 {
			// We've found the end of a number.
			end := x
			number, err := strconv.Atoi(line[start:end])
			if err != nil {
				return nil, nil, err
			}
			//logger.Printf("%v,%v (%v) -> %v", start, end, line[start:end], number)
			ranges = append(ranges, Range{Start: Point{start, y}, End: Point{end-1, y}, N: number})
			start = -1
		}
		if c == '.' {
			// Not a symbol, move on.
			continue
		}
		symbols = append(symbols, Range{Start: Point{x, y}})
	}
	if start != -1 {
		// We've found the end of a number.
		end := len(line)
		number, err := strconv.Atoi(line[start:end])
		if err != nil {
			return nil, nil, err
		}
		//logger.Printf("%v,%v (%v) -> %v", start, end, line[start:end], number)
		ranges = append(ranges, Range{Start: Point{start, y}, End: Point{end-1, y}, N: number})
		start = -1
	}
	if line[len(line)-1] != '.' && !strings.Contains(validNumbers, string(line[len(line)-1])) {
		symbols = append(symbols, Range{Start: Point{len(line)-1, y}})
	}
	//logger.Printf("line %v (%q) -> %v (%v)", y, line, ranges, symbols)
	return ranges, symbols, nil
}

func symbolNearby(symbols map[int][]Range, ranges []Range, xMax int) int {
	total := 0
	for _, r := range ranges {
		total += symbolNearbyRange(symbols, r, xMax)
	}
	return total
}

func symbolNearbyRange(symbols map[int][]Range, r Range, xMax int) int {
	yStart := aoc.MaxInt(0, r.Start.Y-1)
	yEnd := aoc.MinInt(len(symbols), r.End.Y+1)
	xStart := aoc.MaxInt(0, r.Start.X-1)
	xEnd := aoc.MinInt(xMax, r.End.X+1)
	logger.Printf("Looking for symbol in %v,%v - %v,%v", xStart, yStart, xEnd, yEnd)
	for ty := yStart; ty <= yEnd; ty++ {
		ySymbols, ok := symbols[ty]
		if !ok {
			continue
		}
		for _, s := range ySymbols {
			if s.Start.X >= xStart && s.Start.X <= xEnd {
				logger.Printf("%v has %v nearby, adding %v", r, s, r.N)
				return r.N
			}
		}
	}
	return 0
}
