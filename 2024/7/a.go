package day7

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day7A",
		Aliases: []string{"day7a"},
		Usage:   "Day 7 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day7B",
		Aliases: []string{"day7b"},
		Usage:   "Day 7 part B",
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
	var equations []equation
	max := 0
	for _, line := range puzzle {
		e := newEquation(line)
		equations = append(equations, e)
		max = aoc.MaxInt(max, len(e.terms))
	}
	ops := []operator{{"*", multiply}, {"+", add}}
	permutations := calculatePermutations(ops, max)
	return solveAll(equations, permutations), nil
}

func processB(puzzle []string) (int, error) {
	var equations []equation
	max := 0
	for _, line := range puzzle {
		e := newEquation(line)
		equations = append(equations, e)
		max = aoc.MaxInt(max, len(e.terms))
	}
	ops := []operator{{"*", multiply}, {"+", add}, {"||", concatenate}}
	permutations := calculatePermutations(ops, max)
	return solveAll(equations, permutations), nil
}

func calculatePermutations(ops []operator, max int) [][]operator {
	log.Printf("creating operator permutations for %v terms", max)
	var permutations [][]operator
	permutationCount := int(math.Pow(float64(len(ops)), float64(max-1)))
	for permutation := 0; permutation < permutationCount; permutation++ {
		var opPermutation []operator
		b3 := big.NewInt(int64(permutation)).Text(len(ops))
		permutationBase3 := Reverse(strings.Repeat("0", 50-len(b3)) + b3)
		log.Printf("calculating permutation #%v/%v, base3=%v", permutation, permutationCount, permutationBase3)
		for p := 0; p < max-1; p++ {
			i := aoc.MustAtoi(string(permutationBase3[p]))
			//log.Printf("operator #%v for permutation #%v = %v", p, permutation, ops[i])
			opPermutation = append(opPermutation, ops[i])
		}
		permutations = append(permutations, opPermutation)
	}
	log.Printf("permutations: %v", permutations)
	return permutations
}

func solveAll(equations []equation, permutations [][]operator) int {
	sum := 0
	for _, e := range equations {
		log.Printf("--------------------- solving for %v", e)
		if e.solve(permutations) {
			sum += e.answer
			log.Printf("%v (match)", e)
		} else {
			log.Printf("%v (fail)", e)
		}

	}

	return sum
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type equation struct {
	answer    int
	terms     []int
	operators [][]operator
}

func newEquation(s string) equation {
	answerAndTerms := strings.Split(s, ":")
	answer := aoc.MustAtoi(answerAndTerms[0])
	sterms := strings.Split(strings.Trim(answerAndTerms[1], " "), " ")
	terms := aoc.MustAtoiAll(sterms)
	return equation{answer, terms, nil}
}

func (e equation) String() string {
	return fmt.Sprintf("%v = %v", e.terms, e.answer)
}

func (e equation) solve(permutations [][]operator) bool {
	for i := 0; i < len(permutations); i++ {
		answer := 0
		for n, t := range e.terms {
			if n == 0 {
				answer = t
				continue
			}
			//log.Printf("about to eval permutation #%v: %v", i, permutations[i])
			na := permutations[i][n-1].eval(answer, t)
			//log.Printf("%v %v %v = %v", answer, permutations[i][n-1].name, t, na)
			answer = na
		}
		if answer == e.answer {
			return true
		}
	}
	return false
}

type operator struct {
	name string
	eval func(a, b int) int
}

func (o operator) String() string {
	return o.name
}

func multiply(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func concatenate(a, b int) int {
	return aoc.MustAtoi(fmt.Sprintf("%v%v", a, b))
}
