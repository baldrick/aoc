package day12

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
    day = 12
)

var (
    //go:embed puzzle.txt
    puzzle string

    // A is the command to use to run part A for this day.
    A = &cli.Command{
        Name:  "day12A",
        Usage: "Day 12 part A",
        Action: partA,
    }
    B = &cli.Command{
        Name:  "day12B",
        Usage: "Day 12 part B",
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
    total := 0
    for _, line := range puzzle {
        springs, damageCount := decode(line)
        total += arrangements(springs, damageCount)
    }
    return total, nil
}

func processB(puzzle []string) (int, error) {
    return 0, fmt.Errorf("Not yet implemented")
}

func decode(line string) (string, []int) {
    re := regexp.MustCompile(`([\?\.#]*) ([0-9,]*)`)
    matches := re.FindStringSubmatch(line)
    return matches[1], aoc.MustAtoiAll(strings.Split(matches[2], ","))
}

func arrangements(springs string, damageCount []int) int {
    // ??.## 1,2
    // #?#?#? 1
    groups := getGroups(springs)
    log.Printf("%v -> groups: %v", springs, groups)

    return 0
}

type springStatus int

func (ss springStatus) String() string {
    switch ss {
    case unknown: return "unknown"
    case operational: return "operational"
    case broken: return "broken"
    }
    return "?*!"
}

const (
    unknown springStatus = iota
    operational
    broken
)

type groupInfo struct {
    size int
    status springStatus
}

func (gi groupInfo) String() string {
    return fmt.Sprintf("%v %v", gi.size, gi.status)
}

func newGroupInfo(count int, s byte) groupInfo {
    switch s {
    case '.':
        return groupInfo{size: count, status: operational}
    case '#':
        return groupInfo{size: count, status: broken}
    case '?':
        return groupInfo{size: count, status: unknown}
    }
    log.Fatalf("Unknown spring type %q", s)
    return groupInfo{} // shouldn't get here
}

func getGroups(s string) []groupInfo {
    var gi []groupInfo
    count := 1
    last := s[0]
    for i:=1; i<len(s); i++ {
        if s[i] != last {
            gi = append(gi, newGroupInfo(count, last))
            last = s[i]
            count = 1
            continue
        }
        count++
    }
    return append(gi, newGroupInfo(count, last))
}
