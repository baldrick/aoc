# 11621 too high
# 3557 too low

using DataStructures

mutable struct Item
    n::Int
    index::Int
    done::Bool
end

function main()
    f = open(ARGS[1], "r")
    input = []
    index = 0
    while !eof(f)
        line = readline(f)
        item = Item(parse(Int, line), index, false)
        push!(input, item)
        index += 1
    end
    dump(input, "Initial arrangement")
    while true
        movedOne = false
        i = 1
        while i <= length(input)
            #println("====================")
            moved, inc = move(i, input)
            movedOne |= moved
            #print("$i += $inc = ")
            i += inc + 1
            #println(i)
        end
        if !movedOne break end
    end
    dump(input, "complete")
    indexZero = 1
    for i in 1:length(input)
        if input[i].n == 0
            indexZero = i
            break
        end
    end
    println("indexZero = $indexZero")

    ai = indexZero + 1000
    bi = indexZero + 2000
    ci = indexZero + 3000

    if ai > length(input) ai = ai % length(input) end
    if bi > length(input) bi = bi % length(input) end
    if ci > length(input) ci = ci % length(input) end

    if ai == 0 ai = length(input) end
    if bi == 0 bi = length(input) end
    if ci == 0 ci = length(input) end

    a = input[ai]
    b = input[bi]
    c = input[ci]
    println("$(a.n) + $(b.n) + $(c.n) = $(a.n + b.n + c.n) (indices $ai, $bi, $ci)")
end

#=
2 + -3 = -1, want at index 6, i.e. length-1
2 + -10 = -8, want at index 6, i.e. length+(abs(8)%7)*7+-8
=#

function move(i, input)
    item = input[i]
    if item.done
        #println("$(item.n) (at $i) has been moved")
        return false, 0
    end
    if item.n == 0
        #println("$(item.n) (at $i) does not move")
        return false, 0
    end
    dist = abs(item.n) % length(input)
    if item.n < 0
        dist = -dist
    end
    newIndex = i + dist
    if newIndex > length(input)
        newIndex = newIndex % length(input)
    end
    if newIndex <= 0
        newIndex = length(input) + newIndex
    end
    item.done = true
    if newIndex == i 
        # No need to move
        #dump(input, "#$(item.n) stayed at index $newIndex")
    elseif newIndex > i
        #dump(input)
        deleteat!(input, i)
        insert!(input, newIndex-1, item)
        #println()
        #dump(input, "#$i moved $(item.n) $dist steps fwd to $newIndex")
        return true, -1
    else
        insert!(input, newIndex, item)
        deleteat!(input, i+1)
        #dump(input, "#$i moved $(item.n) $dist steps bwd to $newIndex")
        return true, 0
    end
end

function dump(input, message="")
    for n in input
        print("$(n.n), ")
    end
    if message != ""
        println(message)
    end
end

main()
