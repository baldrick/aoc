package day15

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
    day = 15
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day15A",
        Usage: "Day 15 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day15B",
        Usage: "Day 15 part B",
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

func processA(puzzle []string) (int, error) {
    grandTotal := 0
    for _, s := range strings.Split(puzzle[0], ",") {
        grandTotal += hash(s)
    }
    log.Printf("%v -> %v", puzzle[0], grandTotal)
    return grandTotal, nil
}

func hash(s string) int {
    total := 0
    for _, c := range s {
        total += int(c)
        total *= 17
        total = aoc.ModInt(total, 256)
    }
    return total
}

type lens struct {
    focalLength int
    label string
}

func (l *lens) String() string {
    return fmt.Sprintf("{%q:%v}", l.label, l.focalLength)
}

type lenses struct {
    this *lens
    next *lenses
    prev *lenses
}

func (l *lenses) String() string {
    s := ""
    checkLens := l
    for slot:=1; checkLens != nil; slot++ {
        s += fmt.Sprintf("%v:%v", slot, checkLens.this.String())
        checkLens = checkLens.next
    }
    return s
}

func (l *lenses) add(label string, focalLength int) *lenses {
    if l == nil {
        return &lenses{this: &lens{label: label, focalLength: focalLength}}
    }
    checkLens := l
    for ;; {
        if checkLens.this.label == label {
            checkLens.this.focalLength = focalLength
            return l
        }
        if checkLens.next == nil {
            // label not found.
            checkLens.next = &lenses{this: &lens{label: label, focalLength: focalLength}, prev: checkLens}
            return l
        }
        checkLens = checkLens.next
    }
}

func (l *lenses) remove(label string) *lenses {
    if l == nil {
        return nil
    }

    checkLens := l
    for ;; {
        if checkLens.this.label == label {
            if checkLens.prev != nil {
                checkLens.prev.next = checkLens.next
            }
            if checkLens.next != nil {
                checkLens.next.prev = checkLens.prev
            }
            if l == checkLens {
                // We removed the first item in the list, return the next item.
                return checkLens.next
            }
            return l
        }
        if checkLens.next == nil {
            // label not found.
            return l
        }
        checkLens = checkLens.next
    }
}

func processB(puzzle []string) (int, error) {
    boxes := make([]*lenses, 256)
    re := regexp.MustCompile(`([a-z]*)([-=])([0-9]*)`)
    for _, s := range strings.Split(puzzle[0], ",") {
        matches := re.FindStringSubmatch(s)
        label := matches[1]
        h := hash(label)
        switch matches[2] {
        case "=":
            boxes[h] = boxes[h].add(label, aoc.MustAtoi(matches[3]))
        case "-":
            boxes[h] = boxes[h].remove(label)
        default:
            return 0, fmt.Errorf("unhandled operation %q from %v", matches[1], s)
        }
    }
    total := 0
    for n, box := range boxes {
        slot := 1
        for checkLens := box; checkLens != nil; checkLens = checkLens.next {
            log.Printf("%v: (box) %v * (slot) %v * (fl) %v", checkLens.this.label, n+1, slot, checkLens.this.focalLength)
            total += (n+1) * slot * checkLens.this.focalLength
            slot++
        }
    }
    return total, nil
}
