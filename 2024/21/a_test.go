package day21

import (
	"fmt"
	"strings"
	"testing"
)

func Test_day21(t *testing.T) {
	tests := []struct {
		input string
		wantA, wantB int
		ignore bool
	}{
		{
			input:
`
`,
			wantA: 0,
			wantB: 0,
			ignore: false,
		},
	}
	for n, tc := range tests {
		if tc.ignore {
			continue
		}
		t.Run(fmt.Sprintf("Test %v", n), func(t *testing.T) {
			puzzle := strings.Split(tc.input, "\n")
			testPart(t, "A", processA, puzzle, tc.wantA)
			testPart(t, "B", processB, puzzle, tc.wantB)
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
