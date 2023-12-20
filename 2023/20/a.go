package day20

import (
    _ "embed"
    "fmt"
    "log"
    "regexp"
    "strings"

    "github.com/baldrick/aoc/2023/aoc"
    "github.com/urfave/cli"
)

const (
    year = 2023
    day = 20
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day20A",
        Usage: "Day 20 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day20B",
        Usage: "Day 20 part B",
        Action: partB,
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

type operation int

const (
    none operation = iota
    flipFlop
    conjunction
)

func (o operation) String() string {
    switch o {
    case none: return ""
    case flipFlop: return "%"
    case conjunction: return "&"
    }
    return "?"
}

func decodeOperation(s string) operation {
    switch s {
    case "%": return flipFlop
    case "&": return conjunction
    }
    return none
}

type module struct {
    op operation
    next []string
    memory map[string]bool
}

func (m *module) String() string {
    return fmt.Sprintf("%v->%v %v", m.op, strings.Join(m.next, ","), m.memory)
}

type pulse struct {
    from string
    high bool // false => low
    target string
}

func (p pulse) String() string {
    hl := "low"
    if p.high {
        hl = "high"
    }
    return fmt.Sprintf("%v (%v) -> %v", p.from, hl, p.target)
}

func processA(puzzle []string) (int, error) {
    /*
        "broadcaster -> a, b, c",
        "%a -> b",
        "%b -> c",
        "%c -> inv",
        "&inv -> a",
    */
    moduleMap := make(map[string]*module)
    for _, line := range puzzle {
        addModule(moduleMap, line)
    }
    initConjunctionInputs(moduleMap)
    log.Printf("module map: %v", moduleMap)
    thc, tlc := 0, 0
    for n:=0; n<1000; n++ {
        _, hc, lc := pushButton(moduleMap, false, "")
        thc += hc
        tlc += lc
    }
    log.Printf("hi %v * lo %v = %v", thc, tlc, thc * tlc)
    return thc * tlc, nil
}

func addModule(mm map[string]*module, s string) {
    re := regexp.MustCompile(`([&%]?)([a-z]*) -> (.*)`)
    matches := re.FindStringSubmatch(s)
    mm[matches[2]] = &module{
        op: decodeOperation(matches[1]),
        next: strings.Split(matches[3], ", "),
        memory: make(map[string]bool),
    }
}

func initConjunctionInputs(mm map[string]*module) {
    for input, m := range mm {
        for _, target := range m.next {
            if target == "output" {
                continue
            }
            tm, ok := mm[target]
            if !ok {
                log.Printf("WEIRD - target module %q not found", target)
                continue
            }
            if tm.op == conjunction {
                tm.memory[input] = false
            }    
        }
    }
}

func pushButton(mm map[string]*module, partB bool, partBTarget string) (bool, int, int) {
    var pulses []pulse
    pulses = append(pulses, pulse{from: "button", high: false, target: "broadcaster"})
    return handlePulses(mm, pulses, partB, partBTarget)
}

func handlePulses(mm map[string]*module, pulses []pulse, partB bool, partBTarget string) (bool, int, int) {
    hc, lc := 0, 0
    for ;len(pulses)>0; {
        p := pulses[0]
        if p.high {
            hc++
        } else {
            lc++
        }
        pulses = pulses[1:]
        if p.target == "output" {
            log.Printf("output: %v", p.high)
            continue
        }
        if partB && p.target == partBTarget && !p.high {
            return true, -1, -1
        }
        m, ok := mm[p.target]
        if !ok {
            continue
        }
        switch m.op {
        case flipFlop:
            if p.high {
                //log.Printf("%v -%v -> nowhere", p.from, p.high)
            } else {
                s, ok := m.memory["state"]
                if !ok {
                    s = false
                }
                s = !s
                m.memory["state"] = s
                //log.Printf("%v -%v -> %v (%v -> %v)", p.from, p.high, p.target, s, m.next)
                pulses = append(pulses, sendPulse(p.target, s, m.next)...)
            }
        case conjunction:
            m.memory[p.from] = p.high
            send := false
            for _, state := range m.memory {
                if !state {
                    send = true
                    break
                }
            }
            pulses = append(pulses, sendPulse(p.target, send, m.next)...)
            //log.Printf("%v -%v -> %v (%v -> %v)", p.from, p.high, p.target, send, m.next)
        case none:
            pulses = append(pulses, sendPulse(p.target, p.high, m.next)...)
            //log.Printf("%v -%v -> %v (%v -> %v)", p.from, p.high, p.target, p.high, m.next)
        }
    }
    return false, hc, lc
}

func sendPulse(from string, s bool, next []string) []pulse {
    //log.Printf("sending %v to %v", s, next)
    var pulses []pulse
    for _, n := range next {
        pulses = append(pulses, pulse{from: from, high: s, target: n})
    }
    return pulses
}

func processB(puzzle []string) (int, error) {
    moduleMap := getModuleMap(puzzle)
    rxInput := findInputs(moduleMap, "rx")
    if len(rxInput) != 1 {
        panic(fmt.Sprintf("Expecting 1 input for rx, got %v: %v", len(rxInput), rxInput))
    }
    rxInputs := findInputs(moduleMap, rxInput[0])
    var presses []int
    for _, input := range rxInputs {
        moduleMap := getModuleMap(puzzle)
        presses = append(presses, findPressesUntilHigh(moduleMap, input))
    }
    log.Printf("press count: %v", presses)
    return aoc.LCM(presses...), nil
}

func getModuleMap(puzzle []string) map[string]*module {
    moduleMap := make(map[string]*module)
    for _, line := range puzzle {
        addModule(moduleMap, line)
    }
    initConjunctionInputs(moduleMap)
    return moduleMap
}

func findInputs(mm map[string]*module, s string) []string {
    var inputs []string
    for input, m := range mm {
        for _, target := range m.next {
            if target == s {
                inputs = append(inputs, input)
                break
            }
        }
    }
    log.Printf("inputs of %q are: %v", s, inputs)
    return inputs
}

func findPressesUntilHigh(mm map[string]*module, moduleName string) int {
    presses := 0
    finished := false
    for ; !finished ; presses++ {
        finished, _, _ = pushButton(mm, true, moduleName)
        if aoc.ModInt(presses, 1000) == 0 {
            log.Printf("Pressed the button %v times", presses)
        }
    }
    return presses
}