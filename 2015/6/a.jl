f = open(ARGS[1], "r")

function toggle(s)
    (lx, rx, by, ty) = interpret(s)
    global lights[lx:rx, by:ty] = flip.(lights[lx:rx, by:ty])
end

function interpret(s)
    r = r"([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)"
    m = match(r, s)
    return (parse(Int, m[1])+1, parse(Int, m[3])+1, parse(Int, m[2])+1, parse(Int, m[4])+1)
end

function set(s, v)
    (lx, rx, by, ty) = interpret(s)
    global lights[lx:rx, by:ty] .= v
end

function flip(x)
    return !x
end

lights = falses(1000,1000)
while ! eof(f)
    line = readline(f)
    if startswith(line, "turn on")
        set(line[9:end], true)
    elseif startswith(line, "turn off")
        set(line[10:end], false)
    elseif startswith(line, "toggle")
        toggle(line[8:end])
    else
        println("ERROR for $line")
        exit(1)
    end
end

c = 0
for light in lights
    if light
        global c += 1
    end
end

println("There are $c lights on")
