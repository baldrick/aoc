package rangemap

import (
	"fmt"

	"github.com/baldrick/aoc/AdventOfCode/2023/aoc"
)

type TheMap struct {
	Name string
	sourceStart, destStart, length []int
}

func New(name string) *TheMap {
	return &TheMap{Name: name}
}

func (m *TheMap) String() string {
	return fmt.Sprintf("%v source=%v dest=%v length=%v", m.Name, m.sourceStart, m.destStart, m.length)
}

func (m *TheMap) Add(line string) error {
	numbers, err := aoc.ParseNumbers(line)
	if err != nil {
		return err
	}
	m.sourceStart = append(m.sourceStart, numbers[1])
	m.destStart = append(m.destStart, numbers[0])
	m.length = append(m.length, numbers[2])
	return nil
}

func (m *TheMap) Map(n int) int {
	for i, _ := range m.sourceStart {
		if n < m.sourceStart[i] {
			continue
		}
		if n >= m.sourceStart[i] + m.length[i] {
			continue
		}
		return n - m.sourceStart[i] + m.destStart[i]
	}
	return n
}

/*
Rather than list every source number and its corresponding destination number one by one, 
the maps describe entire ranges of numbers that can be converted. Each line within a map contains 
three numbers: the destination range start, the source range start, and the range length.

Consider again the example seed-to-soil map:

50 98 2
52 50 48
The first line has a destination range start of 50, a source range start of 98, and a range length of 2. 
This line means that the source range starts at 98 and contains two values: 98 and 99. 
The destination range is the same length, but it starts at 50, so its two values are 50 and 51. 
With this information, you know that seed number 98 corresponds to soil number 50 and that seed number 99 
corresponds to soil number 51.

The second line means that the source range starts at 50 and contains 48 values: 50, 51, ..., 96, 97. 
This corresponds to a destination range starting at 52 and also containing 48 values: 52, 53, ..., 98, 99. 
So, seed number 53 corresponds to soil number 55.
*/