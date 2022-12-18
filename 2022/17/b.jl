import Printf
using Printf

# Shapes are upside down for ease of checking bottom up.
MINUS=[32+16+8+4]
PLUS=[16,32+16+8,16]
REV_L=[32+16+8,8,8]
I_BEAM=[32,32,32,32]
SQUARE=[32+16,32+16]
FLOOR=256+128+64+32+16+8+4+2+1
NEAR_L_WALL=128
NEAR_R_WALL=2

function main()
    f = open(ARGS[1], "r")
    jets = readline(f)
    chamber = Array{Int}([0, 0, 0, 0, FLOOR])
    jet = 1
    shapes = 0
    rocks = 1_000_000_000_000
    cycleHeight = 0
    while shapes < rocks
        for shape in [MINUS, PLUS, REV_L, I_BEAM, SQUARE]
            jet = drop(chamber, shape, jets, jet)
            shapes += 1
            if shapes == rocks
                break
            end
            if cycleHeight == 0
                mcShapeNumber, mcHeight = cycle(chamber, jet, shape, shapes)
                if mcShapeNumber != 0
                    cycleHeight = length(chamber)-5-mcHeight
                    repetitions = floor((rocks-shapes) / (shapes-mcShapeNumber)) - 1
                    cycleHeight *= repetitions
                    skipShapes = repetitions * (shapes-mcShapeNumber)
                    println("Found cycle from $shapes shapes, chamber height=$(length(chamber)-5) * repetitions=$repetitions = cycle height=$cycleHeight, remaining rocks = $(rocks - skipShapes)")
                    shapes += skipShapes
                end
            end
        end
    end
    println("Chamber height = $(cycleHeight + length(chamber) - 5)") # ignore the headroom (assume it's been added)
    @printf "Chamber height = %d\n" cycleHeight + length(chamber) - 5
    #show_chamber(chamber)
end

m = Dict()
function cycle(chamber, jet, shape, shapeNumber)
    k = (jet, shape, heights_to_rock(chamber))
    println("$shapeNumber: height=$(length(chamber)), jet=$jet, shape=$shape, key=$k")
    if haskey(m, k)
        return m[k]
    end
    global m[k] = (shapeNumber, length(chamber)-5)
    return 0,0
end

function heights_to_rock(chamber)
    heights = []
    for x in [128,64,32,16,8,4,2]
        for y in 1:length(chamber)
            if chamber[y] & x > 0
                push!(heights, y)
                break
            end
        end
    end
    return heights
end

function drop(chamber, shape, jets, jet)
    #println("dropping $shape")
    y = 0 # the top of chamber is y=1, floor moves down as more shapes appear...
    stopped = false
    while !stopped
        y += 1
        #println("start of chamber when dropping $shape")
        #show_chamber(chamber, shape, y)
        shape = jet_around(chamber, y, shape, jets[jet])
        #println("chamber having moved $shape around")
        #show_chamber(chamber, shape, y)
        jet = (jet % length(jets)) + 1
        stopped = has_hit_rock_or_floor(chamber, y, shape)
    end
    add_shape_to_chamber(chamber, y, shape)
    return jet
end

function jet_around(chamber, y, shape, jet)
    if jet == '<'
        #println("Test moving $shape left")
        if !near_left_wall(shape)
            newShape = copy(shape)
            canMove = true
            for shapeRow in 1:length(shape)
                if y <= 0
                    newShape[shapeRow] = shape[shapeRow] << 1
                    y -= 1
                    continue
                end
                newShape[shapeRow] = shape[shapeRow] << 1
                if newShape[shapeRow] & chamber[y] > 0
                    canMove = false
                    break
                end
                y -= 1
            end
            if canMove
                #println("Moving $shape left")
                shape = newShape
            end
        end
    else # jet = ">"
        #println("Test moving $shape right, y=$y, before:")
        #show_chamber(chamber, shape, y)
        sy = y
        if !near_right_wall(shape)
            newShape = copy(shape)
            canMove = true
            for shapeRow in 1:length(shape)
                if sy <= 0
                    newShape[shapeRow] = newShape[shapeRow] >> 1
                    sy -= 1
                    continue
                end
                newShape[shapeRow] = newShape[shapeRow] >> 1
                if newShape[shapeRow] & chamber[sy] > 0
                    #println("$shape cannot move (newShape=$newShape,sy=$sy)")
                    canMove = false
                    break
                end
                sy -= 1
            end
            if canMove
                #println("Moving $shape right")
                shape = newShape
            end
        end
        #println("after possible move right shape=$shape, y=$y:")
        #show_chamber(chamber, shape, y)
    end
    return shape
end

function near_left_wall(shape)
    for shapeRow in shape
        if shapeRow & NEAR_L_WALL != 0
            return true
        end
    end
    return false
end

function near_right_wall(shape)
    for shapeRow in shape
        if shapeRow & NEAR_R_WALL != 0
            return true
        end
    end
    return false
end

function has_hit_rock_or_floor(chamber, y, shape)
    for shapeRow in 1:length(shape)
        if y+2-shapeRow < 1
            return false
        end
        if shape[shapeRow] & chamber[y+2-shapeRow] > 0
            # Shape has hit something so stop.
            return true
        end
    end
    return false
end

function add_shape_to_chamber(chamber, y, shape)
    #println("adding $shape to chamber, current chamber (y=$y):")
    #show_chamber(chamber, shape, y)
    for shapeRow in 1:length(shape)
        if y == 0
            pushfirst!(chamber, shape[shapeRow])
            y += 1
        else
            chamber[y] |= shape[shapeRow]
        end
        y -= 1
    end
    make_room_for_next_shape(chamber)
end

function make_room_for_next_shape(chamber)
    #println("chamber after adding shape:")
    #show_chamber(chamber)
    # Make sure there's a gap of 3 units from top occupied row.
    topOccupiedRow = 1
    for y in 1:length(chamber)
        if chamber[y] == 0
            continue
        end
        topOccupiedRow = y
        break
    end
    rowsToAdd = 5 - topOccupiedRow
    #println("top occupied row = $topOccupiedRow")
    while rowsToAdd > 0
        pushfirst!(chamber, 0)
        rowsToAdd -= 1
    end
end

# bottom row, index = 1
# top of chamber, y = 1
# shape y = bottom of shape
# top of shape = shape y + length - 1
# top of shape in chamber = shape y - length + 1
# chamber y = 1 = top row => shape y == y => draw bottom row
# chamber y = 2 = second row => shape y == y => draw bottom row at 2, next row at 1
# => draw row if chamber y = shape y - length of shape + 1
# and row of shape to draw = length of shape - rows already drawn of shape

function show_chamber(c, shape = nothing, sy = 0)
    println(repeat('=',50))
    #println("sy=$sy")
    shapeRowsDrawn = 0
    if shape != nothing && sy - length(shape) < 0
        shapeRowsDrawn = length(shape) - sy
    end
    seenShape = false
    for y in 1:length(c)
        drawShapeRow = -1
        if shape != nothing
            if y <= sy && y >= sy - length(shape)
                drawShapeRow = length(shape) - shapeRowsDrawn
            end
        end

        for x in [256,128,64,32,16,8,4,2,1]
            if drawShapeRow > 0 && drawShapeRow <= length(shape) && shape[drawShapeRow] & x > 0
                print("@")
                seenShape = true
                continue
            end
            if x == 256 || x == 1
                if y == length(c)
                    print("+")
                else
                    print("|")
                end
            elseif c[y] & x == 0
                print(".")
            elseif y == length(c)
                print("-")
            else
                print("#")
            end
        end
        if seenShape
            shapeRowsDrawn += 1
        end
        println()
    end
end

main()
