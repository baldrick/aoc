import JSON

function main()
    f = open(ARGS[1], "r")
    i = 1
    total = 0
    packets = []
    while !eof(f)
        left = JSON.parse(readline(f))
        push!(packets, left)
        right = JSON.parse(readline(f))
        push!(packets, right)
        if less_than(left, right)
            total += i
        end
        if !eof(f)
            blank = readline(f)
        end
        i += 1
    end
    println(total)
    sp = sort(packets, lt=less_than)
    println("== PACKETS ==")
    display(sp)
    println("=============")
    div1 = JSON.parse("[[2]]")
    div2 = JSON.parse("[[6]]")
    i1 = 0
    i2 = 0
    display(div1)
    display(div2)
    for i in 1:length(sp)
        packet = sp[i]
        #display(packet)
        if packet == div1
            i1 = i
        end
        if packet == div2
            i2 = i
        end
    end
    println("$i1 * $i2 = $(i1 * i2)")
end

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
