package main

import (
    "log"
    "os"

    day1 "github.com/baldrick/aoc/2023/1"
    day9 "github.com/baldrick/aoc/2023/9"
    "github.com/urfave/cli"
)

func main() {
    app := &cli.App{
        Commands: []cli.Command{
            *day1.A, *day1.B,
            *day9.A, *day9.B,
        },
    }
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
