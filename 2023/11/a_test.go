package day11

import (
	"fmt"
	"testing"
)

func Test_day11(t *testing.T) {
	tests := []struct {
		input []string
		wantA int
		expansionRate []int
		wantB []int
	}{
		{
			input: []string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			},
			wantA: 374,
			expansionRate: []int{2, 10, 100},
			wantB: []int{374, 1030, 8410},
		},
		{
			input: []string{
				"#.#",
			},
			wantA: 3,
			expansionRate: []int{1, 2, 10, 100},
			wantB: []int{2, 3, 11, 101},
		},
		{
			input: []string{
				"#..#",
			},
			wantA: 5,
			expansionRate: []int{1, 2, 10, 100, 1000},
			wantB: []int{3, 5, 21, 201, 2001},
		},
		{
			input: []string{
				"#",
				".",
				"#",
			},
			wantA: 3,
			expansionRate: []int{1, 2, 10, 100},
			wantB: []int{2, 3, 11, 101},
		},
		{
			input: []string{
				"#",
				".",
				".",
				"#",
			},
			wantA: 5,
			expansionRate: []int{1, 2, 10, 100},
			wantB: []int{3, 5, 21, 201},
		},
		// {
		// 	input: []string{
		// 		"#..",
		// 		"...",
		// 		"..#",
		// 	},
		// 	wantA: 6,
		// 	expansionRate: []int{0, 1, 10, 100},
		// 	wantB: []int{4, 6, 24, 204},
		// },
		// {
		// 	input: []string{
		// 		"#...",
		// 		"....",
		// 		"....",
		// 		"...#",
		// 	},
		// 	wantA: 10,
		// 	expansionRate: []int{0, 1, 10, 100},
		// 	wantB: []int{6, 10, 46, 406},
		// },
		// {
		// 	input: []string{
		// 		"......",
		// 		".#....",
		// 		"......",
		// 		"...#..",
		// 		"......",
		// 	},
		// 	wantA: 6,
		// 	expansionRate: []int{0, 1, 10, 100},
		// 	wantB: []int{4, 6, 24, 204},
		// },
		// {
		// 	input: []string{
		// 		"#.#",
		// 		"...",
		// 		"..#",
		// 	},
		// 	wantA: 12,
		// 	expansionRate: []int{0, 1, 10, 100},
		// 	wantB: []int{8, 12, 48, 408},
		// },
	}
	for n, tc := range tests {
		t.Run(fmt.Sprintf("Test %v", n), func(t *testing.T) {
			testPart(t, "A", processA, tc.input, 1, tc.wantA)
			for i, er := range tc.expansionRate {
				testPart(t, "B", processB, tc.input, er, tc.wantB[i])
			}
		})
	}
}

func testPart(t *testing.T, name string, process func([]string, int) (int, error), input []string, expansionRate, want int) {
	t.Helper()
	if want == 0 {
		// Assume the test answer is never zero so if we do have
		// that specified, we're only on part A...
		return
	}
	got, err := process(input, expansionRate)
	if err != nil {
		t.Errorf("%v failed: %v", name, err)
	}
	if got != want {
		t.Errorf("%v got %v, want %v", name, got, want)
	}
}
