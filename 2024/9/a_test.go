package day9

import (
	"fmt"
	"testing"
)

func Test_day9(t *testing.T) {
	tests := []struct {
		input        []string
		wantA, wantB int
		ignore       bool
	}{
		{
			input: []string{
				"2333133121414131402",
			},
			wantA:  1928,
			wantB:  2858,
			ignore: false,
		},
		{
			input: []string{
				"213",
			},
			wantA:  1,
			wantB:  0,
			ignore: true,
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

// func Test_move(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		fis       []fileInfo
// 		id        int
// 		freespace []fileInfo
// 		size      int
// 	}{
// 		{
// 			name: "0.11",
// 			fis: []fileInfo{
// 				{id: 0, start: 0, size: 1},
// 				{id: 1, start: 2, size: 2},
// 			},
// 			freespace: []fileInfo{{id: -1, start: 1, size: 1}},
// 			size:      1,
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Logf("tc fis=%v, freespace=%v", tc.fis, tc.freespace)
// 			got := move(tc.fis, tc.id, &tc.freespace, tc.size)
// 			t.Errorf("got fis=%v, freespace=%v", got, tc.freespace)
// 		})
// 	}
// }
