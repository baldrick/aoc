f = open(ARGS[1], "r")

function getRucksack(f)
    line = readline(f)
    r = Set()
    for item in 1:length(line)
        push!(r, line[item])
    end
    return r
end

sum = 0
while ! eof(f)
    elf1 = getRucksack(f)
    elf2 = getRucksack(f)
    elf3 = getRucksack(f)
    all = intersect(elf1, elf2, elf3)
    priority = 0
    ascii_a = Int('a')
    ascii_A = Int('A')
    for item in all
        ascii = Int(item)
        if ascii >= ascii_a
            priority = ascii - ascii_a + 1
        else
            priority = ascii - ascii_A + 27
        end
    end
    global sum += priority
    println("$all - $priority")
end  
println("total priority: $sum")