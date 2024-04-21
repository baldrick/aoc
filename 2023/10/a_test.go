package day10

import (
	"fmt"
	"log"
	"testing"
)

func Test_day10(t *testing.T) {
	tests := []struct {
		input []string
		wantA, wantB int
	}{
		// {
		// 	input: []string{
		// 		"-L|F7",
		// 		"7S-7|",
		// 		"L|7||",
		// 		"-L-J|",
		// 		"L|-JF",
		// 	},
		// 	wantA: 4,
		// 	wantB: -1,
		// },
		// {
		// 	input: []string{
		// 		"7-F7-",
		// 		".FJ|7",
		// 		"SJLL7",
		// 		"|F--J",
		// 		"LJ.LJ",
		// 	},
		// 	wantA: 8,
		// 	wantB: -1,
		// },
		{
			input: []string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			},
			wantA: 4,
			wantB: 1,
		},
		{
			input: []string{
				"......",
				".S--7.",
				".|..|.",
				".|F7|.",
				".LJLJ.",
				"......",
			},
			wantA: -1,
			wantB: -1, //2,
		},
		// {
		// 	input: []string{
		// 		"...........",
		// 		".S-------7.",
		// 		".|F-----7|.",
		// 		".||.....||.",
		// 		".||.....||.",
		// 		".|L-7.F-J|.",
		// 		".|..|.|..|.",
		// 		".L--J.L--J.",
		// 		"...........",
		// 	},
		// 	wantA: -1,
		// 	wantB: 4,
		// },
		// {
		// 	input: []string{
		// 		"..........",
		// 		".S------7.",
		// 		".|F----7|.",
		// 		".||....||.",
		// 		".||....||.",
		// 		".|L-7F-J|.",
		// 		".|..||..|.",
		// 		".L--JL--J.",
		// 		"..........",
		// 	},
		// 	wantA: -1,
		// 	wantB: 4,
		// },
		// {
		// 	input: []string{
		// 		".F----7F7F7F7F-7....",
		// 		".|F--7||||||||FJ....",
		// 		".||.FJ||||||||L7....",
		// 		"FJL7L7LJLJ||LJ.L-7..",
		// 		"L--J.L7...LJS7F-7L7.",
		// 		"....F-J..F7FJ|L7L7L7",
		// 		"....L7.F7||L7|.L7L7|",
		// 		".....|FJLJ|FJ|F7|.LJ",
		// 		"....FJL-7.||.||||...",
		// 		"....L---J.LJ.LJLJ...",
		// 	},
		// 	wantA: -1,
		// 	wantB: 8,
		// },
		// {
		// 	input: []string{
		// 		"FF7FSF7F7F7F7F7F---7",
		// 		"L|LJ||||||||||||F--J",
		// 		"FL-7LJLJ||||||LJL-77",
		// 		"F--JF--7||LJLJ7F7FJ-",
		// 		"L---JF-JLJ.||-FJLJJ7",
		// 		"|F|F-JF---7F7-L7L|7|",
		// 		"|FFJF7L7F-JF7|JL---7",
		// 		"7-L-JL7||F7|L7F-7F7|",
		// 		"L.L7LFJ|||||FJL7||LJ",
		// 		"L7JLJL-JLJLJL--JLJ.L",
		// 	},
		// 	wantA: -1,
		// 	wantB: 10,
		// },
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
	if want == -1 {
		// Assume the test answer is never negative so if we do have
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

func TestIsConnected(t *testing.T) {
	tests := []struct{
		input []string
		pipes [][]pipe
		x, y, x2, y2 int
		want bool
	}{
		{
			// ||
			// ..
			pipes: [][]pipe{
				[]pipe{newPipe(north, south), newPipe(north, south)},
				[]pipe{pipe{pt:ground}, pipe{pt:ground}},
			},
			x: 0,
			y: 1,
			x2: 0,
			y2: 0,
			want: false,
		},
		{
			// F7
			// ..
			pipes: [][]pipe{
				[]pipe{newPipe(east, south), newPipe(west, south)},
				[]pipe{pipe{pt:ground}, pipe{pt:ground}},
			},
			x: 0,
			y: 1,
			x2: 0,
			y2: 0,
			want: true,
		},
		{
			// --
			// ..
			pipes: [][]pipe{
				[]pipe{newPipe(east, west), newPipe(east, west)},
				[]pipe{pipe{pt:ground}, pipe{pt:ground}},
			},
			x: 0,
			y: 1,
			x2: 0,
			y2: 0,
			want: true,
		},
		{
			// .|
			// .|
			pipes: [][]pipe{
				[]pipe{pipe{pt:ground}, newPipe(north, south)},
				[]pipe{pipe{pt:ground}, newPipe(north, south)},
			},
			x: 0,
			y: 0,
			x2: 1,
			y2: 0,
			want: true,
		},
		{
			// .7
			// .J
			pipes: [][]pipe{
				[]pipe{pipe{pt:ground}, newPipe(south, west)},
				[]pipe{pipe{pt:ground}, newPipe(north, west)},
			},
			x: 0,
			y: 0,
			x2: 1,
			y2: 0,
			want: true,
		},
		{
			input: []string{
				"......",
				".S--7.",
				".|..|.",
				".|F7|.",
				".LJLJ.",
				"......",
			},
			x: 0,
			y: 0,
			x2: 1,
			y2: 0,
			want: false,
		},
		{
			// .-
			// .-
			pipes: [][]pipe{
				[]pipe{pipe{pt:ground}, newPipe(east, west)},
				[]pipe{pipe{pt:ground}, newPipe(east, west)},
			},
			x: 0,
			y: 0,
			x2: 1,
			y2: 0,
			want: false,
		},
	}
	for n, tc := range tests {
		t.Run(fmt.Sprintf("Test %v", n), func(t *testing.T) {
			if tc.pipes == nil {
				_, _, tc.pipes = getPipes(tc.input)
			}
			dump(tc.pipes, nil)
			got := connected(tc.pipes, tc.x, tc.y, tc.x2, tc.y2)
			t.Logf("%v: got %v, want %v", tc, got, tc.want)
			if got != tc.want {
				t.Errorf("got %v, want %v for %v", got, tc.want, tc)
			}
		})
	}
}

func TestLocationMap(t *testing.T) {
	m := make(map[location]status)
	m[location{1,2}] = pathPart
	m[location{2,1}] = inside
	m[location{3,4}] = outside
	m[location{3,4}] = maybeInside
	s, ok := m[location{3,4}]
	if !ok {
		t.Errorf("Failed to get status at 3,4")
	}
	if s != maybeInside {
		t.Errorf("Got %v, want %v at 3,4", s, maybeInside)
	}
	log.Printf("map: %v", m)
}