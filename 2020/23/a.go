package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

type Cup struct {
    value int
    i int
    taken bool
}

func (c Cup) String() string {
    if c.taken {
        return fmt.Sprintf("_%v_", c.value)
    }
    return fmt.Sprintf("%v", c.value)
}

type Circle struct {
    cups []*Cup
    firstTaken int
    current int
}

func (c *Circle) append(cup *Cup) {
    cup.i = len(c.cups)
    c.cups = append(c.cups, cup)
}

func (c *Circle) max() int {
    var max int
    for _, cup := range c.cups {
        if cup.value > max && !cup.taken {
            max = cup.value
        }
    }
    return max
}

func (c Circle) String() string {
    var s string
    for _, cup := range c.cups {
        if c.current == cup.value {
            s = fmt.Sprintf("%s (%s)", s, cup)
        } else  {
            s = fmt.Sprintf("%s %s", s, cup)
        }
    }
    return fmt.Sprintf("Cups: %s", s)
}

func (c *Circle) indexOf(cupValue int) int {
    for i, cup := range c.cups {
        if cupValue == cup.value {
            return i
        }
    }
    log.Fatalf("Cup %v not found in %s", cupValue, c)
    return -1
}

func (c *Circle) index(i int) int {
    return i % len(c.cups)
}

func (c *Circle) take(start int, n int, taken bool) {
    var s string
    for i := 0;  i < n;  i++ {
        ci := c.index(i+start)
        c.cups[ci].taken = taken
        s = fmt.Sprintf("%s %s,", s, c.cups[ci])
    }
    c.firstTaken = start
    if taken {
        log.Printf("pick up: %s", s)
    }
}

// The crab selects a destination cup: the cup with a label equal to the current cup's label minus one.
// If this would select one of the cups that was just picked up, the crab will keep subtracting one until it finds a
// cup that wasn't just picked up. If at any point in this process the value goes below the lowest value on any cup's
// label, it wraps around to the highest value on any cup's label instead.
func (c *Circle) destination(current int) *Cup {
    var target int = current - 1
    var targetIndex int = -1
    for ; target > 0; target-- {
        targetIndex = c.indexOf(target)
        log.Printf("Cup %v at %v is %v", target, targetIndex, c.cups[targetIndex])
        if c.cups[targetIndex].taken {
            log.Printf("Cup %v is taken, keep looking", target)
            targetIndex = -1
        } else {
            log.Printf("Found cup %v at %v", target, targetIndex)
            break
        }
    }
    if targetIndex < 0 {
        log.Printf("Looking for max cup %v", c.max())
        targetIndex = c.indexOf(c.max())
    }
    return c.cups[targetIndex]
}

func (c* Circle) placeTakenCups(dest *Cup) {
    log.Printf("destination: %s", dest)
    var newCups []*Cup

    // Copy everything up to and including dest but not taken cups
    for i := 0;  i <= dest.i;  i++ {
        if !c.cups[i].taken {
            newCups = append(newCups, c.cups[i])
        }
    }
    // Insert taken cups
    for i := 0;  i < 3;  i++ {
        var ci int = c.index(c.firstTaken + i)
        if c.cups[ci].taken {
            newCups = append(newCups, c.cups[ci])
        }
    }
    // Insert remaining cups
    for i := dest.i + 1;  i < len(c.cups);  i++ {
        if !c.cups[i].taken {
            newCups = append(newCups, c.cups[i])
        }
    }
    c.take(0, len(c.cups), false)
    c.cups = newCups
}

func (c *Circle) move(m int) {
    log.Printf("-- Move %d --\n%s", m, c)
    currentIndex := c.indexOf(c.current)
    c.take(currentIndex+1, 3, true)
    var dest = c.destination(c.current)
    c.placeTakenCups(dest)
    currentIndex = c.indexOf(c.current) // reindex as circle has shuffled
    log.Printf("cups afer placing taken ones: %s", c)
    log.Printf("current=%v, next current=%v", c.cups[currentIndex], c.cups[c.index(currentIndex+1)])
    c.current = c.cups[c.index(currentIndex+1)].value
}

func (c *Circle) selectCurrent(current int) {
    c.current = current
}

func doit(filename string) {
    var input = readFile(filename)
    for _, line := range input {
        var c Circle
        for i := 0;  i < len(line);  i++ {
            c.append(&Cup{ value: int(line[i] - '0') })
        }
        c.selectCurrent(3)
        for m := 1;  m < 11;  m++ {
            c.move(m)
        }
    }
}

func main() {
    if testMode() {
        doit("test.txt")
    } else {
        doit("input.txt")
    }
}

func testMode() bool {
    if len(os.Args) > 1 {
        if os.Args[1] == "-t" {
            return true
        }
    }
    return false
}

func readFile(f string) []string {
    file, err := os.Open(f)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var input []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return input
}