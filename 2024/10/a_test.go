package day10

import (
	"fmt"
	"testing"
)

func Test_day10(t *testing.T) {
	tests := []struct {
		input        []string
		wantA, wantB int
		ignore       bool
	}{
		{
			input: []string{
				"0123",
				"1234",
				"8765",
				"9876",
			},
			wantA:  1,
			wantB:  0,
			ignore: true,
		},
		{
			input: []string{
				"...0...",
				"...1...",
				"...2...",
				"6543456",
				"7.....7",
				"8.....8",
				"9.....9",
			},
			wantA:  2,
			wantB:  0,
			ignore: true,
		},
		{
			input: []string{
				"..90..9",
				"...1.98",
				"...2..7",
				"6543456",
				"765.987",
				"876....",
				"987....",
			},
			wantA:  4,
			wantB:  0,
			ignore: true,
		},
		{
			input: []string{
				"10..9..",
				"2...8..",
				"3...7..",
				"4567654",
				"...8..3",
				"...9..2",
				".....01",
			},
			wantA:  3,
			wantB:  0,
			ignore: true,
		},
		{
			input: []string{
				"89010123",
				"78121874",
				"87430965",
				"96549874",
				"45678903",
				"32019012",
				"01329801",
				"10456732",
			},
			wantA:  36,
			wantB:  81,
			ignore: false,
		},
		{
			input: []string{
				".....0.",
				"..4321.",
				"..5..2.",
				"..6543.",
				"..7..4.",
				"..8765.",
				"..9....",
			},
			wantA:  0,
			wantB:  3,
			ignore: false,
		},
	}
	for n, tc := range tests {
		if tc.ignore {
			continue
		}
		t.Run(fmt.Sprintf("Test %v", n), func(t *testing.T) {
			testPart(t, "A", processA, tc.input, tc.wantA)
			testPart(t, "B", processB, tc.input, tc.wantB)
		})
	}
}

func testPart(t *testing.T, name string, process func([]string) (int, error), input []string, want int) {
	t.Helper()
	if want == 0 {
		// Assume the test answer is never zero so if we do have
		// that specified, we're only on part A...
		return
	}
	got, err := process(input)
	if err != nil {
		t.Errorf("%v failed: %v", name, err)
	}
	if got != want {
		t.Errorf("%v got %v, want %v", name, got, want)
	}
}
