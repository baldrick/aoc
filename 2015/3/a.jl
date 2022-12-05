f = open(ARGS[1], "r")

function visit(x,y)
    push!(houses, (x,y))
end

function move(x,y,cmd)
    if cmd == '<'
        return (x-1, y)
    elseif cmd == '>'
        return (x+1, y)
    elseif cmd == '^'
        return (x, y+1)
    end
    return (x, y-1)
end

# minor change required for part 1 - just remove rx,ry related stuff.
houses = Set()
line = readline(f)
x=0
y=0
rx=0
ry=0
visit(x,y)
for i in 1:length(line)
    if i%2 == 0
        continue
    end
    global (x,y) = move(x,y,line[i])
    global (rx,ry) = move(rx,ry,line[i+1])
    visit(x,y)
    visit(rx,ry)
end

visited = length(houses)
println("visited $visited houses")