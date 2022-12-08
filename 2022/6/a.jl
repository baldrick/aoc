function main()
    f = open(ARGS[1], "r")
    line = readline(f)
    buffer = line[1:3]
    println("looking for repetition in $line")
    for i in 4:length(line)
        c = line[i]
        # if c is the same as any char in buffer, move on
        cpos = findChar(c, buffer)
        if cpos == 0 && ! repeats(buffer)
            return i
        end
        buffer = buffer[2:3]
        buffer = "$buffer$c"
    end
    return 0
end

# Return the position of the last occurrence of c in buffer or zero if it doesn't occur.
function findChar(c, buffer)
    for i in length(buffer):-1:1
        if buffer[i] == c
            println("found $c in $buffer at $i")
            return i
        end
    end
    println("failed to find $c in $buffer")
    return 0
end

function repeats(buffer)
    for i in 1:length(buffer)-1
        for j in i+1:length(buffer)
            if buffer[i] == buffer[j]
                return true
            end
        end
    end
    return false
end

n = main()
println("$n")