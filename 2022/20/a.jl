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
    ai = (1000+indexZero-1) % length(input) + 1
    bi = (2000+indexZero-1) % length(input) + 1
    ci = (3000+indexZero-1) % length(input) + 1
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
        println("$(item.n) (at $i) has been moved")
        return false, 0
    end
    if item.n == 0
        println("$(item.n) (at $i) does not move")
        return false, 0
    end
    if i+item.n == 0
        newIndex = length(input)
    elseif item.n < 0
        a=abs(i+item.n)
        b=a%length(input)
        c=b*length(input)
        println("i=$i, newIndex = $(i+item.n) + ($a%$(length(input))*$(length(input))) = $(i+item.n) + $b*$(length(input))= $(i+item.n) + $c")
        newIndex = i+item.n + ((abs(i+item.n)%length(input))*length(input))
        if newIndex > length(input)
            newIndex = newIndex % length(input)
        end
        if newIndex < 1
            newIndex = length(input) - newIndex
        end
    else
        println("i=$i, newIndex = $(item.n+1)%$(length(input))+1")
        newIndex = ((item.n + i) % length(input)) + 1
    end
    item.done = true
    if newIndex > i
        #dump(input)
        deleteat!(input, i)
        insert!(input, newIndex-1, item)
        #println()
        #dump(input, "-moved $(item.n)")
        return true, -1
    else
        insert!(input, newIndex, item)
        deleteat!(input, i+1)
        #dump(input, "+moved $(item.n)")
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
