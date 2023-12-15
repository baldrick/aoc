package day15

import (
	"fmt"
	"testing"
)

func Test_day15(t *testing.T) {
	tests := []struct {
		input []string
		wantA, wantB int
	}{
		{
			input: []string{
				"HASH",
			},
			wantA: 52,
			wantB: 0,
		},
		{
			input: []string{
				"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			},
			wantA: 1320,
			wantB: 145,
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

func Test_add(t *testing.T) {
	tests := []struct{
		start *lenses
		label string
		focalLength int
		want *lenses
	} {
		{
			start: (&lenses{this:&lens{label:"a",focalLength:1}}).add("b",2).add("c",3),
			label: "b",
			focalLength: 10,
			want: (&lenses{this:&lens{label:"a",focalLength:1}}).add("b",10).add("c",3),
		},
	}

	for n, tc := range tests {
		t.Run(fmt.Sprintf("Test_remove %v", n), func(t *testing.T) {
			got := tc.start.add(tc.label, tc.focalLength)
			if got.String() != tc.want.String() {
				t.Errorf("Failed: got %v, want %v", got, tc.want)
			}
		})
	}
}

func Test_remove(t *testing.T) {
	tests := []struct{
		start *lenses
		label string
		want *lenses
	} {
		{
			start: (&lenses{this:&lens{label:"a",focalLength:1}}).add("b",2).add("c",3),
			label: "a",
			want: (&lenses{this:&lens{label:"b",focalLength:2}}).add("c",3),
		},
		{
			start: (&lenses{this:&lens{label:"a",focalLength:1}}).add("b",2).add("c",3),
			label: "b",
			want: (&lenses{this:&lens{label:"a",focalLength:1}}).add("c",3),
		},
		{
			start: (&lenses{this:&lens{label:"a",focalLength:1}}).add("b",2).add("c",3),
			label: "c",
			want: (&lenses{this:&lens{label:"a",focalLength:1}}).add("b",2),
		},
		{
			start: &lenses{this:&lens{label:"a",focalLength:1}},
			label: "a",
			want: nil,
		},
	}

	for n, tc := range tests {
		t.Run(fmt.Sprintf("Test_remove %v", n), func(t *testing.T) {
			got := tc.start.remove(tc.label)
			if got.String() != tc.want.String() {
				t.Errorf("Failed: got %v, want %v", got, tc.want)
			}
		})
	}
}