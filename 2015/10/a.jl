function main()
    f = open(ARGS[1], "r")
    line = readline(f)
    for n in 1:50
        line = lookAndSay(line)
        println("$n: length = $(length(line))")
    end
end

function lookAndSay(line)
    newline = []
    i = 1
    while i <= length(line)
        c = line[i]
        n = 1 + countIdentical(c, line, i+1)
        push!(newline, "$n$c")
        i += n
    end
    return join(newline)
end

function countIdentical(c, line, start)
    n = 0
    for t in line[start:end]
        if c != t
            break
        end
        n += 1
    end
    return n
end

main()
