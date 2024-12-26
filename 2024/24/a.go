package day24

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"
	"sync"

	"github.com/baldrick/aoc/common/aoc"
	"github.com/urfave/cli"
)

var (
	//go:embed puzzle.txt
	puzzle string

	// A is the command to use to run part A for this day.
	A = &cli.Command{
		Name:    "day24A",
		Aliases: []string{"day24a"},
		Usage:   "Day 24 part A",
		Action:  partA,
	}
	B = &cli.Command{
		Name:    "day24B",
		Aliases: []string{"day24b"},
		Usage:   "Day 24 part B",
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
	channels := make(map[string]chan int)
	channelRE := regexp.MustCompile(`([a-z0-9]+) ([A-Z]+) ([a-z0-9]+) -> ([a-z0-9]+)`)
	values := make(map[string]int)
	answerWires := aoc.NewStringSet()
	gettingValues := true
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, line := range puzzle {
		if len(line) == 0 {
			gettingValues = false
			continue
		}
		if gettingValues {
			// x01: 1
			wireValue := strings.Split(line, ":")
			wire := wireValue[0]
			value := aoc.MustAtoi(strings.Trim(wireValue[1], " "))
			values[wire] = value
			//log.Printf("%v = %v", wire, value)
			if wire[0] == 'z' {
				//log.Printf("found answer wire %v", wire)
				answerWires.Add(wire)
			}
			c := make(chan int, 1000)
			channels[wire] = c
		} else {
			// x00 AND y00 -> z00
			foo := channelRE.FindStringSubmatch(line)
			//log.Printf("line %q -> %v", line, foo)
			in1 := findchan(channels, foo[1])
			op := foo[2]
			in2 := findchan(channels, foo[3])
			outName := foo[4]
			out := findchan(channels, outName)
			incInputCount(foo[1])
			incInputCount(foo[3])
			if outName[0] == 'z' {
				//log.Printf("found answer wire %v", outName)
				answerWires.Add(outName)
				wg.Add(1)
				go func() {
					defer wg.Done()
					mu.Lock()
					defer mu.Unlock()
					in1, in2, outName, op := in1, in2, outName, op
					v := gate(in1, in2, op)
					//log.Printf("%v %v %v = %v (%v)", foo[1], op, foo[3], foo[4], v)
					values[outName] = v
				}()
			} else {
				wg.Add(1)
				go func() {
					defer wg.Done()
					in1, in2, out, op := in1, in2, out, op
					v := gate(in1, in2, op)
					log.Printf("%v %v %v = %v (%v)", foo[1], op, foo[3], foo[4], v)
					n, ok := inputCount[outName]
					if !ok {
						log.Panicf("Failed to find input count for %q, assume 1", outName)
					}
					//log.Printf("Outputting %v to %v, %v times", v, outName, n)
					for ; n >= 0; n-- {
						out <- v
					}
				}()
			}
		}
	}
	// Inject values.
	for name, value := range values {
		c, ok := channels[name]
		if !ok {
			log.Panicf("Failed to find channel %q", name)
		}
		n, ok := inputCount[name]
		if !ok {
			log.Printf("Failed to find input count for %q, assume 1", name)
			n = 1
		}
		//log.Printf("Inject %v to %v %v times", value, name, n)
		for ; n >= 0; n-- {
			c <- value
		}
		//log.Printf("Injected %v to %v", value, name)
	}
	//log.Printf("Injection complete")
	wg.Wait()
	sum := 0
	answerArray := answerWires.AsArray()
	slices.Sort(answerArray)
	//log.Printf("answers should be in %v", answerArray)
	for bit, name := range answerArray {
		v, ok := values[name]
		if !ok {
			log.Panicf("Failed to find answer %q", name)
		}
		log.Printf("%v = %v", name, v)
		sum += (v << bit)
	}
	return sum, nil
}

var inputCount = make(map[string]int)

func incInputCount(name string) {
	n, ok := inputCount[name]
	if !ok {
		n = 0
	}
	inputCount[name] = n + 1
}

func findchan(channels map[string]chan int, name string) chan int {
	c, ok := channels[name]
	if !ok {
		c = make(chan int, 1000)
		channels[name] = c
	}
	return c
}

// channel for each input/output
// goroutine for each operation
func gate(in1, in2 chan int, op string) int {
	v1 := <-in1
	v2 := <-in2
	switch op {
	case "AND":
		return v1 & v2
	case "OR":
		return v1 | v2
	case "XOR":
		return v1 ^ v2
	}
	log.Panicf("Unknown gate %v", op)
	return -1 // can't get here, silly compiler
}

func processB(puzzle []string) (int, error) {
	return 0, fmt.Errorf("Not yet implemented")
}

/*
Ripple carry adder, inspect puzzle by hand / shell script to find swapped wires.

x00 XOR y00 -> z00 (sumbits)
x00 AND y00 -> prt (n, = carryOut I think)

y01 XOR x01 -> jnj (sumbits)
x01 AND y01 -> kmf (n)
prt XOR jnj -> z01 (sum)
jnj AND prt -> cdb (m)
kmf OR cdb -> qnf (carryOut = n or m)

y02 XOR x02 -> wsm (sumbits)
x02 AND y02 -> jrf (n)
wsm XOR qnf -> z02 (sumbits xor carryIn = sum)
wsm AND qnf -> shr (sumbits and carryIn = m)
jrf OR shr -> vnm (n or m = carryOut)

x03 XOR y03 -> mgr (sumbits)
y03 AND x03 -> pps (n)
vnm XOR mgr -> z03 (sumbits xor carryIn (vnm) = sum)
vnm AND mgr -> phc (carryIn and sumbits = m)
pps OR phc -> jdk (n or m = carryOut)

y04 XOR x04 -> mtb (sumbits)
y04 AND x04 -> dgw (n)
mtb XOR jdk -> z04 (sumbits xor carryIn = sum)
mtb AND jdk -> fmr (sumbits and carryIn = m)
fmr OR dgw -> jtm (carryOut)

y05 XOR x05 -> rqw (sumbits)
y05 AND x05 -> jfq (n)
rqw XOR jtm -> z05 (sumbits xor carryIn = sum)
jtm AND rqw -> tsw (sumbits and carryIn = m)
jfq OR tsw -> ksh (n or m = carryOut)

x06 XOR y06 -> kcc (sumbits)
x06 AND y06 -> qmr (n)
ksh XOR kcc -> z06 (carryIn xor sumbits = sum)
ksh AND kcc -> fjt (carryIn and sumbits = m)
qmr OR fjt -> drd (n or m = carryOut)

y07 XOR x07 -> ncp (sumbits)
y07 AND x07 -> wqt (n)
ncp XOR drd -> z07 (sumbits xor carryIn = sum)
ncp AND drd -> ktn (sumbits and carryIn = m)
wqt OR ktn -> qft (n or m = carryOut)

x08 XOR y08 -> twj (sumbits)
x08 AND y08 -> bqm (n)
twj XOR qft -> z08 (sumbits xor carryIn = sum)
qft AND twj -> qpn (sumbits and carryIn = m)
bqm OR qpn -> vrr (n or m = carryOut)

x09 XOR y09 -> jtc (sumbits)
x09 AND y09 -> nwn (n)
vrr XOR jtc -> z09 (sumbits xor carryIn = sum)
vrr AND jtc -> hwp (sumbits and carryIn = m)
hwp OR nwn -> bkp (n or m = carryOut)

use shell script to find incorrect bits


x16 XOR y16 -> grr (sumbits)
y16 AND x16 -> bss (n)
kcm XOR grr -> fkb (carryIn xor sumbits, should go to z16)
kcm AND grr -> tnn (carryIn and sumbits = m)
tnn OR bss -> z16 (m or n = carryOut)

=> fkb,z16 swapped

fpv OR smh -> hvv (carryIn from bit 20)

x21 XOR y21 -> nnr (sumbits)
y21 AND x21 -> rqf (n)
rqf XOR hvv -> z21 (want nnr xor carryIn = sum)
rqf AND hvv -> jsd (want nnr and carryIn = m)
nnr OR jsd -> sfw ((want) carry out)

=> nnr,rqf swapped

y31 XOR x31 -> tjk (sumbits)
y31 AND x31 -> pct (n)
qsj AND tjk -> z31
qsj XOR tjk -> rdn (expect m but it's carryIn xor sumbits = sum)
rdn OR pct -> vtb (m or n = carry out)

=> z31,rdn swapped

y37 XOR x37 -> gcg (sumbits)
y37 AND x37 -> z37 (should be n)
nbm AND gcg -> vhj (carryIn and sumbits = m)
gcg XOR nbm -> rrn (sumbits xor carryIn = sum)
vhj OR rrn -> jrg (carry out)

=> z37,rrn swapped

=> all swaps are: fkb,nnr,rdn,rqf,rrn,z16,z31,z37

generic ripple carry adder:

sumbits = in1 xor in2
sum = sumbits xor carryIn
n = a and b
m = sumbits and carryIn
carryOut = n or m

*/
