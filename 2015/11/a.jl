function main()
    f = open(ARGS[1], "r")
    while !eof(f)
        line = readline(f)
        pwd = to_numbers(line)
        println("$line = $pwd = $(to_pwd(pwd))")
        println("$line -> $(next_password(pwd))")
    end
end

function to_numbers(line)
    pwd = []
    for i in 1:length(line)
        c = line[i]
        j = Int(c) - Int('a') + 1
        push!(pwd, j)
    end
    for i in 1:length(pwd)
        j = pwd[i]
        if j == 9 || j == 12 || j == 15 # i,o,l
            pwd[i] += 1
            for k in i+1:length(pwd)
                pwd[k] = 1
            end
            return pwd
        end
    end
    return pwd
end

function to_pwd(p)
    pwd = ""
    for i in p
        pwd = "$pwd$(Char(i+Int('a')-1))"
    end
    return pwd
end

function next_password(p)
    while !test(inc(p, length(p)))
    end
    return to_pwd(p)
end

function inc(p, n)
    if rollInc(p, n)
        inc(p, n-1)
    end
    return p
end

function rollInc(p, n)
    if n < 1
        println("rollInc $p at $n !!")
        exit(-1)
    end
    p[n] += 1
    if p[n] == 9 || p[n] == 12 || p[n] == 15 # i, o, l
        p[n] += 1
    elseif p[n] > 26
        p[n] = 1
        return true
    end
    return false
end

#=
Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz.
They cannot skip letters; abd doesn't count.

Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters.

Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
=#
function test(p)
    return testIncStraight(p) && testIOL(p) && testTwoDifferentPairs(p)
end

function testIncStraight(p)
    for i in 1:length(p)-2
        if p[i]+1 == p[i+1] && p[i]+2 == p[i+2]
            return true
        end
    end
    return false
end

# Returns true if p contains i, o or l.
function testIOL(p)
    r = r"[iol]"
    pwd = ""
    for i in p
        pwd="$pwd$(Char(i))"
    end
    return !occursin(r, pwd)
end

function testTwoDifferentPairs(p)
    first_pair = 0
    for i in 1:length(p)-1
        if p[i] == p[i+1]
            first_pair = i
            break
        end
    end
    if first_pair == 0
        return false
    end
    for i in first_pair+2:length(p)-1
        if p[i] == p[i+1] && p[i] != p[first_pair]
            return true
        end
    end
    return false
end

# pwds = ["hijklmmn", "abbceffg", "abbcegjk"]
# for pwd in pwds
#     println("1: $(testIncStraight(pwd))")
#     println("2: $(testIOL(pwd))")
#     println("3: $(testTwoDifferentPairs(pwd))")
# end

main()
