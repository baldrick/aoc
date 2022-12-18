function main()
    maxx = 4_000_000
    if ARGS[1] == "test"
        maxx = 20
    end

    rows = []
    f = open(ARGS[1], "r")
    while !eof(f)
        sx,sy, bx,by = decode(readline(f))
        push!(rows, [sx,sy,bx,by])
    end

    println("Read file")

    for yTarget in 0:maxx
        println("Doing $yTarget")
        targetRow = Set()
        for row in rows
            sx,sy, bx,by = row
            dist = abs(sx-bx) + abs(sy-by)
            overlap = dist-abs(sy-yTarget)
            if overlap <= 0 continue end
            println("$sx,$sy to $bx,$by means dist=$dist, $overlap overlap $yTarget so pushing no beacons at $(sx-overlap) to $(sx+overlap)")
            for x in max(0,sx-overlap):min(sx+overlap,maxx)
                push!(targetRow, x)
            end
        end
        println("$(length(targetRow)): $targetRow")
        if length(targetRow) == maxx
            println("y = $yTarget")
            x = 0
            while x < maxx
                if !in(x, targetRow)
                    println("x = $x, 4m * $x + $yTarget = $(4e6*x+yTarget)")
                    break
                end
                x += 1
            end
        end
    end
end

# Sensor at x=3992558, y=1933059: closest beacon is at x=3748004, y=2000000
function decode(line)
    #println("Matching $line")
    r = r"Sensor at x=([-0-9]+), y=([-0-9]+): closest beacon is at x=([-0-9]+), y=([-0-9]+)"
    m = match(r, line)
    sx = parse(Int, m[1])
    sy = parse(Int, m[2])
    bx = parse(Int, m[3])
    by = parse(Int, m[4])
    return sx, sy, bx, by
end

main()
