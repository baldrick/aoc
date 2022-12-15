function find_boundaries()
    f = open(ARGS[1], "r")
    minx = 999999
    miny = 999999
    maxx = 0
    maxy = 0
    while !eof(f)
        sx, sy, bx, by = decode(readline(f))
        maxx = max(maxx, sx, bx)
        maxy = max(maxy, sy, by)
        minx = min(minx, sx, bx)
        miny = min(miny, sy, by)
    end
    return minx-1, miny-1, maxx-1, maxy-1
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

function populate_grid(minx, miny, maxx, maxy)
    width = maxx - minx + 1
    height = maxy - miny + 1
    grid = zeros(Int8, width, height)
    f = open(ARGS[1], "r")
    while !eof(f)
        sx, sy, bx, by = decode(readline(f))
        ssx = sx - minx
        ssy = sy - miny
        sbx = bx - minx
        sby = by - miny
        #println("Sensor at $ssx,$ssy has beacon at $sbx,$sby")
        grid[ssx, ssy] |= SENSOR
        grid[sbx, sby] |= BEACON
        #no_beacon(grid, ssx, ssy, sbx, sby)
    end
    show_grid(grid, minx, miny)
end

function no_beacon(g, sx, sy, bx, by)
    dx = abs(bx - sx)
    dy = abs(by - sy)
    dist = dx + dy

    yfill = 0
    for x in sx-dist:sx
        for y in sy-yfill:sy+yfill
            mark_no_beacon(g,x,y)
        end
        yfill += 1
    end

    yfill -= 1
    for x in sx+1:sx+dist
        yfill -= 1
        for y in sy-yfill:sy+yfill
            mark_no_beacon(g,x,y)
        end
    end
end

function mark_no_beacon(g, x, y)
    if x < 1 || y < 1 || x > size(g,1) || y > size(g,2)
        return
    end
    if g[x,y] == EMPTY
        g[x,y] |= NO_BEACON
    end
end

EMPTY=0
BEACON=1
SENSOR=2
NO_BEACON=4

function show_grid(g, minx, miny)
    print("   ")
    for x in minx:size(g,1)
        if x % 10 == 0
            print(Int(x/10))
        else
            print(" ")
        end
    end
    println()
    print("   ")
    for x in minx:size(g,1)
        if x % 5 == 0
            print(x%10)
        else
            print(" ")
        end
    end
    println()
    for y in miny:size(g,2)
        if y < 10 print(" ") end
        print("$y ")
        for x in minx:size(g,1)-1
            i = g[x-minx+1, y-miny+1]
            if i == 0
                print(".")
            elseif i == BEACON
                print("B")
            elseif i == SENSOR
                print("S")
            elseif i == NO_BEACON
                print("#")
            else
                print("$i")
            end
        end
        println()
    end
end

minx, miny, maxx, maxy = find_boundaries()
println("$minx,$miny to $maxx,$maxy")
populate_grid(minx, miny, maxx, maxy)
