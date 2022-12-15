function findCavernSize()
    minx=99999
    maxx=0
    miny=99999
    maxy=0
    f = open(ARGS[1], "r")
    while !eof(f)
        rest = readline(f)
        while rest != ""
            x, y, rest = getCoords(rest)
            if x > maxx maxx = x end
            if y > maxy maxy = y end
            if x < minx minx = x end
            if y < miny miny = y end
        end
    end
    return minx, maxx, miny, maxy
end

function main(minx, maxx, miny, maxy)
    f = open(ARGS[1], "r")
    width = maxx - minx + 1
    height = maxy - miny + 1
    cavern = zeros(height, width)
    println("minx=$minx, miny=$miny")
    dx = minx - 1
    dy = miny - 1
    while !eof(f)
        decode(cavern, dx, dy, readline(f))
    end
    i = 0
    while true
        if !dropSand(cavern, width, height, 500 - dx, 1)
            break
        end
        i += 1
    end
    println("$i items of sand")
    showCavern(cavern, width, height)
end

#=
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
=#
# x,y normalizes to right-sized cavern
function decode(cavern, x, y, line)
    #println(line)
    rest = line
    sx, sy, rest = getCoords(rest)
    println("$sx,$sy normalizes to $sx-$x,$sy-$y = $(sx-x),$(sy-y)")
    while rest != ""
        ex, ey, rest = getCoords(rest)
        println("$ex,$ey normalizes to $ex-$x,$ey-$y = $(ex-x),$(ey-y)")
        fill(cavern, sx-x, sy-y, ex-x, ey-y)
        sx = ex
        sy = ey
    end
end

function getCoords(line)
    r = r"([0-9]+),([0-9]+)(.*)"
    m = match(r, line)
    x = parse(Int, m[1])
    y = parse(Int, m[2])
    #println("coords: $x,$y, rest=$(m[3])")
    return x, y, m[3]
end

function fill(cavern, sx, sy, ex, ey)
    dx = sx > ex ? -1 : 1
    dy = sy > ey ? -1 : 1
    println("filling from $sx,$sy (step $dx) - $ex,$ey (step $dy)")
    for x in sx:dx:ex
        for y in sy:dy:ey
            cavern[y, x] = 1
        end
    end
end

function dropSand(c, w, h, x, y)
    moved = true
    while moved
        while moved
            #println("try down")
            nx, ny, moved, ok = down(c, w, h, x, y)
            if !ok return false end
            if moved
                x = nx
                y = ny
            end
        end
        #println("try left")
        nx, ny, moved, ok = left(c, w, h, x, y)
        if !ok return false end
        if moved
            x = nx
            y = ny
        else
            #println("try right")
            nx, ny, moved, ok = right(c, w, h, x, y)
            if !ok return false end
            if moved
                x = nx
                y = ny
            end
        end
    end
    #println("sand stopped at $x,$y")
    if c[y,x] == 0
        c[y,x] = 2
        return true
    end
    return false
end

function down(cavern, w, h, x, y)
    return move(cavern, w, h, x, y, 0, 1)
end

function left(cavern, w, h, x, y)
    return move(cavern, w, h, x, y, -1, 1)
end

function right(cavern, w, h, x, y)
    return move(cavern, w, h, x, y, 1, 1)
end

function move(cavern, w, h, x, y, dx, dy)
    sx = x
    sy = y

    nx = x+dx
    ny = y+dy
    if !inBounds(w, h, nx, ny)
        return x, y, x!=sx || y!=sy, false
    elseif cavern[ny,nx] == 0
        #println("moving from $x,$y to $nx,$ny")
        #showCavern(cavern, w, h, nx, ny)
        x = nx
        y = ny
    end
    return x, y, x!=sx || y!=sy, true
end

function inBounds(w, h, x, y)
    return x > 0 && x <=w && y > 0 && y <= h
end

function showCavern(cavern, x, y, nx=0, ny=0)
    for i in 1:y
        for j in 1:x
            if i==ny && j==nx
                print("*")
            else
                print(cavern[i,j] == 0 ? "." : cavern[i,j] == 1 ? "#" : "o")
            end
        end
        println()
    end
end

minx, maxx, miny, maxy = findCavernSize()
println("bounds $minx,$miny - $maxx,$maxy")
main(minx, maxx, 0, maxy)
