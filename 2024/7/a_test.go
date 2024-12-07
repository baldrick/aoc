package day7

import (
	"fmt"
	"testing"
)

func Test_day7(t *testing.T) {
	tests := []struct {
		input        []string
		wantA, wantB int
	}{
		{
			input: []string{
				"190: 10 19",
				"3267: 81 40 27",
				"83: 17 5",
				"156: 15 6",
				"7290: 6 8 6 15",
				"161011: 16 10 13",
				"192: 17 8 14",
				"21037: 9 7 18 13",
				"292: 11 6 16 20",
			},
			wantA: 3749,
			wantB: 11387,
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

// func Test_Permutations(t *testing.T) {
// 	m := operator{"*", multiply}
// 	a := operator{"+", add}
// 	tests := []struct {
// 		name      string
// 		operators []operator
// 		terms     int
// 		want      [][]operator
// 	}{
// 		{
// 			operators: []operator{m, a},
// 			terms:     2,
// 			want:      [][]operator{{a}, {m}},
// 		},
// 		{
// 			operators: []operator{m, a},
// 			terms:     3,
// 			want:      [][]operator{{a, a}, {m, m}, {a, m}, {m, a}},
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got := calculatePermutations(tc.operators, tc.terms)
// 			t.Errorf("permutations:\n%v", got)
// 		})
// 	}
// }
