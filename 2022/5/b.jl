using DataStructures

function main()
    f = open(ARGS[1], "r")
    stacks = readStacks(f)
    println("$stacks")
    while ! eof(f)
        line = readline(f)
        (count, from, to) = decodeMove(line)
        moveItems(stacks, count, from, to)
    end
    top = topStacks(stacks)
    println("$top")
end

#=
file input looks like:
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 
(blank line)
(instructions)
read the stack starting position into an array of strings reversed...
=#
function readStacks(f)
    stacks = Stack{String}()
    while ! eof(f) # we'll break early
        line = readline(f)
        if line == ""
            return decodeStacks(stacks)
        end
        push!(stacks, line)
    end
    println("Failed to find blank line after stacks!")
    exit(1)
end

# Translate text lines into an array of Stack{String}.
function decodeStacks(inputStacks)
    first = true
    stacks = []
    for line in inputStacks
        if first
            first = false
            items = splitInput(line)
            for item in items
                push!(stacks, Stack{String}())
            end
            continue
        end
        items = splitInput(line)
        println("$items")
        stack = 0
        for item in items
            stack += 1
            if item == "    "
                continue
            end
            push!(stacks[stack], item)
        end
    end
    return stacks
end

# Take a line and split every four characters into an array.
function splitInput(line)
    line = "$line " # each input line misses its final space
    println("processing '$line'")
    numItems = length(line) / 4 # each item has 4 characters
    items = []
    for i in 1:numItems
        n = convert(Int, ((i-1)*4)+1)
        push!(items, line[n:n+3])
    end
    return items
end

# move <count> from <src> to <dest>
function decodeMove(line)
    r = r"move ([0-9]+) from ([0-9]+) to ([0-9]+)"
    m = match(r, line)
    return parse(Int, m[1]), parse(Int, m[2]), parse(Int, m[3])
end

function moveItems(stacks, count, from, to)
    # retain order
    println("move $count from $from to $to")
    temp = Stack{Any}()
    for i in 1:count
        item = pop!(stacks[from])
        push!(temp, item)
    end
    for item in temp
        push!(stacks[to], item)
    end
end

function topStacks(stacks)
    top = []
    for stack in stacks
        push!(top, pop!(stack))
    end
    return top
end

main()
