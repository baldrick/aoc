package day13

import (
	_ "embed"
	"log"
	"math"
	"regexp"
	"sort"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day13A",
		Aliases: []string{"day13a"},
		Usage:   "Day 13 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day13B",
		Aliases: []string{"day13b"},
		Usage:   "Day 13 part B",
		Action:  partB,
	}
)

func partA(ctx *cli.Context) error {
	answer, err := processA(aoc.PreparePuzzle(puzzle))
	if err != nil {
		return err
	}
	log.Printf("Answer A: %v", answer)
	return nil
}

func partB(ctx *cli.Context) error {
	answer, err := processB(aoc.PreparePuzzle(puzzle))
	if err != nil {
		return err
	}
	log.Printf("Answer B: %v", answer)
	return nil
}

func processA(puzzle []string) (int, error) {
	sum := 0
	for n := 0; n < len(puzzle); n += 3 {
		if n >= len(puzzle) {
			break
		}
		for ; n < len(puzzle) && len(puzzle[n]) == 0; n++ {
		}
		if n >= len(puzzle) {
			break
		}
		cm := getClawMachine(puzzle[n : n+3])
		sum += cm.tokensToWinPrizeNonBruteForce() // cm.tokensToWinPrize(0, 100)
	}
	return sum, nil
}

func processB(puzzle []string) (int, error) {
	sum := 0
	for n := 0; n < len(puzzle); n += 3 {
		if n >= len(puzzle) {
			break
		}
		for ; n < len(puzzle) && len(puzzle[n]) == 0; n++ {
		}
		if n >= len(puzzle) {
			break
		}
		cm := getClawMachine(puzzle[n : n+3])
		cm.prize = aoc.NewPairInt(cm.prize.X()+10_000_000_000_000, cm.prize.Y()+10_000_000_000_000)
		sum += cm.tokensToWinPrizeNonBruteForce()
	}
	return sum, nil
}

type clawMachine struct {
	a, b, prize aoc.PairInt
}

/*
"Button A: X+94, Y+34",
"Button B: X+22, Y+67",
"Prize: X=8400, Y=5400",
*/
func getClawMachine(puzzle []string) *clawMachine {
	log.Printf("getting claw machine from %v", puzzle)
	cm := &clawMachine{}
	cm.getButton(puzzle[0])
	cm.getButton(puzzle[1])
	cm.getPrize(puzzle[2])
	log.Printf("claw machine: %v", cm)
	return cm
}

func (cm *clawMachine) getButton(line string) {
	re := regexp.MustCompile(`Button (.): X\+([0-9]+), Y\+([0-9]+)`)
	found := re.FindStringSubmatch(line)
	log.Printf("found %v from %q", found, line)
	switch found[1] {
	case "A":
		cm.a = aoc.NewPairInt(aoc.MustAtoi(found[2]), aoc.MustAtoi(found[3]))
	case "B":
		cm.b = aoc.NewPairInt(aoc.MustAtoi(found[2]), aoc.MustAtoi(found[3]))
	default:
		log.Panicf("unknown button %v found", found[1])
	}
}

func (cm *clawMachine) getPrize(line string) {
	re := regexp.MustCompile("Prize: X=([0-9]+), Y=([0-9]+)")
	found := re.FindStringSubmatch(line)
	log.Printf("prize: %v", found)
	cm.prize = aoc.NewPairInt(aoc.MustAtoi(found[1]), aoc.MustAtoi(found[2]))
}

func (cm *clawMachine) tokensToWinPrize(startPushes int, maxPushes int) int {
	// a moves by x1, y1
	// b moves by x2, y2
	// how to get to prize x, y in fewest moves
	// na*x1 + nb*x2 = prize.x
	// na*y1 + nb*y2 = prize.y
	// brute force up to N pushes per button?
	pushes := aoc.NewIntSet()
	for na := startPushes; na < maxPushes; na++ {
		if na*cm.a.X() > cm.prize.X() || na*cm.a.Y() > cm.prize.Y() {
			break
		}
		for nb := startPushes; nb < maxPushes; nb++ {
			xpos := na*cm.a.X() + nb*cm.b.X()
			ypos := na*cm.a.Y() + nb*cm.b.Y()
			if xpos > cm.prize.X() || ypos > cm.prize.Y() {
				break
			}
			if xpos != cm.prize.X() {
				continue
			}
			if ypos != cm.prize.Y() {
				continue
			}
			pushes.Add(3*na + nb)
			break
		}
	}
	if pushes.Len() == 0 {
		return 0
	}
	sp := pushes.AsArray()
	sort.Ints(sp)
	return sp[0]
}

func (cm *clawMachine) tokensToWinPrizeNonBruteForce() int {
	/*
	   binary search but with two variables - how?
	   guess ~starting point - could be miles off if press A zero times, B lots
	   (1) na * a.x + nb * b.x = prize.X
	   (2) na * a.y + nb * b.y = prize.Y
	   (from 1) na * a.x = prize.X - nb * b.x
	           => na = (prize.X - nb*b.x) / a.x
	   (plug into 2)
	   ((prize.X - nb*b.x) / a.x) * a.y + nb * b.y = prize.Y
	   (a.y * (prize.X - nb*b.x) / a.x) + nb * b.y = prize.Y
	   (a.y * (prize.X - nb*b.x)) + a.x*nb*b.y = a.x*prize.Y
	   a.y * prize.X - a.y*nb*b.x + a.x*nb*b.y = a.x*prize.Y
	   a.x*nb*b.y - a.y*nb*b.x = a.x*prize.Y - a.y*prize.X
	   nb * (a.x*b.y - a.y*b.x) = a.x*prize.Y - a.y*prize.X
	   nb = (a.x*prize.Y - a.y*prize.X) / (a.x*b.y - a.y*b.x)

	   (rearrange 1)
	   na * a.x + nb * b.x = prize.X
	   na * a.x = prize.X - nb*b.x
	   na = (prize.X - nb*b.x) / a.x

	   Button A: X+94, Y+34
	   Button B: X+22, Y+67
	   Prize: X=8400, Y=5400

	   nb = (94*5400 - 34*8400) / (94*67 - 34*22) = 40
	   na = (8400 - 40*22) / 94 = 80

	   works!

	   shold not work for

	   Button A: X+26, Y+66
	   Button B: X+67, Y+21
	   Prize: X=12748, Y=1217

	   nb = (26*1217 - 66*12748) / (26*21 - 66*67) = 208.9 (non-int => not a solution)
	*/
	nb := float64((cm.a.X()*cm.prize.Y() - cm.a.Y()*cm.prize.X())) / float64((cm.a.X()*cm.b.Y() - cm.a.Y()*cm.b.X()))
	if math.Ceil(nb) != nb {
		return 0
	}
	na := (float64(cm.prize.X()) - nb*float64(cm.b.X())) / float64(cm.a.X())
	if math.Ceil(na) != na {
		return 0
	}

	return int(3*na + nb)
}
