package day19

import (
	"fmt"
	"testing"
)

func Test_day19(t *testing.T) {
	tests := []struct {
		input []string
		wantA, wantB int
	}{
		{
			input: []string{
				"px{a<2006:qkq,m>2090:A,rfg}",
				"pv{a>1716:R,A}",
				"lnx{m>1548:A,A}",
				"rfg{s<537:gd,x>2440:R,A}",
				"qs{s>3448:A,lnx}",
				"qkq{x<1416:A,crn}",
				"crn{x>2662:A,R}",
				"in{s<1351:px,qqz}",
				"qqz{s>2770:qs,m<1801:hdj,R}",
				"gd{a>3333:R,R}",
				"hdj{m>838:A,pv}",
				"",
				"{x=787,m=2655,a=1222,s=2876}",
				"{x=1679,m=44,a=2067,s=496}",
				"{x=2036,m=264,a=79,s=2244}",
				"{x=2461,m=1339,a=466,s=291}",
				"{x=2127,m=1623,a=2188,s=1013}",
			},
			wantA: 19114,
			wantB: 0,
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
