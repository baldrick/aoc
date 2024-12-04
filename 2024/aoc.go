package main

import (
    "log"
    "os"

    day2 "github.com/baldrick/aoc/2024/2"
    day3 "github.com/baldrick/aoc/2024/3"
    day4 "github.com/baldrick/aoc/2024/4"
    "github.com/urfave/cli"
)

func main() {
    app := &cli.App{
        Commands: []cli.Command{
            *day2.A, *day2.B,
            *day3.A, *day3.B,
            *day4.A, *day4.B,
        },
    }
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
