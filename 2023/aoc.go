package main

import (
    "log"
    "os"

    day1 "github.com/baldrick/aoc/2023/1"
    day2 "github.com/baldrick/aoc/2023/2"
    day3 "github.com/baldrick/aoc/2023/3"
    "github.com/urfave/cli"
)

func main() {
    app := &cli.App{
        Commands: []cli.Command{
            *day1.A, *day1.B,
            *day2.A, *day2.B,
            *day3.A, *day3.B,
        },
    }
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
