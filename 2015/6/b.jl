f = open(ARGS[1], "r")

function brightness(s, v)
    (lx, rx, by, ty) = interpret(s)
    global lights[lx:rx, by:ty] = changeBrightness.(lights[lx:rx, by:ty], v)
end

function interpret(s)
    r = r"([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)"
    m = match(r, s)
    return (parse(Int, m[1])+1, parse(Int, m[3])+1, parse(Int, m[2])+1, parse(Int, m[4])+1)
end

function changeBrightness(s, v)
    return maximum([s+v,0])
end

lights = zeros(1000,1000)
while ! eof(f)
    line = readline(f)
    if startswith(line, "turn on")
        brightness(line[9:end], 1)
    elseif startswith(line, "turn off")
        brightness(line[10:end], -1)
    elseif startswith(line, "toggle")
        brightness(line[8:end], 2)
    else
        println("ERROR for $line")
        exit(1)
    end
end

c = 0
for light in lights
    global c += light
end

println("Total brightness is $c")
