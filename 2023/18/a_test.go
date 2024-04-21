package day18

import (
	"fmt"
	"testing"
)

func Test_day18(t *testing.T) {
	tests := []struct {
		input []string
		wantA, wantB int
	}{
		{
			input: []string{
				"R 6 (#70c710)",
				"D 5 (#0dc571)",
				"L 2 (#5713f0)",
				"D 2 (#d2c081)",
				"R 2 (#59c680)",
				"D 2 (#411b91)",
				"L 5 (#8ceee2)",
				"U 2 (#caa173)",
				"L 1 (#1b58a2)",
				"U 2 (#caa171)",
				"R 2 (#7807d2)",
				"U 3 (#a77fa3)",
				"L 2 (#015232)",
				"U 2 (#7a21e3)",
			},
			wantA: 62,
			wantB: 952408144115,
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
