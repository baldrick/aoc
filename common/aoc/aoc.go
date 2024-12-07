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
	split := strings.Split(puzzle, "\n")
	re := regexp.MustCompile(`^\s*`)
	var lines []string
	for _, line := range split {
		stripped := re.ReplaceAllString(line, "")
		if len(stripped) == 0 {
			continue
		}
		lines = append(lines, stripped)
	}
	log.Printf("Puzzle contains %v lines:\n%v", len(lines), strings.Join(lines, "\n"))
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

func MakeStringSet(sa []string) *StringSet {
	set := NewStringSet()
	for _, s := range sa {
		set.Add(s)
	}
	return set
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

func (ss *StringSet) Contains(s string) bool {
	_, ok := ss.s[s]
	return ok
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
	for k, _ := range s.s {
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

func MustXtoi(s string) int {
	n, err := strconv.ParseUint(s, 16, 0)
	if err != nil {
		panic(fmt.Sprintf("Failed to translate %q to number: %v", s, err))
	}
	return int(n)
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
	for {
		matches := re.FindStringSubmatch(numberList)
		numbers = append(numbers, MustAtoi(matches[1]))
		if len(matches[2]) == 0 {
			break
		}
		numberList = matches[2]
	}
	return numbers, nil
}

// GCD finds the grreatest common divisor via Euclidean algorithm.
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM finds the lowest common multiple via GCD.
// integers must contain at least two numbers.
func LCM(integers ...int) int {
	if len(integers) < 2 {
		panic(fmt.Sprintf("LCM requires at least two integers not just %v: %v", len(integers), integers))
	}
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type PairInt struct {
	a, b int
}
