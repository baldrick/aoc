cycle = 1
r = 1
s = 0

function main()
    f = open(ARGS[1], "r")
    while !eof(f)
        line = readline(f)
        if line[1:4] == "noop"
            println("#$cycle: noop, r=$r")
            tick()
            continue
        end
        v = parse(Int, line[5:end])
        println("#$cycle: addv, r=$r, v=$v")
        tick()
        tick()
        global r += v
    end
    println("sum: $s")
end

function tick()
    if (cycle-20) % 40 == 0
        println("r * cycle = $r * $cycle = $(r * cycle)")
        global s += r * cycle
    end
    global cycle += 1
end

main()
