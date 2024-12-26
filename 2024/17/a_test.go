package day17

import (
	"fmt"
	"strings"
	"testing"
)

func Test_day17(t *testing.T) {
	tests := []struct {
		input  string
		wantA  string
		wantB  int
		ignore bool
	}{
		{
			input: `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`,
			wantA:  "4,6,3,5,6,3,5,2,1,0",
			ignore: false,
		},
		{
			input: `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
`,
			wantB:  117440,
			ignore: false,
		},
	}
	for n, tc := range tests {
		if tc.ignore {
			continue
		}
		t.Run(fmt.Sprintf("Test %v", n), func(t *testing.T) {
			puzzle := strings.Split(tc.input, "\n")
			testPartA(t, "A", processA, puzzle, tc.wantA)
			testPartB(t, "B", processB, puzzle, tc.wantB)
		})
	}
}

func testPartA(t *testing.T, name string, process func([]string) (string, error), input []string, want string) {
	t.Helper()
	if len(want) == 0 {
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

func testPartB(t *testing.T, name string, process func([]string) (int, error), input []string, want int) {
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

func Test_Discrete(t *testing.T) {
	tests := []struct {
		name       string
		m          machine
		program    string
		want       *machine
		wantOutput string
	}{
		{
			name:    "1",
			m:       machine{c: 9},
			program: "2,6",
			want:    &machine{b: 1, c: 9},
		},
		{
			name:       "2",
			m:          machine{a: 10},
			program:    "5,0,5,1,5,4",
			wantOutput: "0,1,2",
		},
		{
			name:       "3",
			m:          machine{a: 2024},
			program:    "0,1,5,4,3,0",
			wantOutput: "4,2,5,6,7,7,7,7,3,1,0",
			want:       &machine{a: 0, b: -1, c: -1},
		},
		{
			name:    "4",
			m:       machine{b: 29},
			program: "1,7",
			want:    &machine{a: -1, b: 26, c: -1},
		},
		{
			name:    "5",
			m:       machine{b: 2024, c: 43690},
			program: "4,0",
			want:    &machine{a: -1, b: 44354, c: -1},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			output := tc.m.run(strings.Split(tc.program, ","), false)
			if tc.want != nil {
				if (tc.want.a != -1 && tc.want.a != tc.m.a) ||
					(tc.want.b != -1 && tc.want.b != tc.m.b) ||
					(tc.want.c != -1 && tc.want.c != tc.m.c) {
					t.Errorf("run end machine is unexpected, got %q, want %q", tc.m, tc.want)
				}
			}
			if tc.wantOutput != output {
				t.Errorf("run output unexpected, got %q, want %q", output, tc.wantOutput)
			}
		})
	}
}
