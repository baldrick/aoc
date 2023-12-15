package aoc

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func PreparePuzzle(puzzle string) []string {
	lines := strings.Split(puzzle, "\n")
    for ; len(lines[len(lines)-1]) == 0 ; {
        lines = lines[:len(lines)-1]
	}
	log.Printf("Puzzle contains %v lines", len(lines))
    return lines
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
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

func ModInt(a, b int) int {
	return int(math.Mod(float64(a), float64(b)))
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

type IntSet struct {
	s map[int]struct{}
}

func NewIntSet() *IntSet {
	return &IntSet{s: make(map[int]struct{})}
}

func (s *IntSet) Add(i int) {
	s.s[i] = struct{}{}
}

func (s *IntSet) Contains(i int) bool {
	_, ok := s.s[i]
	return ok
}

func (s *IntSet) MapOver(f func(int)) {
	for k,_ := range s.s {
		f(k)
	}
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
