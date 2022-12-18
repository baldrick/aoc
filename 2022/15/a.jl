function main()
    f = open(ARGS[1], "r")
    yTarget = 2_000_000
    if ARGS[1] == "test"
        yTarget = 10
    end
    targetRow = Set()
    beacons = Set()
    while !eof(f)
        sx,sy, bx,by = decode(readline(f))
        if by == yTarget
            push!(beacons, by)
        end
        dist = abs(sx-bx) + abs(sy-by)
        overlap = dist-abs(sy-yTarget)
        if overlap <= 0 continue end
        #println("$sx,$sy to $bx,$by means dist=$dist, $overlap overlap $yTarget so pushing no beacons at $(sx-overlap) to $(sx+overlap)")
        for x in sx-overlap:sx+overlap
            push!(targetRow, x)
        end
    end
    println(length(setdiff(targetRow, beacons)))
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
