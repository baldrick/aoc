package main

import (
    "bufio"
    "log"
    "os"
    "strconv"
    "strings"
)

type mask struct {
    andMask int64
    orMask int64
}

var themask mask
var mem map[int]int64 = make(map[int]int64)

func processMask(line string) {
    if line[:4] == "mask" {
        m := line[7:]
        var bit uint
        themask.andMask = 0
        themask.orMask = 0
        for bit = 0;  bit < 36;  bit++ {
            switch (m[bit:bit+1]) {
            case "X":
                themask.andMask |= 1 << (35-bit)
            case "1":
                themask.orMask |= 1 << (35-bit)
            case "0":
            default:
                log.Fatal("unhandled bit in mask at pos %d: %s", bit, line[bit:bit+1])
            }
        }
        log.Printf("andMask = %s", strconv.FormatInt(themask.andMask, 2))
        log.Printf("orMask =  %s", strconv.FormatInt(themask.orMask, 2))
    }
}

func processMemset(line string) {
    if line[:3] == "mem" {
        var eb = strings.Index(line, "]")
        var addr, err = strconv.Atoi(line[4:eb])
        if err != nil {
            log.Fatal("failed to parse memory position %s", line[4:eb])
        }
        var eq = strings.Index(line, "=")
        var val int64
        val, err = strconv.ParseInt(line[eq+2:], 10, 64)
        val &= themask.andMask
        val |= themask.orMask
        mem[addr] = val
    }
}

func doit(filename string) {
    var input = readFile(filename)
    for _, line := range input {
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