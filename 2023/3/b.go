package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

	"github.com/baldrick/aoc/2023/aoc"
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
	S string
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
	for i:=0; i<len(symbols); i++ {
		s, ok := symbols[i]
		if !ok {
			logger.Fatalf("Symbol %v not found", i)
		}
		total += gearRatio(ranges, s, xMax)
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
		symbols = append(symbols, Range{Start: Point{x, y}, S: string(c)})
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
		symbols = append(symbols, Range{Start: Point{len(line)-1, y}, S: string(line[len(line)-1])})
	}
	//logger.Printf("line %v (%q) -> %v (%v)", y, line, ranges, symbols)
	return ranges, symbols, nil
}

func gearRatio(ranges map[int][]Range, symbols []Range, xMax int) int {
	total := 0
	for _, s := range symbols {
		if s.S != "*" {
			continue
		}
		numbers := numbersNearSymbol(ranges, s, xMax)
		if len(numbers) > 2 {
			logger.Fatalf("Found multiple numbers (%v) for %v", numbers, s)
		}
		if len(numbers) < 2 {
			continue
		}
		total += numbers[0] * numbers[1]
	}
	return total
}

func numbersNearSymbol(ranges map[int][]Range, s Range, xMax int) []int {
	yStart := aoc.MaxInt(0, s.Start.Y-1)
	yEnd := aoc.MinInt(len(ranges), s.Start.Y+1)
	logger.Printf("Looking for numbers next to %v,(%v-%v)", s.Start.X, yStart, yEnd)
	var numbers []int
	for ty := yStart; ty <= yEnd; ty++ {
		yr, ok := ranges[ty]
		if !ok {
			logger.Fatalf("Ranges not found at %v", ty)
		}
		for _, r := range yr {
			xrStart := aoc.MaxInt(0, r.Start.X-1)
			xrEnd := aoc.MinInt(xMax, r.End.X+1)
			logger.Printf("Looking for number at %v-%v,%v", xrStart, xrEnd, ty)
			if s.Start.X >= xrStart && s.Start.X <= xrEnd {
				numbers = append(numbers, r.N)
			}
		}
	}
	return numbers
}
