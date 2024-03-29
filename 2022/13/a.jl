import JSON

function main()
    f = open(ARGS[1], "r")
    i = 1
    total = 0
    while !eof(f)
        left = JSON.parse(readline(f))
        right = JSON.parse(readline(f))
        if less_than(left, right)
            total += i
        end
        if !eof(f)
            blank = readline(f)
        end
        i += 1
    end
    println(total)
end

# Had to resort to reddit for inspiration for this, my previous attempt was nasty.
# However, at least I learnt how to override methods by doing so :-)
function less_than(left::Int, right::Int)
    return left < right
end

function less_than(left::Vector, right::Int)
    return less_than(left, [right])
end

function less_than(left::Int, right::Vector)
    return less_than([left], right)
end

function less_than(left::Vector, right::Vector)
    if isempty(right)
        return false
    end
    return isempty(left) || less_than(left[1], right[1]) || (!less_than(right[1], left[1]) && less_than(left[2:end], right[2:end]))
end

main()
