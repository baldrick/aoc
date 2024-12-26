package day17

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day17A",
		Aliases: []string{"day17a"},
		Usage:   "Day 17 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day17B",
		Aliases: []string{"day17b"},
		Usage:   "Day 17 part B",
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

func processA(puzzle []string) (string, error) {
	a := getRegister(puzzle[0])
	b := getRegister(puzzle[1])
	c := getRegister(puzzle[2])
	program := strings.Split(strings.Split(puzzle[4], " ")[1], ",")
	m := machine{a, b, c}
	log.Printf("Machine before = %v", m)
	output := m.run(program, false)
	log.Printf("Machine after = %v", m)
	return output, nil
}

func processB(puzzle []string) (int, error) {
	b := getRegister(puzzle[1])
	c := getRegister(puzzle[2])
	program := strings.Split(puzzle[4], " ")[1]
	programArray := strings.Split(program, ",")
	a := 0 // 4_043_000_000 // 2_810_000_000 // 25_837_000
	log.Printf("Looking for output %q", program)
	for ; a < 5e9; a++ {
		m := machine{a, b, c}
		if aoc.ModInt(a, 10_000_000) == 0 {
			log.Printf("About to run %v", m)
		}
		// output := m.run(programArray, true)
		output := m.run2(programArray)
		if output == "" {
			continue
		}
		log.Printf("A=%v, output=%q", a, output)
		if output == program {
			return a, nil
		}
		//a++
	}
	return 0, nil
}

func getRegister(line string) int {
	split := strings.Split(line, " ")
	return aoc.MustAtoi(split[2])
}

type machine struct {
	a, b, c int // registers
}

func (m machine) String() string {
	return fmt.Sprintf("A=%v, B=%v, C=%v", m.a, m.b, m.c)
}

func (m *machine) run(program []string, exitIfOutputDiffersFromProgram bool) string {
	//log.Printf("Machine: %v, program: %v", m, program)

	var output []string
	programLength := len(program) - 1
	ip := 0
	ipChecked := 0
	for {
		//log.Printf("ip=%v, machine=%v, output=%v", ip, m, output)
		if ip > programLength {
			return strings.Join(output, ",")
		}
		literal := aoc.MustAtoi(program[ip+1])
		switch program[ip] {
		case "0": // adv, a=int(a/2^combo)
			combo := m.getCombo(program[ip+1])
			//log.Printf("0: A=%v/2^%v", m.a, combo)
			m.a = int(float64(m.a) / math.Pow(2, float64(combo)))
		case "1": // bxl, b=b xor literal
			//log.Printf("1: B=%v ^ %v", m.b, literal)
			m.b = m.b ^ literal
		case "2": // bst, b=combo mod 8
			combo := m.getCombo(program[ip+1])
			//log.Printf("2: B=%v mod 8", combo)
			m.b = aoc.ModInt(combo, 8)
		case "3": // jnz, jump to literal iff a!=0
			if m.a != 0 {
				//log.Printf("3: Jumped from %v to %v", ip, literal)
				ip = int(literal)
				continue
			}
		case "4": // bxc, b = b xor c
			//log.Printf("4: B=%v xor %v", m.b, m.c)
			m.b = m.b ^ m.c
		case "5": // out, output combo mod 8
			combo := m.getCombo(program[ip+1])
			//log.Printf("5: Output %v mod 8", combo)
			output = append(output, fmt.Sprintf("%v", aoc.ModInt(combo, 8)))
			if exitIfOutputDiffersFromProgram {
				if ipChecked > programLength {
					log.Printf("Would check beyond length of program: output=%v, ipChecked=%v", output, ipChecked)
					return ""
				}
				if program[ipChecked] != output[ipChecked] {
					//log.Printf("#%v, %v != %v, this run won't work", ipChecked, program[ipChecked], output[ipChecked])
					return ""
				}
				ipChecked++
			}
		case "6": // bdv, b = int(a/2^combo)
			combo := m.getCombo(program[ip+1])
			//log.Printf("0: B=%v/2^%v", m.a, combo)
			m.b = int(float64(m.a) / math.Pow(2, float64(combo)))
		case "7": // cdv, c = int(a/2^combo)
			combo := m.getCombo(program[ip+1])
			//log.Printf("0: C=%v/2^%v", m.a, combo)
			m.c = int(float64(m.a) / math.Pow(2, float64(combo)))
		}
		//log.Printf("moving to %v", ip+2)
		ip += 2
	}
}

func (m machine) run2(programArray []string) string {
	var output []string
	ipChecked := 0
	startA := m.a
	for {
		m.b = aoc.ModInt(m.a, 8)
		m.b = m.b ^ 1
		m.c = int(float64(m.a) / math.Pow(2, float64(m.b)))
		m.b = m.b ^ 5
		m.b = m.b ^ m.c
		m.a = int(float64(m.a) / 8)
		output = append(output, fmt.Sprintf("%v", aoc.ModInt(m.b, 8)))
		if ipChecked >= len(programArray) {
			log.Panicf("Got %q, already too long vs %q (startA=%v)", output, programArray, startA)
		}
		if programArray[ipChecked] != output[ipChecked] {
			return ""
		}
		ipChecked++
		if m.a == 0 {
			return strings.Join(output, ",")
		}
	}
}

func (m machine) getCombo(operand string) int {
	switch operand {
	case "0", "1", "2", "3":
		// Do nothing, we've already got the operand.
		return aoc.MustAtoi(operand)
	case "4":
		return m.a
	case "5":
		return m.b
	case "6":
		return m.c
	default:
		log.Panicf("Invalid operand %q", operand)
	}
	return -1
}
