package day2

import (
	"fmt"
	"testing"
)

func Test_day2(t *testing.T) {
	tests := []struct {
		input        []string
		wantA, wantB int
	}{
		{
			input: []string{
				"7 6 4 2 1", // safe
				"1 2 7 8 9", // unsafe
				"9 7 6 2 1", // unsafe
				"1 3 2 4 5", // safe
				"8 6 4 4 1", // safe
				"1 3 6 7 9", // safe
				// mine
				"9 4 3 2 1",            // safe
				"1 2 3 4 9",            // safe
				"1 2 3 4 3",            // safe
				"1 5 4 3 2",            // safe
				"1 2 9 3 9",            // unsafe
				"9 5 4 0 -1",           // unsafe
				"1 9 8 7 6",            // safe
				"1 4 3 5 7",            // safe
				"1 4 3 2 5",            // unsafe
				"9 8 7 6 1",            // safe
				"9 8 1 7 1",            // unsafe
				"1 2 1 2 3",            // unsafe
				"1 2",                  // safe
				"2 1",                  // safe
				"2 1 9",                // safe
				"9 1 2",                // safe
				"1 2 9 8",              // unsafe
				"1 2 3 9",              // safe
				"9 1 2 3",              // safe
				"1 1 2 3 4",            // safe
				"48 46 47 49 51 54 56", // all safe below, thx to reddit
				"1 1 2 3 4 5",
				"1 2 3 4 5 5",
				"5 1 2 3 4 5",
				"1 4 3 2 1",
				"1 6 7 8 9",
				"1 2 3 4 3",
				"9 8 7 6 7",
				"7 10 8 10 11",
				"29 28 27 25 26 25 22 20",
				"7 10 8 10 11",
				"29 28 27 25 26 25 22 20",
				"8 9 10 11", // safe
				"9 8 7 7 7", // unsafe
				"90 89 86 84 83 79",
				"97 96 93 91 85",
				"29 26 24 25 21",
				"36 37 40 43 47",
				"43 44 47 48 49 54",
				"35 33 31 29 27 25 22 18",
				"77 76 73 70 64",
				"68 65 69 72 74 77 80 83",
				"37 40 42 43 44 47 51",
				"70 73 76 79 86",
			},
			wantA: 0,
			wantB: 41,
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
