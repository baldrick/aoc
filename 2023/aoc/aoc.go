package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetPuzzleInput(inputFile string, year, day int) ([]string, error) {
	if !strings.Contains(inputFile, ".txt") {
		inputFile += ".txt"
	}
	if !strings.Contains(inputFile, fmt.Sprintf("AdventOfCode/%v/%v/", year, day)) {
		inputFile = fmt.Sprintf("AdventOfCode/%v/%v/%v", year, day, inputFile)
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