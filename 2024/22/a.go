package day22

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day22A",
		Aliases: []string{"day22a"},
		Usage:   "Day 22 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day22B",
		Aliases: []string{"day22b"},
		Usage:   "Day 22 part B",
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
	var secrets []*secret
	for _, line := range puzzle {
		if len(line) == 0 {
			continue
		}
		secrets = append(secrets, newSecret(line))
	}
	for _, s := range secrets {
		start := s.n
		for i := 0; i < 2000; i++ {
			s.quickNext()
		}
		log.Printf("s=%v -> %v", start, s.n)
	}
	sum := 0
	for _, s := range secrets {
		sum += s.n
	}
	return sum, nil
}

func processB(puzzle []string) (int, error) {
	return 0, fmt.Errorf("Not yet implemented")
}

type secret struct {
	n             int
	prices, diffs []int
}

func newSecret(line string) *secret {
	s := &secret{n: aoc.MustAtoi(line)}
	s.addPrice()
	return s
}

func (s *secret) next() {
	/*
	   Calculate the result of multiplying the secret number by 64.
	   Then, mix this result into the secret number.
	   Finally, prune the secret number.

	   Calculate the result of dividing the secret number by 32.
	   Round the result down to the nearest integer.
	   Then, mix this result into the secret number.
	   Finally, prune the secret number.

	   Calculate the result of multiplying the secret number by 2048.
	   Then, mix this result into the secret number.
	   Finally, prune the secret number.


	   To mix a value into the secret number, calculate the bitwise XOR of the given value
	   and the secret number. Then, the secret number becomes the result of that operation.
	   (If the secret number is 42 and you were to mix 15 into the secret number, the secret
	   number would become 37.)

	   To prune the secret number, calculate the value of the secret number modulo 16777216.
	   Then, the secret number becomes the result of that operation. (If the secret number
	   is 100000000 and you were to prune the secret number, the secret number would become 16113920.)

	   r=n*64
	   n=n xor r
	   n=n xor 2^24

	   r=int(n/32)
	   n=n xor r
	   n=n xor 2^24

	   r=n*2048
	   n=n xor r
	   n=n xor 2^24

	*/
	//log.Printf("starting with s.n=%v", s.n)
	r := s.n * 64
	s.n = s.n ^ r
	s.n = aoc.ModInt(s.n, 16777216)
	//log.Printf("stage 1: %v", s.n)

	r = int(s.n / 32)
	s.n = s.n ^ r
	s.n = aoc.ModInt(s.n, 16777216)
	//log.Printf("stage 2: %v", s.n)

	r = s.n * 2048
	s.n = s.n ^ r
	s.n = aoc.ModInt(s.n, 16777216)
	//log.Printf("stage 3: %v", s.n)

	s.addPrice()
	s.addDiff()
}

// I thought bitwise operations may be quicker but having run it
// a few times I'm not sure.  In any case it doesn't matter because
// part two is quite different...
func (s *secret) quickNext() {
	//log.Printf("%16b", notTwoTo24)
	//log.Printf("starting with s.n=%v", s.n)
	r := s.n << 6
	//log.Printf("1a: r=%v", r)
	s.n = s.n ^ r
	//log.Printf("1b: s.n=%v", s.n)
	s.n = s.n & ls24bits
	//log.Printf("stage 1: %v", s.n)

	r = s.n >> 5
	s.n = s.n ^ r
	s.n = s.n & ls24bits
	//log.Printf("stage 2: %v", s.n)

	r = s.n << 11
	s.n = s.n ^ r
	s.n = s.n & ls24bits
	//log.Printf("stage 3: %v", s.n)

	s.addPrice()
	s.addDiff()
}

var ls24bits = 0b_11111111_11111111_11111111

func (s *secret) addPrice() {
	secret := fmt.Sprintf("%v", s.n)
	units := aoc.MustAtoi(string(secret[len(secret)-1]))
	s.prices = append(s.prices, units)
}

func (s *secret) addDiff() {
	if len(s.prices) < 2 {
		return
	}
	s.diffs = append(s.diffs, s.prices[len(s.prices)-1]-s.prices[len(s.prices)-2])
}

func (s *secret) String() string {
	diff := ""
	if len(s.diffs) > 0 {
		diff = fmt.Sprintf("%v", s.diffs[len(s.diffs)-1])
	}
	return fmt.Sprintf("%v: %v (%v)", s.n, s.prices[len(s.prices)-1], diff)
}
