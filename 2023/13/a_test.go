package day13

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_day13(t *testing.T) {
	tests := []struct {
		input []string
		wantA, wantB int
	}{
		{
			input: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
				"",
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			wantA: 405,
			wantB: 400,
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

func Test_transpose(t *testing.T) {
	tests := []struct {
		input []string
		want []string
	}{
		{
			input: []string{
				"#.##",
				"..#.",
				"##..",
			},
			want: []string{
				"#.#",
				"..#",
				"##.",
				"#..",
			},
		},
	}
	for n, tc := range tests {
		t.Run(fmt.Sprintf("Test %v", n), func(t *testing.T) {
			got := transpose(tc.input)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("transpose diffs:\n%v", diff)
			}
		})
	}
}

func Test_flip(t *testing.T) {
	tests := []struct{
		in []string
		n int // 0-based index
		want []string
	}{
		{
			in: []string{
				"#..#",
				"...#",
			},
			n: 6,
			want: []string{
				"#..#",
				"..##",
			},
		},
	}
	for n, tc := range tests {
		t.Run(fmt.Sprintf("Test flip %v", n), func(t *testing.T) {
			flip(tc.in, tc.n)
			if diff := cmp.Diff(tc.in, tc.want); diff != "" {
				t.Errorf("Flip failed, diff:\n%v", diff)
			}
		})
	}
}