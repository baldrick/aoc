package main

import (
    "log"
    "os"

    day1 "github.com/baldrick/aoc/2024/1"
    day2 "github.com/baldrick/aoc/2024/2"
    day3 "github.com/baldrick/aoc/2024/3"
    day4 "github.com/baldrick/aoc/2024/4"
    day5 "github.com/baldrick/aoc/2024/5"
    day6 "github.com/baldrick/aoc/2024/6"
    day7 "github.com/baldrick/aoc/2024/7"
    day8 "github.com/baldrick/aoc/2024/8"
    day9 "github.com/baldrick/aoc/2024/9"
    day10 "github.com/baldrick/aoc/2024/10"
    day11 "github.com/baldrick/aoc/2024/11"
    day12 "github.com/baldrick/aoc/2024/12"
    day13 "github.com/baldrick/aoc/2024/13"
    day14 "github.com/baldrick/aoc/2024/14"
    day15 "github.com/baldrick/aoc/2024/15"
    day16 "github.com/baldrick/aoc/2024/16"
    day17 "github.com/baldrick/aoc/2024/17"
    day18 "github.com/baldrick/aoc/2024/18"
    day19 "github.com/baldrick/aoc/2024/19"
    day20 "github.com/baldrick/aoc/2024/20"
    day21 "github.com/baldrick/aoc/2024/21"
    day22 "github.com/baldrick/aoc/2024/22"
    day23 "github.com/baldrick/aoc/2024/23"
    day24 "github.com/baldrick/aoc/2024/24"
    day25 "github.com/baldrick/aoc/2024/25"
    "github.com/urfave/cli"
)

func main() {
    app := &cli.App{
        Commands: []cli.Command{
            *day1.A, *day1.B,
            *day2.A, *day2.B,
            *day3.A, *day3.B,
            *day4.A, *day4.B,
            *day5.A, *day5.B,
            *day6.A, *day6.B,
            *day7.A, *day7.B,
            *day8.A, *day8.B,
            *day9.A, *day9.B,
            *day10.A, *day10.B,
            *day11.A, *day11.B,
            *day12.A, *day12.B,
            *day13.A, *day13.B,
            *day14.A, *day14.B,
            *day15.A, *day15.B,
            *day16.A, *day16.B,
            *day17.A, *day17.B,
            *day18.A, *day18.B,
            *day19.A, *day19.B,
            *day20.A, *day20.B,
            *day21.A, *day21.B,
            *day22.A, *day22.B,
            *day23.A, *day23.B,
            *day24.A, *day24.B,
            *day25.A, *day25.B,
        },
    }
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
