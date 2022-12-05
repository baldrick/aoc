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

function hasRepeatedNonOverlappingLetterPairs(s)
    for i in 1:length(s)-1
        l=s[i]
        r=s[i+1]
        p="$l$r"
        if findnext(p, s, i+2) != nothing
            return true
        end
    end
    return false
end

function hasRepetitionOneLetterApart(s)
    for i in 1:length(s)-2
        if s[i] == s[i+2]
            return true
        end
    end
    return false
end

#nice if
# It contains a pair of any two letters that appears at least twice in the string without overlapping,
# like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).

naughty = 0
nice = 0
while ! eof(f)
    line = readline(f)
    if startswith(line, "#")
        continue
    end
    if ! hasRepetitionOneLetterApart(line)
        global naughty += 1
        continue
    end
    if ! hasRepeatedNonOverlappingLetterPairs(line)
        global naughty += 1
        continue
    end
    global nice += 1
end
println("There are $nice nice words, $naughty naughty ones")
