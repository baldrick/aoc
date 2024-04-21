package day21

import (
	"fmt"
	"testing"

    "github.com/baldrick/aoc/2023/aoc"
)

func Test_day21(t *testing.T) {
	tests := []struct {
		input []string
		steps, wantA, wantB int
	}{
		{
			input: aoc.PreparePuzzle(`
				...........
				.....###.#.
				.###.##..#.
				..#.#...#..
				....#.#....
				.##..S####.
				.##..#...#.
				.......##..
				.##.#.####.
				.##..##.##.
				...........
			`),
			steps: 6,
			wantA: 16,
			wantB: 0,
		},
	}
	for n, tc := range tests {
		t.Run(fmt.Sprintf("Test %v", n), func(t *testing.T) {
			testPart(t, "A", processA, tc.input, tc.steps, tc.wantA)
			testPart(t, "B", processB, tc.input, tc.steps, tc.wantB)
		})
	}
}

func testPart(t *testing.T, name string, process func([]string, int) (int, error), input []string, steps, want int) {
	t.Helper()
	if want == 0 {
		// Assume the test answer is never zero so if we do have
		// that specified, we're only on part A...
		return
	}
	got, err := process(input, steps)
	if err != nil {
		t.Errorf("%v failed: %v", name, err)
	}
	if got != want {
		t.Errorf("%v got %v, want %v", name, got, want)
	}
}
