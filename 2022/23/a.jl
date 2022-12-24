l = 500

mutable struct Elf
    y::Int
    x::Int
    nameIndex::Int
end

function main()
    f = open(ARGS[1], "r")
    map = zeros(l,l)
    ox,oy = Int(l/2), Int(l/2)
    y = 1
    elves = []
    elf = 1
    while !eof(f)
        line = readline(f)
        for x in 1:length(line)
            if line[x] == '.'
                continue
            end
            map[oy+y,ox+x] = elf
            push!(elves, Elf(oy+y,ox+x,elf))
            elf += 1
        end
        y += 1
    end
    #dump(map, "start")
    process(map, elves)
end

function process(map, elves)
    instructions = [proposeNorth, proposeSouth, proposeWest, proposeEast]

    n = 1
    while n <= 10
        proposals = propose(map, elves, instructions)
        if length(proposals) == 0
            break
        end

        #display(proposals)
        move(map, proposals, elves)
        #dump(map, "after $n")

        #=
        Finally, at the end of the round, the first direction the Elves considered is moved to the end of
        the list of directions. For example, during the second round, the Elves would try proposing a move
        to the south first, then west, then east, then north. On the third round, the Elves would first
        consider west, then east, then north, then south.
        =#
        i = instructions[1]
        deleteat!(instructions, 1)
        push!(instructions, i)

        n += 1
    end
    println("Empty ground between elves = $(emptyGround(map))")
end

function emptyGround(map)
    minx = findx(map, 1, size(map,1))
    maxx = findx(map, size(map,1), 1)
    miny = findy(map, 1, size(map,1))
    maxy = findy(map, size(map,1), 1)
    println("$minx,$miny - $maxx,$maxy")
    ground = 0
    for y in miny:maxy
        for x in minx:maxx
            if map[y,x] == 0
                ground += 1
            end
        end
    end
    return ground
end

function propose(map, elves, instructions)
    proposals = Dict()
    for elf in elves
        if noNeighbours(map, elf)
            #println("$elf does nothing")
            continue
        end

        for i in instructions
            if i(map, proposals, elf)
                break
            end
        end
    end
    return proposals
end

# If there is no Elf in the N, NE, or NW adjacent positions, the Elf proposes moving north one step.
function proposeNorth(map, proposals, elf)
    proposedPosition = addPosition(elf, 0, -1)
    if !elfAt(map, proposedPosition, 1, 0)
        #println("$elf proposes N")
        addProposal(proposals, proposedPosition, elf)
        return true
    end
    return false
end

# If there is no Elf in the S, SE, or SW adjacent positions, the Elf proposes moving south one step.
function proposeSouth(map, proposals, elf)
    proposedPosition = addPosition(elf, 0, 1)
    if !elfAt(map, proposedPosition, 1, 0)
        #println("$elf proposes S")
        addProposal(proposals, proposedPosition, elf)
        return true
    end
    return false
end

# If there is no Elf in the W, NW, or SW adjacent positions, the Elf proposes moving west one step.
function proposeWest(map, proposals, elf)
    proposedPosition = addPosition(elf, -1, 0)
    if !elfAt(map, proposedPosition, 0, 1)
        #println("$elf proposes W")
        addProposal(proposals, proposedPosition, elf)
        return true
    end
    return false
end    

# If there is no Elf in the E, NE, or SE adjacent positions, the Elf proposes moving east one step.
function proposeEast(map, proposals, elf)
    proposedPosition = addPosition(elf, 1, 0)
    if !elfAt(map, proposedPosition, 0, 1)
        #println("$elf proposes E")
        addProposal(proposals, proposedPosition, elf)
        return true
    end
    return false
end


function noNeighbours(map, elf)
    for y in -1:1
        for x in -1:1
            if x == 0 && y == 0
                continue
            end
            if map[elf.y+y,elf.x+x] != 0
                return false
            end
        end
    end
    return true
end

function addProposal(proposals, proposedPosition, elf)
    if haskey(proposals, proposedPosition)
        push!(proposals[proposedPosition], elf)
    else
        proposals[proposedPosition] = [elf]
    end
end

function addPosition(elf, dx, dy)
    return elf.x+dx, elf.y+dy
end

function elfAt(map, proposedPosition, dx, dy)
    x,y = proposedPosition
    return map[y,x] != 0 || map[y+dy, x+dx] != 0 || map[y-dy, x-dx] != 0
end

function move(map, proposals, elves)
    for p in proposals
        if length(p.second) > 1
            #println("Proposal $p ignored, $(length(p.second)) elves proposed it")
            continue
        end
        elf = p.second[1]
        map[elf.y, elf.x] = 0
        elf.x, elf.y = p.first
        map[elf.y, elf.x] = elf.nameIndex
    end
end

function dump(map, message)
    println("=== $message =====")
    minx = findx(map, 1, size(map,1))
    maxx = findx(map, size(map,1), 1)
    miny = findy(map, 1, size(map,1))
    maxy = findy(map, size(map,1), 1)
    println("$minx,$miny - $maxx,$maxy")
    for y in miny:maxy
        for x in minx:maxx
            if map[y,x] == 0
                print(".")
            else
                print(Int(map[y,x]))
            end
        end
        println()
    end
end

function findx(map, from, to)
    step = 1
    if from > to
        step = -1
    end
    for x in from:step:to
        for y in 1:size(map,2)
            if map[y,x] != 0
                return x
            end
        end
    end
    println("ERROR: could not find x edge from $from to $to")
    return -1          
end

function findy(map, from, to)
    step = 1
    if from > to
        step = -1
    end
    for y in from:step:to
        for x in 1:size(map,2)
            if map[y,x] != 0
                return y
            end
        end
    end
    println("ERROR: could not find y edge from $from to $to")
    return -1          
end

main()
