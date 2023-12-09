package main

import (
    "log"
    "os"

    day1 "github.com/baldrick/aoc/2023/1"
    day2 "github.com/baldrick/aoc/2023/2"
    day3 "github.com/baldrick/aoc/2023/3"
    day4 "github.com/baldrick/aoc/2023/4"
    day5 "github.com/baldrick/aoc/2023/5"
    day6 "github.com/baldrick/aoc/2023/6"
    day7 "github.com/baldrick/aoc/2023/7"
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
        },
    }
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
