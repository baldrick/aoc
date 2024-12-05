package day4

import (
	"fmt"
	"testing"
)

func Test_day4(t *testing.T) {
	tests := []struct {
		input        []string
		wantA, wantB int
	}{
		{
			input: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			wantA: 18,
			wantB: 9,
		},
		{
			input: []string{
				"M.MM.S",
				".A..A.",
				"S.SM.S",
			},
			wantA: 0,
			wantB: 2,
		},
		{
			input: []string{
				"S.SS.M",
				".A..A.",
				"M.MS.M",
			},
			wantA: 0,
			wantB: 2,
		},
		{
			input: []string{
				".SS.M",
				"A..A.",
				".MS.M",
			},
			wantA: 0,
			wantB: 1,
		},
		{
			input: []string{
				"S.SS.",
				".A..A",
				"M.MS.",
			},
			wantA: 0,
			wantB: 1,
		},
		{
			input: []string{
				"S.SS.M",
				".A..A.",
				"M.MM.S",
			},
			wantA: 0,
			wantB: 1,
		},
		{
			input: []string{
				"S.SS.M",
				".A..A.",
				"M.MM.M",
			},
			wantA: 0,
			wantB: 1,
		},
		{
			input: []string{
				"S.SS.S",
				".A..A.",
				"M.MM.S",
			},
			wantA: 0,
			wantB: 1,
		},
	}
	for n, tc := range tests {
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
