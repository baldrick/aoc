function main()
    f = open(ARGS[1], "r")
    total = 0
    overlapTotal = 0
    while ! eof(f)
        line = readline(f)
        if ignore(line)
            continue
        end
        (a, b) = ranges(line)
        enc = false
        if enclosed(a, b) || enclosed(b, a)
            total += 1
            #println("$a,$b enclosed")
            enc = true
        end
        if enc || overlap(a, b)
            if ! enc
                println("$a,$b overlap")
            end
            overlapTotal += 1
        end
    end
    println("enclosed = $total, overlap = $overlapTotal")
end

function ignore(line)
    return line == "" || startswith(line, "#")
end

function ranges(line)
    r = r"([0-9]+)-([0-9]+),([0-9]+)-([0-9]+)"
    m = match(r, line)
    return ((parse(Int, m[1]), parse(Int, m[2])), (parse(Int, m[3]), parse(Int, m[4])))
end

function enclosed(a, b)
    (mina, maxa) = a
    (minb, maxb) = b
    return mina >= minb && maxa <= maxb
end

# 3-5,5-6 maxa <= minb = true (5<=5)
# 5-6,3-5 maxb <= mina = true (5<=5)
# 7-10,3-9 mina=7, maxa=10, minb=3, maxb=9
# 56-56,16-55 mina=max=56, minb=16, maxb=55
function overlap(a, b)
    (mina, maxa) = a
    (minb, maxb) = b
    return (mina <= minb && maxa >= minb) || (minb <= maxa && maxb >= mina)
end

main()
