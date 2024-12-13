package day13

import (
	"fmt"
	"testing"
)

func Test_day13(t *testing.T) {
	tests := []struct {
		input        []string
		wantA, wantB int
		ignore       bool
	}{
		{
			input: []string{
				"Button A: X+94, Y+34",
				"Button B: X+22, Y+67",
				"Prize: X=8400, Y=5400",
				"",
				"Button A: X+26, Y+66",
				"Button B: X+67, Y+21",
				"Prize: X=12748, Y=12176",
				"",
				"Button A: X+17, Y+86",
				"Button B: X+84, Y+37",
				"Prize: X=7870, Y=6450",
				"",
				"Button A: X+69, Y+23",
				"Button B: X+27, Y+71",
				"Prize: X=18641, Y=10279",
			},
			wantA:  480,
			wantB:  0,
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
