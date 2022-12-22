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
    el = Int(my/4)
    println("edge length = $el")
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
                x, y, f = move(map, el, x, y, parse(Int, i), 0, facing)
            elseif facing == LEFT
                x, y, f = move(map, el, x, y, -parse(Int, i), 0, facing)
            elseif facing == UP
                x, y, f = move(map, el, x, y, 0, -parse(Int, i), facing)
            elseif facing == DOWN
                x, y, f = move(map, el, x, y, 0, parse(Int, i), facing)
            end
            facing = f
            if map[y,x] != OK
                println("stopped on a non-ok ($(map[y,x])) space after instruction $i !!")
                dump(map,my,mx,x,y,facing)
                exit()
            end
            println("moved $i to $x,$y facing $facing")
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

function move(map, el, x, y, dx, dy, facing)
    stepx = 0
    stepy = 0
    if dx != 0
        stepx = 1
        if dx < 0
            stepx = -1
        end
    end
    if dy != 0
        stepy = 1
        if dy < 0
            stepy = -1
        end
    end
    for m in 1:(abs(dx)+abs(dy))
        ftx, fty, nstepx, nstepy, f = wrap(x, y, stepx, stepy, el, facing)
        tx = Int(ftx)
        ty = Int(fty)
        if map[ty,tx] == WALL
            break
        elseif map[ty,tx] == OK
            x = tx
            y = ty
            stepx = nstepx
            stepy = nstepy
            facing = f
        else # shouldn't encounter empty spaces with a cube
            println("ERROR - empty space found from $x,$y moving $stepx,$stepy to $tx,$ty by $(abs(dx)+abs(dy)) !")
            exit()
        end
    end
    return x, y, facing
end

#=
cube

 12
 3
45
6

A down from 6 goes to top of 2 moving down
B down from 2 goes to rhs of 3 moving left
C left from 6 goes to top of 1 moving down
D left from 1 goes to lhs 4 moving right
E right from 2 goes to rhs 5 moving left
F missed this out by accident when folding up the paper cube ;-)
G left of 3 goes to top of 4 moving down
H down from 5 goes to rhs of 6 moving left

=#

# takes x,y of flattened cube, stepx,stepy and its edge length
# returns tx, ty, nstepx, nstepy
function wrap(x, y, stepx, stepy, el, facing)
    # A x=1-50, y=200, dx=0, dy=1 => x+=100, y=1, dx=0, dy=1
    if x >= 1 && x <= el && y == 4*el && stepx == 0 && stepy == 1
        println("A")
        return x+2*el, 1, 0, 1, DOWN
    # A'
    elseif x >= 2*el+1 && x <= 3*el && y == 1 && stepx == 0 && stepy == -1
        println("A'")
        return x-2*el, 4*el, 0, -1, UP
    # B x=101-150, y=50, dx=0, dy=1 => x=100, y=x-50, dx=-1, dy=0
    elseif x >= 2*el+1 && x <= 3*el && y == el && stepx == 0 && stepy == 1
        println("B")
        return 2*el, x-el, -1, 0, LEFT
    # B'
    elseif x == 2*el && y >= el+1 && y <= 2*el && stepx == 1 && stepy == 0
        println("B'")
        return y+el, el, 0, -1, UP
    # C x=1, y=151-200, dx=-1, dy=0 => x=y-100, y=1, dx=0, dy=1
    elseif x == 1 && y >= 3*el+1 && y <= 4*el && stepx == -1 && stepy == 0
        println("C")
        return y-2*el, 1, 0, 1, DOWN
    # C'
    elseif x >= el+1 && x <= 2*el && y == 1 && stepx == 0 && stepy == -1
        println("C'")
        return 1, x+2*el, 1, 0, RIGHT
    # D x=51, y=1-50, dx=-1, dy=0 => x=1, y = 151-y, dx=1, dy=0
    elseif x == el+1 && y >= 1 && y <= el && stepx == -1 && stepy == 0
        println("D")
        return 1, 3*el+1 - y, 1, 0, RIGHT
    # D'
    elseif x == 1 && y >= el*2+1 && y <= el*3 && stepx == -1 && stepy == 0
        println("D' returning $(el+1), $y-3*$el-1, 1, 0, $RIGHT")
        return el+1, 3*el+1 - y, 1, 0, RIGHT
    # E x=150, y=1-50, dx=1, dy=0 => x=100, y=151-y, dx=-1, dy=0
    elseif x == 3*el && y >= 1 && y <= el && stepx == 1 && stepy == 0
        println("E")
        return 2*el, 3*el+1-y, -1, 0, LEFT
    # E'
    elseif x == 2*el && y >= 2*el+1 && y <= 3*el && stepx == 1 && stepy == 0
        println("E'")
        return 150, 3*el+1-y, -1, 0, LEFT
    # G x=51, y=51-100, dx=-1, dy=0 => x=y-50, y=101, dx=0, dy=1
    elseif x == el+1 && y >= el+1 && y <= 2*el && stepx == -1 && stepy == 0
        println("G")
        return y-el, 2*el+1, 0, 1, DOWN
    # G'
    elseif x >= 1 && x <= el && y == 2*el+1 && stepx == 0 && stepy == -1
        println("G'")
        return 51, x+el, 1, 0, RIGHT
    # H x=51-100, y=150, dx=0, dy=1 => y=x+100, x=50, dx=-1, dy=0
    elseif x >= el+1 && x <= 2*el && y == 3*el && stepx == 0 && stepy == 1
        println("H")
        return el, x+2*el, -1, 0, LEFT
    # H'
    elseif x == el && y >= 3*el+1 && y <= 4*el && stepx == 1 && stepy == 0
        println("H'")
        return y-2*el, 3*el, 0, -1, UP
    else
        println("else")
        return x+stepx, y+stepy, stepx, stepy, facing
    end
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
