f = open(ARGS[1], "r")

function hasNaughtyPairs(s)
    r = r"ab|cd|pq|xy"
    return match(r, s) != nothing
end

function countVowels(s)
    r = r"[b-df-hj-np-tv-z]"
    v = replace(s, r => "")
    return length(v)
end

function hasDoubleLetters(s)
    for i in 1:length(s)-1
        if s[i] == s[i+1]
            return true
        end
    end
    return false
end

naughty = 0
nice = 0
while ! eof(f)
    line = readline(f)
    if startswith(line, "#")
        continue
    end
    if hasNaughtyPairs(line)
        global naughty += 1
        continue
    end
    if countVowels(line) < 3
        global naughty += 1
        continue
    end
    if ! hasDoubleLetters(line)
        global naughty += 1
        continue
    end
    global nice += 1
    #ab, cd, pq, or xy
end
println("There are $nice nice words, $naughty naughty ones")
