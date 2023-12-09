package day6

import (
	"fmt"
	"testing"
)

func Test_day6(t *testing.T) {
	tests := []struct {
		input []string
		wantA, wantB int
	}{
		{
			input: []string{
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			wantA: 288,
			wantB: 71503,
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
