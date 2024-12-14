package day14

import (
	"fmt"
	"testing"

	"github.com/baldrick/aoc/common/aoc"
)

func Test_day14(t *testing.T) {
	tests := []struct {
		input        []string
		wantA, wantB int
		size         aoc.PairInt
		ignore       bool
	}{
		{
			input: []string{
				"p=0,4 v=3,-3",
				"p=6,3 v=-1,-3",
				"p=10,3 v=-1,2",
				"p=2,0 v=2,-1",
				"p=0,0 v=1,3",
				"p=3,0 v=-2,-2",
				"p=7,6 v=-1,-3",
				"p=3,0 v=-1,-2",
				"p=9,3 v=2,3",
				"p=7,3 v=-1,2",
				"p=2,4 v=2,-3",
				"p=9,5 v=-3,-3",
			},
			size:   aoc.NewPairInt(11, 7),
			wantA:  12,
			wantB:  0,
			ignore: false,
		},
	}
	for n, tc := range tests {
		if tc.ignore {
			continue
		}
		t.Run(fmt.Sprintf("Test %v", n), func(t *testing.T) {
			testPart(t, "A", processA, tc.input, tc.size, tc.wantA)
			testPart(t, "B", processB, tc.input, tc.size, tc.wantB)
		})
	}
}

func testPart(t *testing.T, name string, process func([]string, aoc.PairInt) (int, error), input []string, size aoc.PairInt, want int) {
	t.Helper()
	if want == 0 {
		// Assume the test answer is never zero so if we do have
		// that specified, we're only on part A...
		return
	}
	got, err := process(input, size)
	if err != nil {
		t.Errorf("%v failed: %v", name, err)
	}
	if got != want {
		t.Errorf("%v got %v, want %v", name, got, want)
	}
}
