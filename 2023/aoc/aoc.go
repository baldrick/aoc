package aoc

import (
	"bufio"
	"os"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

var (
	Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "f",
			Value: "test",
			Usage: "input file for the puzzle",
		},
	}
)

func GetPuzzleInput(inputFile string, year, day int) ([]string, error) {
	if !strings.Contains(inputFile, ".txt") {
		inputFile += ".txt"
	}
	if !strings.Contains(inputFile, fmt.Sprintf("%v/%v/", year, day)) {
		inputFile = fmt.Sprintf("%v/%v/%v", year, day, inputFile)
	}
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var puzzle []string
	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, line)
	}
	return puzzle, nil
}

func PreparePuzzle(puzzle string) []string {
	lines := strings.Split(puzzle, "\n")
    for ; len(lines[len(lines)-1]) == 0 ; {
        lines = lines[:len(lines)-1]
	}
	log.Printf("Puzzle contains %v lines", len(lines))
    return lines
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type StringSet struct {
	s map[string]struct{}
}

func NewStringSet() *StringSet {
	return &StringSet{s: make(map[string]struct{})}
}

func (ss *StringSet) Add(s string) {
	if len(s) == 0 {
		return
	}
	ss.s[s] = struct{}{}
}

func (ss *StringSet) String() string {
	var sb strings.Builder
	for n, _ := range ss.s {
		sb.WriteString(n)
		sb.WriteString(",")
	}
	return sb.String()
}

func (ss *StringSet) Intersect(other *StringSet) *StringSet {
	i := NewStringSet()
	for n, _ := range ss.s {
		_, ok := other.s[n]
		if ok {
			i.Add(n)
		}
	}
	return i
}

func (ss *StringSet) Len() int {
	return len(ss.s)
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to translate %q to number: %v", s, err))
	}
	return n
}

func MustAtoiAll(s []string) []int {
	var numbers []int
	for _, n := range s {
		numbers = append(numbers, MustAtoi(n))
	}
	return numbers
}

func StripAndParse(strip, line string) ([]int, error) {
	re := regexp.MustCompile(strip + ` *(.*)`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 2 {
		return nil, fmt.Errorf("Failed to parse %q: %v matches found, want 2", line, len(matches))
	}
	return ParseNumbers(matches[1])
}

func ParseNumbers(numberList string) ([]int, error) {
	re := regexp.MustCompile(` *([0-9]*)(.*)`)
	var numbers []int
	for ;; {
		matches := re.FindStringSubmatch(numberList)
		numbers = append(numbers, MustAtoi(matches[1]))
		if len(matches[2]) == 0 {
			break
		}
		numberList = matches[2]
	}
	return numbers, nil
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
