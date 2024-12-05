package day3

import (
	"fmt"
	"testing"
)

func Test_day3(t *testing.T) {
	tests := []struct {
		inputA, inputB []string
		wantA, wantB   int
	}{
		{
			inputA: []string{
				"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			inputB: []string{
				"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			wantA: 161,
			wantB: 48,
		},
		{
			inputB: []string{
				"don't()mul(2,3)do()mul(4,5)don't()mul(6,7)do()mul(9,8)",
			},
			wantB: 92,
		},
		{
			inputB: []string{
				"do()mul(2,3)do()mul(4,5)don't()don't()mul(6,7)do()don't()mul(9,8)do()",
			},
			wantB: 26,
		},
		{
			inputB: []string{
				"do()mul(2,3)do()mul(4,5)don't()don't()mul(6,7)do()don't()mul(9,8)do()",
			},
			wantB: 26,
		},
		{
			inputB: []string{
				"mmul(3,3)",
			},
			wantB: 9,
		},
		{
			inputA: []string{
				"xmul(8,5)&mul[3,7]!^don't()don't()do()don't()_mul(500,15)don't()+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			inputB: []string{
				"xmul(8,5)&mul[3,7]!^don't()don't()do()don't()_mul(500,15)don't()+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			wantA: 7668,
			wantB: 80,
		},
	}
	for n, tc := range tests {
		t.Run(fmt.Sprintf("Test %v", n), func(t *testing.T) {
			testPart(t, fmt.Sprintf("A-%v", n), processA, tc.inputA, tc.wantA)
			testPart(t, fmt.Sprintf("B-%v", n), processB, tc.inputB, tc.wantB)
		})
	}
}

func testPart(t *testing.T, name string, process func([]string) (int, error), input []string, want int) {
	t.Helper()
	t.Logf("########### %v ###########", name)
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
