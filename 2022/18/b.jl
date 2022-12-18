using DataStructures

# 2206 too low

function main()
    f = open(ARGS[1], "r")
    cubes = Dict()
    while !eof(f)
        line = readline(f)
        cubes[decode(line)] = 6
    end
    process_cubes(cubes)
    boundaries = find_boundaries(cubes)
    println("boundaries = $boundaries")
    flood_fill(cubes, boundaries, (0,0,0))
    total = 0
    for cube in cubes
        total += cube.second
    end
    println("Total = $total")
    remove_internal_surfaces(cubes, boundaries)
    nonEnclosedCubes = 0
    for cube in cubes
        nonEnclosedCubes += 1
    end
    println("Enclosed cubes = $(boundaries[1] * boundaries[2] * boundaries[3] - nonEnclosedCubes)")
    etotal = 0
    for cube in cubes
        etotal += cube.second
    end
    println("Total ex enclosed = $etotal")
end

function decode(line)
    r = r"([0-9]+),([0-9]+),([0-9]+)"
    m = match(r, line)
    return parse(Int, m[1]), parse(Int, m[2]), parse(Int, m[3])
end

function remove_internal_surfaces(cubes, boundaries)
    visited = Dict()
    for cube in cubes
        if cube.second == 0
            continue
        end
        # Once completed I realised I could simplify the below with:
        # for (dx,dy,dz) in [(-1,0,0),(1,0,0),(0,-1,0),(0,1,0),(0,0,-1),(0,0,1)]
        for dx in -1:1
            for dy in -1:1
                for dz in -1:1
                    if dx == 0 && dy == 0 && dz == 0 continue end
                    if abs(dx) + abs(dy) + abs(dz) > 1 continue end
                    x,y,z = cube.first
                    x += dx
                    y += dy
                    z += dz
                    if in_bounds(boundaries, (x,y,z)) && !haskey(cubes, (x,y,z))
                        cubes[cube.first] -= 1
                    end
                end
            end
        end
    end
end

function in_bounds(boundaries, pos)
    x,y,z = pos
    if x<0 || y<0 || z<0 return false end
    bx,by,bz = boundaries
    if x>bx || y>by || z>bz return false end
    return true
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

function find_boundaries(cubes)
    mx = 0
    my = 0
    mz = 0
    for cube in cubes
        x,y,z = cube.first
        mx = max(mx, x)
        my = max(my, y)
        mz = max(mz, z)
    end
    return (mx,my,mz)
end

struct qItem
    x::Int
    y::Int
    z::Int
end

function flood_fill(cubes, boundaries, pos)
    q = Queue{qItem}()
    enqueue!(q, qItem(pos[1], pos[2], pos[3]))
    while length(q) > 0
        item = dequeue!(q)
        #println("checking $item")
        x = item.x
        y = item.y
        z = item.z
        if x<0 || y<0 || z<0
            #println("$x,$y,$z out of low bounds")
            continue
        end
        if x>boundaries[1] || y>boundaries[2] || z>boundaries[3]
            #println("$x,$y,$z out of high bounds")
            continue
        end
        if haskey(cubes, (x,y,z))
            #println("cubes[$x,$y,$z] = $(cubes[(x,y,z)])")
            continue
        end
        #println("filled $x,$y,$z")
        cubes[(x,y,z)] = 0 # steam
        maybe_enqueue(cubes, q, x+1,y,z)
        maybe_enqueue(cubes, q, x-1,y,z)
        maybe_enqueue(cubes, q, x,y+1,z)
        maybe_enqueue(cubes, q, x,y-1,z)
        maybe_enqueue(cubes, q, x,y,z+1)
        maybe_enqueue(cubes, q, x,y,z-1)
    end
end

function maybe_enqueue(cubes, q, x,y,z)
    pos=(x,y,z)
    if haskey(cubes, pos) && cubes[pos] > 0
        # Don't flood fill over lava...
        return
    end
    #println("queueing to fill $x,$y,$z")
    enqueue!(q, qItem(x,y,z))
end

main()
