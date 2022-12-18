function main()
    f = open(ARGS[1], "r")
    cubes = Dict()
    while !eof(f)
        line = readline(f)
        cubes[decode(line)] = 6
    end
    process_cubes(cubes)
    display(cubes)
    total = 0
    for cube in cubes
        total += cube.second
    end
    println("Total = $total")
end

function decode(line)
    r = r"([0-9]+),([0-9]+),([0-9]+)"
    m = match(r, line)
    return parse(Int, m[1]), parse(Int, m[2]), parse(Int, m[3])
end

function process_cubes(cubes)
    visited = Dict()
    for cube in cubes
        if haskey(visited, cube)
            continue
        end
        visited[cube] = true
        if neighbour(cubes, cube, 1, 0, 0) cubes[cube.first] -= 1 end
        if neighbour(cubes, cube, -1, 0, 0) cubes[cube.first] -= 1 end
        if neighbour(cubes, cube, 0, 1, 0) cubes[cube.first] -= 1 end
        if neighbour(cubes, cube, 0, -1, 0) cubes[cube.first] -= 1 end
        if neighbour(cubes, cube, 0, 0, 1) cubes[cube.first] -= 1 end
        if neighbour(cubes, cube, 0, 0, -1) cubes[cube.first] -= 1 end
    end
end

function neighbour(cubes, cube, dx, dy, dz)
    x,y,z = cube.first
    x += dx
    y += dy
    z += dz
    return haskey(cubes, (x,y,z))
end

main()
