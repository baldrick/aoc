package day14

import (
	"fmt"
//	"log"
	"testing"

	"github.com/baldrick/aoc/2023/grid"
)

func Test_day14(t *testing.T) {
	tests := []struct {
		input []string
		wantA, wantB int
	}{
		{
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			wantA: 0, //136,
			wantB: 0, //64,
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

func TestCycle(t *testing.T) {
	tests := []struct{
		g *grid.Grid
		dx, dy []int
		want *grid.Grid
	}{
		{
			g: grid.New([]string{"O.."}),
			dx: []int{1},
			dy: []int{0},
			want: grid.New([]string{"..O"}),
		},
		{
			g: grid.New([]string{"O..#..O."}),
			dx: []int{1},
			dy: []int{0},
			want: grid.New([]string{"..O#...O"}),
		},
		{
			g: grid.New([]string{"..O"}),
			dx: []int{-1},
			dy: []int{0},
			want: grid.New([]string{"O.."}),
		},
		{
			g: grid.New([]string{
				"O",
				".",
				".",
			}),
			dx: []int{0},
			dy: []int{1},
			want: grid.New([]string{
				".",
				".",
				"O",
			}),
		},
		{
			g: grid.New([]string{
				".",
				".",
				"O",
			}),
			dx: []int{0},
			dy: []int{-1},
			want: grid.New([]string{
				"O",
				".",
				".",
			}),
		},
		{
			g: grid.New([]string{
				"..#",
				".#.",
				"O.O",
			}),
			dx: []int{0},
			dy: []int{-1},
			want: grid.New([]string{
				"O.#",
				".#O",
				"...",
			}),
		},
		{
			g: grid.New([]string{
				"..#",
				".#.",
				"O.O",
			}),
			dx: []int{0, 1},
			dy: []int{-1, 0},
			want: grid.New([]string{
				".O#",
				".#O",
				"...",
			}),
		},
	}
	for n, tc := range tests {
		t.Run(fmt.Sprintf("TestCycle #%v", n), func(t *testing.T) {
			//log.Printf("Start:\n%v", tc.g)
			for n:=0; n<len(tc.dx); n++ {
				cycle(tc.g, tc.dx[n], tc.dy[n])
				//log.Printf("After %v,%v:\n%v", tc.dx[n], tc.dy[n], tc.g)
			}
			if tc.g.Key() != tc.want.Key() {
				t.Errorf("got\n%v, want\n%v", tc.g, tc.want)
			}
		})
	}
}

func TestFullCycle(t *testing.T) {
	tests := []struct{
		g *grid.Grid
		cycleCount int
		want *grid.Grid
	}{
		{
			g: grid.New([]string{
				".....",
				".###.",
				".#O#.",
				".###.",
				".....",
			}),
			cycleCount: 1e5,
			want: grid.New([]string{
				".....",
				".###.",
				".#O#.",
				".###.",
				".....",
			}),
		},
		{
			g: grid.New([]string{
				".....",
				".#.#.",
				".#O#.",
				".###.",
				".....",
			}),
			cycleCount: 1e5,
			want: grid.New([]string{
				".....",
				".#.#.",
				".#.#.",
				".###.",
				"....O",
			}),
		},
		{
			g: grid.New([]string{
				".....",
				".###.",
				"..O#.",
				".###.",
				".....",
			}),
			cycleCount: 1e5,
			want: grid.New([]string{
				".....",
				".###.",
				"...#.",
				".###.",
				"....O",
			}),
		},
		{
			g: grid.New([]string{
				".....",
				".###.",
				".#O#.",
				".#.#.",
				".....",
			}),
			cycleCount: 1e5,
			want: grid.New([]string{
				".....",
				".###.",
				".#.#.",
				".#.#.",
				"....O",
			}),
		},
		{
			g: grid.New([]string{
				".....",
				".###.",
				".#O..",
				".###.",
				".....",
			}),
			cycleCount: 1e5,
			want: grid.New([]string{
				".....",
				".###.",
				".#...",
				".###.",
				"....O",
			}),
		},
	}
	for n, tc := range tests {
		t.Run(fmt.Sprintf("TestFullCycle #%v", n), func(t *testing.T) {
			for cycles := 0; cycles < tc.cycleCount; cycles++ {
				for _, s := range []struct{
					dx int
					dy int
				}{
					{0,-1},
					{-1,0},
					{0, 1},
					{1, 0},
				} {
					cycle(tc.g, s.dx, s.dy)
				}
			}
		})
		if tc.g.Key() != tc.want.Key() {
			t.Errorf("got\n%v, want\n%v", tc.g, tc.want)
		}
	}
}