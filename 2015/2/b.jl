f = open(ARGS[1], "r")

function getRibbon(x, y, z)
    s1 = x*y
    s2 = x*z
    s3 = y*z
    m1 = minimum([x,y,z])
    m2 = 0
    if m1 == x
        m2 = minimum([y,z])
    elseif m1 == y
        m2 = minimum([x,z])
    else
        m2 = minimum([x,y])
    end
    minPerimeter = 2*m1 + 2*m2
    volume = x*y*z
    println("$x x $y x $z = p$m1,$m2 v=$volume")
    return volume + minPerimeter
end

r=r"([0-9]+)x([0-9]+)x([0-9]+)"
totalRibbon = 0
while ! eof(f)
    line = readline(f)
    m = match(r, line)
    global totalRibbon += getRibbon(parse(Int, m[1]), parse(Int, m[2]), parse(Int, m[3]))
end
println("Total ribbon = $totalRibbon")