package day5

import (
	"fmt"
	"testing"
)

func Test_day5(t *testing.T) {
	tests := []struct {
		input        []string
		wantA, wantB int
	}{
		{
			input: []string{
				"47|53",
				"97|13",
				"97|61",
				"97|47",
				"75|29",
				"61|13",
				"75|53",
				"29|13",
				"97|29",
				"53|29",
				"61|53",
				"97|53",
				"61|29",
				"47|13",
				"75|47",
				"97|75",
				"47|61",
				"75|61",
				"47|29",
				"75|13",
				"53|13",
				"75,47,61,53,29",
				"97,61,53,29,13",
				"75,29,13",
				"75,97,47,61,53",
				"61,13,29",
				"97,13,75,29,47",
			},
			wantA: 143,
			wantB: 123,
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
