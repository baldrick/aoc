package day3

import (
	_ "embed"
	"log"
	"regexp"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle2.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day3A",
		Aliases: []string{"day3a"},
		Usage:   "Day 3 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day3B",
		Aliases: []string{"day3b"},
		Usage:   "Day 3 part B",
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
	for _, line := range puzzle {
		log.Print(line)
		re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		muls := re.FindAllString(line, -1)
		log.Print(muls)
		for _, mul := range muls {
			re = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
			calcs := re.FindAllStringSubmatch(mul, -1)
			for _, calc := range calcs {
				sum += aoc.MustAtoi(calc[1]) * aoc.MustAtoi(calc[2])
			}
			log.Print(calcs)
		}
	}
	return sum, nil
}

func processB(puzzle []string) (int, error) {
	allMuls := ""
	enabled := true
	doNot := regexp.MustCompile(`don't\(\)`)
	do := regexp.MustCompile(`do\(\)`)
	line := puzzle[0]
	for {
		if enabled {
			// already enabled, find where we start being disabled.
			disabledIndex := doNot.FindStringIndex(line)
			if disabledIndex == nil {
				allMuls += line
				break
			}
			log.Printf("%v disabled from %v", line, disabledIndex)
			allMuls += line[:disabledIndex[0]]
			line = line[disabledIndex[1]:]
			enabled = false
		} else {
			// disabled, find where it's next enabled.
			enabledIndex := do.FindStringIndex(line)
			if enabledIndex == nil {
				break
			}
			log.Printf("%v enabled from %v", line, enabledIndex)
			line = line[enabledIndex[1]:]
			enabled = true
		}
	}

	log.Printf("%v -> all muls: %v", puzzle[0], allMuls)
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	calcs := re.FindAllStringSubmatch(allMuls, -1)
	sum := 0
	for _, calc := range calcs {
		sum += aoc.MustAtoi(calc[1]) * aoc.MustAtoi(calc[2])
	}
	log.Print(calcs)
	return sum, nil
}

// func processB(puzzle []string) (int, error) {
// 	// don't()
// 	// do()
// 	log.Print("=================== B ==================")
// 	sum := 0
// 	var allMuls []string
// 	for _, line := range puzzle {
// 		line = fmt.Sprintf(".%v.", line)
// 		log.Print(line)
// 		enabled := true
// 		start := 0
// 		n := 1
// 		for n < 5 {
// 			n++
// 			if enabled {
// 				log.Printf("finding disabled from %v: %v", start, line[start:])
// 				re := regexp.MustCompile(`(.*)don't\(\)(.*)`)
// 				matches := re.FindAllStringSubmatchIndex(line[start:], -1)
// 				if len(matches) == 0 {
// 					log.Printf("No more disabled, add %v", line[start:])
// 					allMuls = append(allMuls, line[start:])
// 					break
// 				}
// 				enabled = false
// 				// [[0 73 0 20 27 73]]
// 				// entire start,end
// 				// start, end of enabled section
// 				// end of don't()
// 				// end of string
// 				startEnabled := start + matches[0][2]
// 				endEnabled := start + matches[0][3]
// 				addMul := line[startEnabled:endEnabled]
// 				log.Printf("%v, add mul (%v-%v) %v", matches, startEnabled, endEnabled, addMul)
// 				allMuls = append(allMuls, addMul)
// 				start += matches[0][4]
// 				log.Printf("%v, disabled, start=%v: %v", matches, start, line[start:])
// 			} else {
// 				log.Printf("finding enabled from %v: %v", start, line[start:])
// 				re := regexp.MustCompile(`.*do\(\)(.*)`)
// 				matches := re.FindAllStringSubmatchIndex(line[start:], -1)
// 				if len(matches) == 0 {
// 					log.Printf("No matches after #%v: %v", start, line[start:])
// 					break
// 				}
// 				enabled = true
// 				// [[0 73 63 73]]
// 				// entire start,end
// 				// start, end of enabled section
// 				start += matches[0][2]
// 				log.Printf("%v, enabled, start=%v: %v", matches, start, line[start:])
// 			}
// 		}
// 		// re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
// 		// muls := re.FindAllString(line, -1)
// 	}

// 	log.Printf("all muls: %v", allMuls)
// 	for _, mul := range allMuls {
// 		re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
// 		calcs := re.FindAllStringSubmatch(mul, -1)
// 		for _, calc := range calcs {
// 			sum += aoc.MustAtoi(calc[1]) * aoc.MustAtoi(calc[2])
// 		}
// 		log.Print(calcs)
// 	}
// 	return sum, nil
// }
