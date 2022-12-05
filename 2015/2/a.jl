f = open(ARGS[1], "r")

function getPaper(x, y, z)
    s1 = x*y
    s2 = x*z
    s3 = y*z
    area = 2 * s1 + 2 * s2 + 2 * s3
    return (area, minimum([s1,s2,s3]))
end

r=r"([0-9]+)x([0-9]+)x([0-9]+)"
totalArea = 0
while ! eof(f)
    line = readline(f)
    m = match(r, line)
    (area, slack) = getPaper(parse(Int, m[1]), parse(Int, m[2]), parse(Int, m[3]))
    global totalArea += area + slack
end
println("Total area = $totalArea")