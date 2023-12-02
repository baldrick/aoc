function main()
    f = open(ARGS[1], "r")
    x = 0
    y = 0
    direction = 0
    seen = Dict()
    seen[(x,y)] = true
    while !eof(f)
        line = readline(f)
        while true
            r = r"([LR]+)([0-9]+)[, ]*(.*)"
            m = match(r, line)
            if m == nothing
                break
            end
            steps = parse(Int, m[2])
            if m[1] == "R"
                direction = (direction + 90 + 360) % 360
            else
                direction = (direction - 90 + 360) % 360
            end
            if direction == 0
                x, y, done = move(seen, x, y, 0, 1, steps)
            elseif direction == 90
                x, y, done = move(seen, x, y, 1, 0, steps)
            elseif direction == 180
                x, y, done = move(seen, x, y, 0, -1, steps)
            elseif direction == 270
                x, y, done = move(seen, x, y, -1, 0, steps)
            else
                println("ERROR $(m)")
            end
            if done
                println("Visited $x,$y already, $(m[3]) left")
                break
            end
            println("Moved $steps on $direction to $x,$y")
            line = m[3]
        end
    end
    println("$x,$y dist = $(abs(x)+abs(y))")
end

function move(seen, x, y, dx, dy, steps)
    for step in 1:steps
        x += dx
        y += dy
        if haskey(seen, (x,y))
            return x, y, true
        else
            seen[(x,y)] = true
        end
    end
    return x, y, false
end

main()
