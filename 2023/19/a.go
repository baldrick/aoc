package day19

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
    day = 19
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day19A",
        Usage: "Day 19 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day19B",
        Usage: "Day 19 part B",
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
    noop operation = iota
    lt
    gt
)

func (op operation) String() string {
    switch op {
    case lt: return "<"
    case gt: return ">"
    case noop: return ""
    }
    return "unhandled"
}

type workflow struct {
    v string
    op operation
    n int
    target string
}

func (w workflow) String() string {
    if len(w.v) > 0 {
        return fmt.Sprintf("%v %v %v -> %v,", w.v, w.op, w.n, w.target)
    }
    return w.target
}

type part struct {
    name string
    value int
}

func processA(puzzle []string) (int, error) {
    // bh{x<2396:sv,x<2996:ccd,a>2397:cll,hcv}
    workflows := make(map[string][]workflow)
    totalRating := 0
    processingWorkflows := true
    for _, line := range puzzle {
        if processingWorkflows {
            if len(line) == 0 {
                processingWorkflows = false
                continue
            }
            addWorkflow(workflows, line)
        } else {
            totalRating += processPart(workflows, line)
        }
    }
    return totalRating, nil
}

func addWorkflow(workflows map[string][]workflow, line string) {
    re := regexp.MustCompile(`([a-zA-Z]*){(.*)}`)
    matches := re.FindStringSubmatch(line)
    workflows[matches[1]] = decodeWorkflows(matches[2])
}

func decodeWorkflows(line string) []workflow {
    //"px{a<2006:qkq,m>2090:A,rfg}",
    var workflows []workflow
    sections := strings.Split(line, ",")
    re := regexp.MustCompile(`([a-zA-Z]*)([<>])([0-9]*):([a-zA-Z]*)`)
    for _, section := range sections {
        //a<2006:qkq
        matches := re.FindStringSubmatch(section)
        log.Printf("Adding workflow %v -> %v", section, matches)
        if len(matches) == 0 {
            workflows = append(workflows, workflow{target:section})
            break
        }
        workflows = append(workflows, workflow{v: matches[1], op: decodeOperation(matches[2]), n: aoc.MustAtoi(matches[3]), target: matches[4]})
    }
    return workflows
}

func decodeOperation(op string) operation {
    switch op {
    case "<": return lt
    case ">": return gt
    }
    panic(fmt.Sprintf("Operation %q not handled", op))
}

func processPart(workflows map[string][]workflow, line string) int {
    //{x=555,m=1815,a=230,s=2491}
    log.Printf("workflows:\n%v", workflows)
    re := regexp.MustCompile(`{(.*)}`)
    matches := re.FindStringSubmatch(line)
    parts := strings.Split(matches[1], ",")
    partValues := make(map[string]int)
    for _, part := range parts {
        re := regexp.MustCompile(`([a-zA-Z]*)=([0-9]*)`)
        matches := re.FindStringSubmatch(part)
        partValues[matches[1]] = aoc.MustAtoi(matches[2])
    }
    return runWorkflow(workflows, partValues)
}

func runWorkflow(workflows map[string][]workflow, parts map[string]int) int {
    w := "in"
    log.Printf("processing %v", parts)
    for ;; {
        if w == "R" {
            log.Printf("%v rejected", parts)
            return 0
        }
        if w == "A" {
            total := 0
            for _, partValue := range parts {
                total += partValue
            }
            log.Printf("%v accepted (%v)", parts, total)
            return total
        }
        pw, ok := workflows[w]
        if !ok {
            panic(fmt.Sprintf("Workflow %q not found", w))
        }
        log.Printf("Processing %v for %v", pw, parts)
        var nextWorkflow bool
        for _, rule := range pw {
            if len(rule.v) == 0 {
                w = rule.target
                log.Printf("Empty rule %q for workflow %q, target now %q", rule, pw, w)
                break
            }
            partValue, ok := parts[rule.v]
            if !ok {
                panic(fmt.Sprintf("Failed to find value for part %q for workflow rule %q", rule.v, rule))
            }
            w, nextWorkflow = getNextTarget(rule, partValue)
            if nextWorkflow {
                break
            }
        }
    }
}

func getNextTarget(rule workflow, partValue int) (string, bool) {
    switch rule.op {
    case lt:
        if partValue < rule.n {
            log.Printf("%v < %v -> %v", partValue, rule.n, rule.target)
            return rule.target, true
        }
    case gt:
        if partValue > rule.n {
            log.Printf("%v > %v -> %v", partValue, rule.n, rule.target)
            return rule.target, true
        }
    case noop:
        log.Printf("noop -> %v", rule.target)
        return rule.target, true
    default:
        panic(fmt.Sprintf("Operation %v not handled", rule.op))
    }
    return "", false
}

func processB(puzzle []string) (int, error) {
    return 0, fmt.Errorf("Not yet implemented")
}
