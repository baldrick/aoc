# What?

adventofcode.com is an advent calendar of puzzles best solved with a computer programme.

Brute forcing puzzles is often an option but there are generally more elegant solutions that are _much_ faster.

# Why?

It's a bit of fun.  Unfortunately also frustration at times.

Really it's an excuse to learn a new language.  That hasn't often been the case for me - I've tended to stick with javascript for the most part although 2020 did see some usage of golang.  I also don't tend to have a lot of time to do these puzzles ... the first couple of years were generally attacked whilst with the in-laws over Christmas and catching up dozens of puzzles in 3-4 days hasn't ever resulted in me solving them all.

This code is most often knocked out in the evening after work and once I've got a working solution I don't bother to refine it.  So don't expect clean code ;-)  And I sometimes just fiddle with part A to get part B done which means code needs unfiddling to re-run part A...

Finally, as I progress through the puzzles I'll change my approach (e.g. use proper tests instead of simply a different input file) but generally only apply it from that point onwards.  Re-running older solutions may therefore require work.

# 2021

I used Python for this year.  The language has been around for years but the v2/v3 debacle put me off learning it.  Until now.  Make sure `aoc/2021` is included in your `PYTHONPATH` for these solutions to work.  Python >=3.10 is required.

# 2022

I gave [Julia](https://julialang.org/) a crack this year for no real reason other than I read about it a week or two beforehand...  The 1-based indexing caused me grief a few times just because I'm used to 0-based.  I find the documentation awkward to use too but overall I can see it'd be a nice language to use with more familiarity.  Having a REPL is a big plus.

# 2023

Back to golang but with the addition of Bazel because I wanted to learn about it (and how it might differ from Google's internal version, blaze).
