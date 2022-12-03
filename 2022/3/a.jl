f = open(ARGS[1], "r")

sum = 0
while ! eof(f)
    line = readline(f)
    itemCount = length(line)
    c1 = Set()
    c2 = Set()
    mid = floor(Int, itemCount/2)
    for itemIndex in 1:mid
        push!(c1, line[itemIndex])
    end
    for itemIndex in mid+1:itemCount
        push!(c2, line[itemIndex])
    end
    both = intersect(c1, c2)
    priority = 0
    ascii_a = Int('a')
    ascii_A = Int('A')
    for item in both
        ascii = Int(item)
        if ascii >= ascii_a
            priority = ascii - ascii_a + 1
        else
            priority = ascii - ascii_A + 27
        end
    end
    global sum += priority
    println("$both - $priority")
end  
println("total priority: $sum")