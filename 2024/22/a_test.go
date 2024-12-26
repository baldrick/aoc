package day22

import (
	"fmt"
	"strings"
	"testing"
)

func Test_day22(t *testing.T) {
	tests := []struct {
		input        string
		wantA, wantB int
		ignore       bool
	}{
		{
			input: `1
10
100
2024
`,
			wantA:  37327623,
			wantB:  23,
			ignore: true,
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

func TestSecretGenerator(t *testing.T) {
	tests := []struct {
		name, start string
		iterations  int
		want        []int
	}{
		{
			start:      "123",
			iterations: 10,
			want: []int{
				15887950,
				16495136,
				527345,
				704524,
				1553684,
				12683156,
				11100544,
				12249484,
				7753432,
				5908254,
			},
		},
	}

	for _, tc := range tests {
		s := newSecret(tc.start)
		for n := 0; n < tc.iterations; n++ {
			s.quickNext()
			if s.n != tc.want[n] {
				t.Errorf("%v: %vth iteration want %v got %v", tc.name, n, tc.want[n], s.n)
			}
			t.Logf("%v", s)
		}
	}
}
