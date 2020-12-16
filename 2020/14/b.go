package main

import (
    "bufio"
    "log"
    "os"
    "strconv"
    "strings"
)

type mask struct {
    floating int64
    zeroMask int64
    oneMask  int64
}

var themask mask
var mem map[int64]int64 = make(map[int64]int64)

func processMask(line string) {
    if line[:4] == "mask" {
        m := line[7:]
        log.Printf("%s (%d)", m, len(m))
        var bit uint
        themask.floating = 0
        themask.zeroMask = 1<<36-1 // 35 bits set
        log.Printf("zeroMask = %s", strconv.FormatInt(themask.zeroMask, 2))
        themask.oneMask = 0
        for bit = 0;  bit < 36;  bit++ {
            switch (m[bit:bit+1]) {
            case "X":
                themask.floating |= 1 << (35-bit)
                themask.zeroMask -= (1 << (35-bit))
            case "1":
                themask.oneMask |= 1 << (35-bit)
            case "0":
                // leave unchanged
            default:
                log.Fatal("unhandled bit in mask at pos %d: %s", bit, line[bit:bit+1])
            }
        }
        log.Printf("zeroMask = %s", strconv.FormatInt(themask.zeroMask, 2))
        log.Printf("oneMask =  %s", strconv.FormatInt(themask.oneMask, 2))
        log.Printf("floating= %s", strconv.FormatInt(themask.floating, 2))
        calculateFloatingCombinations()
    }
}

var combinations []int64

func calculateFloatingCombinations() {
    log.Printf("calculating combinations for floating= %s", strconv.FormatInt(themask.floating, 2))
    combinations = nil
    combinations = append(combinations, 0)
    var bit uint
    for bit = 0;  bit < 36;  bit++ {
        if (themask.floating & (1 << bit)) != 0 {
            addCombination(bit)
        }
    }
}

func addCombination(bit uint) {
    cl := len(combinations)
    for ci := 0;  ci < cl;  ci++ {
        log.Printf("adding combination %d/%d", ci+1, cl)
        combinations = append(combinations, combinations[ci] | 1<<bit)
    }
}

func processMemset(line string) {
    if line[:3] == "mem" {
        var eb = strings.Index(line, "]")
        var addr, err = strconv.ParseInt(line[4:eb], 10, 64)
        if err != nil {
            log.Fatal("failed to parse memory position %s", line[4:eb])
        }
        var eq = strings.Index(line, "=")
        var val int64
        val, err = strconv.ParseInt(line[eq+2:], 10, 64)
        log.Printf("pre-masking mem[%s/%d]=%d", strconv.FormatInt(addr,2), addr, val)
        addr &= themask.zeroMask
        addr |= themask.oneMask
        log.Printf("mem[%s/%d]=%d + floating mask", strconv.FormatInt(addr,2), addr, val)
        memset(addr, val)
    }
}

func memset(addr int64, val int64) {
    for _, c := range combinations {
        log.Printf("mem[%s/%d] = %d for %s", strconv.FormatInt(addr|c, 2), addr|c, val, strconv.FormatInt(c, 2))
        mem[addr | c] = val
    }
}

func doit(filename string) {
    var input = readFile(filename)
    for _, line := range input {
        log.Printf("processing %s", line)
        processMask(line)
        processMemset(line)
    }
    var sum int64
    for _, v := range mem {
        sum += v
    }
    log.Printf("sum is %v", sum)
}

func main() {
    if testMode() {
        doit("testb.txt")
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