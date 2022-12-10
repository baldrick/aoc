cycle = 1
r = 1
s = 0
crt = ""

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
    println(crt)
end

function tick()
    if (cycle-20) % 40 == 0
        println("r * cycle = $r * $cycle = $(r * cycle)")
        global s += r * cycle
    end
    draw()
    global cycle += 1
end

function draw()
    # r is the sprite position
    crtx = (cycle % 40)-1
    if crtx >= r-1 && crtx <= r+1
        println("#$cycle r=$r => lit")
        global crt = "$crt#"
    else
        global crt = "$crt "
    end
    if cycle % 40 == 0
        global crt = "$crt\n"
    end
end

main()
