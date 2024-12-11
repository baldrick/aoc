package day11

import (
	"fmt"
	"testing"
)

func Test_day11(t *testing.T) {
	tests := []struct {
		input        []string
		blinks       int
		wantA, wantB int
		ignore       bool
	}{
		{
			input: []string{
				"0 1 10 99 999",
			},
			blinks: 1,
			wantA:  7,
			wantB:  0,
			ignore: true,
		},
		{
			input: []string{
				"125 17",
			},
			blinks: 6,
			wantA:  22,
			wantB:  0,
			ignore: false,
		},
		{
			input: []string{
				"125 17",
			},
			blinks: 25,
			wantA:  55312,
			wantB:  0,
			ignore: false,
		},
	}
	for n, tc := range tests {
		if tc.ignore {
			continue
		}
		t.Run(fmt.Sprintf("Test %v", n), func(t *testing.T) {
			testPart(t, "A", processA, tc.input, tc.blinks, tc.wantA)
			testPart(t, "B", processB, tc.input, tc.blinks, tc.wantB)
		})
	}
}

func testPart(t *testing.T, name string, process func([]string, int) (int, error), input []string, blinks, want int) {
	t.Helper()
	if want == 0 {
		// Assume the test answer is never zero so if we do have
		// that specified, we're only on part A...
		return
	}
	got, err := process(input, blinks)
	if err != nil {
		t.Errorf("%v failed: %v", name, err)
	}
	if got != want {
		t.Errorf("%v got %v, want %v", name, got, want)
	}
}
