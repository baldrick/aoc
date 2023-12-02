function main()
    f = open(ARGS[1], "r")
    x = 0
    y = 0
    direction = 0
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
            println("Moving $steps on $direction")
            if direction == 0
                y += steps
            elseif direction == 90
                x += steps
            elseif direction == 180
                y -= steps
            elseif direction == 270
                x -= steps
            else
                println("ERROR $(m)")
            end
            line = m[3]
        end
    end
    println("dist = $(abs(x)+abs(y))")
end

main()
