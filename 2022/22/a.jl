EMPTY=0.0
OK=1.0
WALL=2.0

RIGHT=0
DOWN=1
LEFT=2
UP=3

FACINGS=[">","v","<","^"]

function main()
    mx, my = size()
    println("Creating $mx x $my map")
    map = zeros(my,mx)
    f = open(ARGS[1], "r")
    y = 1
    startx=0
    line = "."
    while line != ""
        line = readline(f)
        for x in 1:length(line)
            if line[x] == '.'
                map[y,x] = OK
                if startx == 0
                    startx = x
                end
            elseif line[x] == '#'
                map[y,x] = WALL
            end
        end
        y += 1
    end
    instructions = decode_instructions(readline(f))
    #dump(map,my,mx)
    #println("$instructions, start at $startx,1")
    x = startx
    y = 1
    facing = RIGHT
    for i in instructions
        #println("==========================")
        #dump(map,my,mx,x,y,facing)
        if i == "L"
            facing -= 1
            if facing < RIGHT facing = UP end
        elseif i == "R"
            facing += 1
            if facing > UP facing = RIGHT end
        else
            if facing == RIGHT
                x, y = move(map, mx, my, x, y, parse(Int, i), 0)
            elseif facing == LEFT
                x, y = move(map, mx, my, x, y, -parse(Int, i), 0)
            elseif facing == UP
                x, y = move(map, mx, my, x, y, 0, -parse(Int, i))
            elseif facing == DOWN
                x, y = move(map, mx, my, x, y, 0, parse(Int, i))
            end
            if map[y,x] != OK
                println("stopped on a non-ok ($(map[y,x])) space after instruction $i !!")
                dump(map,my,mx,x,y,facing)
                exit()
            end
        end
    end
    #println("==========================")
    #dump(map,my,mx,x,y,facing)
    #println("==========================")
    println("row=$y, column=$x, facing=$facing; 1000*$y + 4*$x + $facing = $(1000*y+4*x+facing)")
end

function size()
    f = open(ARGS[1], "r")
    line = readline(f)
    height = 0
    width = 0
    while line != ""
        height += 1
        width = max(width, length(line))
        line = readline(f)
    end
    return width, height
end

function decode_instructions(line)
    instructions = []
    while line != ""
        r = r"([0-9]+|[RL])(.*)"
        m = match(r, line)
        push!(instructions, m[1])
        line = m[2]
    end
    return instructions
end

function move(map, mx, my, x, y, dx, dy)
    step = 1
    if dx == 0
        if dy < 0 step = -1 end
        for s in 1:abs(dy)
            ty = wrap(y+step, my)
            if map[ty,x] == WALL
                break
            elseif map[ty,x] == OK
                y = ty
            else # empty space, wrap around
                ox, oy = find_other_side(map, mx, my, x, ty, 0, step)
                if oy == -1 # we hit a wall at the other side, stop here
                    break
                end
                y = oy
            end
        end
    elseif dy == 0
        if dx < 0 step = -1 end
        for s in 1:abs(dx)
            tx = wrap(x+step, mx)
            if map[y, tx] == WALL
                break
            elseif map[y, tx] == OK
                x = tx
            else # empty space, wrap around
                ox, oy = find_other_side(map, mx, my, tx, y, step, 0)
                if ox == -1 # we hit a wall at the other side, stop here
                    break
                end
                x = ox
            end
        end
    else
        println("ERROR: dx == dy == 0!")
        exit()
    end
    return x, y
end

function find_other_side(map, mx, my, x, y, dx, dy)
    #println("find other side from $x,$y with $dx,$dy size $mx x $my")
    tx = wrap(x+dx, mx)
    ty = wrap(y+dy, my)
    while map[ty,tx] == EMPTY
        tx = wrap(tx+dx, mx)
        ty = wrap(ty+dy, my)
    end
    if map[ty,tx] == WALL
        println("hit wall on the other side: at $tx,$ty from $x,$y moving $dx,$dy")
        return -1, -1
    end
    #println("found other side: $tx,$ty")
    return tx, ty
end

function wrap(p,mp)
    if p > mp return 1 end
    if p <= 0 return mp end
    return p
end

function dump(map,my,mx,px=-1,py=-1,facing=-1)
    for y in 1:my
        for x in 1:mx
            if x==px && y==py
                print(FACINGS[facing+1])
            elseif map[y,x] == OK
                print(".")
            elseif map[y,x] == WALL
                print("#")
            else
                print(" ")
            end
        end
        println()
    end
end

main()
